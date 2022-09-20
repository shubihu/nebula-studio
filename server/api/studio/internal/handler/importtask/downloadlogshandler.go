// Code generated by goctl. DO NOT EDIT.
package importtask

import (
	"net/http"

	"github.com/vesoft-inc/go-pkg/validator"
	"github.com/vesoft-inc/nebula-studio/server/api/studio/pkg/ecode"

	"github.com/vesoft-inc/nebula-studio/server/api/studio/internal/logic/importtask"
	"github.com/vesoft-inc/nebula-studio/server/api/studio/internal/svc"
	"github.com/vesoft-inc/nebula-studio/server/api/studio/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadLogsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadLogsRequest
		if err := httpx.Parse(r, &req); err != nil {
			err = ecode.WithCode(ecode.ErrParam, err)
			svcCtx.ResponseHandler.Handle(w, r, nil, err)
			return
		}
		if err := validator.Struct(req); err != nil {
			svcCtx.ResponseHandler.Handle(w, r, nil, err)
			return
		}

		l := importtask.NewDownloadLogsLogic(r.Context(), svcCtx)
		err := l.DownloadLogs(req)
		svcCtx.ResponseHandler.Handle(w, r, nil, err)
	}
}
