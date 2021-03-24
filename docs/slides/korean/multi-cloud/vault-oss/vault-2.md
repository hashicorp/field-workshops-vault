name: Chapter-2
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false



# Chapter 2      
## Vault와 상호작용

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???
Chapter 2 focuses on interacting with Vault

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: Interacting-With-Vault
# Vault와 상호작용

Vault를 사용하기위해 다음과 같은 여러 형태의 메커니즘을 지원 합니다. :
* The Vault [CLI](https://www.vaultproject.io/docs/commands/index.html)
* The Vault [UI](https://learn.hashicorp.com/vault/getting-started/ui)
* The Vault [API](https://www.vaultproject.io/api-docs/index/)

???

* Chapter 2 focuses on interacting with Vault

---
name: Vault-CLI
# Vault CLI
* Vault CLI는 Go 로 작성되었습니다.
* macOS, Windows, Linux 등 여러 OS 환경을 지원합니다.
* 언제든지 최신 버전을 확인하고 다운받을 수 있습니다. (https://www.vaultproject.io/downloads/)

???
* The Vault CLI is distributed as a Go binary.
* It runs on multiple operating systems.

---
name: installing-Vault-CLI
# Vault CLI 설치
* 여러분의 환경에 Vault CLI 설치 방법은 매우 간단합니다. :
  * Zip 파일을 다운 받습니다.
  * 압축을 풀어 `vault` 바이너리를 확인합니다.
  * 어떤 위치에서든 사용하려면 `Path`에 등록합니다.

[tutorial](https://learn.hashicorp.com/vault/getting-started/install) 에서 좀더 자세한 설명을 확인할 수 있습니다.

???
Installing Vault is easy.

---
name: some-cli-commands
# 기본적인 Vault CLI 명령어
* `vault` 는 자체적으로 CLI 명령 목록을 제공합니다.
* `vault version` 은 실행중인 Vault 버전을 알려줍니다.
* `vault read` 는 시크릿 정보를 읽어오는데 사용됩니다.
* `vault write` 는 시크릿 정보를 쓰거나 설정하는데 사용됩니다.

 `-h`, `-help`,  `--help` 플래그를 추가하여 Vault CLI 명령에 대한 도움말 확인이 가능합니다.

???
Let's discuss some of the basic Vault CLI commands.

---
name: vault-server-modes
# Vault Server 모드
Vault 서버는 다음 두가지 형태로 실행할 수 있습니다. :
* "Dev" 모드는 개발 전용입니다.
* "Prod" 모드는 운영 및 QA에서 사용하는 모드입니다.

???
* Discuss Vault's two server modes

---
name: vault-dev-server
# Vault "Dev" Mode
* 안전하지 않습니다.
* 모든것은 메모리에 저장됩니다.(휘발성)
* 자동으로 Unseal 됩니다.
* 실행 전에 루트 토큰을 지정할 수 있습니다.

**"Dev" 모드로 실행되는 환경에 실제 시크릿 정보를 저장하지 마세요.**

???
* Discuss limitations of Vault's "Dev" mode.
* Warn students to never store real secrets on a Dev server.

---
name: Vault-UI
# Vault UI
* Vault UI를 사용하려면, 일단 로그인해야 합니다.
* Vault는 로그인을 위한 여러가지 인증방식을 지원합니다.
* 새로 구성된 Vault는 token 기반 인증방식만 활성화되어있습니다.

???

* Let's talk about the Vault UI a bit, including ways of signing into it.
* While you used the token "root" in the last challenge, you'll be running a Vault server in "Prod"  mode in the rest of the track and will have to use the token generated when you initialize that server in the next challenge.
---
name: signing-into-the-vault-ui
# Signing into the Vault UI
.center[![:scale 70%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/vault_login_page.png)]

???
* This slide shows a screenshot of the login dialog for the Vault server.

---
name: welcome-to-vault
# "Welcome to Vault" 가이드
.center[![:scale 60%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/take_a_tour.png)]
* 처음 UI에 로그인하면 안내 가이드를 확인할 수 있습니다.

???
* Explain the "Welcome to Vault" tour.
* Explain how to get rid of it.
* Point out that they can restart the tour with the "Restart guide" menu under their user icon in the upper right corner of the UI.

---
name: vault-api-1
# Vault API

* Vault에는 Vault를 구성하고 비밀을 관리하는 데 사용할 수있는 HTTP API가 있습니다.
* JSON 출력 형식을 지정하려면 간단한`curl` 명령과`jq`로 Vault의 상태를 확인할 수 있습니다.

명령어 예 :
```bash
curl http://localhost:8200/v1/sys/health | jq
```
???
* Let's talk about the Vault HTTP API

---
name: vault-api-2
# Vault API

```json
{
  "initialized": true,
  "sealed": false,
  "standby": false,
  "performance_standby": false,
  "replication_performance_mode": "disabled",
  "replication_dr_mode": "disabled",
  "server_time_utc": 1557180149,
  "version": "1.4.3",
  "cluster_name": "vault-cluster-db6f271d",
  "cluster_id": "33e85d7c-63bb-7523-0165-9d1aee722d70"
}
```

???
Here is the output from Vault's sys/health endpoint

---
name: vault-api-3
# Vault API 사용에 대한 인증

* `sys/health`에는 인증이 필요하지 않습니다.
* 하지만 대부분의 API 엔드포인트에는 인증이 필요합니다.
*  `X-Vault-Token` 헤더에 Vault token을 제공하여 인증 받습니다.

???
* Talk about how most Vault HTTP API calls will require authentication with a Vault token.

---
name: chapter-2-review-questions
# 📝 Chapter 2 Review

* Vault와 어떻게 상호 작용할 수 있습니까?
* Vault 명령에 대한 도움말을 얻기 위해 사용할 수있는 옵션은 무엇입니까?
* 두 가지 Vault 서버 모드는 무엇입니까?

???
* Let's review what we learned in this chapter.

---
name: chapter-2-review-answers
# 📝 Chapter 2 Review
* Vault와 어떻게 상호 작용할 수 있습니까?
  * Vault CLI
  * Vault UI
  * Vault API
* Vault 명령에 대한 도움말을 얻기 위해 사용할 수있는 옵션은 무엇입니까?
  * `-h`, `-help`, `--help`
* 두 가지 Vault 서버 모드는 무엇입니까?
  * Dev / Prod

???
* Here are the answers to the review questions.
