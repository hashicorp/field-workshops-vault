name: Chapter-3
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false

# Chapter 3      
## 운영을 위한 Vault 서버 실행

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???

Chapter 3 focuses on running a production Vault server

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: vault-production-serves
# 운영을 위한 Vault 서버 실행
* 운영을 위해 Vault 서버를 실행하면 몇가지 단계가 필요합니다. :
  * 구성 파일로 설정 대상 파일을 지정
  * 서버 시작
  * 서버를 초기화
  * 초기화 시 Unseal(봉인 해제) 키와 초기 root 토큰 발급
  * Unseal 키로 Vault 서버의 봉인을 해제

???
* Describe the steps to run a production Vault server.

---
name: configuring-vault
# Vault 서버 설정
* Vault 구성 파일은 [HCL](https://github.com/hashicorp/hcl) 또는 JSON 형식으로 지정합니다.
* 일반적인 구성 설정 대상은 다음과 같습니다. :
  * listener
  * storage
  * seal
  * log_level
  * ui
  * api_addr
  * cluster_addr

???
* Discuss Vault configuration files and common settings.

---
name: running-vault
# 운영을 위한 Vault 서버 실행
* `vault server` 명령을 사용하여 운영모드의 Vault 서버를 실행합니다.
*  `-dev` 옵션을 사용하지 않습니다.

???
* Describe the command to run a Vault production server.

---
name: initializing-vault
# Vault 클러스터 초기화
* Vault 클러스터는 HA를 위해 구성되며, 여러개의 Vault 서버 집합입니다.
* 각 클러스터는 한번 초기화 해야합니다.
* 초기화는 `vault operator init` 명령으로 실행합니다.
* 공유 키 수는 `-key-shares` 옵션으로, 조합의 임계 값은 `-key-threshold` 옵션으로 지정 가능합니다.
* 초기화 명령은 Unseal 키와 초기 root 토큰을 반환합니다.

???
* Describe Vault's `init` command

---
name: unsealing-vault
# Vault Unseal
* 각 Vault 서버는 시작할 때마다 봉인을 해제해야합니다.
* Unseal 전까지는 서버를 사용할 수 없습니다.
* 이 작업은 클러스터 초기화시 반환 된 Unseal 키를 사용하여`vault operator unseal` 명령으로 수행됩니다.

???
* Describe Vault's `unseal` command.

---
name: vault-status-command
# Vault 서버 상태 확인
* `vault status`명령을 사용하여 Vault 서버의 상태를 가져옵니다.
* Vault 서버가 봉인되었는지 아니면 봉인되지 않았는지 알려줍니다.
* 또한 다음 사항을 알려줍니다. :
   * 키 공유 수 및 키 임계 값
   * HA 모드 (클러스터링) 활성화 여부
   * 서버가 Performance Standby로 실행 중인지 여부.

???
Describe the `vault status` command

---
name: chapter-3-review-questions
# 📝 Chapter 3 Review

* "Prod"모드 Vault 서버를 구성하는 데 사용되는 것은 무엇입니까?
* 새 Vault 클러스터에 대해 어떤 Vault 명령을 한 번 실행해야합니까?
* Vault 서버가 시작될 때마다 어떤 Vault 명령을 실행해야합니까?

???
* Let's review what we learned in this chapter.

---
name: chapter-3-review-answers
# 📝 Chapter 3 Review

* "Prod"모드 Vault 서버를 구성하는 데 사용되는 것은 무엇입니까?
  * 설정 파일
* 새 Vault 클러스터에 대해 어떤 Vault 명령을 한 번 실행해야합니까?
  * `vault operator init`
* Vault 서버가 시작될 때마다 어떤 Vault 명령을 실행해야합니까?
  * `vault operator unseal`

???
* Here are the answers to the review questions.
