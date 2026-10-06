package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func doHello(t *testing.T, query string) (*http.Response, []byte) {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, "/hello"+query, nil)
	rec := httptest.NewRecorder()
	Hello(rec, req)
	res := rec.Result()
	body := rec.Body.Bytes()
	return res, body
}

func TestHelloWithName(t *testing.T) {
	res, body := doHello(t, "?name=Ada")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", res.StatusCode)
	}
	var got greetingResponse
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Message != "Hello, Ada!" {
		t.Fatalf("message = %q, want %q", got.Message, "Hello, Ada!")
	}
}

func TestHelloWithoutName(t *testing.T) {
	res, body := doHello(t, "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", res.StatusCode)
	}
	var got greetingResponse
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Message != "Hello, World!" {
		t.Fatalf("message = %q, want %q", got.Message, "Hello, World!")
	}
}

func TestHelloWithEmptyName(t *testing.T) {
	res, body := doHello(t, "?name=")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", res.StatusCode)
	}
	var got greetingResponse
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Message != "Hello, World!" {
		t.Fatalf("message = %q, want %q", got.Message, "Hello, World!")
	}
}

func TestHelloWithOverlongName(t *testing.T) {
	res, body := doHello(t, "?name="+strings.Repeat("a", 101))
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", res.StatusCode)
	}
	var got errorResponse
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Error == "" {
		t.Fatalf("expected non-empty error message")
	}
}
