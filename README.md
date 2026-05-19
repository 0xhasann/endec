go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
protoc --go_out=. person.proto  
go get google.golang.org/protobuf
