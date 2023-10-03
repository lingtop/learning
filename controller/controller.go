package controller

import (
	"template/service/logger"

	"go.uber.org/zap"
)

type Controller struct {
	RequestId string
	Logger    *zap.SugaredLogger
}

func New(requestId string) Controller {
	controllerObj := Controller{RequestId: requestId}
	controllerObj.Logger = logger.Logger.With(
		"request_id", requestId,
		"part", "controller",
	)

	return controllerObj
}
