# netlify-ns1-service

A service that wraps some features of <http://ns1.com> 's API.

Created as an interview project for Netlify.

## Usage

```bash
# build the project
make build

# start the service
NS1_APIKEY=yourapikey ./bin/netlify-ns1-service
```

## Development Plan

- [x] Generate a OpenAPI/Swagger definition for the API
- [x] Use <https://github.com/deepmap/oapi-codegen> to generate boilerplate
- [ ] Use <https://github.com/ns1/ns1-go> to implement routes
    + [ ] tests? use a mock?
- [ ] Also save zones/records in a sql database
    + [ ] https://github.com/Masterminds/squirrel
    + [ ] https://github.com/mattn/go-sqlite3
- [ ] ~Create a UI for the api - likely CLI~
    + [ ] ~Use OAPI definition to generate a client~

## Development Notes

- This was my first project in Go bigger than toy examples (but not the first REST API I've created).  A lot of time was spent getting familiar with how projects are generally set up and the broader ecosystem of packages.
- I decided to start with an OpenAPI/Swagger specification because it would allow me to generate a lot of the boilerplate (>2000 LOC!) and provide the first layer of documentation
- I decided to copy NS1's API, since the service is mostly a proxy for it.  The first pass was to define it using OpenAPI
- I generated schemas for the requests/responses from the structs in the NS1 go client.  This saves some time.
- The code generator uses the Echo http framework.  Its similar to Node frameworks I've used in the past, so it felt a bit familiar.
- 
