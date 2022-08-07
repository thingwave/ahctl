all:
	go build -o ahctl src/main/main.go src/main/datamodels.go

urun:
	./ahctl --sr="http://127.0.0.1:8443/serviceregistry" --cmd=sr-echo

srun:
	./ahctl --sr="https://127.0.0.1:8443/serviceregistry" --cmd=sr-echo --cafile=/tmp/truststore.pem --cert=/tmp/service_registry.pem --key=/tmp/service_registry.key

deb:
	cp ahctl packages/usr/local/bin/
	dpkg-deb --build packages
	mv packages.deb ahctl_amd64.deb

clean:
	rm ./ahctl


