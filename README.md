# Twilio-OTP-Verification Using Golang

Example of using Twilio to send OTP verification code to a user and verifying the code.
## Project Overview Through a Diagram
![](./Diagram/Twilio-OTP-Diagram.png)

## Setup

1. Clone the repo

2. Create a Twilio account and get your account SID and auth token

3. Create a Twilio verify service and get the service SID

4. Create a .env file in the root of the project and add the in the above credentials from Twilio

```bash
TWILIO_ACCOUNT_SID=<ACCOUNT SID>
TWILIO_AUTHTOKEN=<AUTH TOKEN>
TWILIO_SERVICES_ID=<SERVICE ID>
```

5. Install dependencies

```bash
go mod download
```

6. Run the server

```bash
cd cmd
go run main.go
```

## API Documentation
---
### Send OTP

Send a POST request to the /otp endpoint with the following body to send an OTP to a user's phone number

POST /otp

Request Body

```json
{
  "phoneNumber": "<phone-number-with-country-code>"
}
```

```bash
curl -H "Content-Type: application/json" -X POST -d '{"phoneNumber": "+917735278240"}' http://localhost:8000/otp
```

_Be sure to include the country code in the phone number_

Response

```json
{
  "status": 202,
  "message": "success",
  "data": "OTP sent successfully"
}
```
---
### Verify OTP

Verify a user's OTP by sending a POST request to the /verify endpoint with the following body that contains the phone number and the OTP code received by the user

POST /verifyOTP

Request Body

```json
{
  "user": {
    "phoneNumber": "<phone-number-with-country-code>"
  },
  "code": "<code here>"
}
```

```bash
curl -H "Content-Type: application/json" -X POST -d '{"user": {"phoneNumber": "+917735278240"}, "code":"163221"}' http://localhost:8000/verifyOTP
```

Response

```json
{
  "status": 202,
  "message": "success",
  "data": "OTP verified successfully"
}
```
---