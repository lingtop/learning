package external

import (
	"template/service/logger"

	"go.uber.org/zap"
)

const (
	HTTP_METHOD_POST = "POST"
	HTTP_METHOD_GET  = "GET"
)

type External struct {
	RequestId string
	Logger    *zap.SugaredLogger
}

func New(requestId string) External {
	externalObj := External{RequestId: requestId}
	externalObj.Logger = logger.Logger.With(
		"request_id", requestId,
		"part", "external",
	)
	return externalObj
}
