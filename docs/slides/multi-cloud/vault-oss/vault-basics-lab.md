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
# Doing Labs with Instruqt
* [Instruqt](https://instruqt.com/) is the platform used for HashiCorp workshops.
* Instruqt labs are run in "tracks" that are divided into "challenges".
* If you've never used Instruqt before, start with this [tutorial](https://play.instruqt.com/instruqt/tracks/getting-started-with-instruqt).
* Otherwise, you can skip to the next slide.
* This "Vault Basics" lab covers concepts introduced in chapters 2-6.

???
* We'll be using the Instruqt platform for labs in this workshop.
* Don't worry if you've never used it before: there is an easy tutorial that you can run through in 5-10 minutes.

---
name: lab-vault-basics-challenge-1
# 👩‍💻 Challenge 1: The Vault CLI
* In this lab, you'll run some of the Vault CLI commands.
* You'll do this in the first challenge, "The Vault CLI", of the [Vault Basics](https://play.instruqt.com/hashicorp/invite/qfwncq62zsxu) Instruqt track.

???
* Now, you can try running some Vault CLI commands yourself in the first challenge of our first Instruqt track in this workshop.
* This lab covers concepts introduced in chapters 2-6.

---
name:lab-vault-basics-challenge-1-instructions
# 👩‍💻 Challenge 1: Instructions
* Start the "Vault Basics" track by clicking the purple "Start" button on the "Vault CLI" challenge of the track.
* While the challenge is loading, read the displayed text.
* Click the green "Start" button to start the "Vault CLI" challenge.
* Follow the instructions on the right side of the challenge.
* After completing all the steps, click the green "Check" button to see if you did everything right.
* You can also click the "Check" button for reminders.

???
* Give the students some instructions for starting their first challenge.
* This also includes instructions for checking that they did everything right.
* Students can also click the green "Check" button to get reminders of what they should do next.

---
name: lab-vault-basics-challenge-2
# 👩‍💻 Challenge 2: Run a Vault "Dev" Server
* In this lab, you'll run your first Vault server in "Dev" mode.
* You'll also write your first secret to Vault and use the UI.
* Instructions:
  * If the track does not do it for you, click the "Your First Secret" challenge of the "Vault Basics" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Your First Secret" challenge of the "Vault Basics" track.
* This challenge has them run a Dev server, write a secret to the KV v2 secrets engine that was automatically enabled, and use the Vault UI.

---
name: lab-vault-basics-challenge-3
# 👩‍💻 Challenge 3: Use the Vault HTTP API
* In this lab, you'll use the Vault HTTP API.
* You'll first check the health of your Vault server.
* You'll then read your `my-first-secret` secret from Vault.
* Instructions:
  * If the track does not do it for you, click the "The Vault API" challenge in the "Vault Basics" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the challenge, "The Vault API", in the "Vault Basics" track.

---
name: lab-vault-basics-challenge-4
# 👩‍💻 Challenge 4: Run a Vault "Prod" Server
* In this lab, you'll run your first Vault server in "Prod" mode.
* You'll learn how to initialize and unseal a Vault server.
* Instructions:
  * If the track does not do it for you, click the "Run a Production Server" challenge of the "Vault Basics" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Run a Production Server" challenge of the "Vault Basics" track.
* This challenge has them examine a Vault server configuration file, run a Prod server, initialize it, and unseal it.
* Remind students to save their unseal key and root token.

---
name: lab-vault-basics-challenge-5
# 👩‍💻 Challenge 5: Use the KV v2 Secrets Engine
* In this lab, you'll enable and use the KV v2 secrets engine.
* Note that the path will be `kv` instead of `secret`.
* Instructions:
  * If the track does not do it for you, click the "Use the KV V2 Secrets Engine" challenge of the "Vault Basics" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Use the KV V2 Secrets Engine" challenge of the "Vault Basics" track.
* This challenge has them enable an instance of the KV v2 secrets engine.
* Emphasize that the path will be `kv` instead of `secret` as was the case for the challenges with the Dev mode server.

---
name: lab-vault-basics-challenge-6
# 👩‍💻 Challenge 6: Userpass Auth Method
* In this lab, you'll enable and use the Userpass auth method.
* Instructions:
  * If the track does not do it for you, click the "Use the Userpass Auth Method" challenge of the "Vault Basics" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Use the Userpass Auth Method" challenge of the "Vault Basics" track.
* This challenge has them enable an instance of the Userpass auth method.
* It also demonstrates that Vault is "deny by default" since the Userpass user that they create will not have any access to secrets yet.

---
name: lab-vault-basics-challenge-7
# 👩‍💻 Challenge 7: Vault Policies
* In this lab, you'll use Vault policies to grant different users access to different secrets.
* Instructions:
  * If the track does not do it for you, click the "Use Vault Policies" challenge of the "Vault Basics" track.
  * Then click the green "Start" button.
  * Follow the challenge's instructions.
  * Click the green "Check" button when finished.

???
* Instruct the students to do the "Use Vault Policies" challenge of the "Vault Basics" track.
* This challenge has them create a second user and create and associate policies with their 2 users.
* It then has them valdiate that each user can only access their own secrets.
