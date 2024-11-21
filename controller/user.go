package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhushan-aruto/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	dbRepo model.DatabaseRepository
}

func NewUserController(dbRepo model.DatabaseRepository) *UserController {
	return &UserController{
		dbRepo: dbRepo,
	}
}
func NewSlotController(dbRepo model.DatabaseRepository) *UserController {
	return &UserController{
		dbRepo: dbRepo,
	}
}

func (c *UserController) DatabaseInit(w http.ResponseWriter, r *http.Request) {
	if err := c.dbRepo.InitDatabse(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occurred while initializing database, Error -> %v\n", err.Error())
		fmt.Fprintf(w, "error occurred with database")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "database initialization successfull")
}

func (c *UserController) CreateUserController(w http.ResponseWriter, r *http.Request) {
	var request model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid json format")
		return
	}

	userId := uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 5)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occurred while generating hash password, Error -> %v\n", err.Error())
		fmt.Fprintf(w, "errro while generating hashed password")
		return
	}

	user := model.User{
		UserId:   userId,
		Name:     request.Name,
		Phone:    request.Phone,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	if err := c.dbRepo.CreateUser(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occurred with database while creating user, Error -> %v\n", err.Error())
		fmt.Fprintf(w, "error occurred with database")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "user created successfully")
}

func (c *UserController) DeleteUserContoller(w http.ResponseWriter, r *http.Request) {
	var request model.DeletUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid json format")
		return
	}

	if err := c.dbRepo.DeleteUser(request.UserId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occurred with database while deleting user, Error -> %v\n", err.Error())
		fmt.Fprintf(w, "error occurred with database")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "user deleted successfully")
}
func (c *UserController) CreateSlotController(w http.ResponseWriter, r *http.Request) {
	var request model.CreateSlotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid json format")
		return
	}

	if request.Rfid == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "rfid cannot be empty")
		return

	}
	if request.SlotId == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "slotid cannot be empty")
		return

	}
	if request.Status < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "status cannot be negative")
		return
	}
	if request.InTime == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Intime cannot be empty")
		return

	}
	if request.OutTime == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Outtime cannot be empty")
		return
	}
	if request.Amount == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Amount cannot be Zero")
		return
	}
	slot := model.Slot{
		Rfid:    request.Rfid,
		SlotId:  request.SlotId,
		Status:  request.Status,
		InTime:  request.InTime,
		OutTime: request.OutTime,
		Amount:  request.Amount,
	}
	if err := c.dbRepo.CreateSlot(&slot); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occurred with database while creating slot, Error -> %v\n", err.Error())
		fmt.Fprintf(w, "error occurred with database")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "slot created successfully")
}
func (c *UserController) DeleteSlotController(w http.ResponseWriter, r *http.Request) {
	var request model.DeletSlotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid json format")
		return
	}
	if err := c.dbRepo.DeleteSlot(request.SlotId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occurred with database while deleting slot, Error -> %v\n", err.Error())
		fmt.Fprintf(w, "error occurred with database")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "slot deleted succesfully")
}

func (c *UserController) UserLoginController(w http.ResponseWriter, r *http.Request) {
	var request model.UserLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid json format")
		return
	}

	if request.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "email cannot be empty")
		return
	}

	if request.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "password cannot be empty")
		return
	}

	isEmailExists, err := c.dbRepo.CheckUserEmailExists(request.Email)

	if err != nil {
		log.Printf("error occurred with database while checking email exits, Error -> %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occurred with database")
		return
	}

	if !isEmailExists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "user email not exists")
		return
	}

	password, err := c.dbRepo.GetUserPassword(request.Email)

	if err != nil {
		log.Printf("error occurred with database while checking email exits, Error -> %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occurred with database")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(request.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "incorrect password")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, request.Email)
}

func (c *UserController) SlotBookingController(w http.ResponseWriter, r *http.Request) {
	var request model.UserBookingRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid json format")
		return
	}

	//validation write here

	if request.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "email cannot be empty")
		return
	}

	if request.ArriveTime == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, " Arrival cannot be empty")
		return
	}

	slots, err := c.dbRepo.GetSlots()

	if err != nil {
		log.Printf("error occurred with database while getting slots, Error -> %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occurred with database")
		return
	}

	var slotId string
	for _, slot := range slots {
		if slot.Status == 0 {
			slotId = slot.SlotId
			break
		}
	}

	if slotId == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "slots are full")
		return
	}

	userId, err := c.dbRepo.GetUserIdByEmail(request.Email)

	if err != nil {
		log.Printf("error occurred with database while getting user id by email, Error -> %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occurred with database")
		return
	}

	if err := c.dbRepo.BookSlot(slotId, userId); err != nil {
		log.Printf("error occurred with database while booking the slot, Error -> %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occurred with database")
		return
	}

	go func() {
		time.Sleep(time.Second * time.Duration(request.ArriveTime))
		slotStatus, err := c.dbRepo.GetSlotStatus(slotId)

		if err != nil {
			log.Printf("error occurred while checking the slot status, Error -> %v", err.Error())
			return
		}

		if slotStatus == 2 {
			if err := c.dbRepo.CancelBooking(slotId, userId); err != nil {
				log.Printf("error occurred while canceling the booking, err -> %v", err.Error())
				return
			}
		}
	}()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "booking successfull")

}
