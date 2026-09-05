package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	rec := httptest.NewRecorder()

	newRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rec.Code,
		)
	}

	if !strings.Contains(
		rec.Body.String(),
		`"status":"ok"`,
	) {
		t.Fatalf(
			"unexpected response body: %s",
			rec.Body.String(),
		)
	}
}

func TestHealthHandlerRejectsPOST(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/health",
		nil,
	)

	rec := httptest.NewRecorder()

	newRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf(
			"expected status 405, got %d",
			rec.Code,
		)
	}
}