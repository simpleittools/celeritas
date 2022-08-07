package celeritas

import (
	"github.com/justinas/nosurf"
	"net/http"
	"strconv"
)

func (c *Celeritas) SessionLoad(next http.Handler) http.Handler {
	c.InfoLog.Println("SessionLoad called")
	return c.Session.LoadAndSave(next)
}

func (c *Celeritas) NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	secure, _ := strconv.ParseBool(c.config.cookie.secure)

	//csrfHandler.ExemptGlob("/someapi/*")
	csrfHandler.ExemptGlob("/api/*")

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   secure,
		SameSite: http.SameSiteStrictMode,
		Domain:   c.config.cookie.domain,
	})

	return csrfHandler
}
