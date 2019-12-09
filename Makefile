
all:tcpServer


tcpServer:
	go build -o tcpServer main.go msgHandle.go


.PHONY:clean
clean:
	rm -rf tcpServer
