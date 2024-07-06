package dto

type KeyDTO struct {
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	MembershipIds string `json:"membership_ids"`
	LockIds       string `json:"lock_ids"`
}
