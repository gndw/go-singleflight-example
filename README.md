# go-singleflight-example

Package singleflight provides a duplicate function call suppression mechanism.
https://pkg.go.dev/golang.org/x/sync/singleflight

- do singleflight when you have multiple call with same parameter and same expected result at the same time
- don't do singleflight to mutate data
- be careful when data returned by singleflight is a pointer, all callers will receive the same pointer, doing mutation to that pointer will affect all callers

### Notes about this Example
- inside all endpoints, external call is simulated using time.sleep
- whenever concurrent external calls is happening, we are adding additional sleep time to simulate heavy process due to concurrency
- for testing purpose, this example will provide endpoint to count number of external call, max processing time, & average processing time

### How to Run
- run app by `make run-app`

### Testing without SingleFlight
- hit 100 request with 20 concurrent to endpoint **without singleflight** using apache benchmark with `make hit-parallel-ok-with-data` will result in 16 seconds-ish test
```
Concurrency Level:      20
Time taken for tests:   16.031 seconds
Complete requests:      100
Failed requests:        0
Total transferred:      19300 bytes
HTML transferred:       7600 bytes
Requests per second:    6.24 [#/sec] (mean)
Time per request:       3206.164 [ms] (mean)
Time per request:       160.308 [ms] (mean, across all concurrent requests)
Transfer rate:          1.18 [Kbytes/sec] received
```
- receive external call statistic with cmd `get-concurrent-with-data`
```
{
  "id": "r002",
  "notes": [
    "number of database call: 100",
    "max concurrency of database call: 20",
    "maximum processing time: 3002 ms",
    "average processing time: 2962 ms"
  ],
  "details": [
    "external request with 1100 ms",
    "external request with 3000 ms",
    "external request with 3000 ms",
    "external request with 3001 ms",
    "external request with 3001 ms",
    "external request with 3001 ms",
    "external request with 3001 ms",
    "external request with 3001 ms",
    ... // list go on
  ]
}
```
100 calls to database is made, with 20 concurrent request, resulting average process time is increasing up to 3002 ms and average 2962 ms.

### Testing with SingleFlight
- hit 100 concurrent requests to endpoint **with singleflight** using apache benchmark with `make hit-parallel-ok-with-sf` will result in 6 seconds-ish test
```
Concurrency Level:      20
Time taken for tests:   6.628 seconds
Complete requests:      100
Failed requests:        0
Total transferred:      19201 bytes
HTML transferred:       7501 bytes
Requests per second:    15.09 [#/sec] (mean)
Time per request:       1325.638 [ms] (mean)
Time per request:       66.282 [ms] (mean, across all concurrent requests)
Transfer rate:          2.83 [Kbytes/sec] received
```
- receive external call statistic with cmd `get-concurrent-with-sf`
```
{
  "id": "sf002",
  "notes": [
    "number of database call: 6",
    "max concurrency of database call: 1",
    "maximum processing time: 1102 ms",
    "average processing time: 1101 ms"
  ],
  "details": [
    "external request with 1102 ms",
    "external request with 1102 ms",
    "external request with 1100 ms",
    "external request with 1101 ms",
    "external request with 1102 ms",
    "external request with 1101 ms"
  ]
}
```
6 calls to database is made, with 1 concurrent request, resulting average process time is stable around 1102 ms.

### Conclusion
- Singleflight will definitely reduce your external call as long as it was all identical
- Effect of using singleflight will increase exponentially with your RPS, your service will survive higher RPS compared to non-singleflight