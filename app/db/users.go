package db

import (
	"log"

	"github.com/saidalisamed/muxwebappv2/app/models"
)

func (m *Manager) GetUsers() []models.Users {
	rows, err := m.db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	users := []models.Users{}

	for rows.Next() {
		var id int
		var firstName, lastName string
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return nil
		}
		users = append(users, models.Users{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		})
	}
	return users
}
