package requester

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NewHatch ...
func NewHatch(username, password string) (hp *Hatch) {
	var h Hatch
	h.username = username
	h.password = password
	hp = &h
	return
}

const baseURL = "http://prod.usehatchapp.com"

// Hatch ...
type Hatch struct {
	hc http.Client

	username string
	password string
}

func (h *Hatch) login() (err error) {
	payload := map[string]string{
		"username": h.username,
		"password": h.password,
	}

	var bs []byte
	if bs, err = json.Marshal(payload); err != nil {
		return
	}

	var resp *http.Response
	if resp, err = New(&h.hc, baseURL).Post("api/login", bs, nil); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("error: unsuccessful login")
	}

	return
}

func (h *Hatch) getUser() (user interface{}, err error) {
	var resp *http.Response
	if resp, err = New(&h.hc, baseURL).Get("api/users/00000769", nil); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("error: unsuccessful login")
	}

	return
}
