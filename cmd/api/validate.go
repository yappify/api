package main

import "fmt"

func validateCreateUserPayload(payload CreateUserPayload) error {
	if payload.AuthType == "" || payload.Name == "" || payload.Email == "" || payload.Password == "" {
		return fmt.Errorf("auth_type, name, email, and password are required")
	}

	switch payload.AuthType {
	case "credential", "google", "github":
		return nil
	default:
		return fmt.Errorf("invalid auth_type; must be 'credential', 'google', or 'github'")
	}
}
