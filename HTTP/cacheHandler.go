package HTTP

import (
	"github.com/astaxie/beego/logs"
	"io"
	"net/http"
	"strings"
)

type cacheHandler struct {
	*Server
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	method := r.Method
	if method == http.MethodGet {
		v, e := h.Get(key)
		if e != nil {
			logs.Error(e.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(v) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(v)
		return
	}
	if method == http.MethodPut {
		v, _ := io.ReadAll(r.Body)
		if len(v) != 0 {
			e := h.Set(key, v)
			if e != nil {
				logs.Error(e.Error())
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
	if method == http.MethodDelete {
		e := h.Del(key)
		if e != nil{
			logs.Error(e.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}
