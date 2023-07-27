package health

import "net/http"

func Check(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}
