package user

import (
	"fmt"
	"log"
	"net/http"
	password "test-project/cmd/services/auth"
	"test-project/helper"
	"test-project/types"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.Userstore
}

func NewHandler(store types.Userstore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUser
	if err := helper.ParseJson(r, &user); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err)
		return
	}
	u, err := h.store.GetUserByEmail(user.Email)

	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}
	if !password.CompareHashedPassword(u.Password, []byte(user.Password)) {
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "manoj", Value: "kunwar"})
	helper.WriteJson(w, http.StatusAccepted, map[string]any{"status": true, "c_msg": fmt.Sprintf("user logged in! welcome %q", user.Email)})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// ! get json payload;
	var payload types.RegiterPayload
	if err := helper.ParseJson(r, &payload); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err)
	}
	fmt.Println(payload)
	// Check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		log.Println("User already exists!")
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists with email %q", payload.Email))
		return
	}
	// hashedPassowrd, err := password.HashedPassowrd(payload.Password)
	hashedPassowrd, err := password.HashedPassowrd(payload.Password)
	if err != nil {
		log.Println("Unable able to create hash password! :- " + payload.Password)
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("not able to create your account at the moment"))
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassowrd,
	})

	if err != nil {
		log.Println("Not able to create hash password!")
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("not able to create your account at the moment"))
		return
	}

	helper.WriteJson(w, http.StatusCreated, map[string]any{"status": true, "c_msg": "Your account has been created!"})
}
