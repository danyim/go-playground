deps:
	go get github.com/codegangsta/gin

watch:
	gin -a 8081 -p 6060 run main.go
