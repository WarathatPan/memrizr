package handler

import (
	"log"
	"net/http"

	"github.com/WarathatPan/memrizr/account/model"
	"github.com/WarathatPan/memrizr/account/model/apperrors"
	"github.com/gin-gonic/gin"
)

// signupReq is not exported, hence the lowercase name
// it is used for validation and json marshalling
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup handler
func (h *Handler) Signup(c *gin.Context) {
	// define a variable to which we'll bind incoming
	// json body, {email, password}
	var req signupReq

	// Bind incoming json to struct and check for validation errors
	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	// ...data binding and instantiation of User model...
	ctx := c.Request.Context()
	err := h.UserService.Signup(ctx, u)

	// ...error handling ...

	// create token pair as strings
	tokens, err := h.TokenService.NewPairFromUser(ctx, u, "")

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
