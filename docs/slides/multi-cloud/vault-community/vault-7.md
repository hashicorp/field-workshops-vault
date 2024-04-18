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

* Database credentials are normally long-lived.
* Vault's Database Secrets Engine dynamically generates short-lived credentials for databases.
* It supports configuration of database connections and roles with different permissions and time to live (TTL) settings.
* Users or applications request credentials for a specific role from Vault.
* Vault manages the lifecycle of these credentials, automatically deleting them from the database when the TTL expires.

???
* Vault's Database secrets engine supports dynamic generation of short-lived credentials (usernames and passwords) for databases.
* This avoids storing long-lived or permanent credentials on app servers that can easily be compromised.
* Short-lived credentials are much more secure since ex-employees and others are very unlikely to know the current values.

---
name: database-engine-plugins
# Database Secrets Engine: Plugins
* Cassandra
* Elasticsearch
* Influxdb
* HanaDB
* MongoDB
* MSSQL
* MySQL/MariaDB
* PostgreSQL
* Oracle

???
* The database secrets engine has out-of-the-box plugins for many databases.
* Custom plugins can also be built.

---
name: database-engine-workflow
# Database Secrets Engine Workflow
1. Enable an instance of the database secrets engine.
1. Configure it with the correct plugin and connection URL, using a service account created for Vault.
1. Create one or more roles with TTLs and SQL statements that specify required permissions.
1. Applications and users can request credentials from Vault that are valid for the default TTL of the role, but can be renewed up to the max TTL.
1. Vault automatically deletes expired credentials.
1. If credentials are compromised, you can revoke them immediately.

???
* This slide lays out the basic workflow used for all of the Datbase secrets engine plugins.
* All of the plugins work the same basic way.
* A service account with permissions to manage users on the database server is required by each connection.
* User creation and revocation SQL statements are specified for roles to determine the permissions og generated users within various databases.
* Multiple connections and roles can be created for a single secrets engine instance to support connecting to multiple database servers with different levels of access.
* The TTL settings can be tuned to suit your needs.

---
name: sample-web-app
# Lab Environment for Chapters 7 and 8

* In the labs for chapters 7 and 8, we'll use a MySQL database server that runs on the Vault server.
* We will also use a Python web application that stores its data in the MySQL database, but not until chapter 8.
* In the next few slides, we outline many of the steps you will do in the lab.

???
* Discuss the lab environment.

---
name: mysql-configuration-steps
# Configuration Steps for MySQL
1. Enable the database secrets engine on some path.
1. Configure it with the MySQL plugin, connection URL, username, password, and allowed roles.
1. Rotate the "root credentials": Vault modifies the password given in step 2 so that no humans know it anymore.
1. Create roles that can create new credentials that are valid for a specific period of time.

???
* These are the basic steps for configuring the mysql plugin with Vault's database secrets engine.
* The username and password set on the config path must already exist and have permission to manage users.

---
name: mysql-config-connection
class: compact
# Configuring Connections for MySQL
#### Run these commands to enable the Database secrets engine and configure a connection for use with MySQL:
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
#### This creates a connection called "wsmysqldatabase" against the MySQL server on localhost.

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
# Rotating the Root Credentials for MySQL
#### 1. You should **not** use the actual root user of the database (despite the reference to "root credentials"); instead, create a separate user with sufficient privileges to create users and to change its own password. This can be done by running `GRANT ALL PRIVILEGES on *.* to 'hashicorp'@'%' with grant option;` for the user.
#### 2. The actual username you provide should be for host `'%'`. So, create a user like `'hashicorp'@'%'` rather than `'hashicorp'@'localhost'`.
#### 3. If you don't want to use `'%'` as the host for the user, you can specify `root_rotation_statements` when writing to the path `<database>/config/<connection>`; for instance, you could set this to `"ALTER USER '{{username}}'@'localhost' IDENTIFIED BY '{{password}}';"`.

???
* We want to give some advice about rotating root credentials for the database secrets engine when using MySQL.

---
class:compact
# Configuring Roles for MySQL
#### Run this command to configure a role for MySQL:
```sql
vault write lob_a/workshop/database/roles/workshop-app-long \
    db_name=wsmysqldatabase \
    creation_statements="CREATE USER '{{name}}'@'%' IDENTIFIED BY '{{password}}';
    GRANT ALL ON my_app.* TO '{{name}}'@'%';" \
    default_ttl="1h" \
    max_ttl="24h"
```
#### This defines a role against the "wsmysqldatabase" connection which generates credentials with an initial TTL of 1 hour. But their lifetime can be extended up to 24 hours.

???
* We specified a number of things:
    * The creation statements that define the capabilities of the userd that are created
    * The default time to live for generated users
    * The maximum duration for generated users

---
name: mysql-generate-creds
class:compact
# Generating Database Credentials
#### Run this command to generate actual credentials for the MySQL database against the role that was configured on the previous slide:
```bash
vault read lob_a/workshop/database/creds/workshop-app-long  
```
#### This should return something like:<br>
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
# Renewing and Revoking Database Credentials
#### Run this command to renew credentials, replacing `<lease_id>` with the right lease_id:
```bash
vault write sys/leases/renew lease_id="<lease_id>" increment="120"  
```
#### Run this command to revoke credentials, replacing `<lease_id>` with the right lease_id:
```bash
vault write sys/leases/revoke lease_id="<lease_id>"
```
#### You can also determine the remaining lifetime of the credentials:
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
* In this lab challenge, you'll enable the database engine for MySQL and rotate its root credentials.
* You'll do this in the **Vault Dynamic Database Credentials** Instruqt track using the link provided by your instructor.
* Instructions:
  * Click the "Enable the Database Secrets Engine" challenge of the "Vault Dynamic Database Credentials" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Enable the Database Secrets Engine" challenge of the "Vault Dynamic Database Credentials" track.
* This challenge has them enable the Database secrets engine on the path "lob_a/workshop/database" and rotate the root credentials for it.

---
name: lab-database-challenge-2
# 👩‍💻 Challenge 2: Configure the Engine
* In this lab, you'll configure a connection and two roles for the database.
* Instructions:
  * If the track does not do it for you, click the "Configure the Database Secrets Engine" challenge of the "Vault Dynamic Database Credentials" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Configure the Database Secrets Engine" challenge of the "Vault Dynamic Database Credentials" track.
* This challenge has them configure a connection and two roles for the engine they created in the previous challenge.
* One role will have an initial TTL of 1 hour with a maximum TTL of 24 hours.
* The other will have an initial TTL of 3 minutes with a maximum TTL of 6 minutes.

---
name: lab-database-challenge-3
# 👩‍💻 Challenge 3: Generate Credentials
* In this lab, you'll generate and use credentials against both roles that you configured in the previous challenge.
* Instructions:
  * If the track does not do it for you, click the "Generate and Use Dynamic Database Credentials" challenge of the "Vault Dynamic Database Credentials" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Generate and Use Dynamic Database Credentials" challenge of the "Vault Dynamic Database Credentials" track.
* This challenge has them generate short-lived credentials for the MySQL database.
* They will use the `mysql` utility to connect to the database with those credentials.
* They will see that the credentials are deleted after 3 minutes and that logging into MySQL with them is blocked.

---
name: lab-database-challenge-4
# 👩‍💻 Challenge 4: Renew and Revoke Credentials
* In this lab, you'll renew and revoke credentials generated by the database secrets engine.
* Instructions:
  * If the track does not do it for you, click the "Renew and Revoke Database Credentials" challenge of the "Vault Dynamic Database Credentials" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Renew and Revoke Database Credentials" challenge of the "Vault Dynamic Database Credentials" track.
* This challenge has them extend the liftime of generated credentials with Vault's `sys/leases/renew` endpoint.
* They will also revoke credentials with Vault's `sys/leases/revoke` endpoint.

---
name: chapter-7-review-questions
# 📝 Chapter 7 Review
* What is the main advantage of using Vault's database secrets engine?
* What happens when credentials expire?
* Is the database engine limited to the plugins listed in this chapter?
* Can more than one role be used against a single connection?

???
* Let's review what we learned in this chapter.

---
name: chapter-7-review-answers
# 📝 Chapter 7 Review

* What is the main advantage of using Vault's database secrets engine?
  * The credentials are short-lived and less likely to be compromised.
* What happens when credentials expire?
  * Vault deletes them from the database server.
* Is the database engine limited to the plugins listed in this chapter?
  * No.  Custom plugins can be written.
* Can more than one role be used against a single connection?
  * Yes. This allows different apps to get credentials with different TTLs.

???
* Here are the answers to the review questions.
