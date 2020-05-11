package main

import (
	"github.com/ahmetask/go-con"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type CustomCache struct {
	AppName   string      `json:"appName"`
	Namespace string      `json:"namespace"`
	Port      string      `json:"port"`
	Config    interface{} `json:"config"`
}

func main() {
	e := echo.New()

	// Go con server address
	serviceName := "go-con-server:8080"

	e.HEAD("/config", func(c echo.Context) error {
		err := gocon.RefreshConfigs(serviceName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		} else {
			return c.JSON(http.StatusOK, nil)
		}
	})

	e.HEAD("/config/:appName", func(c echo.Context) error {
		appName := c.Param("appName")
		err := gocon.RefreshConfig(serviceName, appName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		} else {
			return c.JSON(http.StatusOK, nil)
		}
	})

	e.GET("/config/:appName", func(c echo.Context) error {
		appName := c.Param("appName")
		var r gocon.Config
		err := gocon.GetConfig(serviceName, appName, &r)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		} else {
			return c.JSON(http.StatusOK, r)
		}
	})

	e.POST("/config", func(c echo.Context) error {
		ch := new(CustomCache)
		if err := c.Bind(ch); err != nil {
			return err
		}
		request := gocon.AddConfigRequest{AppName: ch.AppName, Content: ch.Config, Port: ch.Port, Namespace: ch.Namespace}
		err := gocon.SaveConfig(serviceName, request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		} else {
			return c.JSON(http.StatusOK, nil)
		}
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.Logger.Fatal(e.Start(":8088"))
}
