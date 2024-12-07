package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/bhushan-aruto/model"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	cache *redis.Client
}

func NewRedisRepository(cache *redis.Client) *RedisRepository {
	return &RedisRepository{
		cache,
	}
}

func (repo *RedisRepository) CreateSlot(slotId string, rfid string) error {
	slotStatus := fmt.Sprintf("%v_status", slotId)
	if _, err := repo.cache.Set(context.Background(), slotStatus, 0, time.Duration(0)).Result(); err != nil {
		return err
	}

	slotRfid := fmt.Sprintf("%v_rfid", slotId)

	if _, err := repo.cache.Set(context.Background(), slotRfid, rfid, time.Duration(0)).Result(); err != nil {
		return err
	}

	slotInTime := fmt.Sprintf("%v_in_time", slotId)

	if _, err := repo.cache.Set(context.Background(), slotInTime, "", time.Duration(0)).Result(); err != nil {
		return err
	}

	slotOutTime := fmt.Sprintf("%v_out_time", slotId)

	if _, err := repo.cache.Set(context.Background(), slotOutTime, "", time.Duration(0)).Result(); err != nil {
		return err
	}

	return nil

}

func (repo *RedisRepository) DeleteSlot(slotId string) error {

	slotStatus := fmt.Sprintf("%v_status", slotId)

	slotRfid := fmt.Sprintf("%v_rfid", slotId)

	slotInTime := fmt.Sprintf("%v_in_time", slotId)

	slotOutTime := fmt.Sprintf("%v_out_time", slotId)

	_, err := repo.cache.Del(context.Background(), slotStatus, slotRfid, slotInTime, slotOutTime).Result()

	return err
}

func (repo *RedisRepository) GetlSlots(slotdIds ...string) ([]*model.Slot, error) {

	var slots []*model.Slot

	for _, slotId := range slotdIds {

		slotStatusKey := fmt.Sprintf("%v_status", slotId)

		slotRfidKey := fmt.Sprintf("%v_rfid", slotId)

		slotInTimeKey := fmt.Sprintf("%v_in_time", slotId)

		slotOutTimeKey := fmt.Sprintf("%v_out_time", slotId)

		slotStatusStr, err := repo.cache.Get(context.Background(), slotStatusKey).Result()

		if err != nil {
			return nil, err
		}

		slotStatusInt, err := strconv.Atoi(slotStatusStr)

		if err != nil {
			return nil, err
		}

		slotStatusInt32 := int32(slotStatusInt)

		slotRfid, err := repo.cache.Get(context.Background(), slotRfidKey).Result()

		if err != nil {
			return nil, err
		}

		slotInTime, err := repo.cache.Get(context.Background(), slotInTimeKey).Result()

		if err != nil {
			return nil, err
		}

		slotOutTime, err := repo.cache.Get(context.Background(), slotOutTimeKey).Result()

		if err != nil {
			return nil, err
		}

		slot := &model.Slot{
			SlotId:  slotId,
			Status:  slotStatusInt32,
			Rfid:    slotRfid,
			InTime:  slotInTime,
			OutTime: slotOutTime,
		}

		slots = append(slots, slot)

	}

	return slots, nil
}

func (repo *RedisRepository) OnlineBookSlot(slotId string) error {
	slotStatusKey := fmt.Sprintf("%v_status", slotId)
	_, err := repo.cache.Set(context.Background(), slotStatusKey, 2, time.Duration(0)).Result()
	return err
}

func (repo *RedisRepository) GetSlotStatus(slotId string) (int32, error) {
	slotStatusKey := fmt.Sprintf("%v_status", slotId)

	slotStatusStr, err := repo.cache.Get(context.Background(), slotStatusKey).Result()

	if err != nil {
		return -1, err
	}

	slotStatusInt, err := strconv.Atoi(slotStatusStr)

	if err != nil {
		return -1, err
	}

	slotStatusInt32 := int32(slotStatusInt)

	return slotStatusInt32, nil
}

func (repo *RedisRepository) CancelOnlineBooking(slotId string) error {
	slotStatusKey := fmt.Sprintf("%v_status", slotId)
	_, err := repo.cache.Set(context.Background(), slotStatusKey, 0, time.Duration(0)).Result()
	return err
}
