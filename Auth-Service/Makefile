CURRENT_DIR=$(shell pwd)
DBURL=postgres://postgres:root@localhost:5432/learning_db?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -path migrations -database '$(DBURL)' -verbose up

mig-down:
	migrate -path migrations -database '$(DBURL)' -verbose down

mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir migrations -seq insert_table

gen-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go-grpc_out=. proto/*proto
	protoc --go_out=. --go-grpc_out=. proto/*proto
	
gen-protoAll:
#proto fileni hammasini bittada generatsiya qilish 
	protoc --go_out=./ \
	--go-grpc_out=./ \
	protos/*.proto


doc-run:
	docker run --network="host" --name libCont lib_service

doc-con:
	docker run --network="host" --publish  9090:3030 recruitment-service

doc-network:
	docker network create library

doc-run-allow:
	docker run --network="library" -p 9090:2121 library_service

doc-continer-all-del:
	docker rm $(docker ps -aq)
	docker rmi $(docker ps images)

doc-network-all:
	docker network ls
