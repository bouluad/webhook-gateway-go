package router

import (
	"strings"

	"github.com/gateway/internal/config"
)

func GetJenkinsURL(repoName string) string {
	for key, url := range config.JenkinsMap {
		if strings.HasPrefix(repoName, key) {
			return url
		}
	}
	return ""
}
