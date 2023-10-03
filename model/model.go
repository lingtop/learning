package model

import (
	"template/service/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Model struct {
	RequestId string
	Logger    *zap.SugaredLogger
}

var DecimalPoint int32
var RoundingPoint int32

func InitModel() {
	DecimalPoint = viper.GetInt32("Decimal.DecimalPoint")
	RoundingPoint = viper.GetInt32("Decimal.RoundingPoint")
}

func New(requestId string) Model {
	modelObj := Model{RequestId: requestId}
	modelObj.Logger = logger.Logger.With(
		"request_id", requestId,
		"part", "model",
	)
	return modelObj
}
