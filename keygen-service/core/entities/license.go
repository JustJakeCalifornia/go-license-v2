package entities

const (
	AVAILABLE = true
	USED      = false
)

type License struct {
	// storing in database later
	// then in repo layer, we can use this to check if license is valid
	Key         string `json:"key"`
	IsUsed      bool   `json:"isUsed"`
	CreatedAt   string `json:"createdAt"`
	ExpiresAt   string `json:"expiresAt"`
	CompanyId   string `json:"companyId"`
	CompanyName string `json:"companyName"`
	TeamId      string `json:"teamId"`
	TeamName    string `json:"teamName"`
}
