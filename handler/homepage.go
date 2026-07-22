package handler

import (
	"cv-website/utils"
	"cv-website/view/homepage"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

type HomePageHandler struct{}

func (h *HomePageHandler) Show(c fiber.Ctx) error {
	data := utils.CVData
	data.Year = strconv.Itoa(time.Now().Year())

	c.Set("Content-Type", "text/html; charset=utf-8")
	return homepage.Show(data).Render(c.Context(), c.Response().BodyWriter())
}
