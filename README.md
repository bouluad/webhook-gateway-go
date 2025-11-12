# 🚀 GitHub → Jenkins Webhook Gateway

A lightweight **gateway service written in Go** that receives GitHub webhooks, validates them securely, and forwards them to the correct Jenkins instance based on the repository or organization.  
Designed for infrastructures with **multiple Jenkins masters**, acting as a centralized webhook entry point.

---

## 🧩 Features

✅ **Signature validation** — verifies GitHub HMAC signatures using your secret key  
✅ **Dynamic routing** — maps GitHub repositories or orgs to different Jenkins instances  
✅ **Transparent forwarding** — passes webhook payloads directly to Jenkins endpoints  
✅ **Lightweight and fast** — built with [Gin](https://github.com/gin-gonic/gin)  
✅ **Easy deployment** — Docker-ready, minimal configuration  

## 🧠 How It Works

GitHub sends a POST request to /webhook

The gateway verifies the signature using GITHUB_SECRET

It extracts the repository name (repository.full_name)

Looks up the target Jenkins URL in jenkins_map.yaml

Forwards the entire webhook payload (with headers) to the Jenkins endpoint
---

## 🏗️ Architecture

GitHub → [ Gateway (Go) ] → Jenkins #1
↘︎ → Jenkins #2
↘︎ → Jenkins #3

---

## 📁 Project Structure

github-jenkins-gateway/
│
├── cmd/
│ └── server/
│ └── main.go # Entry point (starts HTTP server)
│
├── internal/
│ ├── config/ # Loads env vars and YAML mapping
│ ├── handler/ # Webhook endpoint logic
│ ├── router/ # Routing logic (GitHub repo → Jenkins)
│ ├── security/ # Signature verification (HMAC)
│ └── forwarder/ # Forwards request to Jenkins
│
├── jenkins_map.yaml # Repository → Jenkins URL mapping
├── Dockerfile
├── go.mod
├── go.sum
└── README.md


---

## ⚙️ Configuration

### 1️⃣ Environment variables

| Variable | Description | Default |
|-----------|--------------|----------|
| `GITHUB_SECRET` | GitHub webhook secret (for signature validation) | `mysecret` |
| `JENKINS_MAP_FILE` | Path to YAML file defining repository mappings | `jenkins_map.yaml` |

---

### 2️⃣ Jenkins mapping file (`jenkins_map.yaml`)

Define which GitHub repository or organization maps to which Jenkins webhook endpoint:

```yaml
org1/infra: "https://jenkins.infra.company.com/github-webhook/"
org2/app: "https://jenkins.app.company.com/github-webhook/"
org3/ci-pipeline: "https://jenkins.sg.company.com/github-webhook/"

## 🚀 Run Locally

### 🧠 Prerequisites

Go 1.22+

A valid GitHub webhook secret

Jenkins servers with /github-webhook/ endpoints

###  Run directly:
```yaml
export GITHUB_SECRET="supersecret"
go run ./cmd/server
