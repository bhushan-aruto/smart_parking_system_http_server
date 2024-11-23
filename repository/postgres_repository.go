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

func (repo *PostgresRepository) GetSlots() ([]*model.Slot, error) {
	query := `select * from slots;`

	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, err
	}

	var slots []*model.Slot

	for rows.Next() {
		var slot model.Slot
		if err := rows.Scan(&slot.SlotId, &slot.Rfid, &slot.Status, &slot.Amount, &slot.InTime, &slot.OutTime); err != nil {
			return nil, err
		}

		slots = append(slots, &slot)
	}

	return slots, nil
}

func (repo *PostgresRepository) GetUserIdByEmail(email string) (string, error) {
	query := `select user_id from users where email = $1`
	var userId string
	err := repo.db.QueryRow(query, email).Scan(&userId)
	return userId, err
}

func (repo *PostgresRepository) OnlineBookSlot(slotdId string, userId string) error {

	query1 := `update slots set status=2 where slot_id=$1`
	query2 := `insert into bookings (user_id) values ($1)`

	tx, err := repo.db.Begin()

	if err != nil {
		return err
	}

	if _, err := tx.Exec(query1, slotdId); err != nil {
		return err
	}

	if _, err := tx.Exec(query2, userId); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *PostgresRepository) OfflineBookSlot(slotId string) error {
	query := `update slots set status=2 where slot_id=$1`
	_, err := repo.db.Exec(query, slotId)
	return err
}

func (repo *PostgresRepository) GetSlotStatus(slotId string) (int32, error) {
	query := `select status from slots where slot_id = $1`

	var status int32

	if err := repo.db.QueryRow(query, slotId).Scan(&status); err != nil {
		return -1, err
	}

	return status, nil
}

func (repo *PostgresRepository) CancelOnlineBooking(slotId string, userId string) error {
	query1 := `update slots set status = 0 where slot_id = $1`
	query2 := `delete from bookings where user_id = $1`

	tx, err := repo.db.Begin()

	if err != nil {
		return err
	}

	if _, err := tx.Exec(query1, slotId); err != nil {
		return err
	}

	if _, err := tx.Exec(query2, userId); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *PostgresRepository) CancelOfflineBooking(slotId string) error {
	query := `update slots set status=0 where slot_id=$1`
	_, err := repo.db.Exec(query, slotId)
	return err
}
