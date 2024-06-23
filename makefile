all:
	go build -o hellow main.go

.PHONY
clean:
	rm hellow