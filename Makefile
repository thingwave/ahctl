all:
	go build -o cmd/ahctl cmd/main/main.go

urun:
	./cmd/ahctl --sr="http://127.0.0.1:8443/serviceregistry/echo" --cmd=test-sr

srun:
	./cmd/ahctl --cmd=test-sr --cafile=/tmp/truststore.pem --cert=/tmp/service_registry.pem --key=/tmp/service_registry.key

clean:
	rm ./cmd/ahctl


