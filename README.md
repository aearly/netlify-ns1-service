# netlify-ns1-service

A service that wraps some features of <http://ns1.com> 's API.

Created as an interview project for Netlify.

## Auto-generated Swagger Docs

<https://app.swaggerhub.com/apis/ae-sandbox/netlify-ns1-service/0.0.1>

## Usage

```bash
# build the project
make build

# start the service
NS1_APIKEY=yourapikey ./bin/netlify-ns1-service
```

Example (using HTTPie):
```
$ http :5000/zones/asdfae.com/www.asdfae.com/A

HTTP/1.1 200 OK
Content-Length: 217
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Oct 2019 23:33:42 GMT

{
    "answers": [
        {
            "answer": [
                "1.2.3.5"
            ],
            "id": "5d97d53ff2abd000010222cf"
        }
    ],
    "domain": "www.asdfae.com",
    "filters": [],
    "id": "5d96eceac94a9000019d308c",
    "meta": {},
    "ttl": 3600,
    "type": "A",
    "use_client_subnet": true,
    "zone": "asdfae.com"
}

```

## Development Plan

- [x] Generate a OpenAPI/Swagger definition for the API
- [x] Use <https://github.com/deepmap/oapi-codegen> to generate boilerplate
- [x] Use <https://github.com/ns1/ns1-go> to implement routes
    + [ ] tests? use a mock?
- [ ] Also save zones/records in a sql database
    + [ ] https://github.com/Masterminds/squirrel
    + [ ] https://github.com/mattn/go-sqlite3
- [ ] ~Create a UI for the api - likely CLI~
    + [ ] ~Use OAPI definition to generate a client~

## Development Notes

- This was my first project in Go bigger than toy examples (but not the first REST API I've created).  A lot of time was spent getting familiar with the language, how projects are generally set up, and the broader ecosystem of packages.
- I decided to start with an OpenAPI/Swagger specification because it would allow me to generate a lot of the boilerplate (>2000 LOC!) and provide the first layer of documentation
- I decided to copy NS1's API, since the service is mostly a proxy for it.  The first pass was to define it using OpenAPI
- I generated schemas for the requests/responses from the structs in the NS1 go client.  This saves some time.
- The code generator uses the Echo http framework.  Its similar to Node frameworks I've used in the past, so it felt a bit familiar.
- Implementing handlers was pretty straightforward, converting request bodies into structs, and then passing them to the NS1 api client
- I ran out of time after getting the basics working.  Unfortunately, it's a glorified proxy, but it's a good starting point for the next steps.

### Things cut for time
- Testing
    + I was planing on creating a `handlers_test.go` file where I would mock the context (and NS1 API client, and eventually the database connection) and test all the handler methods
- The database integration
    + I was planning on creating a simple SQLite database, that updated on write operations after the NS1 requests completed.  Each request would get a DB connection where the handler could make queries using a query builder
- Docs
    + Docs are automatically generated from the OpenAPI spec with swagger hub
    + This is a basic level, was hoping to flesh them out with more examples and a guide on how/why you'd use this API
- CLI UI
    + nowhere near enough time
