name: vault-basics-lab
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false

# Vault Basics Lab  

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???
These slides introduce the Vault Basics track.

---
name: getting-started-with-instruqt
# Instruqt로 실습하기
* [Instruqt](https://instruqt.com/)은 HashiCorp 워크숍에 사용되는 플랫폼입니다.
* Instruqt 랩은 "challenges"으로 구분 된 "tracks"에서 실행됩니다.
* 이전에 Instruqt를 사용한 적이 없다면이 [튜토리얼](https://play.instruqt.com/instruqt/tracks/getting-started-with-instruqt)으로 시작하세요.
* 그렇지 않으면 다음 슬라이드로 건너 뛸 수 있습니다.
* 이 "Vault Basic"실습에서는 2-6 장에 소개 된 개념을 다룹니다.

???
* We'll be using the Instruqt platform for labs in this workshop.
* Don't worry if you've never used it before: there is an easy tutorial that you can run through in 5-10 minutes.
---
name: lab-vault-basics-challenge-1
# 👩‍💻 Challenge 1: The Vault CLI
* 이 실습에서는 일부 Vault CLI 명령을 실행합니다.
* [Vault Basics](https://play.instruqt.com/hashicorp/invite/qfwncq62zsxu) Instruqt 트랙의 첫 번째 도전 인 "Vault CLI"에서이 작업을 수행합니다.

???
* Now, you can try running some Vault CLI commands yourself in the first challenge of our first Instruqt track in this workshop.
* This lab covers concepts introduced in chapters 2-6.

---
name:lab-vault-basics-challenge-1-instructions
# 👩‍💻 Challenge 1: Instructions
* 트랙의 "Vault CLI"챌린지에서 보라색 "시작"버튼을 클릭하여 "Vault Basics"트랙을 시작합니다.
* 챌린지가로드되는 동안 표시된 텍스트를 읽으십시오.
* "Vault CLI" 챌린지를 시작하려면 녹색 "시작"버튼을 클릭합니다.
* 도전 오른쪽에있는 지침을 따르십시오.
* 모든 단계를 완료 한 후 녹색 "확인"버튼을 클릭하여 모든 작업을 올바르게 수행했는지 확인하십시오.
* 또한 "확인"버튼을 클릭하여 미리 알림을받을 수도 있습니다.

???
* Give the students some instructions for starting their first challenge.
* This also includes instructions for checking that they did everything right.
* Students can also click the green "Check" button to get reminders of what they should do next.

---
name: lab-vault-basics-challenge-2
# 👩‍💻 Challenge 2: Run a Vault "Dev" Server
* 이 실습에서는 "Dev"모드에서 첫 번째 Vault 서버를 실행합니다.
* 또한 첫 번째 비밀을 Vault에 쓰고 UI를 사용하게됩니다.
* 지침 :
   * 트랙이 적합하지 않으면 "Vault Basics"트랙의 "Your First Secret"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.
* ???

* Instruct the students to do the "Your First Secret" challenge of the "Vault Basics" track.
* This challenge has them run a Dev server, write a secret to the KV v2 secrets engine that was automatically enabled, and use the Vault UI.

---
name: lab-vault-basics-challenge-3
# 👩‍💻 Challenge 3: Use the Vault HTTP API
* 이 실습에서는 Vault HTTP API를 사용합니다.
* 먼저 Vault 서버의 상태를 확인합니다.
* 그런 다음 Vault에서 'my-first-secret'비밀을 읽습니다.
* 지침 :
   * 트랙이 적합하지 않은 경우 "Vault Basics"트랙에서 "The Vault API"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the challenge, "The Vault API", in the "Vault Basics" track.

---
name: lab-vault-basics-challenge-4
# 👩‍💻 Challenge 4: Run a Vault "Prod" Server
* 이 실습에서는 "Prod"모드에서 첫 번째 Vault 서버를 실행합니다.
* Vault 서버를 초기화하고 봉인 해제하는 방법을 배웁니다.
* 지침 :
   * 트랙이 적합하지 않으면 "Vault Basics"트랙의 "Run a Production Server"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Run a Production Server" challenge of the "Vault Basics" track.
* This challenge has them examine a Vault server configuration file, run a Prod server, initialize it, and unseal it.
* Remind students to save their unseal key and root token.

---
name: lab-vault-basics-challenge-5
# 👩‍💻 Challenge 5: Use the KV v2 Secrets Engine
* 이 실습에서는 KV v2 비밀 엔진을 활성화하고 사용합니다.
* 경로는 `secret`이 아닌 `kv`가됩니다.
* 지침 :
   * 트랙이 적합하지 않은 경우 "Vault Basics"트랙의 "Use the KV V2 Secrets Engine"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Use the KV V2 Secrets Engine" challenge of the "Vault Basics" track.
* This challenge has them enable an instance of the KV v2 secrets engine.
* Emphasize that the path will be `kv` instead of `secret` as was the case for the challenges with the Dev mode server.

---
name: lab-vault-basics-challenge-6
# 👩‍💻 Challenge 6: Userpass Auth Method
* 이 실습에서는 Userpass 인증 방법을 활성화하고 사용합니다.
* 지침 :
   * 트랙이 적합하지 않으면 "Vault Basics"트랙의 "Use the Userpass Auth Method"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Use the Userpass Auth Method" challenge of the "Vault Basics" track.
* This challenge has them enable an instance of the Userpass auth method.
* It also demonstrates that Vault is "deny by default" since the Userpass user that they create will not have any access to secrets yet.

---
name: lab-vault-basics-challenge-7
# 👩‍💻 Challenge 7: Vault Policies
* 이 실습에서는 Vault 정책을 사용하여 다른 사용자에게 다른 보안 비밀에 대한 액세스 권한을 부여합니다.
* 지침 :
   * 트랙이 적합하지 않은 경우 "Vault Basics"트랙의 "Use Vault Policies"챌린지를 클릭합니다.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Use Vault Policies" challenge of the "Vault Basics" track.
* This challenge has them create a second user and create and associate policies with their 2 users.
* It then has them valdiate that each user can only access their own secrets.
