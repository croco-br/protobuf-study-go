generate:
	protoc --go_out=. --go-grpc_out=. proto/*.proto
run:
	go run ./main.go
clear:
	kill $(lsof -t -i :50051) 
