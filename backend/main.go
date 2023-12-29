package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"painh.com/echo-vue-google-auth/config"
	"painh.com/echo-vue-google-auth/utils"
	"text/template"
	"time"
)

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func createSessionKey(email string) string {
	currentTime := time.Now().String()
	input := email + currentTime
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func createEchoInstance() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Renderer = renderer

	return e
}

func configureCORS(e *echo.Echo, cfg map[string]interface{}) {
	domains := cfg["CORSDomains"].([]interface{})
	domainsStringArray := make([]string, len(domains))
	for i, domain := range domains {
		domainsStringArray[i] = domain.(string)
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: domainsStringArray,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}

func configureSession(e *echo.Echo, cfg map[string]interface{}) {
	secretPassword := cfg["secretPassword"].(string)
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secretPassword))))
}

func configureRoutes(e *echo.Echo) {
	e.POST("/login", func(c echo.Context) error {
		var request map[string]interface{}
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
		}

		code, ok := request["code"].(string)
		if !ok || code == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "code is missing")
		}

		ret, err := utils.GetGoogleOauthToken(code)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		user, err := utils.GetGoogleUser(ret.Access_token, ret.Id_token)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		sessionKey := createSessionKey(user.Email)
		utils.SessionRedisSet(sessionKey, user)

		// sessionKey를 추가해서 반환
		return c.JSON(http.StatusOK, map[string]interface{}{
			"sessionKey": sessionKey,
			"user":       user,
		})
	})

	// 인증이 필요한 라우터 경로 그룹 생성
	g := e.Group("/auth")

	// 미들웨어를 사용하여 인증 확인
	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get("session", c)
			if sess.Values["user"] == nil {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	})

	// 인증이 필요한 라우터 경로 그룹에 핸들러 추가
	g.GET("/echo", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo, World!")
	})
}

func main() {
	e := createEchoInstance()

	cfg := config.ReadConfig()

	configureCORS(e, cfg)
	configureSession(e, cfg)
	configureRoutes(e)

	e.Static("/", "static")

	if cfg["TLS"].(bool) == false {
		e.Logger.Fatal(e.Start(cfg["Address"].(string)))
	} else {
		e.Logger.Fatal(e.StartTLS(cfg["Address"].(string), cfg["TLSCertFile"].(string), cfg["TLSKeyFile"].(string)))
	}
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
