build: parser
	go build ./...

parser: parser/parser.go.y
	goyacc -o parser/parser.go $^

test: parser
	go test -v ./...
