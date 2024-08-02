package handler

import (
	register1 "api/internal/authorization/register"
	"api/internal/device"
	"api/internal/models"
	"api/protos/deviceproto"
	"api/publisher"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	PUB *publisher.Publisher
	Reg *register1.Register
	Dev *device.Devices
}

func NewHandler() *Handler {
	a := publisher.NewPub()
	b := register1.NewReg()
	d := device.NewDevice()
	return &Handler{PUB: a, Reg: b, Dev: d}
}

// @Router 			/users/register [post]
// @Summary			To register
// @Description 	New User
// @Security		BearerAuth
// @Tags			Users
// @Accept 			json
// @Produce			json
// @Param 			body body models.RegistrationRequest true "RegistrationRequest"
// @Success			201  {object} models.RegistrationResponse
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "decoding error", http.StatusInternalServerError)
		return
	}
	if err := u.Reg.Register(&req); err != nil {
		http.Error(w, fmt.Sprintf("Cashing register error %v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(models.RegistrationResponse{Message: "check your email to get a code"})
}

// @Router 			/users/verify [post]
// @Summary			To verify
// @Description 	New User
// @Security		BearerAuth
// @Tags			Users
// @Accept 			json
// @Produce			json
// @Param 			body body models.VerifyRequest true "VerifyRequest"
// @Success			201  {object} models.VerifyResponse
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "decoding error", http.StatusInternalServerError)
		return
	}
	if err := u.Reg.UserVerify(&req); err != nil {
		http.Error(w, fmt.Sprintf("Cashing register error %v", err), http.StatusInternalServerError)
		return
	}
	if err := u.Reg.Sendtodatabase(req.Email); err != nil {
		http.Error(w, fmt.Sprintf("saving to the database %v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(models.VerifyResponse{Message: "Successfully verified"})
}

// @Router 			/users/login [post]
// @Summary			To LogIn
// @Description 	New User
// @Security		BearerAuth
// @Tags			Users
// @Accept 			json
// @Produce			json
// @Param 			body body models.LoginRequest true "LoginRequest"
// @Success			201  {object} models.LoginResponse
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "decoding error", http.StatusInternalServerError)
		return
	}
	token, err := u.Reg.Login(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("login error %v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(models.LoginResponse{Message: token})
}

// @Router 			/users/profile/{email} [get]
// @Summary			To see profile
// @Description 	user profile
// @Security		BearerAuth
// @Tags			Users
// @Accept 			json
// @Produce			json
// @Param 			email path string true "UserProfileRequest"
// @Success			200  {object} models.UserProfileResponse
// @Success			201  {object} models.UserProfileResponse
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) UserProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req = models.UserProfileRequest{Email: r.PathValue("email")}
	fmt.Println("here it is")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	res, err := u.Reg.UserProfile(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	fmt.Println("userprofile", res)
	json.NewEncoder(w).Encode(res)
}

// @Router 			/users/update [put]
// @Summary			To Update a user
// @Description 	user profile
// @Security		BearerAuth
// @Tags			Users
// @Accept 			json
// @Produce			json
// @Param 			body body models.UpdateRequest true "UserProfileRequest"
// @Success			200  {object} models.UserProfileResponse
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	res, err := u.Reg.Updaterequest(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// @Router 			/users/logout/{email} [put]
// @Summary			To logout
// @Description 	user profile
// @Security		BearerAuth
// @Tags			Users
// @Accept 			json
// @Produce			json
// @Param 			email path string true "UserProfileRequest"
// @Success			200  {object} models.LogoutResponse
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := r.PathValue("email")
	res, err := u.Reg.Logout(&models.LogoutRequest{Email: email})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// @Router 			/users/delete/{email} [delete]
// @Summary			To delete a profile
// @Description 	user profile
// @Security		BearerAuth
// @Tags			Users
// @Accept 			json
// @Produce			json
// @Param 			email path string true "UserProfileRequest"
// @Success			200  {object} models.LogoutResponse
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := r.PathValue("email")
	res, err := u.Reg.Logout(&models.LogoutRequest{Email: email})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// @Router 			/devices [get]
// @Summary			To see All available devices
// @Description 	user profile
// @Security		BearerAuth
// @Tags			Devices
// @Produce			json
// @Success			200  {object} []models.AllDevices
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) AllDevices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("alldevices")
	w.Header().Set("Content-Type", "application/json")
	res, err := u.Dev.Alldevice(&deviceproto.Deviceslist{})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	fmt.Println("res", res)
	json.NewEncoder(w).Encode(res)
}

// @Router 			/devices/add [post]
// @Summary			To add a new thing
// @Description 	user home
// @Security		BearerAuth
// @Tags			Devices
// @Accept 			json
// @Produce			json
// @Param 			body body models.AddDeviceRequest true "AddDeviceRequest"
// @Success			200  {object} []models.AllDevices
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) AddDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.AddDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	res, err := u.Dev.AddDevice(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// @Router 			/devices/{user} [get]
// @Summary			To see user's device
// @Description 	user device
// @Security		BearerAuth
// @Tags			Devices
// @Accept 			json
// @Produce			json
// @Param 			user path string true "user devices"
// @Success			200  {object} models.GetDevicesResponse
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) GetDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := r.PathValue("user")
	fmt.Println(user)
	res, err := u.Dev.GetDevices(&models.GetDevicesRequest{User: user})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// @Router 			/devices/{device} [delete]
// @Summary			To remove something from home
// @Description 	user profile
// @Security		BearerAuth
// @Tags			Devices
// @Accept 			json
// @Produce			json
// @Param 			device path string true "user devices"
// @Success			200  {object} models.DeleteDeviceResponse
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) DeleteDev(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	device := r.PathValue("device")
	res, err := u.Dev.DeleteDevice(&models.DeleteDeviceRequest{Name: device})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// @Router 			/devices/command/speaker [put]
// @Summary			To give a command fo speaker
// @Description 	Commands
// @Security		BearerAuth
// @Tags			Devices
// @Accept 			json
// @Produce			json
// @Param 			body body models.Speaker true "user devices"
// @Success			201  {object} string
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) Speaker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var speaker models.Speaker
	if err := json.NewDecoder(r.Body).Decode(&speaker); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	byte, err := json.Marshal(speaker)
	if err != nil {
		log.Println(err)
	}
	if err := u.PublishAdjust("speaker", byte); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "succesfull")
}

// @Router 			/devices/command/vaccum [put]
// @Summary			To give a command fo vaccum
// @Description 	Commands
// @Security		BearerAuth
// @Tags			Devices
// @Accept 			json
// @Produce			json
// @Param 			body body models.Vaccum true "user devices"
// @Success			201  {object} string
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) Vaccum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Vaccum models.Vaccum
	if err := json.NewDecoder(r.Body).Decode(&Vaccum); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	byte, err := json.Marshal(Vaccum)
	if err != nil {
		log.Println(err)
	}
	if err := u.PublishAdjust("vaccum", byte); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "succesfull")
}

// @Router 			/devices/command/alarm [put]
// @Summary			To give a command fo alarm
// @Description 	Commands
// @Security		BearerAuth
// @Tags			Devices
// @Accept 			json
// @Produce			json
// @Param 			body body models.Alarm true "user devices"
// @Success			200  {object} string
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) Alarm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var alarm models.Alarm
	if err := json.NewDecoder(r.Body).Decode(&alarm); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	byte, err := json.Marshal(alarm)
	if err != nil {
		log.Println(err)
	}
	if err := u.PublishAdjust("alarm", byte); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "succesfull")
}

// @Router 			/devices/command/door [put]
// @Summary			To give a command fo alarm
// @Description 	Commands
// @Security		BearerAuth
// @Tags			Devices
// @Accept 			json
// @Produce			json
// @Param 			body body models.Door true "user devices"
// @Success			200  {object} string
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) Door(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var door models.Door
	if err := json.NewDecoder(r.Body).Decode(&door); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	byte, err := json.Marshal(door)
	if err != nil {
		log.Println(err)
	}
	if err := u.PublishAdjust("door", byte); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "succesfull")
}

// @Router 			/devices/speaker/{dname} [get]
// @Summary			To get information about speaker
// @Description 	Commands
// @Security		BearerAuth
// @Tags			Devices
// @Produce			json
// @Param 			dname path string true "user devices"
// @Success			200  {object} models.SpeakerGet
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) SpeakerGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dname := r.PathValue("dname")

	res, err := u.Dev.SpeakerGet(&models.GetDevice{Device: dname})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)

}

// @Router 			/devices/vaccum/{dname} [get]
// @Summary			To get information about vaccum
// @Description 	Commands
// @Security		BearerAuth
// @Tags			Devices
// @Produce			json
// @Param 			dname path string true "user devices"
// @Success			200  {object} models.Vaccum
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) VaccumGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dname := r.PathValue("dname")

	res, err := u.Dev.VaccumGet(&models.GetDevice{Device: dname})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)

}

// @Router 			/devices/alarm/{dname} [get]
// @Summary			To get information about alarm
// @Description 	Commands
// @Security		BearerAuth
// @Tags			Devices
// @Produce			json
// @Param 			dname path string true "user devices"
// @Success			200  {object} models.Alarm
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) AlarmGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dname := r.PathValue("dname")

	res, err := u.Dev.AlarmGet(&models.GetDevice{Device: dname})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)

}

// @Router 			/devices/door/{dname} [get]
// @Summary			To get information about door
// @Description 	Commands
// @Security		BearerAuth
// @Tags			Devices
// @Produce			json
// @Param 			dname path string true "user devices"
// @Success			200  {object} models.Door
// @Failure			400  {object} models.StandartError
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError
func (u *Handler) DoorGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dname := r.PathValue("dname")

	res, err := u.Dev.DoorGet(&models.GetDevice{Device: dname})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)

}

func (u *Handler) PublishAdjust(type1 string, device []byte) error {
	var req = models.Message{
		Type:    type1,
		Payload: device,
	}
	if err := u.PUB.Adjust(req); err != nil {
		return err
	}
	return nil
}
