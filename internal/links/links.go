package links

import (
	"log"

	database "github.com/phantranhieunhan/demo-graphql/internal/pkg/db/migrations/postgresql"
	"github.com/phantranhieunhan/demo-graphql/internal/users"
)

// #1
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// #2
func (link Link) Save() int64 {
	//#3
	row := database.Db.QueryRow(`INSERT INTO links(title,address,user_id) VALUES($1,$2,$3)  RETURNING id`, link.Title, link.Address, link.User.ID)
	//#4
	var id int64
	err := row.Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func GetAll() []Link {
	rows, err :=database.Db.Query("select l.*, u.username from links l left join users u on l.user_id = u.id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		i.User =&users.User{}
		if err := rows.Scan(
			&i.ID, &i.Title, &i.Address, &i.User.ID, &i.User.Username,
		); err != nil {
			log.Fatal(err)
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return items
	// defer stmt.Close()
	// rows, err := stmt.Query()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// var links []Link
	// for rows.Next() {
	// 	var link Link
	// 	err := rows.Scan(&link.ID, &link.Title, &link.Address)
	// 	if err != nil{
	// 		log.Fatal(err)
	// 	}
	// 	links = append(links, link)
	// }
	// if err = rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
