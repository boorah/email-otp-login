package helpers

import (
	"math/rand"
	"strconv"
	"time"
)

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateOTP() string {
	otp := ""
	for range 6 {
		digit := r.Intn(10)
		otp += strconv.Itoa(digit)
	}

	return otp
}
