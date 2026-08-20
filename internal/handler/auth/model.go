package auth

type RegisterRequest struct {
	Email string `json:"email" validate:"required" example:"test@mail.ru"`
}

type VerifyRequest struct {
	Token    string `json:"token" validate:"required" example:"a17a89b6-eff8-4b05-9be7-407d62b46c4a"`
	Username string `json:"username" validate:"required" example:"test"`
	Password string `json:"password" validate:"required" example:"Test_1234"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required" example:"test@mail.ru"`
	Password   string `json:"password" validate:"required" example:"Test_1234"`
}

type LoginResponce struct {
	AccessToken string `json:"accessToken" example:"gG3d83bJk5jawy31...WE4u"`
}
