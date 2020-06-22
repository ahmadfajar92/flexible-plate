scaffold-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

scaffold-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

run-first-osx:
	make scaffold-osx && ./scaffold-osx

run-osx: 
	rm scaffold-osx && make scaffold-osx && ./scaffold-osx

run-linux-first: 
	make scaffold-linux && ./scaffold-linux

run-linux: 
	rm scaffold-linux && make scaffold-linux && ./scaffold-linux

run-lambda:
	make scaffold-linux && sam local start-api --debug --env-vars env-lambda.json
