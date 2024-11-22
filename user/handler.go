package user

import (
	"assignment_task/builders"
	"assignment_task/models"
	"assignment_task/requests"
	"assignment_task/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Handler struct {
	store Storer
}

type Storer interface {
	GetUserByEmail(email string) models.User
	RegisterUser(models.User) error
}

func (h *Handler) Register(c echo.Context) error {
	registerRequest := new(requests.RegisterRequest)

	if err := c.Bind(registerRequest); err != nil {
		return err
	}

	user := h.store.GetUserByEmail(registerRequest.Email)

	if user.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "User already exists")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(registerRequest.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	newUser := builders.NewUserBuilder().SetEmail(registerRequest.Email).
	SetPassword(string(encryptedPassword)).
	SetName(registerRequest.Name).
	Build()

	h.store.RegisterUser(newUser)

	return responses.MessageResponse(c, http.StatusCreated, "User successfully created")
}

func (h Handler) Login(c echo.Context) error {
	loginRequest := new(requests.LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	user := h.store.GetUserByEmail(loginRequest.Email)

	if user.ID == 0 || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil) {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
	}

	accessToken, exp, err := createAccessToken(&user)
	if err != nil {
		return err
	}
	refreshToken, err := createRefreshToken(&user)
	if err != nil {
		return err
	}
	res := responses.NewLoginResponse(accessToken, refreshToken, exp)

	c.SetCookie(&http.Cookie{
			Name:     "jwt_token",
			Value:    accessToken,
			Expires:  time.Unix(exp, 0),
			HttpOnly: true, 
			Path:     "/",  
	})

	
	return responses.Response(c, http.StatusOK, res)
}