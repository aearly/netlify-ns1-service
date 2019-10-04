package main

import (
	"encoding/json"
	"fmt"
	. "github.com/aearly/netlify-ns1-service/openapi"
	"github.com/alecthomas/jsonschema"
	"github.com/labstack/echo/v4"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
	"gopkg.in/ns1/ns1-go.v2/rest/model/filter"
	"net/http"
	"os"
)

func main2() {
	schema := jsonschema.Reflect(&dns.Record{})
	result, _ := json.MarshalIndent(schema, "", "  ")
	os.Stdout.Write(result)
}

type Handlers struct {
	apiKey string
}

type ErrorMessage struct {
	Message string `json:"message"`
}

// get a list of registered zones// (GET /zones)
func (h Handlers) GetZones(ctx echo.Context) error {
	client := ctx.(*CustomContext).Ns1Client
	zones, _, err := client.Zones.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, zones)
}

// create a new domain record// (PUT /zones/{zone})
func (h Handlers) CreateZone(ctx echo.Context, zoneName Zone) error {
	client := ctx.(*CustomContext).Ns1Client

	zone := new(dns.Zone)

	if err := ctx.Bind(zone); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorMessage{Message: err.Error()})
	}

	zone.Zone = string(zoneName)

	fmt.Println(zone.Zone)

	_, err := client.Zones.Create(zone)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, struct{}{})
}

// delete a zone// (DELETE /zones/{zone})
func (h Handlers) DeleteZone(ctx echo.Context, zoneName Zone) error {
	client := ctx.(*CustomContext).Ns1Client

	_, err := client.Zones.Delete(string(zoneName))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: err.Error()})
	}
	return ctx.String(http.StatusNoContent, "")
}

// get full info about a zone// (GET /zones/{zone})
func (h Handlers) GetZone(ctx echo.Context, zoneName Zone) error {
	client := ctx.(*CustomContext).Ns1Client

	zone, res, err := client.Zones.Get(string(zoneName))
	if res.StatusCode == 404 {
		return ctx.JSON(http.StatusNotFound, ErrorMessage{Message: err.Error()})
	}
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, zone)
}

// get full info about a domain record// (GET /zones/{zone}/{domain}/{recordType})
func (h Handlers) GetDomain(ctx echo.Context, zoneName Zone, domain Domain, recordType RecordType) error {
	client := ctx.(*CustomContext).Ns1Client

	record, res, err := client.Records.Get(string(zoneName), string(domain), string(recordType))

	if res.StatusCode == 404 {
		return ctx.JSON(http.StatusNotFound, ErrorMessage{Message: err.Error()})
	}
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, record)
}

// create a new domain record// (PUT /zones/{zone}/{domain}/{recordType})
func (h Handlers) CreateDomain(ctx echo.Context, zoneName Zone, domain Domain, recordType RecordType) error {
	client := ctx.(*CustomContext).Ns1Client

	record := new(dns.Record)

	if err := ctx.Bind(record); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorMessage{Message: err.Error()})
	}

	record.Zone = string(zoneName)
	record.Domain = string(domain)
	record.Type = string(recordType)
	if record.Filters == nil {
		record.Filters = []*filter.Filter{}
	}

	_, err := client.Records.Create(record)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, struct{}{})
}
