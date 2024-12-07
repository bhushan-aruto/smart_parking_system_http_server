package model

type DatabaseRepository interface {
	InitDatabse() error
	CreateUser(user *User) error
	DeleteUser(userid string) error
	CheckUserEmailExists(email string) (bool, error)
	GetUserPassword(email string) (string, error)
	GetUserIdByEmail(email string) (string, error)
	OnlineBookSlot(userId string) error
	CancelOnlineBooking(userId string) error
}

type CacheRepository interface {
	CreateSlot(slotId string, rfid string) error
	DeleteSlot(slotId string) error
	GetlSlots(slotdIds ...string) ([]*Slot, error)
	OnlineBookSlot(slotId string) error
	GetSlotStatus(slotId string) (int32, error)
	CancelOnlineBooking(slotId string) error
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
	Rfid   string `json:"rfid"`
	SlotId string `json:"slot_id"`
}

type Slot struct {
	Rfid    string `json:"rfid"`
	SlotId  string `json:"slot_id"`
	Status  int32  `json:"status"`
	InTime  string `json:"in_time"`
	OutTime string `json:"out_time"`
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

type GetSlotUsageRequest struct {
	Rfid string `json:"rfid"`
}

// responses
type GateOpenResponse struct {
	Status bool   `json:"status"`
	SlotId string `json:"slot_id"`
}
