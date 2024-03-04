package test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/luissantanaa/online-store/models"
	"github.com/luissantanaa/online-store/pkg/db"
	"github.com/luissantanaa/online-store/tools"

	"github.com/luissantanaa/online-store/client-service/routes"
)

func init() {
	db.ConnectDb()
}

var client = models.Client{
	Username: "test",
	Password: "test",
}

var app *fiber.App

func init() {
	app = fiber.New()
	routes.SetupRoutes(app)
}

func TestAddClient(t *testing.T) {

	tests := []struct {
		description   string
		route         string // input route
		method        string // input method
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "add client SUCCESS",
			route:         "/api/v1/signup",
			method:        "POST",
			body:          strings.NewReader(client.String()),
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "add client FAIL - Invalid Body",
			route:         "/api/v1/signup",
			method:        "POST",
			body:          strings.NewReader("{\"Name\": \"item\"}"),
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "add client FAIL - Empty Body",
			route:         "/api/v1/signup",
			method:        "POST",
			body:          strings.NewReader("{\"Username\": \"Password\"}"),
			expectedError: false,
			expectedCode:  400,
		},
	}

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

func TestLoginClient(t *testing.T) {

	tests := []struct {
		description   string
		route         string // input route
		method        string // input method
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "login client SUCCESS",
			route:         "/api/v1/login",
			method:        "POST",
			body:          strings.NewReader(client.String()),
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "login client FAIL - Invalid Body",
			route:         "/api/v1/login",
			method:        "POST",
			body:          strings.NewReader("{\"Name\": \"item\"}"),
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "login client FAIL - Empty Body",
			route:         "/api/v1/login",
			method:        "POST",
			body:          strings.NewReader("{\"Username\": \"Password\"}"),
			expectedError: false,
			expectedCode:  400,
		},
	}

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

func TestPermissionsClientFail(t *testing.T) {

	login_route := "http://localhost:8083/api/v1/login"
	get_clients_route := "/api/v1/clients"

	login_resp, login_err := http.Post(login_route, "application/json", strings.NewReader(client.String()))
	if login_err != nil {
		log.Fatalf("An Error Occured %v", login_err)
	}
	defer login_resp.Body.Close()

	login_resp_body, err := io.ReadAll(login_resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data map[string]interface{}
	decode_err := json.Unmarshal([]byte(login_resp_body), &data)

	if decode_err != nil {
		panic(decode_err)
	}

	token := data["token"]

	// Create a new http request with the route from the test case.
	test_req := httptest.NewRequest("GET", get_clients_route, strings.NewReader(""))
	test_req.Header.Set("Content-Type", "application/json")
	test_req.Header.Set("Authorization", token.(string))

	// Perform the request plain with the app.
	resp, err := app.Test(test_req, -1) // the -1 disables request latency
	if err != nil {
		log.Fatal(err)
	}

	// Verify, if the status code is as expected.
	assert.Equalf(t, http.StatusForbidden, resp.StatusCode, tools.ReadResponse(resp))
}
