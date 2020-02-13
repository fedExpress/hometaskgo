package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"test-api/ds"
	"test-api/redis"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	newEvent := ds.Event{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	newEvent.SaveToRedis(redis.GetRedisConnection(), "")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	event := ds.GetFromRedis(redis.GetRedisConnection(), eventID)
	_ = json.NewEncoder(w).Encode(event)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent ds.Event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &updatedEvent)

	updatedEvent.SaveToRedis(redis.GetRedisConnection(), eventID)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	ds.DeleteFromRedis(redis.GetRedisConnection(), eventID)
	fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)

}
