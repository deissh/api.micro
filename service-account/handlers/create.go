package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

// CreateRequest request query params
type CreateRequest struct {
	Hash      string `json:"password" binding:"required,min=6,max=20"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Nickname  string `json:"nickname" binding:"required"` // unique
	Email     string `json:"email" binding:"required"`    // unique
	Sex       int    `json:"sex"`                         // 1 – female; 2 – male.
	BDate     string `json:"bdate"`
	Picture   string `json:"picture"`
	Desc      string `json:"desc"`
	Status    string `json:"status"`
}

// CreateResponse response structure
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
// @Param firstname body string true "user firstname"
// @Param lastname body string true "user lastname"
// @Param nickname body string true "user nickname"
// @Param email body string true "user email"
// @Param password body string true "user password"
// @Param sex body int false "user sex"
// @Param bdate body string false "user bdate"
// @Param picture body string false "user picture"
// @Param desc body string false "user desc"
// @Param status body string false "user status"
// @Success 200 {object} handlers.CreateResponse
// @Failure 400 {object} handlers.ResponseData
// @Router /account.create [Get]
func (h Handler) AccountCreate(c *gin.Context) {

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   err.Error(),
		})
		return
	}
	var user models.User
	h.db.Where(&models.User{Nickname: r.Nickname}).Or(&models.User{Email: r.Email}).First(&user)
	if user.Email == r.Email || user.Nickname == r.Nickname {
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
		Role:      "user",
	}

	if err := us.SetPassword(r.Hash); err != nil {
		log.Error(err)
	}
	h.db.Create(&us)

	// sending activation email
	activateToken, err := helpers.GenerateRandomString(128)
	if err != nil {
		log.Error(err)
	}
	h.db.Create(&models.ActivateTokens{
		Token: activateToken,
		Email: us.Email,
	})
	_ = helpers.SendEmail(
		helpers.CREATE_ACCOUNT,
		us.Email,
		map[string]string{
			"activate_url": "https://anibe.ru/account/activate?token=" + activateToken,
			"first_name":   us.FirstName,
			"last_name":    us.LastName,
		},
	)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		User:    us.View(),
	})
}
