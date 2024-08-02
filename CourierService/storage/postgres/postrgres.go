package postgres

import (
	"database/sql"
	"fmt"
	
	"learning-service/config"
	st "learning-service/storage"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
	learning st.Learning
}

func NewpostgresStorage() (st.InitRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{db:db, learning: &LearningStorage{db}}, nil
}

func (s *PostgresStorage) Learning() st.Learning {
	if s.learning == nil {
		s.learning = &LearningStorage{s.db}
	}
	return s.learning
}

