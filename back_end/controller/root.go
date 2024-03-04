package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/laptop-shop.api/logger"
)

var logr = logger.GetLogger()

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	Msg         string `json:"msg"`
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}

func GetPagination(c *fiber.Ctx) (pageNo, pageSize int) {
	ps := c.Query("page_size")
	pn := c.Query("page")
	pageSize, pageNo = 10, 1

	if len(ps) > 0 {
		psInt, err := strconv.Atoi(ps)
		if err != nil {
			logr.Error(err)
		} else {
			pageSize = psInt
		}
	}

	if len(pn) > 0 {
		pnInt, err := strconv.Atoi(pn)
		if err != nil {
			logr.Error(err)
		} else {
			pageNo = pnInt
		}
	}

	return pageNo, pageSize
}
