package auth

type CredentialsRegisterLinkSendRequest struct {
	Email string `json:"email" validate:"required" example:"test@mail.ru"`
}

type CredentialsRegisterLinkConfirmRequest struct {
	Token          string `json:"token" validate:"required" example:"a17a89b6-eff8-4b05-9be7-407d62b46c4a"`
	Username       string `json:"username" validate:"required" example:"test"`
	Password       string `json:"password" validate:"required" example:"Test_1234"`
	PasswordRepeat string `json:"passwordRepeat" validate:"required" example:"Test_1234"`
}

type CredentialsLoginRequest struct {
	Username string `json:"username" validate:"required" example:"test@mail.ru / test"`
	Password string `json:"password" validate:"required" example:"Test_1234"`
}

type LoginResponce struct {
	AccessToken string `json:"accessToken" example:"gG3d83bJk5jawy31...WE4u"`
}

type TokenRefreshResponce struct {
	AccessToken string `json:"accessToken" example:"gG3d83bJk5jawy31...WE4u"`
}
