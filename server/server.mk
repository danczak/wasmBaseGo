server: *.go
	go build -o ../app

clean:
	rm -f ../app