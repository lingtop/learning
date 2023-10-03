package integration

import (
	"template/service/external/ad_token"
	"template/service/global_variable"
	"template/service/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	KEY_INTEGRATION_OTP  = "otp"
	THRESHOLD_TOKEN_TIME = int64(60)
)

type IntegrationInfo struct {
	Token       string
	ExpiredTime int64
	UseToken    bool
	Url         string
	Scope       string
}

type IntegrationConfig struct {
	Key      string
	UseToken bool
	Url      string
	Scope    string
}

var IntegrationInfoMap map[string]*IntegrationInfo
var MaxRetry int

type Integration struct {
	RequestId string
	Logger    *zap.SugaredLogger
}

func New(requestId string) Integration {
	integrationObj := Integration{RequestId: requestId}
	integrationObj.Logger = logger.Logger.With(
		global_variable.KEY_REQUEST_ID, requestId,
		global_variable.KEY_PART, "integration",
	)
	return integrationObj
}

func InitIntegrationInfo() {
	ad_token.InitAdToken()

	MaxRetry = viper.GetInt("Integration.MaxRetry")

	integrationConfigs := make([]IntegrationConfig, 0)
	viper.UnmarshalKey("Integration.Source", &integrationConfigs)
	IntegrationInfoMap = make(map[string]*IntegrationInfo, len(integrationConfigs))

	for _, integrationConfig := range integrationConfigs {
		IntegrationInfoMap[integrationConfig.Key] = &IntegrationInfo{
			UseToken: integrationConfig.UseToken,
			Url:      integrationConfig.Url,
			Scope:    integrationConfig.Scope,
		}
	}
}
