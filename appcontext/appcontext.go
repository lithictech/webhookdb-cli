package appcontext

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lithictech/go-aperitif/logctx"
	"github.com/lithictech/webhookdb-cli/client"
	"github.com/lithictech/webhookdb-cli/config"
	"github.com/lithictech/webhookdb-cli/prefs"
	"github.com/sirupsen/logrus"
	"os"
)

type AppContext struct {
	Config      config.Config
	Resty       *resty.Client
	GlobalPrefs *prefs.GlobalPrefs
	Prefs       prefs.Prefs
	Auth        client.Auth
	logger      *logrus.Entry
	//ActiveOrg    string (should be derived from a prefs dir)
}

func (ac AppContext) Logger() *logrus.Entry {
	return ac.logger
}

func New(command string, cfg config.Config) (ac AppContext, err error) {
	ac.Config = cfg
	if ac.GlobalPrefs, err = prefs.Load(); err != nil {
		return
	}
	ac.Prefs = ac.GlobalPrefs.GetNS(cfg.PrefsNamespace)
	ac.Auth = client.Auth{Cookie: ac.Prefs.AuthCookie}
	ac.Resty = newResty(cfg, ac.Prefs)
	if ac.logger, err = logctx.NewLogger(logctx.NewLoggerInput{
		Level:     cfg.LogLevel,
		Format:    cfg.LogFormat,
		File:      cfg.LogFile,
		BuildSha:  config.BuildSha,
		BuildTime: config.BuildTime,
		Fields:    logrus.Fields{"command": command},
	}); err != nil {
		return
	}
	// TODO: For now, always use stderr instead of stdout because we are running this as a CLI,
	// not an application, and want the caller to be able to collect logs easily.
	// But we should make this better configurable, maybe with a change to logctx.
	ac.logger.Logger.SetOutput(os.Stderr)
	return
}

func NewTestContext() AppContext {
	cfg := config.LoadConfig()

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	pr := &prefs.GlobalPrefs{}
	ac := AppContext{
		logger:      logger.WithFields(nil),
		Config:      cfg,
		Resty:       newResty(cfg, pr.GetNS("")),
		GlobalPrefs: pr,
		Prefs:       pr.GetNS(""),
	}
	return ac
}

const ctxKey = "appcontext"

func InContext(parent context.Context, ac AppContext) context.Context {
	return context.WithValue(parent, ctxKey, ac)
}

func FromContext(c context.Context) AppContext {
	return c.Value(ctxKey).(AppContext)
}

func newResty(cfg config.Config, pr prefs.Prefs) *resty.Client {
	r := resty.New().
		SetHostURL(cfg.ApiHost).
		SetHeader(
			"User-Agent",
			fmt.Sprintf("WebhookdbCLI/%s built %s", config.BuildSha, config.BuildTime),
		).SetHeader("Cookie", string(pr.AuthCookie))
	r.Debug = cfg.Debug
	return r
}
