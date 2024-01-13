compilerun:
	go build -o getGo main.go
	alias sget="./getGo"

clean:
	rm -rf getGo
	rm -rf *.jpg