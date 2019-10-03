package main

import (
	. "github.com/aearly/netlify-ns1-service/openapi"
	"github.com/labstack/echo/v4"
)

type Handlers struct{}

// get a list of registered domains// (GET /domains)
func (h Handlers) GetDomains(ctx echo.Context) error {
	return nil
}

// create a new domain record// (PUT /domains)
func (h Handlers) CreateDomain(ctx echo.Context) error {
	return nil
}

// get full info about a domain// (GET /domains/{domain})
func (h Handlers) GetDomain(ctx echo.Context, domain Domain) error {
	return nil
}

// get full info about a domain record// (GET /domains/{domain}/{record}/{recordType})
func (h Handlers) GetRecord(ctx echo.Context, domain Domain, record Record, recordType RecordType) error {
	return nil
}

// create a new domain record// (PUT /domains/{domain}/{record}/{recordType})
func (h Handlers) CreateRecord(ctx echo.Context, domain Domain, record Record, recordType RecordType) error {
	return nil
}
