.PHONY: build plan

build:
	GOOS=darwin go build -o bin/terraform-provider-cronitor-darwin
	GOOS=linux go build -o bin/terraform-provider-cronitor-linux

plan: build
	terraform init
	terraform plan