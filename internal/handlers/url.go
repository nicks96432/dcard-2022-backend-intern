package handlers

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

import (
	"dcard-2022-backend-intern/internal/base62"
	"dcard-2022-backend-intern/internal/storage"
)

type payload struct {
	Url      string    `json:"url"`
	ExpireAt time.Time `json:"expireAt"`
}

func (h *Handlers) postNewUrl(c *fiber.Ctx) error {
	ctx := context.TODO()
	p := payload{}

	// parse request body
	if err := c.BodyParser(&p); err != nil {
		return c.Status(http.StatusBadRequest).SendString("bad request")
	}
	e := storage.Entry{
		Id:       0,
		Url:      p.Url,
		ExpireAt: p.ExpireAt,
	}

	// log request body
	if j, err := json.MarshalIndent(p, "", "    "); err == nil {
		log.Println("POST /api/v1/urls", string(j[:]))
	} else {
		log.Println("json marshal failed: ", err.Error())
	}

	// generate random id
	var max big.Int
	max.SetUint64(^uint64(0))
	for {
		random, err := rand.Int(rand.Reader, &max)
		id := int64(random.Uint64())
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("cannot generate random number: %v\n", err))
		}

		if result, err := h.Db.Query(ctx, id); err == nil {
			if !result.ExpireAt.Before(time.Now()) {
				continue
			} else {
				if deleted, err := h.Db.Delete(ctx, id); err != nil {
					log.Printf("cannot delete database entry: %v\n", err)
				} else if !deleted {
					log.Println("database entry not found")
				}
			}
		}
		e.Id = id
		break
	}

	// add entry to database
	if err := h.Db.Add(ctx, &e); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("connot insert database: %v\n", err))
	}

	// encode id & create response body
	id := base62.Encode(uint64(e.Id))

	m := fiber.Map{"id": id}
	if h.Config.Port == 80 {
		m["shortUrl"] = fmt.Sprintf("http://%s/%s", h.Config.Hostname, id)
	} else {
		m["shortUrl"] = fmt.Sprintf("http://%s:%d/%s", h.Config.Hostname, h.Config.Port, id)
	}
	log.Println("add successfully")
	return c.JSON(m)
}
