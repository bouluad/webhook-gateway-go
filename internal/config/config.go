package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	GitHubSecret string
	JenkinsMap   map[string]string
)

func LoadConfig() {
	GitHubSecret = getEnv("GITHUB_SECRET", "mysecret")
	mapFile := getEnv("JENKINS_MAP_FILE", "jenkins_map.yaml")

	file, err := os.ReadFile(mapFile)
	if err != nil {
		log.Fatalf("❌ Failed to read Jenkins map file: %v", err)
	}

	if err := yaml.Unmarshal(file, &JenkinsMap); err != nil {
		log.Fatalf("❌ Failed to parse Jenkins map: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
