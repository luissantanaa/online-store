package test

import (
	"io"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/luissantanaa/online-store/models"
	"github.com/luissantanaa/online-store/pkg/db"
	"github.com/luissantanaa/online-store/tools"

	"github.com/luissantanaa/online-store/store-service/routes"
)

func init() {
	db.ConnectDb()
}

var item = models.Item{
	Name:     "test",
	Quantity: 99,
}

var app *fiber.App

func init() {
	app = fiber.New()
	routes.SetupRoutes(app)
}

func TestAddItems(t *testing.T) {

	tests := []struct {
		description   string
		route         string // input route
		method        string // input method
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "add item",
			route:         "/api/v1/item",
			method:        "POST",
			body:          strings.NewReader(item.String()),
			expectedError: false,
			expectedCode:  200,
		},
	}

	log.Print(strings.NewReader(item.String()))
	log.Print(item.String())

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case.
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, err := app.Test(req, -1) // the -1 disables request latency
		if err != nil {
			log.Fatal(err)
		}

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses,
		// the next test case needs to be processed.
		if test.expectedError {
			continue
		}

		// Verify, if the status code is as expected.
		assert.Equalf(t, test.expectedCode, resp.StatusCode, tools.ReadResponse(resp))
	}
}
