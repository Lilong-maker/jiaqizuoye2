package service

import (
	"context"
	"jiaqizuoye2/src/basic/config"
	"jiaqizuoye2/src/handler/model"

	__ "jiaqizuoye2/src/basic/proto"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {

	var user model.User
	err := user.FindUser(config.DB, in.Name)
	if err != nil {
		return &__.LoginResp{
			Msg:  "用户不存在",
			Code: 400,
		}, nil
	}
	if user.Password != in.Password {
		return &__.LoginResp{
			Msg:  "登录错误",
			Code: 400,
		}, nil
	}
	return &__.LoginResp{
		Msg:  "登录成功",
		Code: 200,
	}, nil

}

// SayHello implements helloworld.GreeterServer
func (s *Server) SpAdd(_ context.Context, in *__.SpAddReq) (*__.SpAddResp, error) {

	var sp model.Sp
	err := sp.FingSp(config.DB, in.Name)
	if err != nil {
		return &__.SpAddResp{
			Msg:  "视频不存在",
			Code: 400,
		}, nil
	}
	sp.Name = in.Name
	sp.SpLen = in.SpLen
	sp.SpNum = in.SpNum
	err = sp.SpAdd(config.DB)
	if err != nil {
		return &__.SpAddResp{
			Msg:  "视频添加失败",
			Code: 400,
		}, nil
	}
	return &__.SpAddResp{
		Msg:  "商品添加成功",
		Code: 200,
	}, nil
}
