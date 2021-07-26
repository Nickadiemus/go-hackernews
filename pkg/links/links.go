package links

import (
	"log"

	database "github.com/nickadiemus/go-hackernews/pkg/db/postgres"
	"github.com/nickadiemus/go-hackernews/pkg/users"
)

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *users.User
}

func (l Link) Save() {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES($1,$2)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(l.Title, l.Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Row inserted!")

	//#5
	// id, err := res.LastInsertId()
	// if err != nil {
	// 	log.Fatal("Error:", err.Error())
	// }
	// return id
}
