all:
	cd check ; GOOS=linux GOARCH=386 go build -o check *.go
	cd in    ; GOOS=linux GOARCH=386  go build -o in *.go
	cd out   ; GOOS=linux GOARCH=386  go build -o out *.go
	sudo docker build --tag lmlt/email-resource:latest .