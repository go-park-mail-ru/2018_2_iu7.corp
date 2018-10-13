all: build

install:
	go get -v .

generate:
	easyjson -all ./models/*.go
	protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. ./services/rpc/*.proto

build: generate install
	go build -v -o ./build/profile-service .

test: generate install
	go test -v .

clean:
	rm -r -f ./build
	rm -f ./models/*_easyjson.go
	rm -f ./services/rpc/*.pb.go
	rm -f ./services/rpc/*.micro.go
