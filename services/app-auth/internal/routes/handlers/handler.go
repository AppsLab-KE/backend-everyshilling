// Package handlers provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Generate OTP and send it to email and phone number
	// (POST /login)
	PostLogin(c *gin.Context)
	// A POST request to registering new users
	// (POST /register)
	Register(c *gin.Context)
	// Send password reset OTP
	// (POST /reset)
	PostReset(c *gin.Context)
	// Change Password
	// (POST /reset/{request-id}/change)
	PostResetRequestIdChange(c *gin.Context, requestId string)
	// Verify OTP
	// (POST /reset/{request-id}/verify)
	PostResetRequestIdVerify(c *gin.Context, requestId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostLogin operation middleware
func (siw *ServerInterfaceWrapper) PostLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostLogin(c)
}

// Register operation middleware
func (siw *ServerInterfaceWrapper) Register(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.Register(c)
}

// PostReset operation middleware
func (siw *ServerInterfaceWrapper) PostReset(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostReset(c)
}

// PostResetRequestIdChange operation middleware
func (siw *ServerInterfaceWrapper) PostResetRequestIdChange(c *gin.Context) {

	var err error

	// ------------- Path parameter "request-id" -------------
	var requestId string

	err = runtime.BindStyledParameter("simple", false, "request-id", c.Param("request-id"), &requestId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter request-id: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostResetRequestIdChange(c, requestId)
}

// PostResetRequestIdVerify operation middleware
func (siw *ServerInterfaceWrapper) PostResetRequestIdVerify(c *gin.Context) {

	var err error

	// ------------- Path parameter "request-id" -------------
	var requestId string

	err = runtime.BindStyledParameter("simple", false, "request-id", c.Param("request-id"), &requestId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter request-id: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostResetRequestIdVerify(c, requestId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {

	errorHandler := options.ErrorHandler

	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/login", wrapper.PostLogin)

	router.POST(options.BaseURL+"/register", wrapper.Register)

	router.POST(options.BaseURL+"/reset", wrapper.PostReset)

	router.POST(options.BaseURL+"/reset/:request-id/change", wrapper.PostResetRequestIdChange)

	router.POST(options.BaseURL+"/reset/:request-id/verify", wrapper.PostResetRequestIdVerify)

	return router
}