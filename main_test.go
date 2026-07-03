package rssclient_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ChrisPJohnstone/go-rss-client"
)

func TestFetchFeed(t *testing.T) {
	entries, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			dir := filepath.Join("testdata", entry.Name())
			respBody, err := os.ReadFile(filepath.Join(dir, "response.xml"))
			if err != nil {
				t.Fatal(err)
			}
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(respBody)
			}))
			defer testServer.Close()
			got, err := rssclient.FetchFeed(testServer.URL)
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
