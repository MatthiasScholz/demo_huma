module := demo_huma
init:
	go mod init $(module)

deps:
	go mod tidy

port := 8888
run:
	$(info "Starts to serve the API at port: $(port)")
	go run main.go

test:
	restish :$(port)

host := localhost
docs:
	open http://$(host):$(port)/docs

openapi:
	open http://$(host):$(port)/openapi.json
	open http://$(host):$(port)/schemas
	open http://$(host):$(port)/docs

localhost := host.docker.internal
apimock:
	$(info Requires a running service to get the OpenAPI spec. -> call 'make run' before.)
	docker pull danielgtaylor/apisprout
	docker run -p 8000:8000 danielgtaylor/apisprout http://$(localhost):$(port)/openapi.json

apimock-test:
	$(info FIXME incomplete example)
	open http://$(host):8000
