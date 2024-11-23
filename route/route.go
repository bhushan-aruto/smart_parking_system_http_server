package route

import (
	"github.com/bhushan-aruto/controller"
	"github.com/bhushan-aruto/model"
	"github.com/gorilla/mux"
)

func NewRouter(dbRepo model.DatabaseRepository) *mux.Router {
	router := mux.NewRouter()

	userController := controller.NewUserController(dbRepo)
	machineController := controller.NewMachineController(dbRepo)

	router.HandleFunc("/db/init", userController.DatabaseInit).Methods("GET")
	router.HandleFunc("/create/user", userController.CreateUserController).Methods("POST")
	router.HandleFunc("/delete/user", userController.DeleteUserContoller).Methods("POST")
	router.HandleFunc("/create/slot", userController.CreateSlotController).Methods("POST")
	router.HandleFunc("/delete/slot", userController.DeleteSlotController).Methods("POST")
	router.HandleFunc("/login/user", userController.UserLoginController).Methods("POST")

	userRouter := router.PathPrefix("/user").Subrouter()
	machineRouter := router.PathPrefix("/machine").Subrouter()

	userRouter.HandleFunc("/book", userController.SlotBookingController).Methods("POST")
	machineRouter.HandleFunc("/book", machineController.GetOpenController).Methods("GET")

	return router
}
