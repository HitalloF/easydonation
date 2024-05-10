package models

type Family struct {
    ID               int    `json:"id"`
    Name             string `json:"name"`
    LocalDonationID  int    `json:"local_donation_id"`
    LastDonationID   int    `json:"last_donation_id"`
}
