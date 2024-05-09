package user

import (
	"net/http"
)

func (ist *instance) Update(w http.ResponseWriter, r *http.Request) {

	r.Response.StatusCode = http.StatusOK
	return
}
