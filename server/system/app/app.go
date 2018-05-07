package app

import (
	"log"
	"net/http"
	"../router"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"os"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Start(){
	if err := godotenv.Load("config.ini"); err != nil {
		log.Fatal(err)
	}

	s.port = ":" + os.Getenv("PORT")

	r := mux.NewRouter()
	router.InitRoutes(r)
	log.Fatal(http.ListenAndServe(s.port, r))
}