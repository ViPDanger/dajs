package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/app"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
	"github.com/ViPDanger/dajs/go-api/pkg/config"
	log "github.com/ViPDanger/dajs/go-api/pkg/logger/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Default values
const cfgPath = "../../cmd/config.ini"

var cfg config.Config
var host string
var mongoUser string
var mongoPassword string
var mongoURI string

func init() {
	cfg = config.NewConfig(cfgPath)
	host = cfg.String("server.host", ":8080")
	mongoURI = "mongodb://" + cfg.String("mongo.ip", ":27017")
	mongoPassword = cfg.String("mongo.password", "password")
	mongoUser = cfg.String("mongo.user", "user")
	os.Setenv("MONGODB_URI", mongoURI)

}

func TestRegistration(t *testing.T) {
	cred := options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}
	_ = mtest.Setup(mtest.NewSetupOptions().SetURI(mongoURI))
	clientOpts := options.Client().
		ApplyURI(mongoURI).
		SetAuth(cred)
	opts := mtest.NewOptions().ClientOptions(clientOpts).ClientType(mtest.Default)
	mtest.New(t, opts).Run("", func(mt *mtest.T) {

		ctx, cancel := context.WithCancel(context.Background())
		_ = mt.DB.Drop(ctx)
		defer func() {
			time.Sleep(500 * time.Millisecond)
			cancel()
			time.Sleep(500 * time.Millisecond)
		}()
		app.Run(ctx, log.Initialization("", ""), app.APIConfig{
			Host:           cfg.String("server.host", ":8080"),
			DB:             mt.DB,
			AuthMiddleware: true,
		})
		tokens := dto.TokensDTO{}
		dataUser, _ := json.Marshal(users[1])
		// register
		resp, err := http.Post("http://"+host+"/register", "application/json", bytes.NewBuffer(dataUser))
		if err != nil || resp == nil {
			assert.Fail(t, err.Error())
			return
		}
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusCreated, string(body))
		// login
		resp, err = http.Post("http://"+host+"/login", "application/json", bytes.NewBuffer(dataUser))
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		//=====	get tokens
		_ = json.Unmarshal(body, &tokens)
		refreshToken, _ := json.Marshal(tokens.Refresh)
		// refresh
		resp, err = http.Post("http://"+host+"/refresh", "application/json", bytes.NewBuffer(refreshToken))
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		//====== protected
		resp, err = http.Get("http://" + host + "/")
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
	})
}

// ======================CHARACTER TEST====================================================

func TestCharacter(t *testing.T) {

	cred := options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}
	_ = mtest.Setup(mtest.NewSetupOptions().SetURI(mongoURI))
	clientOpts := options.Client().
		ApplyURI(mongoURI).
		SetAuth(cred)
	opts := mtest.NewOptions().ClientOptions(clientOpts).ClientType(mtest.Default)
	mtest.New(t, opts).Run("", func(mt *mtest.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			time.Sleep(500 * time.Millisecond)
			cancel()
			time.Sleep(500 * time.Millisecond)
		}()
		app.Run(ctx, log.Initialization("", ""), app.APIConfig{
			Host:           cfg.String("server.host", ":8080"),
			DB:             mt.DB,
			AuthMiddleware: true,
		})

		head := http.Header{}
		// USER LOGIN
		dataUser, _ := json.Marshal(users[1])
		resp, err := http.Post("http://"+host+"/login", "application/json", bytes.NewBuffer(dataUser))
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		tokens := dto.TokensDTO{}

		_ = json.Unmarshal(body, &tokens)
		for _, item := range characters {
			// POST
			head.Add("Authorization", tokens.Access.Str)
			resp, err = DoRequest("POST", "http://"+host+"/char/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusCreated, string(body))
			var payload struct {
				ID string `json:"id"`
			}
			err = json.Unmarshal(body, &payload)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			head.Add("id", payload.ID)

			// GETBYID
			resp, err = DoRequest("GET", "http://"+host+"/char/:id", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			// PUT
			resp, err = DoRequest("PUT", "http://"+host+"/char/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			//	Get All
			resp, err = DoRequest("GET", "http://"+host+"/char/all", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			// GET BY CREATOR ID
			resp, err = DoRequest("GET", "http://"+host+"/char/", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		}
		resp, err = DoRequest("DELETE", "http://"+host+"/char/:id", head, "")
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
	})

}

// ===================== PLAYERCHAR TEST ==================================================
func TestPlayerChar(t *testing.T) {

	cred := options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}
	_ = mtest.Setup(mtest.NewSetupOptions().SetURI(mongoURI))
	clientOpts := options.Client().
		ApplyURI(mongoURI).
		SetAuth(cred)
	opts := mtest.NewOptions().ClientOptions(clientOpts).ClientType(mtest.Default)
	mtest.New(t, opts).Run("", func(mt *mtest.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			time.Sleep(500 * time.Millisecond)
			cancel()
			time.Sleep(500 * time.Millisecond)
		}()
		app.Run(ctx, log.Initialization("", ""), app.APIConfig{
			Host:           cfg.String("server.host", ":8080"),
			DB:             mt.DB,
			AuthMiddleware: true,
		})

		head := http.Header{}
		// USER LOGIN

		dataUser, _ := json.Marshal(users[1])
		resp, err := http.Post("http://"+host+"/login", "application/json", bytes.NewBuffer(dataUser))
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		tokens := dto.TokensDTO{}
		_ = json.Unmarshal(body, &tokens)
		for _, item := range playerChars {
			// POST
			head.Add("Authorization", tokens.Access.Str)
			resp, err = DoRequest("POST", "http://"+host+"/pchar/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusCreated, string(body))
			var payload struct {
				ID string `json:"id"`
			}
			err = json.Unmarshal(body, &payload)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			head.Add("id", payload.ID)

			// GETBYID
			resp, err = DoRequest("GET", "http://"+host+"/pchar/:id", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			// PUT
			resp, err = DoRequest("PUT", "http://"+host+"/pchar/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))

			//	Get All
			resp, err = DoRequest("GET", "http://"+host+"/pchar/all", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			// GET BY CREATOR ID
			resp, err = DoRequest("GET", "http://"+host+"/pchar/", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		}
		resp, err = DoRequest("DELETE", "http://"+host+"/char/:id", head, "")
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
	})
}
func TestNPCChar(t *testing.T) {

	cred := options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}
	_ = mtest.Setup(mtest.NewSetupOptions().SetURI(mongoURI))
	clientOpts := options.Client().
		ApplyURI(mongoURI).
		SetAuth(cred)
	opts := mtest.NewOptions().ClientOptions(clientOpts).ClientType(mtest.Default)
	mtest.New(t, opts).Run("", func(mt *mtest.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			time.Sleep(500 * time.Millisecond)
			cancel()
			time.Sleep(500 * time.Millisecond)
		}()
		app.Run(ctx, log.Initialization("", ""), app.APIConfig{
			Host:           cfg.String("server.host", ":8080"),
			DB:             mt.DB,
			AuthMiddleware: true,
		})

		head := http.Header{}
		// USER LOGIN

		dataUser, _ := json.Marshal(users[1])
		resp, err := http.Post("http://"+host+"/login", "application/json", bytes.NewBuffer(dataUser))
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		tokens := dto.TokensDTO{}
		_ = json.Unmarshal(body, &tokens)
		for _, item := range npcs {
			// POST
			head.Add("Authorization", tokens.Access.Str)
			resp, err = DoRequest("POST", "http://"+host+"/npc/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusCreated, string(body))
			var payload struct {
				ID string `json:"id"`
			}
			err = json.Unmarshal(body, &payload)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			head.Add("id", payload.ID)

			// GETBYID
			resp, err = DoRequest("GET", "http://"+host+"/npc/:id", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			// PUT
			resp, err = DoRequest("PUT", "http://"+host+"/npc/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			//	Get All
			resp, err = DoRequest("GET", "http://"+host+"/npc/all", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			// GET BY CREATOR ID
			resp, err = DoRequest("GET", "http://"+host+"/npc/", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		}
		resp, err = DoRequest("DELETE", "http://"+host+"/char/:id", head, "")
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
	})
}

func TestMonsterChar(t *testing.T) {

	cred := options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}
	_ = mtest.Setup(mtest.NewSetupOptions().SetURI(mongoURI))
	clientOpts := options.Client().
		ApplyURI(mongoURI).
		SetAuth(cred)
	opts := mtest.NewOptions().ClientOptions(clientOpts).ClientType(mtest.Default)
	mtest.New(t, opts).Run("", func(mt *mtest.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			time.Sleep(500 * time.Millisecond)
			cancel()
			time.Sleep(500 * time.Millisecond)
		}()
		app.Run(ctx, log.Initialization("", ""), app.APIConfig{
			Host:           cfg.String("server.host", ":8080"),
			DB:             mt.DB,
			AuthMiddleware: true,
		})

		head := http.Header{}
		// USER LOGIN
		dataUser, _ := json.Marshal(users[1])
		resp, err := http.Post("http://"+host+"/login", "application/json", bytes.NewBuffer(dataUser))
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		tokens := dto.TokensDTO{}
		_ = json.Unmarshal(body, &tokens)
		for _, item := range monsters {
			// POST
			head.Add("Authorization", tokens.Access.Str)
			resp, err = DoRequest("POST", "http://"+host+"/monster/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusCreated, string(body))
			var payload struct {
				ID string `json:"id"`
			}
			err = json.Unmarshal(body, &payload)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			head.Add("id", payload.ID)

			// GETBYID
			resp, err = DoRequest("GET", "http://"+host+"/monster/:id", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			// PUT
			resp, err = DoRequest("PUT", "http://"+host+"/monster/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			//	Get All
			resp, err = DoRequest("GET", "http://"+host+"/monster/all", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			// GET BY CREATOR ID
			resp, err = DoRequest("GET", "http://"+host+"/monster/", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		}
		resp, err = DoRequest("DELETE", "http://"+host+"/char/:id", head, "")
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
	})
}
func TestItem(t *testing.T) {

	cred := options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}
	_ = mtest.Setup(mtest.NewSetupOptions().SetURI(mongoURI))
	clientOpts := options.Client().
		ApplyURI(mongoURI).
		SetAuth(cred)
	opts := mtest.NewOptions().ClientOptions(clientOpts).ClientType(mtest.Default)
	mtest.New(t, opts).Run("", func(mt *mtest.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			time.Sleep(500 * time.Millisecond)
			cancel()
			time.Sleep(500 * time.Millisecond)
		}()
		app.Run(ctx, log.Initialization("", ""), app.APIConfig{
			Host:           cfg.String("server.host", ":8080"),
			DB:             mt.DB,
			AuthMiddleware: true,
		})

		head := http.Header{}
		// USER LOGIN

		dataUser, _ := json.Marshal(users[1])
		resp, err := http.Post("http://"+host+"/login", "application/json", bytes.NewBuffer(dataUser))
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		tokens := dto.TokensDTO{}
		_ = json.Unmarshal(body, &tokens)
		for _, item := range items {
			// POST
			head.Add("Authorization", tokens.Access.Str)
			resp, err = DoRequest("POST", "http://"+host+"/item/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusCreated, string(body))
			var payload struct {
				ID string `json:"id"`
			}
			err = json.Unmarshal(body, &payload)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			head.Add("id", payload.ID)

			// GETBYID
			resp, err = DoRequest("GET", "http://"+host+"/item/:id", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			// PUT
			resp, err = DoRequest("PUT", "http://"+host+"/item/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			//	Get All
			resp, err = DoRequest("GET", "http://"+host+"/item/all", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		}
		resp, err = DoRequest("DELETE", "http://"+host+"/char/:id", head, "")
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
	})

}
func TestGlossary(t *testing.T) {

	cred := options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}
	_ = mtest.Setup(mtest.NewSetupOptions().SetURI(mongoURI))
	clientOpts := options.Client().
		ApplyURI(mongoURI).
		SetAuth(cred)
	opts := mtest.NewOptions().ClientOptions(clientOpts).ClientType(mtest.Default)
	mtest.New(t, opts).Run("", func(mt *mtest.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			time.Sleep(500 * time.Millisecond)
			cancel()
			time.Sleep(500 * time.Millisecond)
		}()
		app.Run(ctx, log.Initialization("", ""), app.APIConfig{
			Host:           cfg.String("server.host", ":8080"),
			DB:             mt.DB,
			AuthMiddleware: true,
		})

		head := http.Header{}
		// USER LOGIN

		dataUser, _ := json.Marshal(users[1])
		resp, err := http.Post("http://"+host+"/login", "application/json", bytes.NewBuffer(dataUser))
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		tokens := dto.TokensDTO{}
		_ = json.Unmarshal(body, &tokens)
		for _, item := range glossarys {
			// POST
			head.Add("Authorization", tokens.Access.Str)
			resp, err = DoRequest("POST", "http://"+host+"/item/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusCreated, string(body))
			var payload struct {
				ID string `json:"id"`
			}
			err = json.Unmarshal(body, &payload)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			head.Add("id", payload.ID)

			// GETBYID
			resp, err = DoRequest("GET", "http://"+host+"/item/:id", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			// PUT
			resp, err = DoRequest("PUT", "http://"+host+"/item/", head, item)
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
			//	Get All
			resp, err = DoRequest("GET", "http://"+host+"/item/all", head, "")
			if err != nil {
				assert.Fail(mt, err.Error())
				return
			}
			body, _ = io.ReadAll(resp.Body)
			assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		}
		resp, err = DoRequest("DELETE", "http://"+host+"/char/:id", head, "")
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
	})

}
func DoRequest(method string, url string, headers http.Header, object any) (*http.Response, error) {
	body, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for header, value := range headers {
		req.Header.Add(header, value[0])
	}
	req.Header = headers
	return http.DefaultClient.Do(req)
}
