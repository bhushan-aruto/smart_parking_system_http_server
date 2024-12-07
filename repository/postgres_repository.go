package repository

import (
	"database/sql"

	"github.com/bhushan-aruto/model"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (repo *PostgresRepository) InitDatabse() error {
	query1 := `create table if not exists users (
				user_id varchar(255) primary key,
				name varchar(255) not null,
				phone varchar(255) not null,
				email varchar(255) not null unique,
				password varchar(255) not null
			);`

	if _, err := repo.db.Exec(query1); err != nil {
		return err
	}

	query3 := `create table if not exists bookings (
				user_id varchar(255) not null,
				foreign key (user_id) references users(user_id) on delete cascade
			 );
			 `

	if _, err := repo.db.Exec(query3); err != nil {
		return err
	}

	return nil
}

func (repo *PostgresRepository) CreateUser(user *model.User) error {
	query := `insert into users (user_id,name,phone,email,password) values ($1,$2,$3,$4,$5)`
	if _, err := repo.db.Exec(query, user.UserId, user.Name, user.Phone, user.Email, user.Password); err != nil {
		return err
	}

	return nil
}
func (repo *PostgresRepository) DeleteUser(userId string) error {
	query1 := `delete from users where user_id=$1`
	if _, err := repo.db.Exec(query1, userId); err != nil {
		return err
	}
	return nil

}

func (repo *PostgresRepository) CheckUserEmailExists(email string) (bool, error) {
	query := `select exists(select 1 from users where email=$1)`
	var emailExists bool
	err := repo.db.QueryRow(query, email).Scan(&emailExists)
	return emailExists, err
}

func (repo *PostgresRepository) GetUserPassword(email string) (string, error) {
	query := `select password from users where email=$1`
	var password string
	err := repo.db.QueryRow(query, email).Scan(&password)
	return password, err
}

func (repo *PostgresRepository) GetUserIdByEmail(email string) (string, error) {
	query := `select user_id from users where email = $1`
	var userId string
	err := repo.db.QueryRow(query, email).Scan(&userId)
	return userId, err
}

func (repo *PostgresRepository) OnlineBookSlot(userId string) error {

	query2 := `insert into bookings (user_id) values ($1)`

	_, err := repo.db.Exec(query2, userId)

	return err
}

func (repo *PostgresRepository) CancelOnlineBooking(userId string) error {

	query := `delete from bookings where user_id = $1`

	_, err := repo.db.Exec(query, userId)

	return err
}
