package handler

import (
	dblayer "filestore-server/db"
	"filestore-server/util"
	"io/ioutil"
	"net/http"
)

const (
	pwd_salt = "*#890"
)

func SighupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if len(username) < 3 || len(password) < 5 {
		w.Write([]byte("invalid parameter"))
		return
	}
	enc_passwd := util.Sha1([]byte(password + pwd_salt))
	suc := dblayer.UserSignup(username, enc_passwd)
	if suc {
		w.Write([]byte("successed"))
	} else {
		w.Write([]byte("failed"))
	}
}
