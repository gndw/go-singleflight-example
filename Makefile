run-app:
	@echo "Running App..."
	@go run main.go

hit-ok:
	@curl -i "localhost:4000/ok"

hit-ok-with-data:
	@curl -i "localhost:4000/ok-with-data?id=r001"

hit-ok-with-sf:
	@curl -i "localhost:4000/ok-with-sf?id=sf001"

hit-parallel-ok-with-data:
	ab -n 100 -c 20 -l localhost:4000/ok-with-data?id=r002

hit-parallel-ok-with-sf:
	ab -n 100 -c 20 -l localhost:4000/ok-with-sf?id=sf002

hit-ok-with-sf-pointer:
	@curl -i "localhost:4000/ok-with-sf-pointer"

get-concurrent-with-data:
	@curl "localhost:4000/get-concurrent?id=r002" | jq .

get-concurrent-with-sf:
	@curl "localhost:4000/get-concurrent?id=sf002" | jq .