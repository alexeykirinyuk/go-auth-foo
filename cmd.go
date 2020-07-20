package main

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/alexeykirinyuk/tech-task-go/data/postgres"
	"github.com/alexeykirinyuk/tech-task-go/infrastructure/auth"
	"github.com/alexeykirinyuk/tech-task-go/routes"
	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	abrenderer "github.com/volatiletech/authboss-renderer"
	"github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"
	"github.com/volatiletech/authboss/v3/defaults"
	_ "github.com/volatiletech/authboss/v3/register"
	"net/http"
	"regexp"
)

const sessionCookieName = "boss_session"

var boss = authboss.New()

func main() {
	db, err := postgres.NewPostgresDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := configureAuth(db); err != nil {
		fmt.Println(err)
		return
	}

	// TODO: add logging
	mux := chi.NewMux()
	mux.Use(boss.LoadClientStateMiddleware)
	mux.Mount("/auth", http.StripPrefix("/auth", boss.Config.Core.Router))

	routes.RegisterAllRoutes(boss, mux, db)

	err = http.ListenAndServe("localhost:5000", mux)
	if err != nil {
		panic(err)
	}
}


func configureAuth(db data.IDatabase) error {
	boss.Config.Storage.Server = auth.NewServerStore(db.GetUserStorage())

	// TODO: use vault or something similar for store secure keys
	var cookieStoreSecureKey = []byte("9a1b7cfe-4f27-4f68-9932-a45ff9daf9a8")
	var sessionStoreBlockKey = []byte("462212d9-87a4-4224-bffb-fcbb677d1ff9")
	boss.Config.Storage.CookieState = abclientstate.NewCookieStorer(cookieStoreSecureKey, nil)

	sessionStore := abclientstate.NewSessionStorer(sessionCookieName, sessionStoreBlockKey, nil)
	cstore := sessionStore.Store.(*sessions.CookieStore)
	cstore.Options.HttpOnly = false
	cstore.Options.Secure = false
	boss.Config.Storage.SessionState = sessionStore

	boss.Config.Paths.Mount = "/auth"
	boss.Config.Paths.RootURL = "http://localhost:5000"
	boss.Config.Modules.ResponseOnUnauthed = authboss.RespondRedirect

	boss.Config.Core.ViewRenderer = abrenderer.NewHTML("/auth", "views")
	boss.Config.Core.MailRenderer = abrenderer.NewHTML("/auth", "views")
	boss.Config.Modules.RegisterPreserveFields = []string{"email", "name"}

	defaults.SetCore(&boss.Config, true, false)

	configureBodyReader()

	if err := boss.Init(); err != nil {
		return err
	}

	return nil
}

func configureBodyReader() {
	emailRule := defaults.Rules{
		FieldName: "email", Required: true,
		MatchError: "Must be a valid e-mail address",
		MustMatch:  regexp.MustCompile(`.*@.*\.[a-z]+`),
	}
	passwordRule := createRequiredRule("password", 8)
	boss.Config.Core.BodyReader = defaults.HTTPBodyReader{
		ReadJSON: false,
		Rulesets: map[string][]defaults.Rules{
			"register": {
				emailRule,
				passwordRule,
				createRequiredRule("first_name", 1),
				createRequiredRule("last_name", 1),
			},
			"recover_end": {passwordRule},
		},
		Confirms: map[string][]string{
			"register":    {"password", authboss.ConfirmPrefix + "password"},
			"recover_end": {"password", authboss.ConfirmPrefix + "password"},
		},
		Whitelist: map[string][]string{
			"register": {"email", "first_name", "last_name", "password"},
		},
	}
}

func createRequiredRule(field string, minLength int) defaults.Rules {
	return defaults.Rules{
		FieldName: "last_name",
		Required:  true,
		MinLength: minLength,
	}
}