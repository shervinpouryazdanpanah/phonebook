package phone

type StoreRequest struct {
	Number string `form:"number" binding:"required,email,min=3,max=100"`
}
