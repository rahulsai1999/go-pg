package db

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Sslmode  string
	Timezone string
}

func getConnStr() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	sslmode := "disable"
	timezone := "Asia/Calcutta"
	conn := Connection{host, user, password, dbname, port, sslmode, timezone}

	template, _ := template.New("xxx").Parse("host={{.Host}} user={{.User}} password={{.Password}} dbname={{.Dbname}} port={{.Port}} sslmode={{.Sslmode}} TimeZone={{.Timezone}}")
	buf := &bytes.Buffer{}
	err := template.Execute(buf, conn)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func DBClient() *gorm.DB {
	dsn := getConnStr()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
