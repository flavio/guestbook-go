package models

import (
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
)

const MSGS_KEY = "messages"

type Message struct {
	Index int64  `json:"index"`
	Data  string `json:"data"`
}

func GetMessages(db *redis.Client) (messages []Message, err error) {
	messagesCount, err := db.LLen(MSGS_KEY).Result()

	if err != nil {
		return
	}

	for i := int64(0); i < messagesCount; i++ {
		var data string
		data, err = db.LIndex(MSGS_KEY, i).Result()
		if err != nil {
			return
		}

		messages = append(
			messages,
			Message{
				Index: i,
				Data:  data,
			})
	}

	return
}

func PutMessage(db *redis.Client, entry string) (index int64, err error) {
	index, err = db.LPush(MSGS_KEY, entry).Result()

	return
}

func DeleteMessage(db *redis.Client, index int64) (err error) {
	// there's no easy way to delete an element from a redis list by using its index

	// calculate a unique value
	newValue := uuid.NewV4()

	// change the value of the item to remove to a unique value
	_, err = db.LSet(MSGS_KEY, index, newValue.String()).Result()
	if err != nil {
		return err
	}

	// remove the element
	_, err = db.LRem(MSGS_KEY, 0, newValue.String()).Result()
	if err != nil {
		return err
	}

	return
}
