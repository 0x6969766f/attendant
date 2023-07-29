package image

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func (i *ImageAPI) GetImage(w http.ResponseWriter, r *http.Request) error {
	image := filepath.Join("./images", chi.URLParam(r, "filename"))
	http.ServeFile(w, r, image)
	return nil
}
