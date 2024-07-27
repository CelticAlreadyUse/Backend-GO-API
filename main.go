package main

// RSS AGREGATOR
import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in evironment")  //log.Fatal wil be broke the program and return the message immedietly
	}

	fmt.Println("Port:",portString)

	router := chi.NewRouter()
	
	v1Router := chi.NewRouter()
	v1Router.Get("/ready",handlerReadliness)
	v1Router.Get("/err",handleErr)

	router.Mount("/v1",v1Router)
	httpServer := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}
	log.Printf("Server starting on port: %v",portString)
	err := httpServer.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))


}