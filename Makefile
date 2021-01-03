clean:
	rm -R dist
compile:
	GOARM=6 GOARCH=arm GOOS=linux go build -o dist/linux/arm6/iothome cmd/main.go
	GOARM=7 GOARCH=arm GOOS=linux go build -o dist/linux/arm7/iothome cmd/main.go
create-checksum:
	echo `openssl sha256 dist/linux/iothome | awk '{print $2}'` 
	openssl sha256 dist/windows/iothome.exe | awk '{print $2}' > dist/windows/iothome.sha256.txt
build: compile create-checksum
run:
	go run cmd/main.go
