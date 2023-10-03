package controller

import (
	"template/service/custom_error"
	"template/service/model"
)

type CreateApplicationInput struct {
	ApplicationName string `json:"application_name" binding:"required"`
	CallbackUrl     string `json:"callback_url"`
}

type CreateApplicationOutput struct {
	ApplicationId     uint64 `json:"application_id"`
	ApplicationSecret string `json:"application_secret"`
}

func (controller Controller) CreateApplication(input CreateApplicationInput) (CreateApplicationOutput, error) {
	output := CreateApplicationOutput{}

	modelObj := model.New(controller.RequestId)
	application, err := modelObj.CreateApplication(input.ApplicationName, input.CallbackUrl)

	if err != nil {
		controller.Logger.Errorf("Creating application failed cuz: %s", err.Error())
		customError := custom_error.SystemError{
			Code:    custom_error.DatabaseError,
			Message: err.Error(),
		}
		return output, customError
	}

	output.ApplicationId = application.Id
	output.ApplicationSecret = application.Secret

	controller.Logger.Info("Create Application completed")
	return output, nil
}
