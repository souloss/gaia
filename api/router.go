package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	v1 "github.com/witchc/gaia/api/v1"
)

const DEFAULT_API_VERSION string = "v1"

func StartService() error {
	router := gin.Default()
	v1.SetRouterGroup(router)

	endPoint := fmt.Sprintf(":%d", 8085)
	rTimeout := time.Duration(30) * time.Second
	wTimeout := time.Duration(30) * time.Second

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    rTimeout,
		WriteTimeout:   wTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
