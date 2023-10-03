package migration

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var createApplicationTableMigration = &Migration{
	Number: 1,
	Name:   "Create sign_set_applications table",
	Forwards: func(db *gorm.DB) error {

		const dropSql = `
			DROP TABLE IF EXISTS sign_set_applications CASCADE;
		`
		err := db.Exec(dropSql).Error
		if err != nil {
			return errors.Wrap(err, "unable to drop loan_running_numbers table")
		}

		const sql = `
			CREATE TABLE "sign_set_applications" (
				"id" BIGSERIAL NOT NULL,
				"name" TEXT NOT NULL,
				"secret" TEXT NOT NULL,
				"callback_url" TEXT NOT NULL,
				"updated_time" BIGINT NOT NULL,
				"created_time" BIGINT NOT NULL,
				PRIMARY KEY ("id")
			);
		`

		err = db.Exec(sql).Error
		if err != nil {
			return errors.Wrap(err, "unable to create sign_set_applications table")
		}

		return nil
	},
}

func init() {
	Migrations = append(Migrations, createApplicationTableMigration)
}
