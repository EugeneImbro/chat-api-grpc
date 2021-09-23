package postgres

import (
	"context"
	"fmt"
	"os"
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
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/EugeneImbro/chat-backend/internal/repository"
)

var (
	db  *sqlx.DB
	ctx = context.Background()
	r   repository.Repository
)

func TestMain(m *testing.M) {
	shutdown := setup()

	r = New(db)

	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() func() {
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
	c, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		logrus.WithError(err).Fatal("cannot initialize db container")
	}
	mappedPort, err := c.MappedPort(ctx, nat.Port(port))
	if err != nil {
		logrus.WithError(err).Fatal("cannot get mapped port")
	}

	logrus.Println("postgres container ready and running at port: ", mappedPort)

	db, err = sqlx.Open("postgres", dbURL(mappedPort))

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

	shutdownFn := func() {
		if c != nil {
			c.Terminate(ctx)
		}
	}

	return shutdownFn
}

func TestUserRepository_GetById(t *testing.T) {
	query := "INSERT INTO users (id, nickname) VALUES (1, 'Richard Cheese')"
	_, err := db.Exec(query)
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err := db.ExecContext(ctx, `DELETE FROM users`)
		require.NoError(t, err)
	})

	tt := []struct {
		name     string
		input    int32
		expected *repository.User
		err      error
	}{
		{
			name:     "OK",
			input: 1,
			expected: &repository.User{Id: 1, NickName: "Richard Cheese"},
		},
		{
			name: "NOT FOUND",
			input: 999,
			err:  repository.ErrNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			user, err := r.GetUserByID(context.Background(), tc.input)
			assert.Equal(t, user, tc.expected)
			assert.Equal(t, err, tc.err)
		})

	}

}
