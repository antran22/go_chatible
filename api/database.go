package api

import (
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"go_chatible/model/user"
)

func ConnectDB() {
	var dbOpt = &pg.Options{
		Addr:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASS"),
		Database: os.Getenv("DB_NAME"),
	}
	DB = pg.Connect(dbOpt)
}

func isDuplicateObjectError(err error) bool {
	return err.Error()[7:12] == "42710"
}

func CreateSchema() error {
	if _, err := DB.Exec(user.GenderTypeCommand); err != nil {
		if !isDuplicateObjectError(err) {
			return err
		}
	}
	if _, err := DB.Exec(user.PrefTypeCommand); err != nil {
		if !isDuplicateObjectError(err) {
			return err
		}
	}
	userModel := (*user.User)(nil)
	opt := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	if err := DB.CreateTable(userModel, opt); err != nil {
		return err
	}
	return nil
}
