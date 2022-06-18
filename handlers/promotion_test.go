package handlers_test

import (
	"errors"
	"fmt"
	"go-unit-test/handlers"
	"go-unit-test/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrage
		amount := 100
		expected := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)
		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		//Act
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)
		fmt.Println(res)

		// Assert

		// assert.Equal(t, fiber.StatusOK, res.StatusCode)
		// body, _ := io.ReadAll(res.Body)
		// assert.Equal(t, strconv.Itoa(90), string(body))
		// ==================================================
		// assert.Equal(t, fiber.StatusOK, res.StatusCode)
		// body, _ := io.ReadAll(res.Body)
		// assert.Equal(t, strconv.Itoa(expected), string(body))
		// ==================================================
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}

	})

	t.Run("error bad request", func(t *testing.T) {
		// Arrage
		amount := "10a"
		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(0, 400)
		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		//Act
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)
		// fmt.Println(fiber.StatusBadRequest,res)

		// Assert
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	t.Run("error not found", func(t *testing.T) {
		// Arrage
		amount := 0
		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(0,  errors.New(""))
		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		//Act
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)
		// body, _ := io.ReadAll(res.Body)
		// fmt.Println(fiber.StatusNotFound, res,string(body))

		// Assert
		assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
	})
}
