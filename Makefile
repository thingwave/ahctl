SRC=src/main/main.go src/main/datamodels.go

all:
	go build -o ahctl $(SRC)

all-arm64:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o ahctl src/main/main.go src/main/datamodels.go

urun:
	./ahctl --sr="http://127.0.0.1:8443/serviceregistry" --cmd=sr-echo

srun:
#	./ahctl --sr="https://127.0.0.1:8443/serviceregistry" --cmd=sr-echo --cafile=/tmp/truststore.pem --cert=/tmp/service_registry.pem --key=/tmp/service_registry.key
	./ahctl --cafile=../eclipse-arrowhead/certificates/testcloud2/testcloud2.pem --cert=../eclipse-arrowhead/DataManager/certificates/datamanager.pem --key=../eclipse-arrowhead/DataManager/certificates/datamanager.key --cmd=sr-echo


deb:
	mkdir -p packages/usr/local/bin
	cp ahctl packages/usr/local/bin/
	dpkg-deb --build packages
	mv packages.deb ahctl_amd64.deb
	#mv packages.deb ahctl_arm64.deb

clean:
	rm ./ahctl
