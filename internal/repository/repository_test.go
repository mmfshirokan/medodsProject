package repository

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/jackc/pgx/v5/pgxpool"
// 	"github.com/mmfshirokan/medodsProject/internal/model"
// 	"github.com/ory/dockertest/v3"
// 	"github.com/ory/dockertest/v3/docker"
// 	log "github.com/sirupsen/logrus"
// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	conn *Postgres

// 	sameUID = uuid.New()

// 	testTokens = [4]model.RefreshToken{
// 		{
// 			ID:         uuid.New(),
// 			UserID:     uuid.New(),
// 			Hash:       "hash1",
// 			Expiration: time.Now().Add(time.Hour * 1).UTC(),
// 		},
// 		{
// 			ID:         uuid.New(),
// 			UserID:     uuid.New(),
// 			Hash:       "hash2",
// 			Expiration: time.Now().Add(time.Hour * 1).UTC(),
// 		},
// 		{
// 			ID:         uuid.New(),
// 			UserID:     sameUID,
// 			Hash:       "hash3",
// 			Expiration: time.Now().Add(time.Hour * 1).UTC(),
// 		},
// 		{
// 			ID:         uuid.New(),
// 			UserID:     sameUID,
// 			Hash:       "hash4",
// 			Expiration: time.Now().Add(-time.Second * 30).UTC(),
// 		},
// 	}

// 	testUsers = [3]model.User{
// 		{
// 			ID:       testTokens[0].UserID,
// 			IP:       "192.168.2.1",
// 			Name:     "test1",
// 			Email:    "test1",
// 			Password: "test1",
// 		},
// 		{
// 			ID:       testTokens[1].UserID,
// 			IP:       "192.169.2.1",
// 			Name:     "test2",
// 			Email:    "test2",
// 			Password: "test2",
// 		},
// 		{
// 			ID:       testTokens[2].UserID,
// 			IP:       "192.169.7.1",
// 			Name:     "test3",
// 			Email:    "test3",
// 			Password: "test3",
// 		},
// 	}
// )

// func TestMain(m *testing.M) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
// 	defer cancel()

// 	pool, err := dockertest.NewPool("")
// 	if err != nil {
// 		log.Error("Could not construct pool: ", err)
// 		return
// 	}

// 	err = pool.Client.Ping()
// 	if err != nil {
// 		log.Error("Could not connect to docker: ", err)
// 		return
// 	}

// 	pg, err := pool.RunWithOptions(&dockertest.RunOptions{
// 		Hostname:   "postgres_test",
// 		Repository: "postgres",
// 		Tag:        "latest",
// 		Env: []string{
// 			"POSTGRES_PASSWORD=password",
// 			"POSTGRES_USER=user",
// 			"POSTGRES_DB=db",
// 			"listen_addresses = '*'",
// 		},
// 	}, func(config *docker.HostConfig) {
// 		config.AutoRemove = true
// 		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
// 	})
// 	if err != nil {
// 		log.Error("Could not start resource: ", err)
// 		return
// 	}

// 	postgresHostAndPort := pg.GetHostPort("5432/tcp")
// 	postgresUrl := fmt.Sprintf("postgres://user:password@%s/db?sslmode=disable", postgresHostAndPort)

// 	log.Info("Connecting to database on url: ", postgresUrl)

// 	var dbpool *pgxpool.Pool
// 	if err = pool.Retry(func() error { // remove retry? (not nessesary)
// 		dbpool, err = pgxpool.New(ctx, postgresUrl)
// 		if err != nil {
// 			dbpool.Close()
// 			log.Error("can't connect to the pgxpool: %w", err)
// 		}
// 		return dbpool.Ping(ctx)
// 	}); err != nil {
// 		log.Fatalf("Could not connect to docker: %s", err)
// 	}

// 	commandArr := []string{
// 		"-url=jdbc:postgresql://" + postgresHostAndPort + "/db",
// 		"-user=user",
// 		"-password=password",
// 		"-locations=filesystem:../../migrations/",
// 		"-schemas=medods",
// 		"-connectRetries=60",
// 		"migrate",
// 	}
// 	cmd := exec.Command("flyway", commandArr[:]...)

// 	err = cmd.Run()
// 	if err != nil {
// 		log.Error(fmt.Print("error: ", err))
// 	}

// 	pool.MaxWait = 120 * time.Second
// 	conn = New(dbpool)

// 	code := m.Run()

// 	if err := pool.Purge(pg); err != nil {
// 		log.Fatalf("Could not purge resource: %s", err)
// 	}

// 	os.Exit(code)
// }

// func TestAddUsr(t *testing.T) {
// 	ctx := context.Background()

// 	for i, tu := range testUsers {
// 		err := conn.AddUsr(ctx, tu)
// 		assert.Nil(t, err, "test case: %d", i)
// 	}

// 	log.Info("TestAddUsr passed!")
// }

// func TestAdd(t *testing.T) {
// 	ctx := context.Background()

// 	for i, tt := range testTokens {
// 		err := conn.Add(ctx, tt)
// 		assert.Nil(t, err, "test case: %d", i)
// 	}

// 	log.Info("TestAdd passed!")
// }

// func TestGetWithUserID(t *testing.T) {
// 	ctx := context.Background()

// 	testTable := []struct {
// 		testMsg  string
// 		usrID    uuid.UUID
// 		expected []model.RefreshToken
// 	}{
// 		{
// 			testMsg:  "GetWithUserID test 1",
// 			usrID:    testTokens[0].UserID,
// 			expected: []model.RefreshToken{testTokens[0]},
// 		},
// 		{
// 			testMsg:  "GetWithUserID test 2",
// 			usrID:    testTokens[1].UserID,
// 			expected: []model.RefreshToken{testTokens[1]},
// 		},
// 		{
// 			testMsg:  "GetWithUserID test 3",
// 			usrID:    testTokens[2].UserID,
// 			expected: []model.RefreshToken{testTokens[2], testTokens[3]},
// 		},
// 		{
// 			testMsg:  "GetWithUserID nothing (non existed ID)",
// 			usrID:    uuid.New(),
// 			expected: []model.RefreshToken{},
// 		},
// 	}

// 	for _, test := range testTable {
// 		actual, err := conn.GetWithUserID(ctx, test.usrID)

// 		assert.Nil(t, err, test.testMsg)
// 		for i := range test.expected {
// 			assert.Equal(t, test.expected[i].UserID, actual[i].UserID, test.testMsg)
// 			assert.Equal(t, test.expected[i].ID, actual[i].ID, test.testMsg)
// 			assert.Equal(t, test.expected[i].Hash, actual[i].Hash, test.testMsg)
// 		}
// 	}

// 	log.Info("TestGetWithUserID passed!")
// }

// func TestGetHash(t *testing.T) {
// 	ctx := context.Background()

// 	testTable := []struct {
// 		testMsg  string
// 		id       uuid.UUID
// 		expected string
// 		hasError bool
// 	}{
// 		{
// 			testMsg:  "ValidateRFT test 1",
// 			id:       testTokens[0].ID,
// 			expected: testTokens[0].Hash,
// 			hasError: false,
// 		},
// 		{
// 			testMsg:  "ValidateRFT test 2",
// 			id:       testTokens[1].ID,
// 			expected: testTokens[1].Hash,
// 			hasError: false,
// 		},
// 		{
// 			testMsg:  "ValidateRFT test 3",
// 			id:       testTokens[2].ID,
// 			expected: testTokens[2].Hash,
// 			hasError: false,
// 		},
// 		{
// 			testMsg:  "ValidateRFT test 4",
// 			id:       testTokens[3].ID,
// 			expected: testTokens[3].Hash,
// 			hasError: true,
// 		},
// 		{
// 			testMsg:  "ValidateRFT nothing",
// 			id:       uuid.New(),
// 			expected: "something",
// 			hasError: true,
// 		},
// 	}

// 	for _, test := range testTable {
// 		actual, err := conn.GetHash(ctx, test.id)
// 		if test.hasError {
// 			assert.Error(t, err, test.testMsg)
// 			assert.Empty(t, actual, test.testMsg)
// 			continue
// 		}
// 		assert.Nil(t, err, test.testMsg)

// 		assert.Equal(t, test.expected, actual, test.testMsg)
// 	}

// 	log.Info("ValidateRFT passed!")
// }

// func TestGetPwd(t *testing.T) {
// 	ctx := context.Background()

// 	for i, tu := range testUsers {
// 		act, err := conn.GetPwd(ctx, tu.ID)
// 		assert.Nil(t, err, "test case: %d", i)
// 		assert.Equal(t, tu.Password, act, "test case: %d", i)
// 	}
// }

// func TestDelete(t *testing.T) {
// 	ctx := context.Background()

// 	for i, tt := range testTokens {
// 		err := conn.Delete(ctx, tt.ID)
// 		assert.Nil(t, err, "test case: %d", i)
// 	}

// 	log.Info("Delete passed!")
// }
