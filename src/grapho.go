package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/artshpakov/grapho/src/schema"
	"github.com/go-pg/pg"
	"github.com/graphql-go/graphql"
)

type Config struct {
	DBAddr string
	DBUser string
	DBName string
}

func loadConfig() (config Config) {
	configfile := "./config/config.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("Errors: %v", result.Errors)
	}
	return result
}

func main() {
	var config = loadConfig()
	db := pg.Connect(&pg.Options{Addr: config.DBAddr, User: config.DBUser, Database: config.DBName})

	schema, err := schema.DefineSchema(db)
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	http.HandleFunc("/graph", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		result := executeQuery(r.URL.Query()["q"][0], schema)
		json.NewEncoder(w).Encode(result)
	})
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "src/assets/index.html")
	})

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
