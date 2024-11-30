package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhushan-aruto/model"
)

type MachineController struct {
	dbRepo model.DatabaseRepository
}

func NewMachineController(dbRepo model.DatabaseRepository) *MachineController {
	return &MachineController{
		dbRepo,
	}
}

func (c *MachineController) GetOpenController(w http.ResponseWriter, r *http.Request) {
	slots, err := c.dbRepo.GetSlots()

	if err != nil {
		log.Printf("error occurred with database while getting the slots, Error -> %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occurred with database")
		return
	}

	for _, slot := range slots {
		if slot.Status == 0 {
			if err := c.dbRepo.OfflineBookSlot(slot.SlotId); err != nil {
				log.Printf("error occurred with database while updating the slot status, Error -> %v\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "error occurred with database")
				return
			}

			go func() {
				time.Sleep(time.Second * 20)
				slotStatus, err := c.dbRepo.GetSlotStatus(slot.SlotId)

				if err != nil {
					log.Printf("error occurred with database while getting the slot status, Error -> %v\n", err.Error())
					return
				}

				if slotStatus == 2 {
					if err := c.dbRepo.CancelOfflineBooking(slot.SlotId); err != nil {
						log.Printf("error occurred with database while updating the slot status, Error -> %v\n", err.Error())
						return
					}
				}
			}()

			w.WriteHeader(http.StatusOK)
			response := &model.GateOpenResponse{
				Status: true,
				SlotId: slot.SlotId,
			}

			json.NewEncoder(w).Encode(response)
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	response := &model.GateOpenResponse{
		Status: false,
		SlotId: "",
	}

	json.NewEncoder(w).Encode(response)
}

func (c *MachineController) GetSlotUsageAmount(w http.ResponseWriter, r *http.Request) {
	var request model.GetSlotUsageRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid json format")
		return
	}

	

}
