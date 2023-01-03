package user

type RegsiterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `Account:"email" binding:"required,email"`
	Password   string `Account:"password" binding:"required"`
}