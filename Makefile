.PHONY: build plan

build:
	go build -o terraform-provider-cronitor

plan: build
	terraform init
	terraform plan