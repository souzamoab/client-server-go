all: cliente servidor

cliente: cliente.go
	go build cliente.go

servidor: servidor.go
	go build servidor.go

clean:
	rm -f cliente servidor
