package request

type Login struct {
	Name     string `form:"name"   binding:"required"`
	Password string `form:"password"  binding:"required"`
}
type Sp struct {
	Name  string `form:"name"   binding:"required"`
	SpLen string `form:"splen"  binding:"required"`
	SpNum string `form:"spnum"  binding:"required"`
}
