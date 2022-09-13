package fruit

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"log"
)

type Service struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewFruitService(conn *pgx.Conn) *Service {
	return &Service{
		db:  conn,
		ctx: context.Background(),
	}
}

func (s *Service) ListFruits() []Fruit {
	var fruits []Fruit
	if s.db == nil {
		return fruits
	}
	rows, err := s.db.Query(s.ctx, "select id, name, quantity from fruit")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var id uuid.UUID
		var name string
		var quantity string
		if err := rows.Scan(&id, &name, &quantity); err != nil {
			log.Println("failed to read fruit", err)
		}
		fruits = append(fruits, Fruit{
			Id:       id.String(),
			Name:     name,
			Quantity: quantity,
		})
	}
	return fruits
}

func (s *Service) Create(fruit Fruit) error {
	if s.db == nil {
		return errors.New("DB Connection is not created")
	}
	if fruit.Name == "" {
		return errors.New("Missing fruit name")
	}
	if _, err := s.db.Exec(s.ctx, "insert into fruit (id, name, quantity) values ($1, $2, $3)", fruit.Id, fruit.Name, fruit.Quantity); err != nil {
		log.Println("failed to create a fruit", err)
		return err
	}
	return nil
}

func (s *Service) Update(fruit Fruit) error {
	if s.db == nil {
		return errors.New("DB Connection is not created")
	}
	if fruit.Name == "" {
		return errors.New("Missing fruit name")
	}
	if _, err := s.db.Exec(s.ctx, "update fruit set name = $1, quantity = $2 where id = $3", fruit.Name, fruit.Quantity, fruit.Id); err != nil {
		log.Println("failed to update the fruit", err)
		return err
	}
	return nil
}

func (s *Service) FindByID(id string) (*Fruit, error) {
	fruit := &Fruit{}
	if s.db == nil {
		return nil, errors.New("DB Connection is not created")
	}
	if err := s.db.QueryRow(s.ctx, "select id, name, quantity from fruit where id = $1", id).Scan(&fruit.Id, &fruit.Name, &fruit.Quantity); err != nil {
		log.Println(fmt.Sprintf("failed to select the fruit %v", id), err)
		return nil, err
	}
	return fruit, nil
}

func (s *Service) DeleteByID(id string) error {
	if s.db == nil {
		return errors.New("DB Connection is not created")
	}
	var tag pgconn.CommandTag
	var err error
	if tag, err = s.db.Exec(s.ctx, "delete from fruit where id = $1", id); err != nil {
		log.Println(fmt.Sprintf("failed to select the fruit %v", id), err)
		return err
	}
	fmt.Println("update tag : ", tag)
	return nil
}
