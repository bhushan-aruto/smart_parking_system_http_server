package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
func (s *UserController) CreateSlotController(w http.ResponseWriter, r *http.Request) {
	var request model.CreateSlotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid json format")
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
	if err := s.dbRepo.CreateSlot(&slot); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occurred with database while creating slot, Error -> %v\n", err.Error())
		fmt.Fprintf(w, "error occurred with database")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "slot created successfully")
}
func (s *UserController) DeleteSlotController(w http.ResponseWriter, r *http.Request) {
	var request model.DeletSlotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid json format")
		return
	}
	if err := s.dbRepo.DeleteSlot(request.SlotId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occurred with database while deleting slot, Error -> %v\n", err.Error())
		fmt.Fprintf(w, "error occurred with database")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "slot deleted succesfully")
}
