package main

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"
	"AstraFlow-go/pkg/config"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	config.InitConfig()
	database.InitDB()

	username := "demo@astraflow.com"
	inputPassword := "Demo@123"

	fmt.Printf("\n--- Debugging User: %s ---\n", username)

	// 1. Check by Username
	var userByUsername model.User
	err := database.DB.Where("username = ?", username).First(&userByUsername).Error
	if err != nil {
		fmt.Printf("[Check 1] Find by Username: FAILED (%v)\n", err)
	} else {
		fmt.Printf("[Check 1] Find by Username: FOUND (ID: %d, RoleID: %d, PwdHash: %s...)\n", 
			userByUsername.ID, userByUsername.RoleID, userByUsername.Password[:10])
	}

	// 2. Check by Email
	var userByEmail model.User
	// Note: user_repository.go actually queries the 'username' field when FindByUsername is called, 
	// but let's check if the 'email' field in DB is actually populated.
	err = database.DB.Where("email = ?", username).First(&userByEmail).Error
	if err != nil {
		fmt.Printf("[Check 2] Find by Email column: FAILED (%v)\n", err)
	} else {
		fmt.Printf("[Check 2] Find by Email column: FOUND (ID: %d)\n", userByEmail.ID)
	}

	// 3. Verify Password
	var targetUser *model.User
	if userByUsername.ID != 0 {
		targetUser = &userByUsername
	} else if userByEmail.ID != 0 {
		targetUser = &userByEmail
	}

	if targetUser != nil {
		err = bcrypt.CompareHashAndPassword([]byte(targetUser.Password), []byte(inputPassword))
		if err != nil {
			fmt.Printf("[Check 3] Password Verification: FAILED (%v)\n", err)
			
			// Generate what the hash SHOULD be
			hashed, _ := bcrypt.GenerateFromPassword([]byte(inputPassword), bcrypt.DefaultCost)
			fmt.Printf("          Expected similar to: %s\n", string(hashed))
			fmt.Printf("          Actual in DB:        %s\n", targetUser.Password)
		} else {
			fmt.Printf("[Check 3] Password Verification: SUCCESS\n")
		}
	} else {
		fmt.Println("[Check 3] Skipping password check (User not found)")
	}
}
