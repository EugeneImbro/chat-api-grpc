package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/EugeneImbro/chat-backend/internal/model"
)

func TestUserRepository_GetById(t *testing.T) {
	ctx := context.Background()
	c, db, _ := createPreparedDBContainer()
	defer c.Terminate(ctx)

	r := New(db)

	t.Run("OK", func(t *testing.T) {
		rollback := "DELETE FROM users WHERE id=1"
		defer func() {
			if _, err := db.Exec(rollback); err != nil {
				logrus.WithError(err).Fatal("cannot execute rollback query")
			}
		}()

		query := "INSERT INTO users (id, nickname) VALUES (1, 'Richard Cheese')"

		if _, err := db.Exec(query); err != nil {
			logrus.WithError(err).Fatal("cannot execute query")
		}

		expected := &model.User{Id: 1, NickName: "Richard Cheese"}

		user, err := r.GetUserByID(context.Background(), 1)
		if err != nil {
			logrus.WithError(err).Fatal("cannot get user from repository")
		}

		assert.Equal(t, user, expected)
	})

}

func createPreparedDBContainer() (testcontainers.Container, *sqlx.DB, error) {
	user := "postgres"
	password := "password"
	dbName := "postgres"

	var env = map[string]string{
		"POSTGRES_USER":     user,
		"POSTGRES_PASSWORD": password,
		"POSTGRES_DB":       dbName,
	}
	var port = "5432/tcp"
	dbURL := func(port nat.Port) string {
		return fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable",
			user, password, port.Port(), dbName)
	}

	ctx := context.Background()

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:latest",
			ExposedPorts: []string{port},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env:          env,
			WaitingFor:   wait.ForSQL(nat.Port(port), "postgres", dbURL).Timeout(time.Second * 5),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		logrus.WithError(err).Fatal("cannot initialize db container")
	}
	mappedPort, err := container.MappedPort(ctx, nat.Port(port))
	if err != nil {
		logrus.WithError(err).Fatal("cannot get mapped port")
	}

	logrus.Println("postgres container ready and running at port: ", mappedPort)

	db, err := sqlx.Open("postgres", dbURL(mappedPort))

	migrator, err := migrate.New(
		fmt.Sprintf("file://%s", "../../../migrations"),
		dbURL(mappedPort),
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create migrator")
	}
	defer migrator.Close()

	if err := migrator.Up(); err != nil {
		logrus.WithError(err).Fatal("failed to migrate")
	}

	return container, db, nil
}
