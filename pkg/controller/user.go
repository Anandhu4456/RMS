package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Anandhu4456/rms/pkg/model"
)

// signup

func Signup(w http.ResponseWriter, r *http.Request) {
	var userDetails model.UserSignup
	// parse json request
	if err := json.NewDecoder(r.Body).Decode(&userDetails); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if userDetails.PasswordHash != userDetails.ConfirmPasswordHash {
		http.Error(w, "password doesnt match", http.StatusBadRequest)
		return
	}
	// TODO : store in db
}

// login
func Login() {

}

// get jobs
func GetJob() {

}

// get job by id
func GetJobById() {

}

// upload resume

func UploadResume() {

}
