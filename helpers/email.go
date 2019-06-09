package helpers

import (
	"errors"
	"github.com/labstack/gommon/log"
	"gopkg.in/resty.v1"
	"net/http"
)

// CREATE_ACCOUNT contain SendGrid template id
const CREATE_ACCOUNT = "d-8dcb61a96c1142fab385d1f2e8e0dbc3"

//const ACTIVATED_ACCOUNT = "d-3b2da2e5cd324497bd23cf8cac87464b"
//const PASSWORD_RESTORE = "d-07c3179af9794a099dd7811b133da955"

type sendGridReq struct {
	From             map[string]string `json:"from"`
	Personalizations []struct {
		To []struct {
			Email string `json:"email"`
		} `json:"to"`
		DynamicTemplateData map[string]string `json:"dynamic_template_data"`
	} `json:"personalizations"`
	TemplateId string `json:"template_id"`
}

// SendEmail send email activation
func SendEmail(template string, email string, params map[string]string) error {
	resp, err := resty.R().
		SetAuthToken(GetEnvWithPanic("SENDGRID_TOKEN")).
		SetBody(sendGridReq{
			From: map[string]string{
				"email": "support@anibe.ru",
			},
			Personalizations: []struct {
				To []struct {
					Email string `json:"email"`
				} `json:"to"`
				DynamicTemplateData map[string]string `json:"dynamic_template_data"`
			}{
				{
					To: []struct {
						Email string `json:"email"`
					}{
						{Email: email},
					},
					DynamicTemplateData: params,
				},
			},
			TemplateId: template,
		}).
		Post("https://api.sendgrid.com/v3/mail/send")
	if err != nil {
		log.Error(err)
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		// auth is not required
		return nil
	}
	return errors.New("sendgrid error")
}
