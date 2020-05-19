name: Chapter-2
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false



# Chapter 2      
## Interacting With Vault

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???
Chapter 2 focuses on interacting with Vault

---
layout: true

.footer[
- Copyright © 2019 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: Interacting-With-Vault
# Interacting With Vault

Vault provides several mechanisms for interacting with it:
* The Vault [CLI](https://www.vaultproject.io/docs/commands/index.html)
* The Vault [UI](https://learn.hashicorp.com/vault/getting-started/ui)
* The Vault [API](https://www.vaultproject.io/api-docs/index/)

???

* Chapter 2 focuses on interacting with Vault

---
name: Vault-CLI
# The Vault CLI
* The Vault CLI is a Go application.
* It runs on macOS, Windows, Linux, and other operating systems.
* You can download the latest version [here](https://www.vaultproject.io/downloads/).

???
* The Vault CLI is distributed as a Go binary.
* It runs on multiple operating systems.

---
name: installing-Vault-CLI
# Installing the Vault CLI
* Installing Vault on your laptop is easy:
  * Simply download the zip file.
  * Unpack the `vault` binary.
  * Place the binary in your path.

See this [tutorial](https://learn.hashicorp.com/vault/getting-started/install) for more details.

???
Installing Vault is easy.

---
name: some-cli-commands
# Some Basic Vault CLI Commands
* `vault` by itself will give you a list of many Vault CLI commands.
  * The list starts with the most common ones.
* `vault version` tells you the version of Vault you are running.
* `vault read` is used to read secrets from Vault.
* `vault write` is used to write secrets to Vault.

The `-h`, `-help`, and `--help` flags can be added to get help for any Vault CLI command.

???
Let's discuss some of the basic Vault CLI commands.

---
name: getting-started-with-instruqt
# Doing Labs with Instruqt
* [Instruqt](https://instruqt.com/) is the platform used for HashiCorp workshops.
* Instruqt labs are run in "tracks" that are divided into "challenges".
* If you've never used Instruqt before, start with this [tutorial](https://play.instruqt.com/instruqt/tracks/getting-started-with-instruqt).
* Otherwise, you can skip to the next slide.

???
* We'll be using the Instruqt platform for labs in this workshop.
* Don't worry if you've never used it before: there is an easy tutorial that you can run through in 5-10 minutes.
---
name: lab-vault-basics-challenge-1
# 👩‍💻 Lab Challenge 2.1: The Vault CLI
* In this lab, you'll run some of the Vault CLI commands.
* You'll do this in the first challenge, "The Vault CLI", of the [Vault Basics](https://play.instruqt.com/hashicorp/invite/qfwncq62zsxu) Instruqt track.
* You'll continue to work through this Instruqt track in chapters 2-6.

???
* Now, you can try running some Vault CLI commands yourself in the first challenge of our first Instruqt track in this workshop.
* We'll be running the Instruqt track "Vault Basics" in the chapters 2-6.

---
name:lab-vault-basics-challenge-1-instructions
# 👩‍💻 Lab Challenge 2.1: Instructions
* Start the "Vault Basics" track by clicking the purple "Start" button on the "Vault CLI" challenge of the track.
* While the challenge is loading, read the displayed text.
* Click the green "Start" button to start the "Vault CLI" challenge.
* Follow the instructions on the right side of the challenge.
* After completing all the steps, click the green "Check" button to see if you did everything right.
* You can also click the "Check" button for reminders.

???
* Give the students some instructions for starting their first challenge.
* This also includes instructions for checking that they did everything right.
* Students can also click the green "Check" button to get reminderd of what they should do next.

---
name: vault-server-modes
# Vault Server Modes
Vault servers can be run in two different modes:
* "Dev" mode that is only intended for development
* "Prod" mode that can be used in QA and production

???
* Discuss Vault's two server modes

---
name: vault-dev-server
# Vault's "Dev" Mode
* It is not secure.
* It stores everything in memory.
* Vault is automatically unsealed.
* The root token can be specified before launching.

**Please never store actual secrets on a server run in "Dev" mode.**

???
* Discuss limitations of Vault's "Dev" mode.
* Warn students to never store real secrets on a Dev server.

---
name: lab-vault-basics-challenge-2
# 👩‍💻 Lab Challenge 2.2: Run a Vault "Dev" Server
* In this lab, you'll run your first Vault server in "Dev" mode.
* You'll also write your first secret to Vault and use the UI.
* Instructions:
  * Click the "Your First Secret" challenge of the "Vault Basics" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Your First Secret" challenge of the "Vault Basics" track.
* This challenge has them run a Dev server, write a secret to the KV v2 secrets engine that was automatically enabled, and use the Vault UI.

---
name: Vault-UI
# The Vault UI
* In order to use the Vault UI, you must sign in.
* Vault supports multiple authentication methods.
* A new Vault server will only have the Token auth method enabled.
* In the challenge you just completed, you used the Token auth method and specified "root" as the token.

???

* Let's talk about the Vault UI a bit, including ways of signing into it.
* While you used the token "root" in the last challenge, you'll be running a Vault server in "Prod"  mode in the rest of the track and will have to use the token generated when you initialize that server in the next challenge.
---
name: signing-into-the-vault-ui
# Signing into the Vault UI
.center[![:scale 70%](images/vault_login_page.png)]

???
* This slide shows a screenshot of the login dialog for the Vault server.

---
name: welcome-to-vault
# The "Welcome to Vault" Tour
.center[![:scale 60%](images/take_a_tour.png)]
* When you first login to Vault, you can take a tour.

???
* Explain the "Welcome to Vault" tour.
* Explain how to get rid of it.
* Point out that they can restart the tour with the "Restart guide" menu under their user icon in the upper right corner of the UI.

---
name: vault-api-1
# The Vault API

* Vault has an HTTP API that you can use to configure Vault and manage your secrets.
* You can check Vault's health with a simple `curl` command followed by `jq` to format the JSON output.

Command:
```bash
curl http://localhost:8200/v1/sys/health | jq
```
???
* Let's talk about the Vault HTTP API

---
name: vault-api-2
# The Vault API

```json
{
  "initialized": true,
  "sealed": false,
  "standby": false,
  "performance_standby": false,
  "replication_performance_mode": "disabled",
  "replication_dr_mode": "disabled",
  "server_time_utc": 1557180149,
  "version": "1.1.1",
  "cluster_name": "vault-cluster-db6f271d",
  "cluster_id": "33e85d7c-63bb-7523-0165-9d1aee722d70"
}
```

???
Here is the output from Vault's sys/health endpoint

---
name: vault-api-3
# Authenticating Against the Vault API

* The sys/health endpoint didn't require any authentication.
* But most Vault API calls do require authentication.
* This is done with a Vault token that is provided with the `X-Vault-Token` header.

???
* Talk about how most Vault HTTP API calls will require authentication with a Vault token.

---
name: lab-vault-basics-challenge-3
# 👩‍💻 Lab Challenge 2.3: Use the Vault HTTP API
* In this lab, you'll use the Vault HTTP API.
* You'll first check the health of your Vault server.
* You'll then read your `my-first-secret` secret from Vault.
* Instructions:
  * Click the challenge called "The Vault API" in the "Vault Basics" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the challenge, "The Vault API", in the "Vault Basics" track.

---
name: chapter-2-review-questions
# 📝 Chapter 2 Review

* How can you interact with Vault?
* What options can you use to get help for Vault commands?
* What are the two Vault server modes?

???
* Let's review what we learned in this chapter.

---
name: chapter-2-review-answers
# 📝 Chapter 2 Review
* How can you interact with Vault?
  * The Vault CLI
  * The Vault UI
  * The Vault API
* What options can you use to get help for Vault commands?
  * `-h`, `-help`, and `--help`
* What are the two Vault server modes?
  * Dev and Prod

???
* Here are the answers to the review questions.
