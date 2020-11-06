package table

type ActivationToken string

type EmailActivationToken struct {
	ActivationToken ActivationToken `json:"activation_token" gorm:"primary_key"`
	UserId          UserId          `json:"user_id"`
	ExpiresAt       int64           `json:"expires_at"`
	UpdatedAt       string          `json:"updated_at" sql:"DEFAULT:current_timestamp"`
	CreatedAt       string          `json:"created_at" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}
