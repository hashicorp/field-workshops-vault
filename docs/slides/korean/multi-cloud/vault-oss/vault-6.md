name: chapter-6
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false

# Chapter 6      
## Vault 정책

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???

* Chapter 6 introduces Vault Policies

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: vault-policies
# Vault 정책
* Vault 정책은 사용자나 애플리케이션이 액세스 할 수있는 시크릿을 제한합니다.
* Vault는 기본적으로 액세스를 *거부* 하는 최소 권한 부여 원칙을 따릅니다.
* Vault 관리자는 정책 설명을 사용하여 사용자 및 애플리케이션에 특정 경로에 대한 액세스 권한을 명시 적으로 부여해야합니다.
* 정책은 경로를 지정하는 것 외에도 해당 경로에 대한 기능 정의를 지정합니다.
* 정책은 HCL (HashiCorp 구성 언어)로 작성됩니다.

---
name: vault-policy-example
# Vault 정책 예제
* 다음은 Vault 정책의 예입니다. :
```hcl
# Allow tokens to look up their own properties
path "auth/token/lookup-self" {
    capabilities = ["read"]
}
```
* 이 정책은 토큰이 자신의 속성을 변경하는 것을 허용하지 않습니다.
???
* This policy allows tokens to look up their own properties

---
name: vault-policy-paths-capabilities
# 정책 경로와 기능
* 정책 경로는 Vault API 경로에 매핑됩니다.
* 부여되는 가장 일반적인 기능은 POST 및 GET과 같은 HTTP 동사에 해당하는`create`,`read`,`update`,`delete`, `list`입니다.
* 다른 두 기능은 HTTP 동사에 해당하지 않습니다.
   *  `sudo`는 루트로 보호 된 경로에 대한 액세스를 허용합니다.
   *  `deny`는 경로에 대한 액세스를 거부하고 다른 기능보다 우선합니다.

???
* Explain policy paths and capabilities

---
name: policies-for-lobs
# LOB에 대한 정책 구성
* Line of business(LOB)는 기업용 비즈니스 애플리케이션의 통칭입니다.
* 많은 조직에서 LOB 및 부서별로 Vault 비밀을 구성합니다.
* 다음은 LOB A, 부서 1에 대한 예제 정책입니다.

```hcl
path "lob_a/dept_1/*" {
    capabilities = ["read", "list", "create", "delete", "update"]
}
```

* 이 정책은 glob 문자 (`*`)를 사용하여`lob_a/dept_1/`아래에 마운트 된 모든 보안 비밀에 모든 표준 기능을 부여합니다.

???
* Talk about how many organizations organize Vault secrets by line of business and department.
* Explain the policy including the glob character and that it can only be used at the end of a path.

---
name: vault-policy-commands
# Vault 정책에 대한 CLI 명령어
* Vault 정책은 Vault의 CLI, UI 또는 API를 사용하여 Vault 서버에 추가 할 수 있습니다.
* CLI로 정책을 추가하는 명령은`vault policy write`입니다.
* 다음은 HCL 파일 "lob-A-dept-1-policy.hcl"에서 "lob-A-dept-1"이라는 정책을 생성하는 명령입니다. <br>
   `vault policy write lob-A-dept-1 lob-A-dept-1-policy.hcl`
* 다음은이 정책을 Userpass 사용자와 연결하는 명령입니다. <br>
   `vault write auth/userpass/users/joe/policies policies=lob-A-dept-1`

???
Describe the most important Vault CLI commands for policies.

---
name: chapter-6-review-questions
# 📝 Chapter 6 Review
* Vault는 기본적으로 시크릿에 대한 액세스 권한을 부여하나요?
* HTTP 동사에 해당하는 정책 기능은 무엇입니까?
* Vault에 정책을 추가하는 데 사용할 수있는 CLI 명령은 무엇입니까?

???
* Let's review what we learned in this chapter.

---
name: chapter-6-review-answers
# 📝 Chapter 6 Review

* Vault는 기본적으로 시크릿에 대한 액세스 권한을 부여하나요?
  * 아니오.
* HTTP 동사에 해당하는 정책 기능은 무엇입니까?
  * `create`, `read`, `update`, `delete`, `list`
* Vault에 정책을 추가하는 데 사용할 수있는 CLI 명령은 무엇입니까?
  * `vault policy write`

???
* Here are the answers to the review questions.
