build:
	echo "Building..."
	go run main.go
run:
	echo "Running..."
	go run main.go

compile:
	echo "Compiling for someone else..."
	GOOS=linux GOARCH=arm64 go build -o bin/linux-arm64 main.go
	#GOOS=darwin GOARCH=amd64 go build -o bin/darwin-intel main.go
	#GOOS=windows GOARCH=amd64 go build -o bin/windows-intel.exe main.go

deploy:
	echo "Deploying..."
	go build -o /var/www/api/backend main.go
	sudo systemctl restart api.service

all:
	build
