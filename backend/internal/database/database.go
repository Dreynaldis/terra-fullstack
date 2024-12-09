package database

import (
	"backend/internal/config"
	"backend/internal/model"
	"backend/internal/utils"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Service interface {
	Health() map[string]string
	Close() error
	Queries() *model.Queries
}
type service struct {
	db      *sql.DB
	queries *model.Queries
}

var (
	dbInstance *service
	once       sync.Once
)

func NewDatabase(cfg *config.Config) *service {
	once.Do(func() {
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSchema)

		db, err := sql.Open("pgx", connStr)
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("Database Connection test failed: %v", err)
		}

		log.Println("Connected to the database")

		queries := model.New(db)
		dbInstance = &service{
			db:      db,
			queries: queries,
		}

		//seed initial data for demo purposes, comment after first run
		if err := dbInstance.seedInitialUsers(); err != nil {
			log.Printf("seeding error :%v", err)
		}
	})

	return dbInstance
}
func (s *service) Queries() *model.Queries {
	return s.queries
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("Db down: %v", err)
		return stats
	}
	stats["status"] = "up"
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)

	return stats
}

func (s *service) Close() error {
	if s.db != nil {
		log.Println("Closing database connection")
		return s.db.Close()
	}
	return nil
}

func (s *service) seedInitialUsers() error {
	users := []struct {
		Username string
		Email    string
		Password string
		Provider string
	}{
		{"TerraAdmin", "terra@gmail.com", "admin123", "local"},
	}

	for _, u := range users {
		hashedPasswords, err := utils.HashPassword(u.Password)
		if err != nil {
			log.Printf("Failed to hash password for %s: %v", u.Username, err)
			continue
		}

		_, err = s.queries.GetUserByEmail(context.Background(), u.Email)
		if err == nil {
			log.Printf("User %s already exists, skipping...", u.Username)
			continue
		}

		err = s.queries.CreateUser(context.Background(), model.CreateUserParams{
			Username: u.Username,
			Email:    u.Email,
			Password: hashedPasswords,
			Provider: "local",
		})
		if err != nil {
			log.Printf(" user %s: %v", u.Username, err)
			continue
		}
		log.Printf("User %s created successfully", u.Username)
	}
	return nil
}
