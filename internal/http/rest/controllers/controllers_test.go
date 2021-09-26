package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func toJsonString(i interface{}) string {
	j, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(j)
}

func testPath(t *testing.T, app *fiber.App, method, path string, body interface{}, expectedStatusCode int, expectedBody interface{}) error {
	var req *http.Request
	if body != nil {
		rBody, err := json.Marshal(body)
		if err != nil {
			t.Logf(err.Error())
			return err
		}
		req = httptest.NewRequest(method, path, bytes.NewReader(rBody))
	} else {
		req = httptest.NewRequest(method, path, nil)
	}
	req.Header.Add(`Content-Type`, `application/json`)

	res, err := app.Test(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	assert.Equal(t, expectedStatusCode, res.StatusCode)
	assert.Equal(t, expectedBody, string(resBody))
	return nil
}
