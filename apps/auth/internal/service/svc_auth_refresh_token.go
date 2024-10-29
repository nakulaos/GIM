package service

import (
	"GIM/pkg/common/xjwt"
	"GIM/pkg/common/xlog"
	"GIM/pkg/constant"
	"GIM/pkg/proto/pb_auth"
	"context"
)

func (s *authService) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, _ error) {
	resp = &pb_auth.RefreshTokenResp{}
	var (
		refreshToken *xjwt.JwtToken
		accessToken  *xjwt.JwtToken
		sessionId    string
		err          error
	)
	/*
		1. 解析refreshToken
		2. 获取uid，platform对应sessionId，比较sessionId是否一致
		3. 生成新的accessToken
		4. 设置新的accessToken对应sessionId
	*/
	refreshToken, err = xjwt.Decode(req.RefreshToken)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_JWT_TOKEN_ERR, ERROR_AUTH_JWT_TOKEN_ERR)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	sessionId, err = s.authCache.GetRefreshTokenSessionId(refreshToken.Uid, refreshToken.Platform)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_REDIS_GET_FAILED, ERROR_AUTH_REDIS_GET_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	if sessionId != refreshToken.SessionId {
		resp.Set(ERROR_CODE_AUTH_JWT_SESSION_ID_ERR, ERROR_AUTH_JWT_SESSION_ID_ERR)
		return
	}
	accessToken, err = xjwt.CreateToken(refreshToken.Uid, refreshToken.Platform, true, constant.CONST_DURATION_SHA_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED, ERROR_AUTH_GENERATE_TOKEN_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	err = s.authCache.SetAccessTokenSessionId(refreshToken.Uid, refreshToken.Platform, accessToken.SessionId)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_REDIS_SET_FAILED, ERROR_AUTH_REDIS_SET_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	resp.AccessToken = &pb_auth.Token{
		Token:  accessToken.Token,
		Expire: accessToken.Expire,
	}
	return
}
