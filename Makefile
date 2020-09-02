.PHONY: build plan test

build:
	GOOS=darwin go build -o bin/terraform-provider-cronitor-darwin
	GOOS=linux go build -o bin/terraform-provider-cronitor-linux

plan: build
	terraform init
	terraform plan

test:
	TF_ACC=1 go test -v ./cronitor/...
