package ad_token

import (
	"encoding/json"

	"template/service/external"
	"template/service/global_variable"
	"template/service/logger"

	"github.com/spf13/viper"
)

const (
	OAUTH2_GRANT_TYPE = "client_credentials"
)

var Oauth2Url string
var ClientId string
var ClientSecret string

type Oauth2Response struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
}

type Oauth2Request struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

func InitAdToken() {
	Oauth2Url = viper.GetString("Oauth.Url")
	ClientId = viper.GetString("Oauth.ClientID")
	ClientSecret = viper.GetString("Oauth.ClientSecret")
}

func RequestOauthToken(requestId string, scope string) (Oauth2Response, error) {
	oauthLogger := logger.Logger.With(
		global_variable.KEY_REQUEST_ID, requestId,
		global_variable.KEY_PART, "ad_token",
	)
	response := Oauth2Response{}

	request := Oauth2Request{
		ClientId:     ClientId,
		ClientSecret: ClientSecret,
		GrantType:    OAUTH2_GRANT_TYPE,
		Scope:        scope,
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		oauthLogger.Errorf("could not parse struct to json cuz: %s", err.Error())
		return response, err
	}

	externalObj := external.New(requestId)
	respStatusCode, respBytes, err := externalObj.RequestHttp(
		Oauth2Url, requestBytes, external.HTTP_METHOD_POST, "",
	)
	if err != nil {
		oauthLogger.Errorf("failed to request to get oauth2 token cuz: %s", err.Error())
		return response, err
	}

	if respStatusCode != 200 {
		oauthLogger.Errorf("failed to request to get oauth2 cuz: return status is : %d", respStatusCode)
		return response, err
	}

	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		oauthLogger.Errorf("could not parse struct to json cuz: %s", err.Error())
		return response, err
	}

	return response, nil
}
