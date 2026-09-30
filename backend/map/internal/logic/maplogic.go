package logic

import (
	"context"

	"emap/map/internal/svc"
	"emap/map/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MapLogic {
	return &MapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MapLogic) Map(req *types.Request) (resp *types.Response, err error) {
	return &types.Response{Message: l.svcCtx.Config.Amap_Key}, nil
}
