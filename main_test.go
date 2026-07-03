package rssclient_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/ChrisPJohnstone/go-rss-client"
)

type testMeta struct {
	StatusCode    int    `json:"status_code"`
	ExpectedError string `json:"expected_error"`
}

func TestFetchFeed(t *testing.T) {
	entries, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		t.Run(entry.Name(), func(t *testing.T) {
			dir := filepath.Join("testdata", entry.Name())

			meta := testMeta{StatusCode: http.StatusOK}
			metaBytes, err := os.ReadFile(filepath.Join(dir, "meta.json"))
			if err == nil {
				if err := json.Unmarshal(metaBytes, &meta); err != nil {
					t.Fatal(err)
				}
			}

			respBody, err := os.ReadFile(filepath.Join(dir, "response.xml"))
			if err != nil {
				t.Fatal(err)
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(meta.StatusCode)
				if _, err := w.Write(respBody); err != nil {
					t.Fatal(err)
				}
			}))
			defer ts.Close()

			got, err := rssclient.FetchFeed(ts.URL)

			if meta.ExpectedError != "" {
				if err == nil {
					t.Fatalf("expected error %q, got nil", meta.ExpectedError)
				}
				if !strings.Contains(err.Error(), meta.ExpectedError) {
					t.Fatalf("expected error containing %q, got %q", meta.ExpectedError, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}

			expectedBytes, err := os.ReadFile(filepath.Join(dir, "expected.json"))
			if err != nil {
				t.Fatal(err)
			}

			var expected rssclient.Feed
			if err := json.Unmarshal(expectedBytes, &expected); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, expected) {
				t.Errorf("mismatch")
			}
		})
	}
}
