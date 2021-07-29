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

// Inserts a new Link into the database
func (l Link) Save() int64 {
	query := `
		INSERT INTO 
			links(Title,Address) 
		VALUES
			($1,$2) 
		RETURNING id`
	stmt, err := database.Db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// retrieve last id from returning sql statement
	var id int64
	err = stmt.QueryRow(l.Title, l.Address).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func (l Link) GetAll() []Link {
	query := `
		SELECT id, title, address FROM links
	`
	stmt, err := database.Db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var links []Link
	r, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	for r.Next() {
		var link Link
		err := r.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	return links
}
