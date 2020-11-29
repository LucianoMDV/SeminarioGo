package chat

import (
	"fmt"

	"github.com/LucianoMDV/SeminarioGo/Modulo.02/02.3_Ejemplo_Rest_Completo/internal/config"
	"github.com/jmoiron/sqlx"
)

//Message is a object...
type Message struct {
	ID   int64
	Text string
}

//Service is a interface...
type Service interface {
	AddMessage(Message) (Message, error)
	FindByID(int) *Message
	FindAll() []*Message
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New is a...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddMessage(m Message) (Message, error) {
	sqlStatement := "INSERT INTO messages (text) VALUES (?)"

	res, err := s.db.Exec(sqlStatement, m.Text)
	if err != nil {
		return m, err
	}

	m.ID, _ = res.LastInsertId()
	fmt.Println(res.LastInsertId())

	return m, err
}

func (s service) FindByID(int) *Message {
	return nil
}
func (s service) FindAll() []*Message {
	var list []*Message
	// list = append(list, &Message{0, "Hello word"})
	if err := s.db.Select(&list, "SELECT * FROM messages"); err != nil {
		panic(err)
	}
	return list
}
