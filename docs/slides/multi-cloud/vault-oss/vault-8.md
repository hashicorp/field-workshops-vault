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

.center[![:scale 100%](images/adobe.png)]

???
* Adobe has been running Vault Enterprise in production for two years and now the platform services over 130 teams.
* Adobe cloud services used by 67% of the Fortune 50. For example, have you ever:
* Gotten access to HBO?
* Streamed the Superbowl on a mobile device?
...well you’ve used one of the Adobe 23 products that Vault helps keep secure.

---
name: Vault-Transit-Engine

# Vault Transit Engine - Encryption as a Service
.center[![:scale 80%](images/vault-eaas.webp)]

* Vault's Transit Secrets Engine functions as an Encryption-as-a-Service.
* Developers use it to encrypt and decrypt data stored outside of Vault.

???
* Let's talk about Vault's Encryption-as-a-Service, the Transit secrets engine.
* The Challenge is: All application data should be encrypted, but deploying a cryptography and key management infrastructure is expensive, hard to develop against, and not cloud or multi-datacenter friendly
* One of the solutions for that chalange is Vault EaaS. You can use it's centralized key management to simplify encrypting data in transit and at rest across clouds and data centers
* It provides an encryption API & service that are secure, accessible and easy to implement.
* Instead of forcing developers to learn cryptography, we present them with a familiar API that can be used to encrypt and decrypt data that is stored outside of Vault.

---
name: transit-engine-benefits
# Transit Engine Benefits
.center[![:scale 80%](images/eaas_example.png)]

* Vault's Transit Engine provides developers a well-architected EaaS API so that they don't have to become encryption or cryptography experts.
* It provides centralized key management.
* It ensures that only approved ciphers and algorithms are used.

???
* Why would you want to do this?:
* you want to store encrypted data in your application’s primary data store.
* you don’t want to be bothered with the details of implementing encryption algorithms.
* you want stuff, like key rolling, authentication and auditing out-of-the-box.

---
name: Vault-Transit-Engine-1
# Vault Transit - Example Application

* In the next lab we'll use a web application that uses the Transit engine to encrypt and decrypt data.
* The app will store its encrypted data in the same MySQL database we used in Chapter 7.
* It will also get MySQL credentials from the Database secrets engine we configured in that chapter's lab.
* We'll first run the web app without Vault: No records are encrypted.
* We'll then run it with Vault enabled and see that new records are encrypted.

???
* Let me walk you through the steps you will do during the lab. 
* We will use the same MySQL database from chapter 7.
* We will get its MySQL credentials from the Database secrets engine you set up in chapter 7.
* we will first run without Vault and then with it.

---
name: web-app-screenshot
# The Web App
### Here is a screenshot of the Python web app:

.center[![:scale 70%](images/transit_app.png)]

???
* That you what you will see once you start the lab.

---
name: web-app-views
# The Web App's Views
###There are two main sections in the application.
1. **Records View**
  * The Records View displays records in plain text, showing what a logged-in user would see after any encrypted data is decrypted.

1. **Database View**
  * The Database View displays the raw records in the database, showing what SQL commands would return:

---
name: records-view
# The Web App's Records View
.center[![:scale 90%](images/records_view.png)]

* As we would expect an authorized user is able to see some of the sensitive data because the app has decrypted any encrypted data.

???
* Show the records view of the web app.

---
name: Vault-Transit-Engine-6
# The Add User Screen
* In the lab, you will add new users to the database.
.center[![:scale 60%](images/add_user.png)]

???
* By Adding  new records to the database.
* Beccause Vault is not enabled, the data will be stored as plain text.

---
name: database-record-without-vault
# Record in Database View Without Vault Enabled
* After adding a record in the lab, you will be instructed to click on the  **Database View** menu.
* You should see the exact same data that you entered.
* This means that Personally Identifiable Data (PII) is being stored in plain text in our database records.
* How can we improve this? Let's enable Vault's Transit engine and see.

---
name: encrypted-record
# A Database Record Encrypted by Vault
#### Here is a record that was encrypted by Vault's Transit engine.
.center[![:scale 80%](images/database_view_with_encrypted_record.png)]
* Note that the birth_date and social_security_number are encrypted.
???
* Show a record from the database encrypted by Vault's Transit engine.
* Point out that the birth_date and social_security_number field are encrypted as indicated by their starting with "vault:v1".
* Point out that the "v1" indicates that the first version of the encryption key was used.

---
name: encryption-key-rotation
# Rotating Transit Engine Encryption Keys
* The encryption keys of Vaults Transit Engine can be rotated.
* The newest version of the key is used to encrypt new data
* Older versions of the key can still decrypt old data but cannot decrypt new data.
* When we rotate the encryption keys, apps that use the Transit engine are unaware of any changes.
* Data can also be re-encrypted using the `rewrap` endpoint.

---
name: chapter-8-review-questions
# 📝 Chapter 8 Review
* What is the main advantage of using Vault's Transit secrets engine?
* Where does Vault's Transit Engine store encrypted data?
* Was the application still able to decrypt older encrypted records after you rotated the encryption key?
* Is it possible to tell which version of an encryption key was used?

???
* Developers can encrypt data without being experts in cryptography.
* Wherever developers want, but outside of Vault
* Yes
* Yes. The version is indicated by `v1`, `v2`, etc.


---
name: chapter-8-review-answers
# 📝 Chapter 8 Review
* What is the main advantage of using Vault's Transit secrets engine?
  * Developers can encrypt data without being experts in cryptography.
* Where does Vault's Transit Engine store encrypted data?
  * Wherever developers want, but outside of Vault
* Was the application still able to decrypt older encrypted records after you rotated the encryption key?
  * Yes
* Is it possible to tell which version of an encryption key was used?
  * Yes. The version is indicated by `v1`, `v2`, etc.

???
* Here are the answers to the review questions.

---
name: conclusion
# Vault OSS vs Enterprise!
.center[![:scale 50%](images/ent.png)]

---
name: conclusion
# Thank You for Participating!
.center[![:scale 40%](images/vault_logo.png)]

### For more information please refer to the following links:
* https://www.vaultproject.io/docs/
* https://www.vaultproject.io/api/
* https://learn.hashicorp.com/vault

???
* Thank the students for their participation
* Share some Vault links

---
name: Feedback-Survey
# Workshop Feedback Survey
* Your feedback is important to us!
* The survey is short, we promise:
  * http://bit.ly/hashiworkshopfeedback

???
* Ask them to fill out the online survey
