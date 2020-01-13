---
name: vault-aws-basics-first

# Before You Begin

<br>
<br>
If you have not completed the [Vault OSS Basics Workshop](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/vault-oss/#1) yet, please begin there and come back here when you are done.

This workshop is specific to AWS and will not cover the basics that are covered in that workshop.

---
name: vault-aws-table-of-contents

# Table of Contents

- Generating Credentials for AWS using Vault
    - 🧪 Lab - AWS Dynamic Secrets with Vault

https://hashicorp.github.io/field-workshops-vault/slides/aws/vault-oss/index.html

---
name: vault-aws-instruqt-tracks
# Lab Environment Used
* This workshop uses [Instruqt](https://instruqt.com) for hands-on labs.
* Instruqt labs are run in "tracks" that are divided into "challenges".
* This workshop uses the following tracks:
    1. https://instruqt.com/hashicorp/tracks/vault-basics
    1. https://instruqt.com/hashicorp/tracks/vault-aws-dynamic-secrets

---
name: vault-aws-dynamic-secrets-1
# Generating Dynamic Credentials for AWS with Vault
<br>
The AWS secrets engine generates AWS access credentials dynamically, based on IAM policies.

The AWS secrets engine makes working with AWS IAM much easier, since their is no clicking around
a web UI, and everything can be scripted/automated.

---
name: vault-aws-dynamic-secrets-2
# Generating Dynamic Credentials for AWS with Vault
<br>
Additionally, the process is codified and mapped to internal auth methods (such as LDAP). The
AWS IAM credentials are time-based and are automatically revoked when the Vault lease expires.

Vault supports three different types of credentials to retrieve from AWS: `iam_user`,
`assumed_role`, `federation_token`.

---
name: vault-aws-dynamic-secrets-3
# Generating Dynamic Credentials for AWS with Vault
<br>
`iam_user`: Vault will create an IAM user for each lease, attach the managed IAM policies as
specified in the role to the user. Vault will then generate an access key and secret key for
the IAM user and return them to the caller.

---
name: vault-aws-dynamic-secrets-4
# Generating Dynamic Credentials for AWS with Vault
<br>
`assumed_role`: Vault will call `sts:AssumeRole` and return the access key, secret key, and session
token to the caller.

`federation_token`: Vault will call `sts:GetFederationToken` passing in the supplied AWS policy
document and return the access key, secret key, and session token to the caller.

---
name: vault-aws-dynamic-secrets-5
# Generating Dynamic Credentials for AWS with Vault
<br>
Most commonly, users leverate the `iam_user` Vault AWS secrets engine. For more information on the
`assumed_role` and `federation_token` secrets engines, review the
[STS Federation Tokens](https://www.vaultproject.io/docs/secrets/aws/index.html#sts-credentials)
 documentation.

---
name: lab-aws-dynamic-secrets-with-vault
# 🔬 Lab - AWS Dynamic Secrets with Vault
### [Instruqt - AWS Dynamic Secrets with Vault](https://instruqt.com/hashicorp/tracks/vault-aws-dynamic-secrets)

#### Lab Summary
- Enable the AWS Secrets Engine
- Configure the AWS Secrets Engine
- Generate AWS Credentials
- Revoke AWS Credentials