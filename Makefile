clean:
	rm -R dist
compile:
	GOOS=linux go build -o dist/linux/iothome cmd/main.go
	GOOS=windows go build -o dist/windows/iothome.exe cmd/main.go
create-checksum:
	echo `openssl sha256 dist/linux/iothome | awk '{print $2}'` 
	openssl sha256 dist/windows/iothome.exe | awk '{print $2}' > dist/windows/iothome.sha256.txt
build: compile create-checksum
run:
	go run cmd/main.go
