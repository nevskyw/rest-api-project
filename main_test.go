package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	firstContext  *gin.Context
	secondContext *gin.Context
	errorContext  *gin.Context
)

func Test_getAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	firstContext, _ = gin.CreateTestContext(w)
	firstContext.Params = []gin.Param{{Key: "id", Value: "1"}}
	secondContext, _ = gin.CreateTestContext(w)
	secondContext.Params = gin.Params{{Key: "id", Value: "2"}}
	errorContext, _ = gin.CreateTestContext(w)
	errorContext.Params = nil

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "first",
			args: args{
				c: firstContext,
			},
		},
		{
			name: "first",
			args: args{
				c: secondContext,
			},
		},
		{
			name: "first",
			args: args{
				c: errorContext,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getAlbums(tt.args.c)
		})
	}
}

func Test_getAlbumByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	firstContext, _ = gin.CreateTestContext(w)
	firstContext.Params = []gin.Param{{Key: "id", Value: "1"}}
	secondContext, _ = gin.CreateTestContext(w)
	secondContext.Params = gin.Params{{Key: "id", Value: "2"}}
	errorContext, _ = gin.CreateTestContext(w)
	errorContext.Params = nil

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "first",
			args: args{
				c: firstContext,
			},
		},
		{
			name: "first",
			args: args{
				c: secondContext,
			},
		},
		{
			name: "first",
			args: args{
				c: errorContext,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getAlbumByID(tt.args.c)
		})
	}
}

func Test_postAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	firstContext, _ = gin.CreateTestContext(w)
	firstContext.Params = []gin.Param{{Key: "id", Value: "1"}}
	secondContext, _ = gin.CreateTestContext(w)
	secondContext.Params = gin.Params{{Key: "id", Value: "2"}}
	errorContext, _ = gin.CreateTestContext(w)
	errorContext.Params = nil

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "first",
			args: args{
				c: firstContext,
			},
		},
		{
			name: "first",
			args: args{
				c: secondContext,
			},
		},
		{
			name: "first",
			args: args{
				c: errorContext,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postAlbums(tt.args.c)
		})
	}
}

func Test_deleteAlbumID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	firstContext, _ = gin.CreateTestContext(w)
	firstContext.Params = []gin.Param{{Key: "id", Value: "1"}}
	secondContext, _ = gin.CreateTestContext(w)
	secondContext.Params = gin.Params{{Key: "id", Value: "2"}}
	errorContext, _ = gin.CreateTestContext(w)
	errorContext.Params = nil

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "first",
			args: args{
				c: firstContext,
			},
		},
		{
			name: "first",
			args: args{
				c: secondContext,
			},
		},
		{
			name: "first",
			args: args{
				c: errorContext,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteAlbumID(tt.args.c)
		})
	}
}

func Test_updateAlbumID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	firstContext, _ = gin.CreateTestContext(w)
	firstContext.Params = []gin.Param{{Key: "id", Value: "1"}}
	secondContext, _ = gin.CreateTestContext(w)
	secondContext.Params = gin.Params{{Key: "id", Value: "2"}}
	errorContext, _ = gin.CreateTestContext(w)
	errorContext.Params = nil

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "first",
			args: args{
				c: firstContext,
			},
		},
		{
			name: "first",
			args: args{
				c: secondContext,
			},
		},
		{
			name: "first",
			args: args{
				c: errorContext,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateAlbumID(tt.args.c)
		})
	}
}
