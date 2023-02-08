SRC=src/main/main.go src/main/datamodels.go
CAFILE=../eclipse-arrowhead/certificates/testcloud2/testcloud2.pem
CERTFILE=../eclipse-arrowhead/DataManager/certificates/datamanager.pem
KEYFILE=../eclipse-arrowhead/DataManager/certificates/datamanager.key


all:
	go build -o ahctl $(SRC)

all-arm64:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o ahctl $(SRC)

urun:
	./ahctl --sr="http://127.0.0.1:8443/serviceregistry" --cmd=sr-echo

srun:
#	./ahctl --sr="https://127.0.0.1:8443/serviceregistry" --cmd=sr-echo --cafile=/tmp/truststore.pem --cert=/tmp/service_registry.pem --key=/tmp/service_registry.key
	./ahctl --cafile=$(CAFILE) --cert=$(CERTFILE) --key=$(KEYFILE) --cmd=sr-echo

deb:
	mkdir -p packages/usr/local/bin
	cp ahctl packages/usr/local/bin/
	mkdir -p packages/usr/local/share/man/man1/
	cp man/ahctl.1 packages/usr/local/share/man/man1
	dpkg-deb --build packages
	mv packages.deb ahctl_amd64.deb
	#mv packages.deb ahctl_arm64.deb

clean:
	rm ./ahctl
