package handler

import (
	"net/http"

	"emap/map/internal/logic"
	"emap/map/internal/svc"
	"emap/map/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MapHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMapLogic(r.Context(), svcCtx)
		resp, err := l.Map(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
