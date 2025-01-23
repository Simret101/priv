package captha

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// Initialize a global store for CAPTCHA
var store = base64Captcha.DefaultMemStore

// GenerateCaptcha generates a new CAPTCHA and returns its base64 string along with its ID.
func GenerateCaptcha(c *gin.Context) {
	// Configure the CAPTCHA driver with height, width, length, max skew, and dot count
	driver := base64Captcha.NewDriverDigit(100, 240, 4, 0.7, 80)
	// Create a new CAPTCHA with the configured driver and store
	captcha := base64Captcha.NewCaptcha(driver, store)

	// Generate the CAPTCHA and get its ID and base64 string
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		// Return an error response if CAPTCHA generation fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate CAPTCHA"})
		return
	}

	// Create a data URI for embedding the image directly in an HTML tag
	dataURI := fmt.Sprintf("data:image/png;base64,%s", b64s)

	// Return the CAPTCHA ID and image data URI in the response
	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": dataURI, // Data URI for direct embedding
	})
}
