package services

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Upload struct {
	UploadDir string
	BaseURL   string
}

type uploadImagePayload struct {
	Input struct {
		Base64      string `json:"base64"`
		Filename    string `json:"filename"`
		ContentType string `json:"content_type"`
	} `json:"input"`

	SessionVariables map[string]string `json:"session_variables"`
}

var allowedImageTypes = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/webp": ".webp",
}

func (u *Upload) UploadImageAction(c *gin.Context) {

	var payload uploadImagePayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	role := payload.SessionVariables["x-hasura-role"]

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "only admins can upload images",
		})
		return
	}

	ext, ok := allowedImageTypes[payload.Input.ContentType]

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unsupported file type, use JPEG, PNG, or WebP",
		})
		return
	}

	// Strip data URL prefix if the client sent one, e.g. "data:image/png;base64,...."
	raw := payload.Input.Base64

	if idx := strings.Index(raw, ","); strings.HasPrefix(raw, "data:") && idx != -1 {
		raw = raw[idx+1:]
	}

	fileBytes, err := base64.StdEncoding.DecodeString(raw)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid base64 data",
		})
		return
	}

	// 5MB limit
	if len(fileBytes) > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "file too large, max 5MB",
		})
		return
	}

	if err := os.MkdirAll(u.UploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to prepare upload directory",
		})
		return
	}

	uniqueName := fmt.Sprintf("%s-%d%s", uuid.New().String(), time.Now().Unix(), ext)
	fullPath := filepath.Join(u.UploadDir, uniqueName)

	if err := os.WriteFile(fullPath, fileBytes, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save file",
		})
		return
	}

	url := fmt.Sprintf("%s/uploads/%s", u.BaseURL, uniqueName)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Upload successful",
		"url":     url,
	})
}