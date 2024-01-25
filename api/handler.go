package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/panchanandevops/Twilio-OTP-Verification/data"
)

const appTimeout = time.Second * 10

// sendSMS is a Gin middleware function that handles the sending of SMS with OTP.
func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a context with timeout
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		// Initialize payload to hold request data
		var payload data.OTPData
		app.validateBody(c, &payload)

		// Create a new data structure with sanitized input
		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		// Call Twilio service to send OTP
		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		// Respond with success message
		app.writeJSON(c, http.StatusAccepted, "OTP sent successfully")
	}
}

// verifySMS is a Gin middleware function that handles the verification of SMS OTP.
func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a context with timeout
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		// Initialize payload to hold request data
		var payload data.VerifyData
		app.validateBody(c, &payload)

		// Create a new data structure with sanitized input
		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		// Call Twilio service to verify OTP
		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		fmt.Println("err: ", err)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		// Respond with success message
		app.writeJSON(c, http.StatusAccepted, "OTP verified successfully")
	}
}
