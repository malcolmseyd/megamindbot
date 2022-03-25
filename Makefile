# executable will have same name as current directory
all:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build .

clean:
	rm $(shell basename `pwd`)