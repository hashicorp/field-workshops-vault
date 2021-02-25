name: vault-title-slide
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false

# Vault's AWS Secrets Engine

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???
* Let's learn about Vault's AWS secrets engine.

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: vault-aws-basics-first
# Before You Begin

* If you have not completed the [Vault OSS Workshop](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/#1) yet, please begin there and come back here when you are done.

* This workshop is specific to AWS and does not cover the basics of Vault that are covered in that workshop.

???
* Please complete the Vault OSS Workshop first.

---
name: vault-aws-table-of-contents
# Table of Contents

1. Generating Credentials for AWS using Vault
  * 🔬 Lab: AWS Dynamic Secrets with Vault
1. Authenticating to Vault using AWS to Confirm Identity
  * 🔬 Lab: AWS Auth Method for Vault

???
* These are the topics we will cover

---
name: vault-aws-instruqt-tracks
# Lab Environment Used
* This workshop uses [Instruqt](https://instruqt.com) for hands-on labs.
* Instruqt labs are run in "tracks" that are divided into "challenges".
* This workshop uses the following tracks:
    1. [AWS Dynamic Secrets with Vault](https://play.instruqt.com/hashicorp/invite/etextvhbaypp)
    1. [AWS Auth Method for Vault](https://play.instruqt.com/hashicorp/invite/bict3hvtfqx1)

???
* We will be using Instruqt for the labs.

---
name: vault-aws-dynamic-secrets-1
# Generating Credentials for AWS with Vault
* The [AWS Secrets Engine](https://www.vaultproject.io/docs/secrets/aws/) generates AWS access credentials dynamically, based on IAM policies.
* The AWS secrets engine makes working with AWS IAM much easier, since their is no clicking around
a web UI.
* Everything can be scripted and automated.

???
* The AWS secrets engine generates AWS credentials dynamically.

---
name: vault-aws-dynamic-secrets-2
# Generating Credentials for AWS with Vault
* Additionally, the process is codified and mapped to internal auth methods (such as LDAP).
* The AWS IAM credentials are time-based and are automatically revoked when the Vault lease expires.
* Vault can retrieve three types of credentials from AWS:
  * `iam_user`
  * `assumed_role`
  * `federation_token`.

???
* The process of generating AWS creds with Vault can be codified and mapped to internal auth methods such as LDAP.
* The AWS secrets engine can retrieve 3 types of credentials from AWS.

---
name: vault-aws-dynamic-secrets-3
# Generating iam_user Credentials
* Vault will create an IAM user for each lease and attach the managed IAM policies as
specified in the role to the user.
* Vault will then generate an access key and secret key forthe IAM user and return them to the caller.

???
* Here are some details on generating iam_user credentials.

---
name: vault-aws-dynamic-secrets-4
# Generating assumed_role Credentials
* Vault will call `sts:AssumeRole`.
* It will return the access key, secret key, and session
token to the caller.

???
* Here are some details on generating assumed_role  credentials.

---
name: vault-aws-dynamic-secrets-4
# Generating federation_token Credentials
* Vault will call `sts:GetFederationToken`, passing in the supplied AWS policy document.
* It will return the access key, secret key, and session token to the caller.

???
* Here are some details on generating federation_token  credentials.

---
name: vault-aws-dynamic-secrets-5
# Generating Credentials for AWS with Vault
* Most commonly, users leverage the `iam_user` Vault AWS secrets engine.
* This is what we will do in the lab.
* For more information on the
`assumed_role` and `federation_token` secrets engines, review the [STS AssumeRole](https://www.vaultproject.io/docs/secrets/aws/#sts-assumerole) and the [STS Federation Tokens](https://www.vaultproject.io/docs/secrets/aws/#sts-credentials) documentation.

---
name: lab-aws-dynamic-secrets-with-vault
# 🔬 Lab: AWS Dynamic Secrets with Vault
This lab uses the Instruqt track, [AWS Dynamic Secrets with Vault](https://play.instruqt.com/hashicorp/invite/etextvhbaypp)

#### Lab Summary:
- Enable the AWS Secrets Engine
- Configure the AWS Secrets Engine
- Generate AWS Credentials
- Revoke AWS Credentials
