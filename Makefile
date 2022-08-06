all:
	go build -o ahctl src/main/main.go src/main/datamodels.go

urun:
	./ahctl --sr="http://127.0.0.1:8443/serviceregistry" --cmd=echo

srun:
	./ahctl --sr="https://127.0.0.1:8443/serviceregistry" --cmd=echo --cafile=/tmp/truststore.pem --cert=/tmp/service_registry.pem --key=/tmp/service_registry.key

clean:
	rm ./ahctl


