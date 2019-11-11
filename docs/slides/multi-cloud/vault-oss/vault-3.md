name: Chapter-3
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false
# Chapter 3      
## Running a Production Vault Server

???

Chapter 3 focuses on running a production Vault server

---
layout: true

.footer[
- Copyright © 2019 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: vault-production-serves
# Running a Production Vault Server
* Running a Vault server in "Prod" mode involves multiple steps:
  * Specify configuration in a config file.
  * Start the server.
  * Initialize the server to get unseal keys and an initial root token.
  * Unseal the Vault server with the unseal keys.

???
* Describe the steps to run a production Vault server.

---
name: configuring-vault
# Configuring Vault Servers
* Vault configuration files can be specified in [HCL](https://github.com/hashicorp/hcl) or JSON.
* Common configuration settings include:
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
# Starting a Production Vault Server
* You use the `vault server` command to start a Vault Production server.
* But, you do not use the `-dev` option.

???
* Describe the command to run a Vault production server.

---
name: initializing-vault
# Initializing Vault Clusters
* Recall that a Vault cluster runs multiple Vault servers.
* Each Vault cluster must be initialized once.
* This is done with the `vault operator init` command.
* The number of key shares and the key threshold can be specified with the `-key-shares` and `key-threshold` options.
* The command returns the unseal keys and the initial root token for the cluster.

???
* Describe Vault's `init` command

---
name: unsealing-vault
# Unsealing Vault Servers
* Each Vault server must be unsealed each time it is started.
* You cannot use the server until you unseal it.
* This is done with the `vault operator unseal` command, using the unseal keys returned when you initialized the cluster.

???
* Describe Vault's `unseal` command.
---
name: lab-vault-basics-challenge-3
# 👩‍💻 Lab Challenge 3.1: Run a Vault "Prod" Server
* In this lab, you'll run your first Vault server in "Prod" mode.
* You'll learn how to initialize and unseal a Vault server.
* Instructions:
  * Click the "Run a Production Server" challenge of the "Vault Basics" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Run a Production Server" challenge of the "Vault Basics" track.
* This challenge has them examine a Vault server configuration file, run a Prod server, initialize it, and unseal it.
* Remind students to save their unseal key and root token.

---
name: vault-status-command
# Determining the Status of a Vault Server
* Use the `vault status` command to get the status of a Vault server.
* It will tell you if your Vault server is sealed or unsealed.
* It will also tell you the following:
  * the number of key shares and the key threshold
  * whether HA mode (clustering) is enabled
  * whether the server is running as a performance standby.

???
Describe the `vault status` command
