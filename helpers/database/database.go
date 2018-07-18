package database

import (
	"fmt"
	"os"
)

//ConnectionString is our connection string to the database
func ConnectionString() string {
	if addr := os.Getenv("DATABASE_HOST"); addr != "" {
		proto := os.Getenv("DATABASE_PROTOCOL")
		user := os.Getenv("DATABASE_USERNAME")
		pass := os.Getenv("DATABASE_PASSWORD")
		db := os.Getenv("DATABASE_NAME")

		return fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true&loc=%s", user, pass, proto, addr, db, "America%2FChicago")
	}
	return "root:@tcp(127.0.0.1:3306)/demodata?parseTime=true&loc=America%2FChicago"
}
