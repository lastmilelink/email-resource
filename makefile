all:
	cd check ; GOOS=linux GOARCH=386 go build -o check
	cd in    ; GOOS=linux GOARCH=386  go build -o in
	sudo docker build --tag lmlt/email-resource:latest .