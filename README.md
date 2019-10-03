# netlify-ns1-service

A service that wraps some features of <http://ns1.com> 's API.

Created as an interview project for Netlify.

## Creation Process

- [ ] Generate a OpenAPI/Swagger definition for the API
- [ ] Use <https://github.com/deepmap/oapi-codegen> to generate boilerplate
- [ ] Use <https://github.com/ns1/ns1-go> to implement routes
    + [ ] hard code API token for now
    + [ ] tests? use a mock?
- [ ] Also save zones/records in a sql database
    + [ ] https://github.com/Masterminds/squirrel
    + [ ] https://github.com/mattn/go-sqlite3
- [ ] Create a UI for the api - likely CLI
    + [ ] Use OAPI definition to generate a client
