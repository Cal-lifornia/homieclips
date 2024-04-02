package components

type Profile struct {
	Email   string `json:"email" mapstructure:"email"`
	Name    string `json:"name" mapstructure:"name"`
	Picture string `json:"picture" mapstructure:"picture"`
}
