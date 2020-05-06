package main

import (
	"fmt"
	arithmetic "github.com/eddypastika/go-arithmetic"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	a, b := 20, 20

	m := arithmetic.Multiply(a, b)
	fmt.Println("Multiply result: ", m)

	add := arithmetic.Add(a, b)
	fmt.Println("Add result: ", add)

	s := arithmetic.Subtract(a, b)
	fmt.Println("Subtract result: ", s)

	e := echo.New()
	e.GET("/calculate", arithmeticHandler)
	e.Start(":9090")

}

func arithmeticHandler(c echo.Context) error {
	var res int64

	values := c.QueryParam("values")
	op := c.QueryParam("operator")

	arrValues := strings.Split(values, ",")
	arrValuesInt := []interface{}{}
	for _, v := range arrValues {
		param, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		arrValuesInt = append(arrValuesInt, param)
	}

	switch strings.ToLower(op) {
	case "multiply":
		res = arithmetic.Multiply(arrValuesInt...)
	case "add":
		res = arithmetic.Add(arrValuesInt...)
	case "subtract":
		res = arithmetic.Subtract(arrValuesInt...)
	}

	return c.JSON(http.StatusOK, res)
}
