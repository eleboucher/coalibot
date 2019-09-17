package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/genesixx/coalibot/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func AddCommand(event *utils.Message, command string, option string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	usr, err := event.API.GetUserInfo(event.User)
	if err != nil {
		fmt.Println(err)
		return
	}
	user := usr.Name

	var lastInsertId int
	err = db.QueryRow(`INSERT INTO commands(command_name,"user",option) VALUES($1,$2,$3) RETURNING id`, command, user, option).Scan(&lastInsertId)
	if err != nil {
		fmt.Println(err)
	}
}
