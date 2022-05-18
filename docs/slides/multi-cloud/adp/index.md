name: field-workshops-assets-index
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)
<div class="clearfix">
  <div class="img-container">
    <img src="images\Vault_VerticalLogo_ColorWhite_RGB.png" style="width:5%">
  </div>
  <div class="img-container">
    <img src="images\Consul_VerticalLogo_ColorWhite_RGB.png" style="width:5%">
  </div>
  <div class="img-container">
    <img src="images\Boundary_VerticalLogo_ColorWhite_RGB.png" style="width:5%">
  </div>
</div>

# Zero Trust Security & Data Protection
## Trust Nothing. Authenticate and Authorize Everything.

???

This workshop introduces students to Vault and some of its more advanced data
protection capabilities to enable the Zero Trust security model.

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](images/HashiCorp_Icon_Black.svg)
]

---
name: Link-to-Slide-Deck
# The Slide Deck
<br><br>
### Follow along on your own computer at this link:

https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/adp

???

The link to this slide deck

---
name: intros
# Introductions

* Your Name
* Job Title
* Secrets Management Experience

???

* Use this slide to introduce yourself, give a little bit of your background story, then go around the room and have all your participants introduce themselves.

---
name: prerequisites
# Prerequisites

- Some very basic familiarity with the following Vault concepts
  - Interacting with Vault via one of the CLI, GUI, or API
      - The Vault binary
      - Client-Server model
  - Vault Authentication Methods (High-level understanding)
  - Vault Secrets Engines (High-level understanding)

???

* This workshop assumes some basic familiarity with Vault, but we will explain
everything as we go so don't worry too much

---
name: table of contents
# Table of Contents
* HashiCorp Vault Overview
* Dynamic Database Secrets
* API Encryption as a Service (& Tokenization)
* Filesystem/Database/Hardware Encryption as a Service

???

* We're going to be covering the following topics...

---
name: environment
# Lab Environment Used
* This workshop uses Instruqt for hand-on labs.
* Instruqt labs are run in “tracks” that are divided into “challenges”
* This workshop uses the following tracks
  1. Vault Basics
  1. Vault Dynamic Database Credentials
  1. Vault Transform and Transit Secrets Engines (App-Level Encryption)
  1. (Optional) Vault Transform-Tokenization Secrets Engine
  1. (Optional) Vault Key Management Secrets Engine
  1. Vault KMIP Secrets Engine (Filesystem/Database Level Encryption)
* Your instructor will provide the URLs for the tracks.

???

* We'll have some hands-on exercises to go through
* Ask if everyone has the Instruqt invite links, and send out new invites if
needed

---
name: prologue
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Prologue
## HashiCorp and Zero Trust

???

---
name: securing
class: shelf, no-footer
## Securing a datacenter was easy...

* All unauthorized traffic or access could be restricted/blocked
* Networks were trusted and apps and databases could interconnect with ease
* Four walls and trusted network protected secrets and sensitive information
* Resources were mostly static. Rules were based upon IP addresses, credentials were baked into source code or kept in a static file on disk

##### But what happens when your apps and infrastructure extend to multiple datacenters, cloud, or both?

???

* Traditional security models were built on the idea of perimter based security
* (go through each point on the slide)

---
name: securing
class: col-2
# Secure Infrastructure

With each new cloud, network topologies become more complex
#### Your infrastructure is now one IAM API call away from exposure

![](images/slide9.png)

???

* Moving to the multi-cloud or hybrid environments means that fixed perimeter no
longer exists
* Networks are more complicated, and with access configured via APIs, it doesn't
take much to accidentally grant access to the wrong people

---
name: jit
class: center, middle
# Just in Time Access in a Zero Trust World
![](images/slide10.png)

???

* Rather than relying on fixed attributes like IP addresses, we need some form
of Identity-based access, where we verify something's identity before letting it
access things
* Four parts to that (go over those on slide)

---
name: jit
class: center, middle
# Our Focus Today
![](images/slide11.png)

???

* We are mostly going to be focusing on that first problem today, how we can
identify machines, and grant access to secrets and other means of data protection

---
# Applied Zero Trust
Adopting a Zero Trust model requires technical capabilities that...
* Never Trust, Always Verify: Every user or machine must authenticate and be explicitly authorized for access to credentials, APIs, VMs, databases, etc.
* Principle of Least Privilege: Every user or machine accesses only the resources it requires for a defined purpose and time period.
* Assumed Breach: Continuously defend and encrypt critical PII and company data assuming that your network has been breached.

???

* We'll be discussing that in the context of of Zero Trust
* (go over points on slide)

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Chapter 1:
## Vault Basics

???

Chapter 1 introduces Vault

---
# Problems with the Traditional Security Model
- IP Address based rules
- Hardcoded credentials with problems such as:
  - Shared service accounts for apps and users
  - Difficult to rotate, decommission, and determine who has access
  - Revoking compromised credentials can break entire application stacks

???

* Earlier we discussed the traditional perimeter-based security model. Here are
some problems with it:
* (go over each of these points)

---
class: center, middle
# Identity Based Security
![:scale 70%](images/slide16.png)
**[Identity Based Security and Low Trust Networks](https://www.hashicorp.com/identity-based-security-and-low-trust-networks)**

???
* Here we see that Vault has multiple means of authenticating users and applications with its Auth Methods.
* Vault can manage many types of secrets and excels at generating short-lived, dynmamic secrets.
* Vault's ACL policies are associated with tokens that users and applications use to access secrets after authenticating.
* Tokens can only read/write secrets that its policies allow.
* Click on the link to read a white paper about identity-based security in low trust networks.

---
# Identity Based Security
Vault was designed to address the security needs of modern applications. It differs from the traditional approach by using...
* Identity based rules allowing security to stretch across perimeter
  * Never Trust, Always Verify
* Dynamic, short-lived, credentials that are frequently rotated and unique to every client (no shared credentials)
  * Principle of least privilege
  * Ties all actions back to identity
* Credentials and Entities can be easily invalidated to reduce blast radius

???
* This slide discusses how Vault is designed for modern applications.
* Mapping that onto the principles of zero trust... (go over each point)

---
class: middle, center
# Vault Architecture - High Availability
Integrated Raft Storage
![:scale 90%](images/slide18.png)

???
* Vault allows multiple servers to be combined in a highly available cluster within a single cloud region or physical data center.
* In this way, if a availability zone goes down, you still have a quorum of Vault servers available to handle requests

---
class: middle, center
# Vault Architecture - Multi-Region
Replication
![:scale 65%](images/slide19.png)

???
* Vault Enterprise supports replication between clusters across regions and data centers.
* It supports Disaster Recovery and Performance replication.
* These can be used together.

---
# Chapter 1 Review
## What is HashiCorp Vault?
* Vault is a Secrets Management System.
* It is API-driven and cloud agnostic.
* It can be used in untrusted networks.
* It can authenticate users and applications against many identity systems.
* It supports just-in-time generation of short-lived secrets.
* It runs in highly available clusters that can be replicated across regions.

???

* To review... (go through points on slide)

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Chapter 2:
## Basic Operations

---
# Interacting with Vault
Vault provides several mechanisms for interacting with it:
* The Vault CLI
* The Vault GUI
* The Vault API

---
# Basic Vault CLI Commands
* **`vault`** by itself will give you a list of many Vault CLI Commands
  * The list starts with the most common ones
* **`vault version`** tells you the version of Vault you are running
* **`vault read`** is used to read secrets from Vault
* **`vault write`** is used to write secrets to Vault<br><br>
The **`-h`**, **`-help`**, **`--help`** flags can be added to get help for any Vault CLI command

---
# Running a Production Vault Server
Running a Vault server in “Prod” mode involves multiple steps
* Specify a configuration in a config file
* Start the server
* Initialize the server to get unseal keys and an initial root token
* Unseal the Vault server with the unseal keys

---
# Initializing Vault Clusters
* Recall that a Vault cluster runs multiple Vault servers
* Each Vault cluster must be initalized once
* This is done with the **`vault operator init`** command
* The number of key shares and the key threshold can be specified with the **`key-shares`** and **`key-threshold`** CLI options
* The command returns the unseal keys and initial root token for the cluster

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Chapter 3:
## Vault Auth Methods and Static Secrets

---
class: middle, center
# Vault Authentication Methods
![:scale 40%](images/slide27.png)
Vault acts as an Identity Broker for the underlying platform or cloud.<br>
Use the right tool for the job to authenticate your clients!

---
class: middle, center
# Vault Secrets Engines
![:scale 50%](images/slide28.png)
Vault includes many different Secrets Engines

---
# KV Secrets Engine Commands
* Use this command to mount an instance of the KV v2 secrets engine on the default path `kv`:<br>
`vault secrets enable -version=2 kv`
* The `vault kv` commands allow you to interact with KV engines
  * `vault kv list` lists secrets at a specified path
  * `vault kv put` writes a secret at a specified path
  * `vault kv get` reads a secret at a specified path
  * `vault kv delete` deletes a secret at a specified path
* Other `vault kv` subcommands operate on versions of KV v2 secrets

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Lab 1:
## Vault Basics

---
# Doing Labs with Instruqt
* Instruqt is the platform used for HashiCorp's workshops
* Instruqt labs are run in “tracks” that are divided into “challenges”
* If you’ve never used Instruqt before, start with this **[tutorial](https://play.instruqt.com/instruqt/tracks/getting-started-with-instruqt)**
* Otherwise, you can skip to the next slide

---
# Lab 1: Vault Basics
In this lab you will cover some basics of Vault like the CLI, GUI, K/V Secrets Engine, and Userpass Auth Method

Lab: **Vault Basics**

Your instructor will provide the URL for the track.

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Chapter 4:
## Dynamic Secrets

---
class: title-slide, center
.middle[
  ### Why are Dynamic Secrets so Important?
]

---
class: middle, center
# Static Secret Management - Historical
Administrators obtained static, long-lived credentials and manually configured applications
![:scale 100%](images/slide35.png)

---
class: middle, center
# Static Secret Management - Scaling
You shouldn’t share AD credentials with your teammates, so why do so with machines and services?
![:scale 60%](images/slide36.png)

---
class: middle, center
# Dynamic Secrets in Action
Unique, short-lived, just-in-time credentials for each application instance
![:scale 80%](images/slide37.png)

---
# Dynamic Secrets: Protecting Databases
* Database credentials are historically long-lived
* Vault’s Database Secrets Engine dynamically generates short-lived credentials for databases
* It supports configuration of database connections and roles with different permissions and time-to-live (TTL) settings
* Users or applications request credentials for a specific role from Vault
* Vault manages the lifecycle of these credentials, automatically deleting them from the database when the TTL expires
* Auditing is now improved as each application instance has a unique credential

---
class: col-2
# Dynamic Secrets Engine: Plugins
<div>
* Cassandra
* Elasticsearch
* Influxdb
* HanaDB
* MongoDB
* MSSQL
* MySQL/MariaDB
* PostgresQL
* Oracle
* Snowflake
* And more (New plugins are always being added!)
<div>

---
# Dynamic Secrets Engine: Workflow
1. Enable an instance of the database secrets engine
1. Configure it with the correct Plugin and connection URL, using a service account created for Vault
1. Create one or more roles with TTLs and SQL statements that specific required permissions. (Principle of least privilege)
1. Applications and users can request credentials from Vault that are valid for the default TTL of the role, but can be renewed up to the max TTL
1. Vault automatically deletes expired credentials from the database
1. If credentials are compromised, administrators can revoke them immediately. (Break glass)

---
class: middle, center
# Dynamic Secrets Engine: Workflow
![:scale 60%](images/slide41.png)

---
# Other Dynamic Secrets Engines
* Public Key Infrastructure / Venafi
* SSH Keys
* AWS/Azure/GCP/AliCloud API credentials
* Active Directory / LDAP service accounts (Dynamic, check-in/check-out)
* Consul Tokens
* TFE Tokens

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Lab 2:
## Dynamic MySQL credentials

---
# Lab Environment
In this lab we’ll use a MySQL database server that runs on the Vault Server

Lab: **Vault Dynamic Database Credentials**

Your instructor will provide the URL for the track.

---
# Configuration Steps for MySQL
1. Enable the database secrets engine on some path in Vault
1. Configure it with the MySQL plugin, connection URL, username, password, and allowed roles
1. Rotate the “root credentials”: Vault modifies the password given in step 2 so that no humans know it anymore
1. Create roles that can create new credentials that are valid for a specific period of time

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Chapter 5:
## Encryption as a Service

---
# Applied Zero Trust
* Assume Breach: Continuously defend critical PII data and company data with the assumption that your network has been breached
  * The average time to determine adversarial presence within an enterprise is 191 days (Ponemon Institute)
* What are our customers trying to achieve?
  * Encryption of all critical data
  * Universal KMS support for applications in hybrid cloud environments
  * Consolidated workflows for hardware (KMIP) encryption solutions
  * Fips 140-2 (140-3) Compliance

---
# Data Breaches: The Application Layer
* With cloud adoption, the traditional approach to securing customer data breaks down
* Adversaries are not typically breaking into a cloud datacenter and stealing physical hardware
* They are breaking into organizations via phishing attacks, exposed networks, and supply chain attacks
* Once inside the network they are escalating credentials and gaining privileged access to databases and systems

---
class: col-2
# Compromised DBA Creds
Breaches are commonly carried out via attackers who have gained escalated credentials. They were then able to bypass TDE as an example. Credit card numbers are exposed in plaintext
![:scale 100%](images/slide49.png)

---
# Vault's Solutions
1. Vault’s Transform And Transit Secrets Engines provide Encryption-as-a-Service
2. Developers use Vault to encrypt and decrypt data outside of Vault

---
# Solution 1:Transit Secrets Engine
.center[![:scale 90%](images/slide53.png)]

---
# Example with Encryption Enabled
.center[![:scale 90%](images/Transit_enabled.png)]

---
# Solution 2: Transform Secrets Engine
.center[![:scale 90%](images/slide50.png)]

---
# Transform Data Masking
.center[![:scale 90%](images/masking.png)]

---
# Transform (& Transit) Engine Benefits:
* Vault’s Transform Engine provides developers with a well-architected API so they do not have to become encryption or cryptography experts
* Vault is platform agnostic so developers can code against one API
* It ensures approved ciphers and algorithms are used
* It supported automated key rotation and re-wrapping
* If an attacker manages to get access to the encrypted data, they will only see ciphertext that is useless without Vault.
* The Transform Secrets engine is Format Preserving. Thus, it does not require any changes to database structure
  * i.e. 16 digits CCNs are encrypted as 16 digit ciphertext

---
class: col-2
# Tokenization
.smaller[
* Non-Reversible Identification: Protect data pursuant to requirements for data irreversibility (PCI-DSS, GDPR, etc.) with strong forward secrecy
* Integrated Metadata: Supports metadata for identifying data type and purpose
]
.center[![:scale 50%](images/tokenization_high_level.png)]

---
# Tokenization (Diagram)
.center[![:scale 90%](images/tokenization_workflow.png)]

---
# Transform vs. Transit vs. Tokenization
.center[![:scale 60%](images/ADP_matrix.png)]


---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Lab 3:
## Vault Transform & Transit Secrets Engines

---
# Lab Environment
In this next Lab we’ll use a web application that leverages both the Transform and Transit Secrets engines to encrypt and decrypt data
* The “HashiCups” application will leverage Vault’s API to encrypt customer credit card numbers before writing them to the backend database
* This lab will also showcase data-masking

Lab: **Advanced Data Protection with Transform**

Your instructor will provide the URL for the track.

---
# Lab 3 Part 2 (Optional)
In this optional lab, you can leverage a Golang application and the Transform Secrets Engine to tokenize data.
* The Golang application will leverage Vault’s API to tokenize customer Social Security numbers before writing them to the backend database
* This lab will also showcase application code modifications as well

Lab: **Vault ADP with Tokenization**

Your instructor will provide the URL for the track.

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg


# Chapter 6:
## The Vault KMIP Secrets Engine

---
# Data Breach: Filesystem/OS layer
* In this next scenario, an adversary has gained remote access to the underlying database host via SSH/RDP
* They may not have Database credentials; however, they have the ability to inspect files on the host
.center[![:scale 70%](images/slide58.png)]

---
# Data Breach: Filesystem/OS layer
* To provide a defense-in-depth approach, Vault can also act as a KMIP (Key-Management-Interoperability-Protocol) Server
* This enables Vault to act as a KMS server for several different hardware and OS level native encryption technologies
* Some Applications include:
  * Key creation, storage, management
  * Encrypt/decrypt
  * Cryptographic offloads for FDE, volume encryption, secret management, etc

---
# Data Breach: Filesystem/OS layer
* To provide a defense-in-depth approach, Vault can also act as a KMIP (Key-Management-Interoperability-Protocol) Server
* This enables Vault to act as a KMS server for several different hardware and OS level native encryption technologies
* Some Applications include:
  * Key creation, storage, management
  * Encrypt/decrypt
  * Cryptographic offloads for FDE, volume encryption, secrets management, etc

---
# Example (MongoDB):
.center[![:scale 70%](images/Mongodb.png)]

---
# Examples:
Transparent Database Encryption (TDE): Automatically protect data in MySQL MongoDB, and other databases using Vault Enterprise
* MySQL
* MongoDB Enterprise

---
# Examples:
Disk and Volume Protection: Protect volume data on physical (FDE) and virtual (VMDK) infrastructures on prem and in the cloud
* NetApp Storage Encryption, NetApp Full Disk Encryption
* Dell EMC Unity, Data Domain, VMAX, etc.
* HPE Key Manager, Tape Storage

---
# Examples:
Portable Key Management: Protect encryption keys for data including files, virtual machines, and more across on-prem and cloud infrastructures
* IBM Filenet
* Oracle Key Vault
* **VMWare vSphere**
* Cisco UCS
* Rubrik
* Brocade Encryption SAN

---
# Key Management Secrets Engine
This engine provides a consistent workflow for distribution and lifecycle management of cryptographic keys in Key Management Service providers
* Maintain control of keys while taking advantage of cryptographic capabilities native to the KMS providers.
* Manage key lifecycle operations, such as creating, reading, updating, and rotating keys.
* Multi-Cloud organizations can consolidate key management to one central tool.
* Own an original copy of the key material for additional durability.

---
# Example KMS workflow (Azure & AWS GA)
Optional Lab: **Vault Key Management Secrets Engine**
.center[![:scale 100%](images/kms.png)]

Your instructor will provide the URL for the track.

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)
# Lab 4:
## The Vault KMIP Secrets Engine (MongoDB Encryption)

---
# Lab Environment
In this next Lab we’ll leverage Vault to act as a KMIP key management server for MongoDB’s data encryption
* First we will look at how the data collections would look to an adversary that gains access to the host
* We will then show how to protect against this threat vector with Vault’s KMIP integration

Lab: **Vault KMIP Secrets Engine for MongoDB Encryption**

Your instructor will provide the URL for the track.

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Chapter 7:
## Consul Service Mesh

---
class: col-2, title, no-footer
background-image: url(images/Consul_Wallpaper_Background-2.jpg)
<span style="color:white;font-size:36pt;font-weight:bold">Service Networking Across Any Cloud</span>

<span style="color:white;font-size:20pt">Automate network configurations, discover services, and enable secure connectivity across any cloud or runtime</span>
Automate network configurations, discover services, and enable secure connectivity across any cloud or runtime.

.center[![:scale 100%](images/slide66.png)]

---
class: col-2
# Establish the connection  - 1/2
.smaller[* The proxy of the web service uses Consul service discovery APIs to request the location of the DB
* The local agent returns the proxy’s IP address/Port of a healthy DB instance
* The local agent also returns the URI for the expected identity of the service it is connected to
* Proxies between web and database start TLS handshake to authenticate the identity]

.center[![:scale 100%](images/slide67.png)]

---
class: col-2
# Establish the connection  - 1/2
.smaller[* The DB proxy sends the authorization request to its local agent

* The local agent authorizes the connection based on locally cached intention

* Mutual TLS is established]

.center[![:scale 100%](images/Consul_2.png)]

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Chapter 8:
## Boundary

---
# Traditional workflow for access
.center[![:scale 90%](images/slide70.png)]

---
# HashiCorp Boundary
Ephemeral Access Vision
.center[![:scale 70%](images/slide71.png)]

---
# Access Model
.center[![:scale 55%](images/slide72.png)]

---
# Just-in-Time Credentials (Roadmap)
.center[![:scale 65%](images/slide73.png)]

---
class: title, shelf, no-footer, fullbleed
background-image: url(images/HashiCorp-Title-bkg.jpeg)

# Epilogue

---
# Workshop Feedback Survey
Your feedback is important to us!

The survey is short, we promise: **[HashiCorp Workshop Survey](https://docs.google.com/forms/d/1jlyG6dj6zAtvhDorELtyuX7eQwPsKGOpmsXqTQb3aXo/edit)**
