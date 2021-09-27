package input

type C2S_Default struct {
}

type C2S_Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
