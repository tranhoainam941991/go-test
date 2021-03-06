package main

import (
	"errors"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func markPrime(primes []bool, i int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := i * i; j <= n; j += i {
		primes[j] = false
	}
}

func getPrime(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("cap need to be greater than 0")
	}
	if n < 3 {
		return n, nil
	}

	primes := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		primes[i] = true
	}
	var wg sync.WaitGroup
	for i := 2; i*i <= n; i++ {
		wg.Add(1)
		go markPrime(primes, i, n, &wg)
	}
	wg.Wait()
	var result int
	for i := n; i != 2; i-- {
		if primes[i] {
			result = i
			break
		}
	}
	return result, nil
}

func main() {
	app := fiber.New()
	app.Use(recover.New())
	app.Get("/:cap", func(c *fiber.Ctx) error {
		cap, convertError := strconv.Atoi(c.Params("cap"))
		if convertError != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": convertError.Error(),
			})
		}
		result, error := getPrime(cap)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": error.Error(),
			})
		} else {
			return c.JSON(fiber.Map{
				"result": result,
			})

		}

	})

	app.Listen(":3000")
}
