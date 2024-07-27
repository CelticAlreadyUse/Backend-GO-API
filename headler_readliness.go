package main

import "net/http"

func handlerReadliness(w http.ResponseWriter,r *http.Request){
	respondWithJSON(w,200,struct{}{})
}