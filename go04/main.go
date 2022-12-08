package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var ctx = context.Background()

type PokemonResponse struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB") + ":6379",
		Password: "",
		DB:       0,
	})
	var responseObj PokemonResponse
	key := "pokedex"
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
		if err != nil {
			panic(err)
		}
		body, err := io.ReadAll(response.Body)
		json.Unmarshal(body, &responseObj)
		rdb.Set(ctx, key, body, 10*time.Second)
	} else if err != nil {
		panic(err)
	} else {
		json.Unmarshal([]byte(val), &responseObj)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&responseObj)
	if err != nil {
		log.Fatalln("error")
	}
}

type RedisHash struct {
	Name   string `redis:"name"`
	Age    int    `redis:"age"`
	Online bool   `redis:"online"`
}

func RedisHashToStruct() {
	db := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB") + ":6379",
		Password: "",
		DB:       0,
	})
	var rh = RedisHash{
		Name:   "oscar",
		Age:    28,
		Online: true,
	}
	db.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, "rh", "name", rh.Name)
		pipe.HSet(ctx, "rh", "age", rh.Age)
		pipe.HSet(ctx, "rh", "online", rh.Online)
		return nil
	})
	var rh2 RedisHash
	err := db.HGetAll(ctx, "rh").Scan(&rh2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("rh=%+v\n", rh2)
}

func main() {
	RedisHashToStruct()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8090", r)

}
