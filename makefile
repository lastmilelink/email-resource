all:
	cd check ; go build -o check
	cd in ; go build -o in
	sudo docker build --tag email-resource:latest .