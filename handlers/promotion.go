package handlers

import (
	"fmt"
	"go-unit-test/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	promoService services.PromotionService
}

func NewPromotionHandler(promoService services.PromotionService) PromotionHandler {
	return promotionHandler{promoService: promoService}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	//http://localhost:8000/calculate?amount=100

	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Println("amount",amount)
	discount, err := h.promoService.CalculateDiscount(amount)
	fmt.Println(err)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	fmt.Println("Discount",discount)

	return c.SendString(strconv.Itoa(discount))
}
