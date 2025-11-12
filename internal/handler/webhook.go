package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gateway/internal/config"
	"github.com/gateway/internal/forwarder"
	"github.com/gateway/internal/router"
	"github.com/gateway/internal/security"
)

func HandleWebhook(c *gin.Context) {
	signature := c.GetHeader("X-Hub-Signature-256")
	event := c.GetHeader("X-GitHub-Event")

	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read payload"})
		return
	}

	if !security.VerifySignature(payload, signature, config.GitHubSecret) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
		return
	}

	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	repo := ""
	if repoInfo, ok := data["repository"].(map[string]interface{}); ok {
		if fullName, ok := repoInfo["full_name"].(string); ok {
			repo = fullName
		}
	}

	if repo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repository not found"})
		return
	}

	jenkinsURL := router.GetJenkinsURL(repo)
	if jenkinsURL == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "no Jenkins mapping found"})
		return
	}

	status, err := forwarder.ForwardToJenkins(jenkinsURL, event, payload)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{"status": "forwarded", "jenkins_url": jenkinsURL})
}
