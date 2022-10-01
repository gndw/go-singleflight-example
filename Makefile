run-app:
	@echo "Running App..."
	@go run main.go

hit-ok:
	@curl -i "localhost:8080/ok"

hit-parallel-ok-with-sf:
	ab -n 20 -c 20 -l localhost:8080/ok-with-sf

hit-parallel-ok-with-data:
	ab -n 20 -c 20 -l localhost:8080/ok-with-data

hit-ok-with-data:
	@curl -i "localhost:8080/ok-with-data"

hit-ok-with-sf:
	@curl -i "localhost:8080/ok-with-sf"

hit-ok-with-sf-key-1:
	@curl -i "localhost:8080/ok-with-sf-key?id=1"

hit-ok-with-sf-key-2:
	@curl -i "localhost:8080/ok-with-sf-key?id=2"

hit-ok-with-sf-pointer:
	@curl -i "localhost:8080/ok-with-sf-pointer"