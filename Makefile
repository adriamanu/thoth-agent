run:
	go run *.go
lint:
	gofmt -l -w -s *.go
build:
	go build *.go
stress:
	bash stress.sh
stress_stop:
	killall yes
