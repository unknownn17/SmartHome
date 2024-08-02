package router

import (
	_ "api/api/docs"
	"api/api/handler"
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           SMART HOME
// @version         1.0
// @description     Smart Home Pool
// @host            localhost:9000
// @BasePath        /
// @schemes         http
// @securityDefinitions.apiKey BearerAuth
// @in              header
// @name            Authorization/

func Router() {
	r := http.NewServeMux()
	handler := handler.NewHandler()

	r.HandleFunc("POST /users/register", func(w http.ResponseWriter, r *http.Request) {
		handler.Register(w, r)
	})
	r.HandleFunc("POST /users/verify", func(w http.ResponseWriter, r *http.Request) {
		handler.Verify(w, r)
	})
	r.HandleFunc("POST /users/login", func(w http.ResponseWriter, r *http.Request) {
		handler.LogIn(w, r)
	})

	r.HandleFunc("GET /users/profile/{email}", func(w http.ResponseWriter, r *http.Request) {
		handler.UserProfile(w, r)
	})

	r.HandleFunc("PUT /users/update", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateUser(w, r)
	})
	r.HandleFunc("GET /users/logout/{email}", func(w http.ResponseWriter, r *http.Request) {
		handler.Logout(w, r)
	})
	r.HandleFunc("DELETE /users/delete/{email}", func(w http.ResponseWriter, r *http.Request) {
		handler.Delete(w, r)
	})

	

	r.HandleFunc("GET /devices", func(w http.ResponseWriter, r *http.Request) {
		handler.AllDevices(w, r)
	})

	r.HandleFunc("POST /devices/add", func(w http.ResponseWriter, r *http.Request) {
		handler.AddDevice(w, r)
	})

	r.HandleFunc("GET /devices/{user}", func(w http.ResponseWriter, r *http.Request) {
		handler.GetDevices(w, r)
	})

	r.HandleFunc("DELETE /devices/{device}", func(w http.ResponseWriter, r *http.Request) {
		handler.DeleteDev(w, r)
	})

	r.HandleFunc("PUT /devices/command/speaker", func(w http.ResponseWriter, r *http.Request) {
		handler.Speaker(w, r)
	})

	r.HandleFunc("GET /devices/speakers/{dname}", func(w http.ResponseWriter, r *http.Request) {
		handler.SpeakerGet(w, r)
	})
	r.HandleFunc("PUT /devices/command/vaccum", func(w http.ResponseWriter, r *http.Request) {
		handler.Vaccum(w, r)
	})

	r.HandleFunc("GET /devices/vaccum/{dname}", func(w http.ResponseWriter, r *http.Request) {
		handler.VaccumGet(w, r)
	})
	r.HandleFunc("PUT /devices/command/alarm", func(w http.ResponseWriter, r *http.Request) {
		handler.Alarm(w, r)
	})

	r.HandleFunc("GET /devices/alarm/{dname}", func(w http.ResponseWriter, r *http.Request) {
		handler.AlarmGet(w, r)
	})
	r.HandleFunc("PUT /devices/command/door", func(w http.ResponseWriter, r *http.Request) {
		handler.Door(w, r)
	})

	r.HandleFunc("GET /devices/door/{dname}", func(w http.ResponseWriter, r *http.Request) {
		handler.DoorGet(w, r)
	})

	fmt.Println("server started on port 7777")
	r.Handle("/swagger/", httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":7777", r))
}
