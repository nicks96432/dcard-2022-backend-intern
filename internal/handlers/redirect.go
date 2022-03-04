package handlers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

import "dcard-2022-backend-intern/internal/base62"

func (h *Handlers) getRoot(c *fiber.Ctx) error {
	ctx := context.TODO()

	// decode param
	encoded := c.Params("url")
	id, err := base62.Decode(encoded)
	if err != nil {
		log.Printf("cannot decode id: %v\n", err)
		return c.SendStatus(404)
	}

	// query database by id
	e, err := h.Db.Query(ctx, int64(id))
	if err != nil {
		log.Printf("database query failed: %v\n", err)
		return c.SendStatus(404)
	}

	// log request
	if j, err := json.MarshalIndent(e, "", "    "); err == nil {
		log.Printf("GET /%s %s", encoded, string(j[:]))
	} else {
		log.Printf("json marshal failed: %v\n", err)
	}

	// check expire time
	if e.ExpireAt.Before(time.Now()) {
		if deleted, err := h.Db.Delete(ctx, e.Id); err != nil {
			log.Printf("cannot delete database entry: %v\n", err)
		} else if !deleted {
			log.Println("database entry not found")
		}
		log.Println("entry expired")
		return c.SendStatus(404)
	}

	// send redirect url
	log.Println("redirect successfully")
	return c.Redirect(e.Url, 301)
}
