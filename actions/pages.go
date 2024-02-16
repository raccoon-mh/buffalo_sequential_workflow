package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, tr.HTML("home/index.html"))
}

// 로그인 폼 + 로그인폼은 기본 렌더를 따름.
func AuthLoginHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("auth/sign-in.html"))
}
