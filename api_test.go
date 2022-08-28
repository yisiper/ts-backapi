//go:build !integration
// +build !integration

package ts_backapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProcessOrderResponse(t *testing.T) {
	sampleResponse := `{"order_id":"xxxxxx","order_description":"sample description","order_status":"New",
		"last_updated_timestamp":"1642321210439","special_order":false}`
	server := httptest.NewServer(http.TimeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/processOrder":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(sampleResponse))
		default:
			w.WriteHeader(http.StatusOK)
		}
	}), 10*time.Millisecond, "server timeout"))

	defer server.Close()
	client := server.Client()

	var response *http.Response
	var err error
	t.Run("call processOrder", func(t *testing.T) {
		postBody, _ := json.Marshal(RequestProcessOrder{OrderId: "xxxxxx"})
		response, err = client.Post(server.URL+"/processOrder", "application/json", bytes.NewBuffer(postBody))
		assert.NoError(t, err)
	})

	defer response.Body.Close()

	t.Run("decode response", func(t *testing.T) {
		var v ResponseProcessOrder
		err = json.NewDecoder(response.Body).Decode(&v)
		assert.NoError(t, err)
	})
}
