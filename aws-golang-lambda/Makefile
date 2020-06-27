all: build deploy

go_build:
	GOOS=linux GOARCH=amd64 go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o build/bin/app .

tf_init:
	terraform init infra

tf_apply:
	terraform apply --auto-approve infra

tf_destroy:
	terraform destroy --auto-approve infra

.PHONY: build
build: go_build

.PHONY: deploy
deploy: tf_init tf_apply

.PHONY: clean
clean: tf_init tf_destroy
