package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

const (
	url = "redis://miguelesdb@34.68.95.85:6379"
)

func conectar_server(wri http.ResponseWriter, req *http.Request) {
	wri.Header().Set("Access-Control-Allow-Origin", "*")
	wri.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	wri.Header().Set("Content-Type", "application/json")
	if req.Method == "GET" {
		wri.WriteHeader(http.StatusOK)
		wri.Write([]byte("{\"mensaje\": \"ok blue\"}"))
		fmt.Println("aca entre")
		return
	}
	//fmt.Println("es un post")
	datos, _ := ioutil.ReadAll(req.Body)
	//fmt.Println("Respuesta del server: ")
	//fmt.Println(datos)
	json.NewEncoder(wri).Encode("blue deployment")
	bodyString := string(datos)
	log.Print(bodyString)
	publish(bodyString)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", conectar_server)
	fmt.Println("Cliente se levanto en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func publish(mensaje string) {
	c, err := redis.DialURL(url)
	if err != nil {
		fmt.Println(err)
	}
	c.Do("PUBLISH", "vacunados", mensaje)
	set(mensaje)
}

func set(mensaje string) {
	conn, err := redis.DialURL(url)
	if err != nil {
		fmt.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
	}
	a, err := conn.Do("lpush", "personas", mensaje)
	if err != nil {
		fmt.Println(err)
		fmt.Println(a)
	}
}
