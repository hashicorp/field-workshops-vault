name: chapter-8
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false

# Chapter 8    
## Encryption as a Service

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???

* Chapter 8 introduces Vault's Transit secrets engine which functions as Vault's Encryption-as-a-Service (EaaS).

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: Vault-Transit-Engine

# Vault Transit Engine - Encryption as a Service
.center[![:scale 80%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/vault-eaas.webp)]

* Vault의 Transit Secrets Engine은 Encryption-as-a-Service로 작동합니다.
* 개발자는이를 사용하여 Vault 외부에 저장된 데이터를 암호화하고 해독합니다.

???
* Let's talk about Vault's Encryption-as-a-Service, the Transit secrets engine.
* It provides an encryption API & service that are secure, accessible and easy to implement.
* Instead of forcing developers to learn cryptography, we present them with a familiar API that can be used to encrypt and decrypt data that is stored outside of Vault.

---
name: transit-engine-benefits
# Transit Engine Benefits

* Vault의 Transit Engine은 개발자가 암호화 또는 암호화 전문가가 될 필요가 없도록 잘 설계된 EaaS API를 제공합니다.
* 중앙 집중식 키 관리를 제공합니다.
* 승인 된 암호 및 알고리즘 만 사용되도록합니다.
* 자동 키 순환 및 재 래핑을 지원합니다.
* 공격자가 암호화 된 데이터에 액세스 할 수있는 경우 Vault 없이는 쓸모없는 암호문 만 볼 수 있습니다.

???
* There are seveal benefits of using the Transit engine.

---
name: Vault-Transit-Engine-1
# Vault Transit - Example Application

* 다음 실습에서는 Transit 엔진을 사용하여 데이터를 암호화하고 해독하는 웹 애플리케이션을 사용합니다.
* 이 앱은 7 장에서 사용한 것과 동일한 MySQL 데이터베이스에 암호화 된 데이터를 저장합니다.
* 또한 해당 장의 실습에서 구성한 데이터베이스 시크릿 엔진에서 MySQL 자격 증명을 얻습니다.
* 먼저 Vault없이 웹 앱을 실행합니다. 암호화 된 레코드가 없습니다.
* 그런 다음 Vault를 활성화 한 상태에서 실행하고 새 레코드가 암호화되었는지 확인합니다.

???
* Discuss the web app we will be using in this chapter's lab.
* Point out that it will use the same MySQL database from chapter 7.
* Point out that it will get its MySQL credentials from the Database secrets engine students set up in chapter 7.
* Indicate that we will first run without Vault and then with it.

---
name: web-app-screenshot
# Web App
### 다음은 Python 웹 앱의 스크린 샷입니다.:

.center[![:scale 70%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/transit_app.png)]

???
* Show the screenshot of the web app.

---
name: web-app-views
# Web App - Views
### 응용 프로그램에는 두 가지 주요 섹션이 있습니다.

1. **Records View**
  * 레코드보기는 암호화 된 데이터가 해독 된 후 로그인 한 사용자에게 표시되는 내용을 보여주는 일반 텍스트로 레코드를 표시합니다.

2. **Database View**

  * 데이터베이스보기는 데이터베이스의 원시 레코드를 표시하여 SQL 명령이 반환하는 내용을 보여줍니다.

---
name: records-view
# Web App - Records View
.center[![:scale 90%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/records_view.png)]

* 앱이 암호화 된 데이터를 해독했기 때문에 승인 된 사용자는 일부 민감한 데이터를 볼 수 있습니다.

???
* Show the records view of the web app.

---
name: Vault-Transit-Engine-6
# "Add User" 화면
* 실습에서는 데이터베이스에 새 사용자를 추가합니다.
.center[![:scale 60%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/add_user.png)]

???
* Describe the Add User screen that students will use to add new records to the database.
* Point out again that when Vault is enabled, the records will be encrypted in the database.

---
name: database-record-without-vault
# Vault를 활성화하지 않고 데이터베이스 View에 기록
* 실습에 레코드를 추가하면 **Database View** 메뉴를 클릭하라는 메시지가 표시됩니다.
* 입력 한 것과 동일한 데이터가 표시되어야합니다.
* 이것은 개인 식별 데이터 (PII)가 데이터베이스 기록에 일반 텍스트로 저장됨을 의미합니다.
* 어떻게 개선 할 수 있습니까? Vault의 Transit 엔진을 활성화하고 확인해 보겠습니다.

---
name: encrypted-record
# A Database Record Encrypted by Vault
#### 다음은 Vault의 Transit Engine에 의해 암호화 된 기록입니다.
.center[![:scale 80%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/database_view_with_encrypted_record.png)]
* birth_date 및 social_security_number는 암호화되어 있습니다.
???
* Show a record from the database encrypted by Vault's Transit engine.
* Point out that the birth_date and social_security_number field are encrypted as indicated by their starting with "vault:v1".
* Point out that the "v1" indicates that the first version of the encryption key was used.

---
name: encryption-key-rotation
# Transit Engine 암호화 키 교체
* Vaults Transit Engine의 암호화 키는 순환 할 수 있습니다.
* 키의 최신 버전은 새 데이터를 암호화하는 데 사용됩니다.
* 이전 버전의 키는 여전히 이전 데이터를 해독 할 수 있지만 새 데이터는 해독 할 수 없습니다.
* 암호화 키를 교체 할 때 Transit 엔진을 사용하는 앱은 변경 사항을 인식하지 못합니다.
* 데이터는`rewrap` 엔드 포인트를 사용하여 다시 암호화 할 수도 있습니다.

---
name: lab-transit-challenge-1
# 👩‍💻 Challenge 1: Enable the Transit Engine
* 이 실습 챌린지에서는 Transit Engine을 활성화합니다.
* [Vault Encryption as a Service](https://play.instruqt.com/hashicorp/invite/qleasfx1dszc) Instruqt 트랙에서이 작업을 수행합니다.
* 지침 :
   * "Vault Encryption as a Service"트랙의 "Enable the Transit Secrets Engine"챌린지를 클릭합니다.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Enable the Transit Secrets Engine" challenge of the "Vault Encryption as a Service" track.
* This challenge has them enable the Transit secrets engine on the path "lob_a/workshop/transit".

---
name: lab-database-challenge-2
# 👩‍💻 Challenge 2: Create an Encryption Key
* 이 실습에서는 이전 챌린지에서 활성화 한 Transit Engine과 함께 사용할 암호화 키를 생성합니다.
* 지침 :
   * 트랙이 자동으로 수행되지 않는 경우 "Vault Encryption as a Service"트랙의 "Create a Key for the Transit Secrets Engine"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Create a Key for the Transit Secrets Engine" challenge of the "Vault Encryption as a Service" track.
* This challenge has them create an encryption key for use with the Transit engine they enabled in the previous challenge.

---
name: lab-database-challenge-3
# 👩‍💻 Challenge 3: Use the Web App Without Vault
* 이 실습에서는 Vault없이 웹 애플리케이션을 사용합니다.
* 지침 :
   * 트랙이 적합하지 않은 경우 "Vault Encryption as a Service"트랙의 "Use the Web App Without Vault"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.
* ???

* Instruct the students to do the "Use the Web App Without Vault" challenge of the "Vault Encryption as a Service" track.
* This challenge has them use the web application without Vault.
* Point out that the new record they add during this challenge will nto be encrypted.

---
name: lab-database-challenge-4
# 👩‍💻 Challenge 4: Use the Web App With Vault
* 이 실습에서는 Vault와 함께 웹 애플리케이션을 사용합니다.
* 암호화 키도 교체됩니다.
* 지침 :
   * 트랙이 적합하지 않은 경우 "Vault Encryption as a Service"트랙의 "Use the Web App With Vault"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Use the Web App With Vault" challenge of the "Vault Encryption as a Service" track.
* This challenge has them use the web application with Vault.
* Point out that the new record they add will have sensitive fields encrypted by Vault's Transit engine.

---
name: chapter-8-review-questions
# 📝 Chapter 8 Review
* Vault의 Transit Engine 사용의 주요 이점은 무엇입니까?
* Vault의 Transit Engine은 암호화 된 데이터를 어디에 저장합니까?
* 암호화 키를 교체 한 후에도 응용 프로그램에서 이전의 암호화 된 레코드를 해독 할 수 있었습니까?
* 어떤 버전의 암호화 키가 사용되었는지 알 수 있습니까?

???
* Let's review what we learned in this chapter.

---
name: chapter-8-review-answers
# 📝 Chapter 8 Review
* Vault의 Transit Engine 사용의 주요 이점은 무엇입니까?
  * 개발자는 암호화 전문가가 아니어도 데이터를 암호화 할 수 있습니다.
* Vault의 Transit Engine은 암호화 된 데이터를 어디에 저장합니까?
  * 개발자가 원하는 곳, Vault 외부
* 암호화 키를 교체 한 후에도 응용 프로그램에서 이전의 암호화 된 레코드를 해독 할 수 있었습니까?
  * 네
* 어떤 버전의 암호화 키가 사용되었는지 알 수 있습니까?
  * 예. 버전은`v1`,`v2` 등으로 표시됩니다.

???
* Here are the answers to the review questions.

---
name: conclusion
# Thank You for Participating!
.center[![:scale 40%](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/images/vault_logo.png)]

### 자세한 내용은 다음 링크를 참조하십시오.
* https://www.vaultproject.io/docs/
* https://www.vaultproject.io/api/
* https://learn.hashicorp.com/vault

???
* Thank the students for their participation
* Share some Vault links

---
name: Feedback-Survey
# Workshop Feedback Survey
* 귀하의 의견은 우리에게 중요합니다!
* 설문 조사는 짧습니다. :
  * http://bit.ly/hashiworkshopfeedback

???
* Ask them to fill out the online survey
