package api

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

func NewApiServer() {
	mux := mux.NewRouter()
	mux.HandleFunc("/room", getRooms)
}

func getRooms(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{})

}
