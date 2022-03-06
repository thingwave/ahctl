all:
	go build -o cmd/ahtcl cmd/main/main.go

run:
	./cmd/ahctl --cmd=test-sr --cafile=/tmp/truststore.pem --cert=/tmp/service_registry.pem --key=/tmp/service_registry.key

clean:
	rm ./cmd/ahctl


