package handler

import (
	"os"
)

func init() {
	// Set TZ to UTC so that the datetime objects are renderred into UTC by the API.
	os.Setenv("TZ", "UTC")
}
