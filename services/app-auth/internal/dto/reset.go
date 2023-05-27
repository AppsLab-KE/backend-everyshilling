package dto

type ResetReq struct {
	ConfirmPassword string `json:"confirm_password,omitempty"`
	Password        string `json:"password,omitempty"`
	TrackerUUID     string `json:"tracker_uuid"`
}
type ResetRes struct {
}
