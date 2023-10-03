package external

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func (external External) RequestHttp(url string, input []byte, method string, bearer string) (int, []byte, error) {
	external.Logger.Info("Requesting http to url: %s method: %s", url, method)

	responseStatus := 400

	// requestBody := bytes.NewReader(input)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(input))
	if err != nil {
		external.Logger.Errorf("Create request failed to url: %s method: %s cuz: %s", url, method, err.Error())
		return responseStatus, nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	if bearer != "" {
		bearerText := "Bearer " + bearer
		request.Header.Add("Authorization", bearerText)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		external.Logger.Errorf("Request failed to url: %s method: %s cuz: %s", url, method, err.Error())
		return responseStatus, nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		external.Logger.Errorf("Read response body failed to url: %s method: %s cuz: %s", url, method, err.Error())
		return responseStatus, nil, err
	}

	responseStatus = response.StatusCode
	return responseStatus, responseBody, nil
}
