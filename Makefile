gettool:
	go get -u github.com/rakyll/hey

loadtest:
	hey -n 10000 -c 100 -m POST -d '{"name":"john farmer","phone":"0123456789","email":"test@mail.com"}' http://localhost:8080/

buildgo:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -o app -a -ldflags '-w -s' main.go

builddocker:
	docker build -t "mongmx/gotest" .

startdocker:
	docker run -dt -p 8080:8080 --name gotest mongmx/gotest