package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Anandhu4456/rms/pkg/model"
	"gorm.io/gorm"
)

// signup
var db *gorm.DB

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
	query := `insert into users (email,password_hash,confirm_password_hash,address)values(?,?,?,?)`

	if err := db.Exec(query, userDetails.Email, userDetails.PasswordHash, userDetails.ConfirmPasswordHash, userDetails.Address); err != nil {
		http.Error(w, "user signup failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user signup successfully"))
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
