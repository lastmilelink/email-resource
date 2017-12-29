all:
	cd check ; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o check *.go
	cd in    ; CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o in *.go
	cd out   ; CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o out *.go
	docker build --tag lmlt/email-resource:latest .