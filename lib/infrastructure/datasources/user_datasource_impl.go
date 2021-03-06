package datasources

import (
	"fmt"
	"strconv"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/database"
	"go.uber.org/zap"
)

type UserDatasourceImpl struct{}

func (impl UserDatasourceImpl) User() (string, error) {
	logger, _ := zap.NewProduction()
	configHandler := cfg.ConfigHandler{}
	config := configHandler.LoadConfig()

	/// Open connection
	db := new(database.Postgresql)
	session, err := db.ConnectAndGet(config)

	/// Error handler
	if err != nil {
		fmt.Print(err.Error())
		logger.Fatal(err.Error())
		return err.Error(), err
	}

	/// Functional code
	var result int
	session.QueryRow(`SELECT count(*) FROM user;`).Scan(&result)
	fmt.Println("The result is : ", result)
	logger.Info(strconv.Itoa(result))

	/// Close connection
	db.Close(*session)
	return "Here is your user", nil
}
