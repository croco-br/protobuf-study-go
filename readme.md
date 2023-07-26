#Dependencies

sudo apt install -y protobuf-compiler 
export PATH="$PATH:$(go env GOPATH)/bin"
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go install github.com/ktr0731/evans@latest

#generate

make generate

#run

make run

#test grpc calls (with evans installed on GOPATH)

evans -r repl