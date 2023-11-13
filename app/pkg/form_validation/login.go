package formvalidation

// Login login form validation
type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
