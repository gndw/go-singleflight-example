# go-singleflight-example

Package singleflight provides a duplicate function call suppression mechanism.
https://pkg.go.dev/golang.org/x/sync/singleflight

- do singleflight when you have multiple call with same parameter and same expected result at the same time
- don't do singleflight to mutate data
- be careful when data returned by singleflight is a pointer, all callers will receive the same pointer, doing mutation to that pointer will affect all callers

### Notes about this Example
- inside all endpoints, external call is simulated using time.sleep
- whenever concurrent external calls is happening, we are adding additional sleep time to simulate heavy process due to concurrency
### How to Run
- run app by `make run-app`
- hit 20 concurrent requests to endpoint **without singleflight** using apache benchmark with `make hit-parallel-ok-with-data` will result in 4 seconds-ish test
```
Concurrency Level:      20
Time taken for tests:   4.014 seconds
Complete requests:      20
Failed requests:        0
Total transferred:      3860 bytes
HTML transferred:       1520 bytes
Requests per second:    4.98 [#/sec] (mean)
Time per request:       4014.292 [ms] (mean)
Time per request:       200.715 [ms] (mean, across all concurrent requests)
Transfer rate:          0.94 [Kbytes/sec] received
```
- hit 20 concurrent requests to endpoint **with singleflight** using apache benchmark with `make hit-parallel-ok-with-sf` will result in 2 seconds-ish test
```
Concurrency Level:      20
Time taken for tests:   2.213 seconds
Complete requests:      20
Failed requests:        0
Total transferred:      3841 bytes
HTML transferred:       1501 bytes
Requests per second:    9.04 [#/sec] (mean)
Time per request:       2213.152 [ms] (mean)
Time per request:       110.658 [ms] (mean, across all concurrent requests)
Transfer rate:          1.69 [Kbytes/sec] received
```