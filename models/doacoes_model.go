package models

import "time"

type Donation struct {
	DonationID  int       `json:"donation_id"`
	FamilyID    int       `json:"family_id"`
	Date        time.Time `json:"date"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Confirmed   bool      `json:"confirmed"`
}