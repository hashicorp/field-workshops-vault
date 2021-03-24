name: chapter-7
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false

# Chapter 7    
## Dynamic Database Secrets

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???

* Chapter 7 introduces Vault's Database secrets engine which can dynamically generate short-lived credentials for various databases.

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: dynamic-database-secrets
# Dynamic Secrets: Protecting Databases

* 데이터베이스 자격 증명은 일반적으로 오래 지속됩니다.
* Vault의 데이터베이스 비밀 엔진은 데이터베이스에 대한 단기 자격 증명을 동적으로 생성합니다.
* 지정한 권한과 TTL (Time to Live) 설정으로 데이터베이스 연결 및 역할 구성을 지원합니다.
* 사용자나 애플리케이션은 Vault에서 특정 역할에 대한 자격 증명을 요청합니다.
* Vault는 이러한 자격 증명의 수명주기를 관리하여 TTL이 만료되면 데이터베이스에서 자동으로 삭제합니다.

???

* Vault의 데이터베이스 비밀 엔진은 데이터베이스에 대한 단기 자격 증명 (사용자 이름 및 비밀번호)의 동적 생성을 지원합니다.
* 이렇게하면 쉽게 손상 될 수있는 앱 서버에 수명이 길거나 영구적 인 자격 증명이 저장되지 않습니다.
* 단기 자격 증명은 전직 직원과 다른 직원이 현재 값을 알 가능성이 거의 없기 때문에 훨씬 더 안전합니다.

---
name: database-engine-plugins
# Database Secrets Engine - 플러그인
* Cassandra
* Elasticsearch
* Influxdb
* HanaDB
* MongoDB
* MSSQL
* MySQL/MariaDB
* PostgreSQL
* Oracle
* Snowflake

???
* The database secrets engine has out-of-the-box plugins for many databases.
* Custom plugins can also be built.

---
name: database-engine-workflow
# Database Secrets Engine - Workflow
1. 데이터베이스 비밀 엔진의 인스턴스를 활성화합니다.
1. Vault 용으로 생성 된 서비스 계정을 사용하여 올바른 플러그인 및 연결 URL로 구성합니다.
1. 필요한 권한을 지정하는 TTL 및 SQL 문으로 하나 이상의 역할을 만듭니다.
1. 애플리케이션과 사용자는 역할의 기본 TTL에 유효한 자격 증명을 Vault에서 요청할 수 있지만 최대 TTL까지 갱신 할 수 있습니다.
1. Vault는 만료 된 자격 증명을 자동으로 삭제합니다.
1. 자격 증명이 손상된 경우 즉시 취소 할 수 있습니다.

???
* This slide lays out the basic workflow used for all of the Datbase secrets engine plugins.
* All of the plugins work the same basic way.
* A service account with permissions to manage users on the database server is required by each connection.
* User creation and revocation SQL statements are specified for roles to determine the permissions og generated users within various databases.
* Multiple connections and roles can be created for a single secrets engine instance to support connecting to multiple database servers with different levels of access.
* The TTL settings can be tuned to suit your needs.

---
name: sample-web-app
# 7장, 8장을위한 실습 환경

* 7 장과 8 장의 실습에서는 Vault 서버에서 실행되는 MySQL 데이터베이스 서버를 사용합니다.
* 또한 MySQL 데이터베이스에 데이터를 저장하는 Python 웹 애플리케이션도 사용하지만 8 장까지는 사용할 수 없습니다.
* 다음 몇 개의 슬라이드에서는 실습에서 수행 할 여러 단계를 간략하게 설명합니다.

???
* Discuss the lab environment.

---
name: mysql-configuration-steps
# MySQL 구성 가이드
1. 특정 경로에서 데이터베이스 시크릿 엔진을 활성화합니다.
1. MySQL 플러그인, 연결 URL, 사용자 이름, 암호, 허용 된 역할로 구성합니다.
1. "Root credential" 교환 : Vault와 연계된 계정은 더 이상 다른 사람이 알 수 없도록, 앞서 2단계에서 제공된 암호를 수정합니다.
1. 특정 기간 동안 유효한 새 자격 증명을 만들 수있는 역할을 만듭니다.

???
* These are the basic steps for configuring the mysql plugin with Vault's database secrets engine.
* The username and password set on the config path must already exist and have permission to manage users.

---
name: mysql-config-connection
class: compact
# MySQL 구성 가이드
#### 다음 명령을 실행하여 데이터베이스 시크릿 엔진을 활성화하고 MySQL에서 사용할 연결을 구성합니다.
```bash
vault secrets enable -path=lob_a/workshop/database database

vault write lob_a/workshop/database/config/wsmysqldatabase \
    plugin_name=mysql-database-plugin \
    connection_url="{{username}}:{{password}}@tcp(localhost:3306)/" \
    allowed_roles="workshop-app","workshop-app-long" \
    username="hashicorp" \
    password="Password123"

vault write -force lob_a/workshop/database/rotate-root/wsmysqldatabase
```
#### 구성하게 되면, localhost의 MySQL 서버에 대해 "wsmysqldatabase"라는 연결이 생성됩니다.

???
* This slide shows the commands to enable the Database secrets engine and configure a connection for MySQL.
* We specified a number of things in the configuration:
    * The path someone would call: "lob_a/workshop/database"
    * The name of the database the role can interact with: wsmysqldatabase
    * The connection URL
    * The initial username and password
    * The roles that can be used with this connection
* We then rotated the password for the "root" user so that only Vault knows it.

---
name: rotating-root-credentials
class: compact
# MySQL에 대한 root credentials 교체
#### 1. "root credentials"에 대한 참조에도 불구하고 데이터베이스의 실제 루트 사용자를 사용해서는 안됩니다. 대신 사용자를 만들고 자신의 암호를 변경할 수있는 충분한 권한이있는 별도의 사용자를 만듭니다. `GRANT ALL PRIVILEGES on *.* to 'hashicorp'@'%' with grant option;`

#### 2. 제공하는 실제 사용자 이름은 호스트` '%'`용이어야합니다. 따라서` 'hashicorp'@'localhost'` 대신`'hashicorp'@'%'`와 같은 사용자를 만듭니다.
#### 3. 사용자 호스트로` '%'`를 사용하지 않으려면`<database>/config/<connection>`경로에 쓸 때`root_rotation_statements`를 지정할 수 있습니다. 예를 들어이를` "ALTER USER '{{username}}'@'localhost'IDENTIFIED BY '{{password}}';"`로 설정할 수 있습니다.

???
* We want to give some advice about rotating root credentials for the database secrets engine when using MySQL.

---
class:compact
# MySQL에 Configuring Roles 설정
#### 이 명령을 실행하여 MySQL에 대한 Role(역할)을 구성합니다.
```sql
vault write lob_a/workshop/database/roles/workshop-app-long \
    db_name=wsmysqldatabase \
    creation_statements="CREATE USER '{{name}}'@'%' IDENTIFIED BY '{{password}}';
    GRANT ALL ON my_app.* TO '{{name}}'@'%';" \
    default_ttl="1h" \
    max_ttl="24h"
```
#### 이는 초기 TTL이 1 시간 인 자격 증명을 생성하는 "wsmysqldatabase"연결에 대한 역할을 정의합니다. 그러나 TTL은 최대 24 시간까지 연장 할 수 있습니다.

???
* We specified a number of things:
    * The creation statements that define the capabilities of the userd that are created
    * The default time to live for generated users
    * The maximum duration for generated users

---
name: mysql-generate-creds
class:compact
# Database Credentials 생성
#### 이전 슬라이드에서 구성된 역할에 대해 MySQL 데이터베이스에 대한 실제 자격 증명을 생성하려면이 명령을 실행합니다. :
```bash
vault read lob_a/workshop/database/creds/workshop-app-long  
```
#### 이것은 다음과 같은 것을 반환합니다 :<br>
```bash
Key                Value
---                -----
lease_id           lob_a/workshop/database/creds/workshop-app-long/JeUGIL2xD6BzXSneqity8UmF
lease_duration     1h
lease_renewable    true
password           A1a-zy4ENaf2kwpzGk9t
username           v-token-workshop-a-DM0BJ3eMlMhbf
```

???
* Now, we can begin generating credentials for our MySQL database.

---
name: mysql-renew-revoke-creds
class:compact
# Renewing & Revoking Database Credentials
#### 이 명령을 실행하여 자격 증명을 갱신하고`<lease_id>`를 올바른 Lease ID로 바꿉니다.:
```bash
vault write sys/leases/renew lease_id="<lease_id>" increment="120"  
```
#### 이 명령을 실행하여 자격 증명을 갱신하고`<lease_id>`를 올바른 Lease ID로 바꿉니다.:
```bash
vault write sys/leases/revoke lease_id="<lease_id>"
```
#### Credentials의 남은 수명을 확인할 수도 있습니다.
```bash
vault write sys/leases/lookup lease_id="<lease_id>"
```

???
* These are the commands to renew and revoke Vault leases.
* When you run the `renew` command, Vault extends the lifetime of the credentials.
* When you run the `revoke` command, Vault revokes the lease and removes the credentials from the database server.
* It is also possible to determine the remaining lifetime of credentials.

---
name: lab-database-challenge-1
# 👩‍💻 Challenge 1: Enable the Engine
* 이 실습 과제에서는 MySQL 용 데이터베이스 엔진을 활성화하고 루트 자격 증명을 교체합니다.

* [Vault Dynamic Database Credentials](https://play.instruqt.com/hashicorp/invite/sryhqfdm6sgx) Instruqt 트랙에서이 작업을 수행합니다.
* 지침 :
   * "Vault Dynamic Database Credentials"트랙의 "Enable the Database Secrets Engine"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Enable the Database Secrets Engine" challenge of the "Vault Dynamic Database Credentials" track.
* This challenge has them enable the Database secrets engine on the path "lob_a/workshop/database" and rotate the root credentials for it.

---
name: lab-database-challenge-2
# 👩‍💻 Challenge 2: Configure the Engine
* 이 실습에서는 데이터베이스에 대한 연결과 두 가지 역할을 구성합니다.
* 지침 :
   * 트랙이 자동으로 수행되지 않는 경우 "Vault Dynamic Database Credentials"트랙의 "Configure the Database Secrets Engine"챌린지를 클릭합니다.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Configure the Database Secrets Engine" challenge of the "Vault Dynamic Database Credentials" track.
* This challenge has them configure a connection and two roles for the engine they created in the previous challenge.
* One role will have an initial TTL of 1 hour with a maximum TTL of 24 hours.
* The other will have an initial TTL of 3 minutes with a maximum TTL of 6 minutes.

---
name: lab-database-challenge-3
# 👩‍💻 Challenge 3: Generate Credentials
* 이 실습에서는 이전 챌린지에서 구성한 두 역할에 대해 자격 증명을 생성하고 사용합니다.
* 지침 :
   * 트랙에서 자동으로 수행되지 않는 경우 "Vault Dynamic Database Credentials"트랙의 "Generate and Use Dynamic Database Credentials"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Generate and Use Dynamic Database Credentials" challenge of the "Vault Dynamic Database Credentials" track.
* This challenge has them generate short-lived credentials for the MySQL database.
* They will use the `mysql` utility to connect to the database with those credentials.
* They will see that the credentials are deleted after 3 minutes and that logging into MySQL with them is blocked.

---
name: lab-database-challenge-4
# 👩‍💻 Challenge 4: Renew and Revoke Credentials
* 이 실습에서는 데이터베이스 비밀 엔진에서 생성 한 자격 증명을 갱신하고 취소합니다.
* 지침 :
   * 트랙에서 자동으로 수행되지 않는 경우 "Vault Dynamic Database Credentials"트랙의 "Renew and Revoke Database Credentials"챌린지를 클릭하십시오.
   * 그런 다음 녹색 "시작"버튼을 클릭합니다.
   * 도전 지침을 따르십시오.
   * 완료되면 녹색 "확인"버튼을 클릭합니다.

???
* Instruct the students to do the "Renew and Revoke Database Credentials" challenge of the "Vault Dynamic Database Credentials" track.
* This challenge has them extend the liftime of generated credentials with Vault's `sys/leases/renew` endpoint.
* They will also revoke credentials with Vault's `sys/leases/revoke` endpoint.

---
name: chapter-7-review-questions
# 📝 Chapter 7 Review
* Vault의 데이터베이스 시크릿 엔진 사용의 주요 이점은 무엇입니까?
* Credential이 만료되면 어떻게됩니까?
* 데이터베이스 엔진이 이번 장에 나열된 플러그인만으로 제한됩니까?
* 단일 연결에 대해 둘 이상의 Role(역할)을 사용할 수 있습니까?

???
* Let's review what we learned in this chapter.

---
name: chapter-7-review-answers
# 📝 Chapter 7 Review

* Vault의 데이터베이스 시크릿 엔진 사용의 주요 이점은 무엇입니까?
  * 자격 증명은 수명이 짧고 손상 될 가능성이 적습니다.
* Credential이 만료되면 어떻게됩니까?
  * Vault는 데이터베이스 서버에서 삭제합니다.
* 데이터베이스 엔진이 이번 장에 나열된 플러그인만으로 제한됩니까?
  * 아니요. 사용자 지정 플러그인을 작성할 수 있습니다.
* 단일 연결에 대해 둘 이상의 Role(역할)을 사용할 수 있습니까?
  * 예. 이를 통해 서로 다른 앱이 서로 다른 TTL로 자격 증명을 얻을 수 있습니다.

???
* Here are the answers to the review questions.
