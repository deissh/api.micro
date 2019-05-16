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

// AccountCreate godoc
// @Summary Create new account
// @Description Register new account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param firstname query string false "user firstname"
// @Param lastname query string false "user lastname"
// @Param sex query string false "user sex"
// @Param bdate query string false "user bdate"
// @Param picture query string false "user picture"
// @Param desc query string false "user desc"
// @Param status query string false "user status"
// @Param password query string false "user password"
// @Success 200 {object} handlers.CreateResponse
// @Failure 400 {object} handlers.ResponseData
// @Failure 500 {object} handlers.ResponseData
// @Router /account.create [Get]
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
