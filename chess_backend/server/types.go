package server

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Player struct {
	id           string
	Username     string `json:"username"`
	Passwordhash string `json:"passwordhash"`
	Email        string `json:"email"`
	Rating       int    `json:"rating"`
	created_at   time.Time
}

func (u Player) String() string {
	return fmt.Sprintf("ID: %s, Username: %s, Passwordhash: %s, Email: %s, Rating: %d",
		u.id, u.Username, u.Passwordhash, u.Email, u.Rating)
}

func NewPlayer(username, passwordhash, email string) *Player {
	return &Player{
		Username:     username,
		Passwordhash: passwordhash,
		Email:        email,
		Rating:       rand.Intn(2000),
		created_at:   time.Now().UTC(),
	}
}

type Storage interface {
	Init() error
	getAllAccounts() ([]Player, error)
	createAccount(*Player) error
	deleteAccount(string) error
}

type PsqlDB struct {
	db *sql.DB
}

func NewPSQLDB() (*PsqlDB, error) {
	defer timer("Connecting to PSQL")()
	// pass the db credentials into variables
	dbparams := goDotEnvVariable("DBHOST", "DBPORT", "DBUSER", "DBPASS", "DBNAME")

	// create a connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbparams[0], dbparams[1], dbparams[2], dbparams[3], dbparams[4])
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return &PsqlDB{db}, nil
}

func (psqlDB PsqlDB) Init() error {
	return psqlDB.createAccountTable()
}

func (psqlDB PsqlDB) createAccountTable() error {
	query := `create table if not exists player (
		id serial primary key,
		username varchar(50) not null,
		passwordhash varchar(255) not null,
		email varchar(50) not null,
		rating serial not null,
		created_at timestamp
	)`
	_, err := psqlDB.db.Exec(query)
	return err
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

type APIError struct {
	Message string
}

type CreateAccountReq struct {
	username     string
	passwordhash string
	email        string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error
