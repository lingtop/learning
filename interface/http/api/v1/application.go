package v1

import (
	"net/http"

	"template/service/controller"
	"template/service/custom_error"
	"template/service/global_variable"
	"template/service/interface/http/response"
	"template/service/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateApplication(context *gin.Context) {
	contextLogger, _ := context.Get("logger")
	apiLogger := contextLogger.(*zap.SugaredLogger)
	apiLogger.Info("call interface to create application")

	input := controller.CreateApplicationInput{}
	response := response.ResponseOutput{}

	if err := context.ShouldBindJSON(&input); err != nil {
		apiLogger.Errorf("could not bind json body to create application cuz: %s", err.Error())

		isValidation, errorMessage := util.ParseValidtionErrorToString(err)
		if isValidation {
			response.Code = custom_error.InputValidationError
			response.Message = errorMessage
		} else {
			response.Code = custom_error.InvalidJSONString
			response.Message = err.Error()
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	requestId, _ := context.Get("request_id")
	controllerObj := controller.New(requestId.(string))

	output, err := controllerObj.CreateApplication(input)
	if err != nil {
		isBusinessError := util.IsBusinessError(err)
		if !isBusinessError {
			response.Code = err.(custom_error.SystemError).Code
			response.Message = err.Error()
			context.JSON(http.StatusInternalServerError, response)
		} else {
			response.Result.Code = err.(custom_error.BusinessError).Code
			response.Result.Message = err.(custom_error.BusinessError).Message
			context.JSON(http.StatusOK, response)
		}
		return
	}

	response.Message = global_variable.RESULT_SUCCESS
	response.Result.Message = global_variable.RESULT_SUCCESS
	response.Result.Data = output

	context.JSON(http.StatusOK, response)
}

func init() {
	// GET Method
	// methodRoutes[ROUTE_GET]["/orders"] = GetOrders

	// POST Method
	methodRoutes[ROUTE_POST]["/create-application"] = CreateApplication
}
