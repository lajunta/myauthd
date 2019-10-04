package utils

import (
	"database/sql"
	"time"

	// must init mysql
	_ "github.com/go-sql-driver/mysql"

	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

// Config is struct for configuration
type Config struct {
	DbAddress         string
	DbName            string
	TableName         string
	DbUser            string
	DbPass            string
	LoginFieldName    string
	PassFieldName     string
	RealNameFieldName string
	RolesFieldName    string
	CryptMethod       string
}

var (
	configFileName = "config.toml"
	connStr        string
	realName       string
	roles          string
	// CFG a instance of Config
	CFG Config
	err error
	// DB Shared db info
	DB *sql.DB
)

func init() {
	readConfig()
	connStr = CFG.DbUser + ":" + CFG.DbPass + "@(" + CFG.DbAddress + ")/" + CFG.DbName
	connect()
}

// Connect the db server
func connect() {
	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln("Can't Connect Database")
		os.Exit(1)
	}
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Second)
}

func readConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Can't Open User Home Dir")
		os.Exit(1)
	}
	configFilePath := path.Join(home, ".myauthd", configFileName)
	_, err = toml.DecodeFile(configFilePath, &CFG)
	if err != nil {
		log.Fatalln("Can't Open ConfigFile of ConfigFile Wrong!!")
		os.Exit(1)
	}
}

// CheckDB check connect still alive?
func CheckDB() {
	c := time.Tick(30 * time.Minute)
	for range c {
		if err := DB.Ping(); err != nil {
			connect()
		} else {
			log.Println("DB Connect is ok.")
		}
	}
}
