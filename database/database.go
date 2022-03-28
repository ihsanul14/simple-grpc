package database

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

const database = "MySQL"

func ConnectMySQL(log *logrus.Logger) (*gorm.DB, error) {
	var (
		err      error
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASS")
		dbname   = os.Getenv("DB_NAME")
	)
	msqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	log.Debug(msqlInfo)
	Db, err = gorm.Open(mysql.Open(msqlInfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Fatalf("database.database.Open: %v", err.Error())
		return nil, err
	} else {
		log.Info(database, "is Connected")
	}
	Db = Db.Debug()
	return Db, err
}

func GetMySQLDB() *gorm.DB {
	return Db
}

type TestDBError struct {
	TestDB
	Error error
}

type TestDB interface {
	// ConnectMySQL(*gorm.DB, error)
	Table(name string) TestDB
	Where(query interface{}, args ...interface{}) TestDB
	Scan(dest interface{}) *TestDBError
}
