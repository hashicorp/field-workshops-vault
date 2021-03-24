name: chapter-5
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false

# Chapter 5      
## Vault에 사용자 인증을 위한 방법

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???

* Chapter 5 introduces Vault authentication methods
* It focuses on the Userpass method.

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: vault-auth-methods
# Vault에 사용자 인증을 위한 방법
.center[![:scale 45%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/vault_auth_methods.png)]
.center[Vault supports many different authentication methods.]

???
* 인증 방법은 애플리케이션이나 사용자가 자신의 신원을 확인하는 방법입니다.
  * 비슷한 방법으로 호텔 체크인 데스크에서 유효한 신분증을 제시하는 방식입니다.
* 사용자와 애플리케이션은 인증을 위해 일종의 자격 증명이나 토큰을 제공합니다.
* 여러 인증 방법과 동일한 인증 방법의 여러 인스턴스를 활성화 할 수 있습니다.

---
name:vault-auth-methods-2
# 주요 Vault 인증 방식

<div style="float: left; width: 50%;">
<u>사람 기반의 방식</u>
<ul>
<li>Userpass</li>
<li>GitHub</li>
<li>LDAP</li>
<li>JWT/OIDC</li>
<li>Okta</li>
</ul>
</div>
<div style="float: right; width: 50%;">
<u>애플리케이션 기반 방식</u>
<ul>
<li>AppRole</li>
<li>AWS</li>
<li>Azure</li>
<li>Google Cloud</li>
<li>Kubernetes</li>
</ul>
</div>


???
* Userpass - Allows users to authenticate with username and password managed by Vault
* GitHub - Allows users to authenticate with their GitHub personal access tokens
* LDAP - Allows users to authenticate against an LDAP server with their username and password managed by that server.
* JWT/OIDC - Allows users to authenticate against an external OpenID Connect provider or with JSON Web Tokens (JWTs)
* Okta - Allows users to authenticate using Okta single sign-on.
* AppRole - Allows applications to authenticate in automated workflows using a role and a role ID.
* AWS - Allows applications on AWS EC2 instances and Lambda functions to authenticate with IAM credentials or EC2 metadata.
* Azure - Allows applications associated with Azure Managed Service Identities to authenticate using Azure Active Directory credentials.
* Google Cloud - Allows applications in GCP to authenticate using Google Cloud IAM service accounts or Google Compute Engine (GCE) metadata.
* Kubernetes - Allows Kubernetes pods to authenticate with JWT tokens.

---
name: enabling-auth-methods
# 인증 방식 활성화

* 대부분의 Vault 인증 방법은 명시 적으로 활성화해야합니다.
* `vault auth enable` 명령으로 수행됩니다.
* 각 인증 방법에는 기본 경로가 있으며, 여러 인스턴스를 활성화하기 위해 대체 경로를 지정할 수 있습니다. <br>`vault auth enable -path=aws-east aws`
* 사용자 지정 경로는 CLI 명령 및 API 호출에서 지정해야합니다. <br>
   `vault write aws/config/root` <br>
   대신 <br>
   `vault write aws-east/config/root`

???

* Talk about enabling auth methods.
* Talk about default and custom paths
* Explain the examples

---
name: userpass-0
# Vault 인증 방식 - Userpass
.center[![:scale 30%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/userpass_login.png)]
* Userpass 방법은 Vault에서 관리하는 사용자 이름과 비밀번호로 사용자를 인증합니다.

???
* The Userpass method allows users to authenticate with username and password managed by Vault.
* It is not recommended for production, but it's fine for development and lab environments.
* In the real world you'd probably have Vault use your Active Directory, LDAP, GitHub, or other system of record for authentication by users.

---
name: chapter-5-review-questions
# 📝 Chapter 5 Review
* Vault는 어떤 유형의 엔티티를 인증 할 수 있나요?
* Userpass 인증 방법에 대한 자격 증명을 관리하는 시스템은 무엇입니까?
* 기본 정책 이외의 다른 정책이 할당되지 않은 사용자가 모든 비밀에 액세스 할 수 있습니까?

???
* Let's review what we learned in this chapter.

---
name: chapter-5-review-answers
# 📝 Chapter 5 Review

* Vault는 어떤 유형의 엔티티를 인증 할 수 있나요?
  * 사용자와 애플리케이션
* Userpass 인증 방법에 대한 자격 증명을 관리하는 시스템은 무엇입니까?
  * Vault
* 기본 정책 이외의 다른 정책이 할당되지 않은 사용자가 모든 시크릿에 액세스 할 수 있습니까?
  * 불가

???
* Here are the answers to the review questions.
