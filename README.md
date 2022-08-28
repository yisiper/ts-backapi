

This test written in golang, with go built in httptest.Server mocking and ginkgo BDD style testing

- API Test Automation
   - Make POST calls
   - Set headers for a request
   - Set the body for a request
   - Assert response code
   - Assert response from a request
   - Assert/ validate field response

## How to run
1. if golang installed:
  * go test -v
  * ginkgo -v processOrder

2. run with docker image
   * docker build -t testapi:v1 .
   * docker run --rm testapi:v1 go test -v
   * docker run --rm testapi:v1 ginkgo -v processOrder
