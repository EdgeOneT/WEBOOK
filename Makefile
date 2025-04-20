.PHONY: docker
docker:
	@rm webook || true
	@go mod tidy
	@$env:GOOS="linux"; $env:GOARCH="arm"; go build -tags=k8s -o webook .
	@docker rmi -f edge/webook:v0.0.1
	@docker build -t edge/webook:v0.0.1 .


#windows使用docker：中@里面的：一条一条来编译
