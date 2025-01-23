package recaptcha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"

// recaptchaPrivateKey holds the secret key which will be initialized via Init function
var recaptchaPrivateKey string

// check uses the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
func check(remoteip, response string) (r RecaptchaResponse, err error) {
	resp, err := http.PostForm(recaptchaServerName,
		url.Values{"secret": {recaptchaPrivateKey}, "remoteip": {remoteip}, "response": {response}})
	if err != nil {
		log.Printf("Post error: %s\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: could not read body: %s", err)
		return
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Printf("Read error: got invalid JSON: %s", err)
		return
	}
	return
}

// Confirm is the public interface function.
// It calls check, which checks whether the client answered correctly.
func Confirm(remoteip, response string) (result bool, err error) {
	resp, err := check(remoteip, response)
	if err != nil {
		return false, err
	}
	result = resp.Success
	return
}

// Init initializes the recaptcha private key from the environment variable
func Init() {
	recaptchaPrivateKey = os.Getenv("recaptcha.secret")
	if recaptchaPrivateKey == "" {
		log.Fatal("RECAPTCHA_SECRET is not set in the environment variables")
	}
	fmt.Println("ReCaptcha Secret Key Initialized Successfully")
}
