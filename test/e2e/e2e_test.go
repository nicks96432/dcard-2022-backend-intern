package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

import (
	"dcard-2022-backend-intern/internal/config"
	server "dcard-2022-backend-intern/internal/server/http"
	"dcard-2022-backend-intern/internal/storage/sqlite"
)

var app *fiber.App
var c = config.Default()

func TestMain(m *testing.M) {
	db, err := sqlite.New(c)
	if err != nil {
		os.Exit(1)
	}
	app = server.New(db, c)

	os.Exit(m.Run())
}

func TestE2E_normal(t *testing.T) {
	// create request
	url, exp := "http://www.dcard.tw", time.Date(2487, time.June, 3, 10, 16, 0, 0, time.Local)
	body := fmt.Sprintf(`{ "url": "%s", "expireAt": "%s" }`, url, exp.Format(time.RFC3339))
	buf := bytes.NewBufferString(body)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/urls", buf)
	if err != nil {
		t.Fatalf("cannot create new request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// send POST /api/v1/urls request
	res, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("POST /api/v1/urls: expect success, got error %v\n", err)
	}
	if expected, got := 200, res.StatusCode; expected != got {
		t.Fatalf("POST /api/v1/urls: expect status %d, got %d\n", expected, got)
	}
	if expected, got := "application/json", res.Header.Get("Content-Type"); expected != got {
		t.Fatalf("POST /api/v1/urls: expect Content-Type %s, got %s\n", expected, got)
	}

	// parse response
	var v map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		t.Fatalf("POST /api/v1/urls: cannot decode response body: %v\n", err)
	}
	if len(v["id"].(string)) != 11 {
		t.Fatalf("POST /api/v1/urls: field id is empty\n")
	}
	if expected, got := fmt.Sprintf("http://%s:%d/%s", c.Hostname, c.Port, v["id"]), v["shortUrl"]; expected != got {
		t.Fatalf("POST /api/v1/urls: field shortUrl expected %s, got %s\n", expected, got)
	}

	// send GET /id request
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", v["id"].(string)), nil)
	if err != nil {
		t.Fatalf("cannot create new request: %v\n", err)
	}
	res, err = app.Test(req)
	if err != nil {
		t.Fatalf("POST /api/v1/urls: expect success, got error %v\n", err)
	}
	if expected, got := 301, res.StatusCode; expected != got {
		t.Fatalf("POST /api/v1/urls: expect status %d, got %d\n", expected, got)
	}
	if expected, got := url, res.Header.Get("Location"); expected != got {
		t.Fatalf("POST /api/v1/urls: expect Location %s, got %s\n", expected, got)
	}
}

func TestE2E_expired(t *testing.T) {
	// create request
	url := "http://www.dcard.tw"
	exp, err := time.Parse(time.RFC3339, "2021-02-08T09:20:41Z")
	if err != nil {
		t.Fatalf("cannot create new time: %v\n", err)
	}
	body := fmt.Sprintf(`{ "url": "%s", "expireAt": "%s" }`, url, exp.Format(time.RFC3339))
	buf := bytes.NewBufferString(body)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/urls", buf)
	if err != nil {
		t.Fatalf("cannot create new request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// send request
	res, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("POST /api/v1/urls: expect success, got error %v\n", err)
	}
	if expected, got := 200, res.StatusCode; expected != got {
		t.Fatalf("POST /api/v1/urls: expect status %d, got %d\n", expected, got)
	}
	if expected, got := "application/json", res.Header.Get("Content-Type"); expected != got {
		t.Fatalf("POST /api/v1/urls: expect Content-Type %s, got %s\n", expected, got)
	}

	var v map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		t.Fatalf("POST /api/v1/urls: cannot decode response body: %v\n", err)
	}
	if len(v["id"].(string)) != 11 {
		t.Fatalf("POST /api/v1/urls: field id is empty\n")
	}
	if expected, got := fmt.Sprintf("http://%s:%d/%s", c.Hostname, c.Port, v["id"]), v["shortUrl"]; expected != got {
		t.Fatalf("POST /api/v1/urls: field shortUrl expected %s, got %s\n", expected, got)
	}

	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", v["id"].(string)), nil)
	if err != nil {
		t.Fatalf("cannot create new request: %v\n", err)
	}

	res, err = app.Test(req)
	if err != nil {
		t.Fatalf("POST /api/v1/urls: expect success, got error %v\n", err)
	}
	if expected, got := 404, res.StatusCode; expected != got {
		t.Fatalf("POST /api/v1/urls: expect status %d, got %d\n", expected, got)
	}
}
