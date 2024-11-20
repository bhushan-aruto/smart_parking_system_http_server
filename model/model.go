package model

type DatabaseRepository interface {
	InitDatabse() error
	CreateUser(user *User) error
	DeleteUser(userid string) error
	CreateSlot(slot *Slot) error
	DeleteSlot(slotid string) error
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeletUserRequest struct {
	UserId string `json:"user_id"`
}

type User struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type CreateSlotRequest struct {
	Rfid    string `json:"rfid"`
	SlotId  string `json:"slot_id"`
	Status  string `json:"status"`
	InTime  string `json:"in_time"`
	OutTime string `json:"out_time"`
	Amount  string `json:"amount"`
}

type Slot struct {
	Rfid    string `json:"rfid"`
	SlotId  string `json:"slot_id"`
	Status  string `json:"status"`
	InTime  string `json:"in_time"`
	OutTime string `json:"out_time"`
	Amount  string `json:"amount"`
}
type DeletSlotRequest struct {
	SlotId string `json:"slot_id"`
}
