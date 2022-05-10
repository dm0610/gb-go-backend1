package defmux

import (
	"goback1/lesson4/reguser/internal/infrastructure/db/mem/usermemstore"
	"goback1/lesson4/reguser/internal/usecases/app/repos/userrepo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter_CreateUser(t *testing.T) {
	ust := usermemstore.NewUsers()
	us := userrepo.NewUsers(ust)
	rt := NewRouter(us)

	hts := httptest.NewServer(rt)

	r, _ := http.NewRequest("POST", hts.URL+"/create", strings.NewReader(`{"name":"user123"}`))
	r.SetBasicAuth("admin", "admin")

	cli := hts.Client()

	resp, err := cli.Do(r)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Error("status wrong:", resp.StatusCode)
	}

	// (&http.Client{}).Get(httptest.NewServer(nil).URL)
}

func TestRouter_CreateUser2(t *testing.T) {
	ust := usermemstore.NewUsers()
	us := userrepo.NewUsers(ust)
	rt := NewRouter(us)

	h := rt.AuthMiddleware(http.HandlerFunc(rt.CreateUser)).ServeHTTP

	w := &httptest.ResponseRecorder{}
	r := httptest.NewRequest("POST", "/create", strings.NewReader(`{"name":"user123"}`))
	r.SetBasicAuth("admin", "admin")

	h(w, r)

	if w.Code != http.StatusCreated {
		t.Error("status wrong:", w.Code)
	}
}