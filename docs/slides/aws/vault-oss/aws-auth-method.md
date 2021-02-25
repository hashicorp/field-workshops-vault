name: vault-title-slide
class: title, shelf, no-footer, fullbleed
background-image: url(https://hashicorp.github.io/field-workshops-assets/assets/bkgs/HashiCorp-Title-bkg.jpeg)
count: false

# Vault's AWS Auth Method

![:scale 15%](https://hashicorp.github.io/field-workshops-assets/assets/logos/logo_vault.png)

???
* Let's learn about Vault's AWS Auth Method

---
layout: true

.footer[
- Copyright © 2021 HashiCorp
- ![:scale 100%](https://hashicorp.github.io/field-workshops-assets/assets/logos/HashiCorp_Icon_Black.svg)
]

---
name: vault-aws-auth-method-1
# Authenticating to Vault using AWS
* Auth methods are the components in Vault that perform authentication.
* They are responsible for assigning identity and policies to users, programs, or machines.
* Vault's [AWS Auth Method](https://www.vaultproject.io/docs/auth/aws/) provides an automated mechanism to retrieve a Vault token for IAM principals and AWS EC2 instances.
* This method securely provides your AWS-based
applications access to secrets stored in HashiCorp Vault without passwords.
* It includes two methods: `iam` and `ec2`.

???
* What are Vault auth methods?
* What the AWS auth method does

---
name: vault-aws-auth-methods-iam-1
# The IAM Method (1)

* With the `iam` method, a special AWS request signed with AWS IAM credentials is used for authentication.
* The IAM credentials are automatically supplied to AWS
instances in IAM instance profiles, Lambda functions, and others.
* Vault can use this information provided by AWS to authenticate clients.
* The `iam` auth method authenticates AWS IAM principals including IAM users, roles assumed from other accounts, Lambdas, and EC2 instances launched in an IAM profile.

???
* Let's explore the `iam` method.

---
name: vault-aws-auth-methods-iam-2
# The IAM Method (2)
* The `iam` auth method authenticates by having clients provide a specially signed AWS API request.
* The method then passes that to AWS to validate the signature and tell Vault who created it.
* The actual AWS secret access key is never transmitted over the wire.
* The AWS signature algorithm automatically expires requests after 15 minutes, providing simple and robust protection against replay attacks.

???
* How does the `iam` method work?

---
name: vault-aws-auth-methods-ec2-1
# The EC2 Method (1)
* With the `ec2` method, AWS is treated as a Trusted Third Party.
* Cryptographically signed dynamic metadata information that uniquely represents each EC2 instance is used for authentication
* This metadata information is automatically supplied by AWS to all EC2 instances.

#### .center[ The `ec2` auth method only authenticates AWS EC2 instances.]

???
* Let's explore the `ec2` auth method.
---
name: vault-aws-auth-methods-ec2-2
# The EC2 Method (2)

.center[![Vault AWS EC2 Auth Flow](images/vault-aws-ec2-auth-flow.png)]
.center[Explanations of steps on next slide]

???
* This slide illustrates the workflow of the `ec2` method.

---
name: vault-aws-auth-methods-ec2-3
# Vault AWS Auth Method, EC2

- **Step 1:** EC2 instance fetches Identity Document from EC2 Metadata Service. AWS also provides the PKCS#7 signature of the data and the public keys for signature verification.
- **Step 2:** The AWS EC2 instance makes a request to Vault with the PKCS#7 signature that contains the identity document.
- **Step 3:** Vault verifies the signature on the document against AWS.
- **Step 4:** If successful, Vault returns the initial Vault token to the EC2 instance. The token is assigned Vault policies for the associated role.

???
* These are the 4 steps shown on the preceeding slide.

---
name: comparing-aws-auth-methods-summary
class: compact
# AWS Auth Method Summary
* Based on how you attempt to authenticate, Vault will determine if you are attempting to use the `iam` or `ec2` type.
* As you learned, each has a different
authentication workflow and each can solve different use cases.
* Please note the following points:
  * The `ec2` method was implemented before the primitives to implement the iam method were supported by AWS.
  * *The `iam` method is the recommended approach since it is more flexible and uses best practices to perform access control and authentication.*

---
name: lab-aws-auth-method-for-vault
# 🔬 Lab: AWS Auth Method for Vault
This lab uses the Instruqt track, [AWS Auth Method for Vault](https://play.instruqt.com/hashicorp/invite/bict3hvtfqx1).

#### Lab Summary:
- Enable the AWS Auth Method
- Configure the AWS IAM Auth Method
- Log In Using the AWS IAM Auth Method
- Log Out of Vault
