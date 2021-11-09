dev:
	goverwatch -command "go run main.go" -files "**/**/*.go server/chart.html"
run:
	go run *.go
build:
	go build
lint:
	gofmt -l -w -s *.go
stress:
	bash stress.sh
stress_stop:
	killall yes
