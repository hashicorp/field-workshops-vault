name: chapter-1
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false

# Chapter 1  
## HashiCorp Vault 기본 소개

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???
Chapter 1 introduces Vault

---
name: hashiCorp-vault-overview

# HashiCorp Vault 기본 소개

![:scale 10%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

  * HashiCorp Vault 는 API 기반으로 동작하며, 클라우드 종속성이 없는 민감정보 관리 시스템입니다.
  * 하이브리드 클라우드 환경에서 민감한 데이터를 안전하게 저장하고 관리 할 수 있습니다.
  * 또한 Vault를 사용하여 수명이 짧은 동적 자격 증명을 생성하거나, 즉시 애플리케이션 데이터를 암호화 할 수 있습니다.

???
This is meant as a high level overview.  For detailed descriptions or instructions please see the docs, API guide, or learning site:
* https://www.vaultproject.io/docs/
* https://www.vaultproject.io/api/
* https://learn.hashicorp.com/vault/

---
name: the-old-way
layout: false
# 기존 보안 모델
.center[![:scale 70%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/bodiam_castle.jpg)]
.center[Also known as the "Castle and Moat" method.]

???
* This picture shows the traditional castle and moat security model.

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: traditional-security-models
# 기존 보안 모델
* 기존 보안 모델은 `경계 기반 보안`이라는 아이디어를 기반으로 구축되었습니다.
* 방화벽이있고, 그 방화벽 내부는 안전한 것으로 간주되었습니다.
* 데이터베이스와 같은 리소스는 대부분 정적이었습니다. 이러한 규칙은 IP 주소를 기반으로했기 때문에 자격 증명은 소스 코드와 함께 빌드 되거나 디스크의 정적인 파일에 보관되었습니다.

???
This slide discusses the traditional security model

---
name: problems-with-traditional-security-models
# 기존 보안 모델의 과제
* IP 주소 기반 규칙
* 다음과 같은 문제가 있는 하드코딩 된 자격 증명 정보 :
  * 애플리케이션이나 사용자를 위한 공유 서비스 계정
  * 권한있는 사용자가 민감 점보를 교체, 해제하는 의사결정이 어렵고 수동으로 관리
  * 자격증명을 수거하면 서비스 손상이 발생할 수 있음

???
* This slide describes some of the problems with the traditional security model.
---
name: the-new-way
layout: false
# 최근의 시크릿 관리 환경
.center[![:scale 65%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/nomadic_houses.jpg)]
.center[No well defined perimeter; security enforced by identity.]

???
* 더이상 우리의 서비스는 안전한 '성' 안에 있지 않습니다.
* 언제든지 옮겨갈 수 있고 사라질수도, 더 많아질 수도 있습니다.
* 서비스의 높은 이동성과 확장성을 유지하면서 민감한 데이터의 접근을 통제하려면 어떻게 해야할까요?

---
name: identity-based-security-1
#Identity Based Security
.center[![:scale 75%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/identity-triangle.png)]
.center[[Identity Based Security and Low Trust Networks](https://www.hashicorp.com/identity-based-security-and-low-trust-networks)
]

???
* Here we see that Vault has multiple means of authenticating users and applications with its Auth Methods.
* Vault can manage many types of secrets and excels at generating short-lived, dynmamic secrets.
* Vault's ACL policies are associated with tokens that users and applications use to access secrets after authenticating.
* Tokens can only read/write secrets that its policies allow.
* Click on the link to read a white paper about identity-based security in low trust networks.

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: identity-based-security-2
# 인증 기반 보안 환경

Vault는 Cloud-Native 기반 애플리케이션의 보안 요구 사항을 해결하도록 설계되었습니다. 다음과 같이 기존 접근방식과 차이가 있습니다.:

* 보안이 네트워크 경계에 걸쳐 확장되도록 하는 인증 기반 규칙
* 자주, 그리고 동적으로 교체되는 단기 자격 증명
* 추적하기 위한 개별 계정과 엔티티
* 쉽게 무효화 시킬 수 있는 자격 증명과 엔티티 구조

???
* This slide discusses how Vault is designed for modern applications.

---
name: secrets-engines
layout: false
# Vault Secrets Engines
.center[![:scale 60%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/vault-engines.png)]
.center[[Vault Secrets Engines](https://www.vaultproject.io/docs/secrets/)]

???
* Vault provides many out-of-the-box secrets engines.
* Additional custom secrets engines can be added by customers.
* Click on the link to learn more about Vault secrets engines.

---
name: vault-reference-architecture-1
# Vault의 내부 아키텍쳐
.center[![:scale 75%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/vault_arch.png)]
.center[[HashiCorp Vault Internals Architecture](https://www.vaultproject.io/docs/internals/architecture/)
]

???
* Click the link to learn more about the internal's of Vault's architecture.

---
name: vault-reference-architecture-2
# Vault의 고가용성(HA)를 위한 아키텍처
.center[![:scale 60%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/vault-ref-arch-lb.png)]
.center[[Vault High Availability](https://www.vaultproject.io/docs/concepts/ha/)
]

???
* Vault allows multiple servers to be combined in a highly available cluster within a single cloud region or physical data center.
* Click on the link to learn more about Vault's high availability in a single cluster.

---
name: vault-reference-architecture-3
# Vault의 다중 지역을 지원하는 아키텍처
.center[![:scale 70%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/vault-ref-arch-replication.png)]
.center[[Vault Enterprise Replication](https://www.vaultproject.io/docs/enterprise/replication/)
]

???
* Vault Enterprise supports replication between clusters across regions and data centers.
* It supports Disaster Recovery and Performance replication.
* These can be used together.
* Click the link to learn more about Vault's replication.

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: chapter-1-review-question
# 📝 Chapter 1 Review

* HashiCorp Vault가 무엇인가

???
* Let's review what we learned in this chapter.
---
name: chapter-1-review-answer
# 📝 Chapter 1 Review
* HashiCorp Vault가 무엇인가
  * Vault는 민감 데이터 관리 시스템 입니다.
  * API 기반이며 클라우드에 구애받지 않습니다.
  * 신뢰할 수 없는 네트워크에서 사용할 수 있습니다. (Zero-Trust)
  * 많고 다중의 시스템에 대해 사용자와 애플리케이션 인증 할 수 있습니다.
  * 단기적으로 사용하는 동적인 시크릿 생성을 지원합니다.
  * 지역간에 복제 할 수 있는 고가용성 클러스터를 지원합니다.

???
* Here are the answers to the review questions.
