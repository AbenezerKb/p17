include .env
export


run:
	go run ./cmd/rest/main.go

migrateup:
	migrate -path ./initiator/migration -database "cockroachdb://root:@localhost:26257/smsgateway?sslmode=disable" -verbose up
	 
migratedown:
	migrate -path ./initiator/migration -database "cockroachdb://root:@localhost:26257/smsgateway?sslmode=disable" -verbose down 