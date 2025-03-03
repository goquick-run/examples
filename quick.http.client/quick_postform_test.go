package quick

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/jeffotoni/quick"
)

// TestFormValue ensures that FormValue() correctly retrieves single values.
func TestFormValue(t *testing.T) {
	q := quick.New()

	q.Post("/formvalue", func(c *quick.Ctx) error {
		name := c.FormValue("name")
		return c.Status(200).SendString(name)
	})

	ts := httptest.NewServer(q)
	defer ts.Close()

	// Send form-urlencoded data
	form := url.Values{}
	form.Set("name", "Jefferson")

	resp, err := http.Post(ts.URL+"/formvalue", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	buf := new(strings.Builder)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if buf.String() != "Jefferson" {
		t.Errorf("Expected 'Jefferson', got '%s'", buf.String())
	}
}

// TestFormValues ensures that FormValues() correctly retrieves multiple values.
func TestFormValues(t *testing.T) {
	q := quick.New()

	q.Post("/formvalues", func(c *quick.Ctx) error {
		values := c.FormValues()
		jsonData, _ := json.Marshal(values)
		return c.Status(200).Send(jsonData)
	})

	ts := httptest.NewServer(q)
	defer ts.Close()

	form := url.Values{}
	form.Set("name", "Jefferson")
	form.Set("email", "jeff@example.com")

	resp, err := http.Post(ts.URL+"/formvalues", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if result["name"] != "Jefferson" {
		t.Errorf("Expected name 'Jefferson', got '%s'", result["name"])
	}
	if result["email"] != "jeff@example.com" {
		t.Errorf("Expected email 'jeff@example.com', got '%s'", result["email"])
	}
}

// TestFormValuesJSON ensures that FormValues() works correctly with JSON requests.
func TestFormValuesJSON(t *testing.T) {
	q := quick.New()

	q.Post("/formvaluesjson", func(c *quick.Ctx) error {
		values := c.FormValues()
		jsonData, _ := json.Marshal(values)
		return c.Status(200).Send(jsonData)
	})

	ts := httptest.NewServer(q)
	defer ts.Close()

	jsonData := `{"name": "Jefferson", "email": "jeff@example.com"}`
	resp, err := http.Post(ts.URL+"/formvaluesjson", "application/json", bytes.NewBufferString(jsonData))
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if result["name"] != "Jefferson" {
		t.Errorf("Expected name 'Jefferson', got '%s'", result["name"])
	}
	if result["email"] != "jeff@example.com" {
		t.Errorf("Expected email 'jeff@example.com', got '%s'", result["email"])
	}
}

// TestFormValue_Empty ensures that FormValue() returns an empty string when the key is missing.
func TestFormValue_Empty(t *testing.T) {
	q := quick.New()

	q.Post("/formvalue_empty", func(c *quick.Ctx) error {
		value := c.FormValue("missing_key")
		return c.Status(200).SendString(value)
	})

	ts := httptest.NewServer(q)
	defer ts.Close()

	form := url.Values{}

	resp, err := http.Post(ts.URL+"/formvalue_empty", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	buf := new(strings.Builder)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if buf.String() != "" {
		t.Errorf("Expected empty string, got '%s'", buf.String())
	}
}

// TestFormValues_Empty ensures that FormValues() returns an empty map when no data is sent.
func TestFormValues_Empty(t *testing.T) {
	q := quick.New()

	q.Post("/formvalues_empty", func(c *quick.Ctx) error {
		values := c.FormValues()
		jsonData, _ := json.Marshal(values)
		return c.Status(200).Send(jsonData)
	})

	ts := httptest.NewServer(q)
	defer ts.Close()

	form := url.Values{}

	resp, err := http.Post(ts.URL+"/formvalues_empty", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty map, got %v", result)
	}
}
