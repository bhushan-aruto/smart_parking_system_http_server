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
				email varchar(255) not null,
				password varchar(255) not null
			);`

	if _, err := repo.db.Exec(query1); err != nil {
		return err
	}

	query2 := `create table if not exists slots (
				slot_id varchar(255) primary key,
				rfid varchar(255) not null,
				status integer not null,
				amount integer not null,
				in_time varchar(255) not null,
				out_time varchar(255) not null
		     );`

	if _, err := repo.db.Exec(query2); err != nil {
		return err
	}

	query3 := `create table if not exists bookings (
				user_id varchar(255) not null,
				booking_time varchar(255) not null,
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
func (repo *PostgresRepository) CreateSlot(slot *model.Slot) error {
	query := `insert into slots (rfid,slot_id,status,in_time,out_time,amount) values($1,$2,$3,$4,$5,$6)`
	if _, err := repo.db.Exec(query, slot.Rfid, slot.SlotId, slot.Status, slot.InTime, slot.OutTime, slot.Amount); err != nil {
		return err
	}
	return nil
}
func (repo *PostgresRepository) DeleteSlot(slotId string) error {
	query1 := `delete from slots where slot_id=$1`
	if _, err := repo.db.Exec(query1, slotId); err != nil {
		return err
	}
	return nil
}
