go build
go install . 
protoc --proto_path=. --proto_path=./proto --go_out=paths=source_relative:. --go-asynq_out=paths=source_relative:.\ test/example.proto