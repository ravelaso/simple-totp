package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <token>", os.Args[0])
	}

	secret := os.Args[1]
	// Get the current time
	now := time.Now()

	// Generate the OTP
	passcode, err := totp.GenerateCode(secret, now)
	if err != nil {
		log.Fatalf("Error generating OTP: %v", err)
	}

	fmt.Printf("Your OTP is: %s\n", passcode)
}
