name: chapter-4
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false

# Chapter 4      
## Vault 시크릿 엔진

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???

* Chapter 4 introduces Vault secrets engines
* It focuses on the KV v2 engine.

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: vault-secrets-engines-1
# Vault 시크릿 엔진

.center[![:scale 65%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/vault-secrets-engines.png)]
.center[Vault includes many different secrets engines.]

???
* Use this screenshot from the Vault UI to talk about Vault's many secrets engines but note that the next slide lists them too.
* Some are for storing static secrets.
* Others can dynamically generate secrets such as database and cloud credentials.
* There is even one called "Transit" that provides encryption as a service.

---
name:vault-secrets-engines-2
# 주요 Vault 시크릿 엔진
* Key/Value (KV)
* PKI
* SSH
* TOTP
* Databases
* AWS, Azure, and Google
* Active Directory
* Transit

???
Spend some time pointing out what some of these do:
* KV - Used to manage generic, static secrets. KV v2 supports versioning.
* PKI - Used to generate dynamic X.509 certificates
* SSH - Take all the pain and drudgery out of securing your SSH infrastructure. Vault can provide key signing services that make securing SSH a snap.
* TOTP - The TOTP tool allows Vault to either act as a code-generating device for MFA logins or to provide TOTP server capabilities for MFA infrastructure.
* Databases - Generate dynamic, short-lived database credentials.
* Cloud credentials engines - Generate dynamic, short-lived cloud credentials for major clouds.
* Active Directory - Vault can rotate AD passwords.
* Transit - Implement's Vault's encryption-as-a-service. Provides an API that can handle all your encryption and decryption needs, based on policy, so that you don't have to manage a complicated key infrastructure.

---
name: enabling-secrets-engines
# 시크릿 엔진 활성화

* 대부분의 Vault 시크릿 엔진은 명시 적으로 활성화해야합니다.
* `vault secrets enable` 명령으로 수행됩니다.
* 각 시크릿 엔진에는 기본 경로가 있습니다.
* 여러 인스턴스를 활성화하기 위해 대체 경로를 지정할 수 있습니다. <br>`vault secrets enable -path=aws-east aws`
* 사용자 지정 경로는 CLI 명령 및 API 호출시 지정해야합니다. <br>
   `vault write aws/config/root` <br>
   대신 <br>
   `vault write aws-east/config/root`

???

* Talk about enabling secrets engines.
* Talk about default and custom paths
* Explain the examples

---
name: vault-kv-engine
# KV 시크릿 엔진
* Vault의 KV 비밀 엔진에는 두 가지 버전이 있습니다.
  * KV v1 (버전 관리 없음)
  * KV v2 (버전 관리 포함)
* 두 번째 실습 과제에서는 "Dev"모드 Vault 서버에 대해 자동으로 활성화되는 KV v2 엔진의 인스턴스를 사용합니다.
* Vault는 "Prod"모드로 서버 실행 시 KV 비밀 엔진 인스턴스를 활성화하지 않습니다. 따라서 직접 활성화해야합니다.

???
* We already used Vault's Key/Value (KV) engine in the second challenge of the "Vault Basics" Instruqt track that had been automatically enabled for the "Dev" mode server.
* But we'll need to mount it ourselves for the "Prod" mode server.

---
name: vault-kv-commands
# KV 시크릿 엔진 명령어
* 다음 명령어를 사용하여 기본 경로에  `kv` 시크릿 엔진을 활성화 합니다.  :<br>
`vault secrets enable -version=2 kv`
* `vault kv` 명령으로 KV 시크릿 엔진을 사용할 수 있습니다.
  * `vault kv list` - 지정한 경로의 시크릿 목록을 나열합니다.
  * `vault kv put` - 지정한 경로에 시크릿을 기록합니다.
  * `vault kv get` - 지정한 경로의 시크릿을 읽어옵니다.
  * `vault kv delete` - 지정한 경로의 시크릿을 삭제합니다.
* `vault kv`의 하위 명령어는 KV v2에서 동작합니다.

???

* Describe how to mount an instance of the KV v2 secrets engine.
* Describe the various `vault kv` subcommands.

---
name: chapter-4-review-questions
# 📝 Chapter 4 Review

* 여러 인스턴스를 활성화하기 위해`vault secrets enable` 명령에 추가 된 옵션은 무엇입니까?
* KV 비밀 엔진의 두 버전의 차이점은 무엇입니까?
* KV v2 시크릿에 저장되었던 과거 버전을 검색 할 수 있습니까?

???
* Let's review what we learned in this chapter.

---
name: chapter-4-review-answers
# 📝 Chapter 4 Review

* 여러 인스턴스를 활성화하기 위해`vault secrets enable` 명령에 추가 된 옵션은 무엇입니까?
  * `-path=<path>` 옵션을 활용하여 활성화 합니다.
* KV 시크릿 엔진의 두 버전의 차이점은 무엇입니까?
  * KV V2는 시크릿의 버전 관리를 지원합니다.
* KV v2 시크릿에 저장되었던 과거 버전을 검색 할 수 있습니까?
  * 예. 실습과정에서 Vault UI에서 이 작업을 수행했습니다.

???
* Here are the answers to the review questions.
