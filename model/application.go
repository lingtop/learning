package model

import (
	"context"
	"time"

	"template/service/database/postgresql"
	"template/service/util"
)

const APPLICATION_TABLE_NAME = "sign_set_applications"

type Application struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Secret      string `json:"secret"`
	CallbackUrl string `json:"callback_url"`
	UpdatedTime uint64 `json:"updated_time"`
	CreatedTime uint64 `json:"created_time"`
}

func (model Model) CreateApplication(applicationName string, callbackUrl string) (Application, error) {
	model.Logger.Infof("Creating Application")
	output := Application{}

	secret := "daasasdfasdf"
	secretCipher := util.EncryptedText(secret)
	created_time := time.Now().Unix()

	err := postgresql.DatabasePool.QueryRow(context.Background(),
		`
		INSERT INTO `+APPLICATION_TABLE_NAME+` ("name", "secret", "callback_url", "updated_time", "created_time")
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`, applicationName, secretCipher, callbackUrl, created_time, created_time,
	).Scan(&output.Id)

	if err != nil {
		model.Logger.Errorf("Create application failed cuz: %s", err)
		return output, err
	}

	output.Name = applicationName
	output.Secret = secret
	output.CallbackUrl = callbackUrl

	model.Logger.Info("Create Application completed")
	return output, nil
}
