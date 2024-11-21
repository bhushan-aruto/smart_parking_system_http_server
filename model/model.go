package model

type DatabaseRepository interface {
	InitDatabse() error
	CreateUser(user *User) error
	DeleteUser(userid string) error
	CreateSlot(slot *Slot) error
	DeleteSlot(slotid string) error
	CheckUserEmailExists(email string) (bool, error)
	GetUserPassword(email string) (string, error)
	GetSlots() ([]*Slot, error)
	GetUserIdByEmail(email string) (string, error)
	BookSlot(slotdId string, userId string) error
	GetSlotStatus(slotId string) (int32, error)
	CancelBooking(slotId string, userId string) error
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
	Status  int32  `json:"status"`
	InTime  string `json:"in_time"`
	OutTime string `json:"out_time"`
	Amount  int32  `json:"amount"`
}

type Slot struct {
	Rfid    string `json:"rfid"`
	SlotId  string `json:"slot_id"`
	Status  int32  `json:"status"`
	InTime  string `json:"in_time"`
	OutTime string `json:"out_time"`
	Amount  int32  `json:"amount"`
}
type DeletSlotRequest struct {
	SlotId string `json:"slot_id"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserBookingRequest struct {
	Email      string `json:"email"`
	ArriveTime int32  `json:"arrive_time"`
}
