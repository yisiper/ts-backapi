package processOrder_test

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"net/http"
	"time"
	ts_backapi "ts-backapi"
)

var _ = Describe("ProcessOrder", func() {
	var server *ghttp.Server
	var client *http.Client
	var serverUrl string

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = server.HTTPTestServer.Client()
		serverUrl = server.URL()
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Call /processOrder", func() {
		BeforeEach(func() {
			sampleResponse := `{"order_id":"xxxxxx","order_description":"sample description","order_status":"New",
					"last_updated_timestamp":"1642321210439","special_order":false}`
			var r ts_backapi.ResponseProcessOrder
			_ = json.Unmarshal([]byte(sampleResponse), &r)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/processOrder"),
					ghttp.RespondWithJSONEncoded(http.StatusOK, r),
				),
			)
		})

		When("has request", func() {
			var response *http.Response
			var err error
			var result ts_backapi.ResponseProcessOrder
			var reqAt time.Time
			BeforeEach(func() {
				reqAt = time.Now().UTC()
				_ = reqAt

				postBody, _ := json.Marshal(ts_backapi.RequestProcessOrder{OrderId: "xxxxxx"})
				response, err = client.Post(serverUrl+"/processOrder", "application/json", bytes.NewBuffer(postBody))
				err = json.NewDecoder(response.Body).Decode(&result)
			})

			AfterEach(func() {
				defer response.Body.Close()
			})

			It("should no request error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should http status ok", func() {
				Expect(response).To(HaveHTTPStatus(http.StatusOK))
			})
			It("should no json response error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			Context("validate response", func() {
				It("has order id", func() {
					Expect(result.OrderId).ToNot(Equal(""))
				})
				It("should valid unix timestamp", func() {
					Expect(result.LastUpdatedTimestamp.Time.IsZero()).ToNot(Equal(true))
				})

				/*
					It("should greater than timestamp", func() {
						Expect(result.LastUpdatedTimestamp.Time.UnixMilli()).Should(BeNumerically(">=", reqAt.UnixMilli()))
					})
				*/

				It("has order status", func() {
					Expect(result.OrderStatus).ToNot(Equal(""))
				})
			})
		})
	})
})
