package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

type CreateRequest struct {
	FirstName string `form:"firstname" binding:"required"`
	LastName  string `form:"lastname" binding:"required"`
	Nickname  string `form:"nickname" binding:"required"` // unique
	Email     string `form:"email" binding:"required"`    // unique
	Sex       int    `form:"sex"`                         // 1 – female; 2 – male.
	BDate     string `form:"bdate"`
	Picture   string `form:"picture"`
	Desc      string `form:"desc"`
	Status    string `form:"status"`
	Password  string `form:"password" binding:"required"`
}

type CreateResponse struct {
	Version string      `json:"v"`
	User    models.User `json:"user"`
}

func (h Handler) CreateUser(c *gin.Context) {

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	if err := h.db.Where(&models.User{Nickname: r.Nickname, Email: r.Email}).Error; err == nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Nickname or Email already registered",
		})
		return
	}

	d, _ := time.Parse("2006-01-02", r.BDate)

	us := models.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Nickname:  r.Nickname,
		Email:     r.Email,
		Sex:       r.Sex,
		BDate:     d,
		Picture:   r.Picture,
		Desc:      r.Desc,
		Status:    r.Status,
		Badges:    []models.Badges{},
	}

	h.db.Create(&us)
	if err := us.SetPassword(r.Password); err != nil {
		log.Error(err)
	}

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		User:    us,
	})
}
