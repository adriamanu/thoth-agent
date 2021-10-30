run:
	go run *.go
lint:
	gofmt -l -w -s *.go
stress:
	bash stress.sh
stress_stop:
	killall yes