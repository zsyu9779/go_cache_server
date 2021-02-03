package HTTP

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"net/http"
)

type statusHandler struct {
	*Server
}

func (h *statusHandler)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	if r.Method !=http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	status ,err :=json.Marshal(h.GetStat())
	if err != nil {
		logs.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(status)
}
