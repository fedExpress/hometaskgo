package ds

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (event *Event) SaveToRedis(client *redis.Client, id string) {
	if id != "" {
		client.HMSet(fmt.Sprintf("EVENT::%s", id),
			"id", id,
			"title", event.Title,
			"desc", event.Description)
	} else {
		val, _ := client.Get("FREEIDLIST").Result()
		i, _ := strconv.Atoi(val)
		client.HMSet(fmt.Sprintf("EVENT::%s", strconv.Itoa(i+1)),
			"id", strconv.Itoa(i+1),
			"title", event.Title,
			"desc", event.Description)
		_ = client.Set("FREEIDLIST", strconv.Itoa(i+1), 0)
	}
}

func DeleteFromRedis(client *redis.Client, id string) {
	_, _ = client.Del(fmt.Sprintf("EVENT::%s", id)).Result()
}

func GetFromRedis(client *redis.Client, id string) (event Event) {

	evFlds, _ := client.HMGet(fmt.Sprintf("EVENT::%s", id), "id", "title", "desc").Result()
	event.ID = evFlds[0].(string)
	event.Title = evFlds[1].(string)
	event.Description = evFlds[2].(string)
	return
}
