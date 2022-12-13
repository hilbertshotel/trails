package handlers

import (
	"net/http"
	"trails/dep"
)

func wsa(w http.ResponseWriter, r *http.Request, d *dep.Dependencies) {
	addr := "ws://" + d.Cfg.HostAddr + "/sorter"
	w.Write([]byte(addr))
}
