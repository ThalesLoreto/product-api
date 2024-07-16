package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ThalesLoreto/product-api/internal/dto"
	"github.com/ThalesLoreto/product-api/internal/entity"
	"github.com/ThalesLoreto/product-api/internal/infra/database"
	"github.com/go-chi/jwtauth/v5"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, expiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       db,
		Jwt:          jwt,
		JwtExpiresIn: expiresIn,
	}
}

// CreateUser godoc
// @Summary Login user
// @Description Login user
// @Tags users
// @Accept json
// @Produce json
// @Param input body dto.LoginUserInput true "User Credentials"
// @Success 200 {object} dto.LoginUserOutput
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 401 {string} string
// @Router /users/login [post]
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.LoginUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := uh.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = u.ComparePassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := uh.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": jwtauth.ExpireIn(time.Duration(uh.JwtExpiresIn) * time.Minute),
	})
	accessToken := dto.LoginUserOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserInput true "User data"
// @Success 201 {string} string "User created"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = uh.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
