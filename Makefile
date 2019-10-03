OAPI_SCAFFOLD=openapi/service.gen.go

.PHONY: oapi-scaffold
oapi-scaffold: ${OAPI_SCAFFOLD}

${OAPI_SCAFFOLD}: service-oapi.yml
	oapi-codegen service-oapi.yml > ${OAPI_SCAFFOLD}

build:
	go build -o server
