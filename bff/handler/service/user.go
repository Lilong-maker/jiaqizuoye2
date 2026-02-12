package service

import (
	"jiaqizuoye2/bff/basic/config"
	"jiaqizuoye2/bff/handler/request"
	"jiaqizuoye2/pkg"
	__ "jiaqizuoye2/src/basic/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var form request.Login
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	r, err := config.UserClient.Login(c, &__.LoginReq{
		Name:     form.Name,
		Password: form.Password,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
	})
	return
}
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		return
	}
	upload, err := pkg.Upload(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "上传失败",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"data": upload,
	})
	return
}
func SpAdd(c *gin.Context) {
	var form request.Sp
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	r, err := config.UserClient.SpAdd(c, &__.SpAddReq{
		Name:  form.Name,
		SpLen: form.SpLen,
		SpNum: form.SpNum,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
	})
	return
}
