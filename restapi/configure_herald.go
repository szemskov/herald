// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"gopkg.in/yaml.v2"

	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/zerolog/log"
	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/szemskov/herald/models"
	"github.com/szemskov/herald/restapi/operations"
	"github.com/szemskov/herald/restapi/operations/chat"
	"github.com/szemskov/herald/restapi/operations/message"
)

const (
	OK    = "Success"
	ERROR = "Error"
	CONFIG = "Config"
	CONFIGPATH = "./config.yaml"
)

type BotConfig struct {
	ChatId int64 `yaml:"chat"`
	Token  string  `yaml:"token"`
}

//go:generate swagger generate server --target ../../herald --name Herald --spec ../swagger.yaml

func configureFlags(api *operations.HeraldAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HeraldAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.MessageCreateHandler = message.CreateHandlerFunc(func(params message.CreateParams) middleware.Responder {
		config := map[string]BotConfig{}
		ok := true

		ctx := params.HTTPRequest.Context()

		if config, ok = ctx.Value(CONFIG).(map[string]BotConfig); !ok {
			log.Printf("Bot config %s at context init required", CONFIG)
			return message.NewCreateDefault(500)
		}


		if _, ok = config[params.ChatName]; !ok {
			log.Printf("Unexpected chat name %s", params.ChatName)
			return message.NewCreateDefault(404)
		}

		token := config[params.ChatName].Token
		chatId := config[params.ChatName].ChatId

		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Printf("Unauthorized on account %s with token %s", params.ChatName, token)
			return message.NewCreateDefault(500)
		}
		log.Printf("Authorized on account as %s", bot.Self.UserName)

		msg := tgbotapi.NewMessage(chatId, *params.Message.Body)
		if params.Message.Kind != nil {
			msg.ParseMode = *params.Message.Kind
		}

		response, err := bot.Send(msg)
		if err != nil {
			log.Printf("Send message to chat %d problem %s", chatId, err)
			return message.NewCreateDefault(500)
		}

		log.Printf("Send message [%s] to %s", msg.Text, params.ChatName)
		params.Message.ID = int64(response.MessageID)

		return message.NewCreateCreated().WithPayload(params.Message)
	})

	api.ChatListHandler = chat.ListHandlerFunc(func(params chat.ListParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()

		payload := make([]*models.Chat, 0, 10)
		for name, config := range ctx.Value(CONFIG).(map[string]BotConfig) {
			payload = append(payload, &models.Chat{
				ID: config.ChatId,
				Name: name,
				Token: config.Token,
			})
		}

		return chat.NewListOK().WithPayload(payload)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler //addLogging(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return addConfig(handler)
}

func addLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := map[string]interface{}{
			"hostname":  r.Host,
			"url":       r.URL.String(),
			"method":    r.Method,
			"client-ip": r.Header.Get("X-Forwarded-For"),
		}

		msg := ""

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Info().Interface("context", ctx).Interface("request", string(body)).Msg(msg)
		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, r)

		msg = OK
		if rec.Code != 200 {
			msg = ERROR
		}
		log.Info().Interface("context", ctx).Interface("response", rec.Body.String()).Msg(msg)

		_, err = rec.Body.WriteTo(w)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
	})
}

func addConfig(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		botConfigMap := map[string]BotConfig{}

		yamlFile, err := ioutil.ReadFile(CONFIGPATH)
		if err != nil {
			log.Printf("read config.yaml error: #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, botConfigMap)
		if err != nil {
			log.Printf("unmarshal config.yaml error: %v", err)
		}
		log.Printf("Bot config %#v", botConfigMap)

		ctx = context.WithValue(ctx, CONFIG, botConfigMap)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}