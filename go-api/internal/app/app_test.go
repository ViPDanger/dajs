package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

var users = [...]dto.UserDTO{
	{Username: "alice", Password: "password123"},
	{Username: "bob", Password: "qwerty456"},
	{Username: "charlie", Password: "letmein789"},
	{Username: "diana", Password: "passw0rd"},
	{Username: "edward", Password: "secure!456"},
	{Username: "fiona", Password: "helloWorld1"},
	{Username: "george", Password: "abcDEF123"},
	{Username: "hannah", Password: "sunshine99"},
	{Username: "ian", Password: "hunter2"},
	{Username: "julia", Password: "admin@123"},
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

// ======================CHARACTER TEST=======================

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
			_ = mt.DB.Drop(ctx)
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
		// POST
		head.Add("Authorization", tokens.Access.Str)
		resp, err = DoRequest("POST", "http://"+host+"/char/", head, characters[0])
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
		var char dto.CharacterDTO
		err = json.Unmarshal(body, &char)
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		// PUT
		char.Name = "Zatalora"
		resp, err = DoRequest("PUT", "http://"+host+"/char/", head, char)
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		// ADD
		resp, err = DoRequest("POST", "http://"+host+"/char/", head, characters[0])
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusCreated, string(body))

		//	Get All
		resp, err = DoRequest("GET", "http://"+host+"/char/all", head, "")
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		fmt.Println(string(body))
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		// GET BY CREATOR ID
		resp, err = DoRequest("GET", "http://"+host+"/char/", head, "")
		if err != nil {
			assert.Fail(mt, err.Error())
			return
		}
		body, _ = io.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, http.StatusOK, string(body))
		fmt.Println(string(body))
	})

}

var characters = [...]dto.CharacterDTO{
	{
		ID:        "",
		Name:      "Thalor",
		Alignment: "Chaotic Good",
		Status:    dto.CharacterStatusDTO{HP: 35, MaxHP: 35, TemporaryHP: 0, ArmorClass: 17, Speed: 30, Initiative: 3},
		Attributes: dto.AttributesDTO{
			Strength:     16,
			Dexterity:    12,
			Constitution: 14,
			Intelligence: 10,
			Wisdom:       13,
			Charisma:     15,
			SkillsDTO: dto.SkillsDTO{
				Athletics:    dto.SkillDTO{Proficient: true, Modifier: 5},
				Persuasion:   dto.SkillDTO{Proficient: true, Modifier: 4},
				Insight:      dto.SkillDTO{Proficient: false, Modifier: 1},
				Perception:   dto.SkillDTO{Proficient: true, Modifier: 3},
				History:      dto.SkillDTO{Proficient: false, Modifier: 0},
				Intimidation: dto.SkillDTO{Proficient: true, Modifier: 4},
			},
		},
		Abilities: []dto.AbilityDTO{
			{ID: "ab1", Name: "Divine Smite", Description: "Radiant burst damage", LevelGained: 2},
		},
		Spells: []dto.SpellDTO{
			{ID: "sp1", Name: "Bless", Description: "Add d4 to attacks/saves", Level: 1},
			{ID: "sp2", Name: "Cure Wounds", Description: "Restore HP", Level: 1},
		},
		Tags: []string{"tank", "support"},
	},
	{
		ID:        "",
		Name:      "Sylra Moonshade",
		Alignment: "Neutral Good",
		Status:    dto.CharacterStatusDTO{HP: 27, MaxHP: 27, TemporaryHP: 0, ArmorClass: 14, Speed: 35, Initiative: 5},
		Attributes: dto.AttributesDTO{
			Strength:     10,
			Dexterity:    18,
			Constitution: 12,
			Intelligence: 14,
			Wisdom:       15,
			Charisma:     11,
			SkillsDTO: dto.SkillsDTO{
				Stealth:        dto.SkillDTO{Proficient: true, Modifier: 6},
				Nature:         dto.SkillDTO{Proficient: true, Modifier: 4},
				Survival:       dto.SkillDTO{Proficient: true, Modifier: 5},
				Perception:     dto.SkillDTO{Proficient: true, Modifier: 5},
				AnimalHandling: dto.SkillDTO{Proficient: false, Modifier: 2},
			},
		},
		Abilities: []dto.AbilityDTO{
			{ID: "ab2", Name: "Favored Enemy", Description: "Bonus vs chosen creature type", LevelGained: 1},
		},
		Spells: []dto.SpellDTO{
			{ID: "sp3", Name: "Hunter's Mark", Description: "Mark enemy for extra damage", Level: 1},
		},
		Tags: []string{"scout", "ranged"},
	},
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
