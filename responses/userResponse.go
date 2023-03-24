package responses

type AuthResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Role         string `json:"role"`
	IsBlocked    bool   `json:"isBlocked"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type UserResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	IsBlocked bool   `json:"isBlocked"`
}
