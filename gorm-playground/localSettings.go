package main

var (
	DBUser     = "DBUSER"
	DBPassword = "DBPASS"
	DBName     = "DBNAME"
	DBHost     = "127.0.0.1"
	DBPort     = 3306
	DBDriver   = "mysql" //mysql, sqlite, etc
)

type Settings struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	DBDriver   string
}

func GetSettings() *Settings {
	return &Settings{
		DBUser:     DBUser,
		DBPassword: DBPassword,
		DBHost:     DBHost,
		DBPort:     DBPort,
		DBName:     DBName,
		DBDriver:   DBDriver,
	}
}
