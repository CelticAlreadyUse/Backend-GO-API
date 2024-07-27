package main

import (
	"encoding/json"
	"log"
	"net/http"
)
func respondWithError(w http.ResponseWriter,code int,msg string){
	if code > 499 {
		// why on 499? karna range 400an adalah client side error a
		//atau kesalahan yang dilakukan oleh client dan 
		//bukan applikasi	
		log.Println("Responding with 5XX error",msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w,code,errResponse{
		Error: msg,
	})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON reponse : %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
