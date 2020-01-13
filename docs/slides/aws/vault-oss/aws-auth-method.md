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

- Authenticating to Vault using AWS to Confirm Identity
    - 🔬 Lab - AWS Auth Method for Vault

https://hashicorp.github.io/field-workshops-vault/slides/aws/vault-oss/aws-auth-method.html

---
name: vault-aws-instruqt-tracks
# Lab Environment Used
* This workshop uses [Instruqt](https://instruqt.com) for hands-on labs.
* Instruqt labs are run in "tracks" that are divided into "challenges".
* This workshop uses the following tracks:
    1. https://instruqt.com/hashicorp/tracks/vault-basics
    1. https://instruqt.com/hashicorp/tracks/vault-aws-auth-method

---
name: vault-aws-auth-method-1
# Authenticating to Vault using AWS to Confirm Identity

Auth methods are the components in Vault that perform authentication
and are responsible for assigning identity and a set of policies to a user.

The `aws` auth method provides an automated mechanism to retrieve a Vault token
for IAM principals and AWS EC2 instances. Securely provide your AWS-based
applications password-less access to secrets stored in HashiCorp Vault.

---
name: vault-aws-auth-method-2
# Authenticating to Vault using AWS to Confirm Identity

<br>
<br>
There are two authentication methods present for the
[`aws` auth method](https://www.vaultproject.io/docs/auth/aws.html) in Vault:
`iam` and `ec2`.

---
name: vault-aws-auth-methods-iam-1
# Vault AWS Auth Method, IAM

<br>
With the `iam` method, a special AWS request signed with AWS IAM credentials is
used for authentication. The IAM credentials are automatically supplied to AWS
instances in IAM instance profiles, Lambda functions, and others, and it is
this information already provided by AWS which Vault can use to authenticate
clients.

The `iam` auth method authenticates AWS IAM principals (including IAM users /
roles assumed from other accounts, Lambdas, or EC2 instances launched in an IAM
profile.

---
name: vault-aws-auth-methods-iam-2
# Vault AWS Auth Method, IAM

<br>
The `iam` auth method authenticates by having clients provide a specially signed
AWS API request which the method then passes on to AWS to validate the signature
and tell Vault who created it.

The actual secret (i.e., the AWS secret access key) is never transmitted over the wire,
and the AWS signature algorithm automatically expires requests after 15 minutes,
providing simple and robust protection against replay attacks.

---
name: vault-aws-auth-methods-ec2-1
# Vault AWS Auth Method, EC2

<br>
With the `ec2` method, AWS is treated as a Trusted Third Party and cryptographically
signed dynamic metadata information that uniquely represents each EC2 instance
is used for authentication. This metadata information is automatically supplied
by AWS to all EC2 instances.

The `ec2` auth method authenticates only AWS EC2 instances and is specialized to handle
EC2 instances.

.center[_Authentication flow reference diagram on next slide._]

---
name: vault-aws-auth-methods-ec2-2
# Vault AWS Auth Method, EC2

<br>
.center[![Vault AWS EC2 Auth Flow](images/vault-aws-ec2-auth-flow.png)]
.center[_Legend on next slide._]

---
name: vault-aws-auth-methods-ec2-3
# Vault AWS Auth Method, EC2

- 1) EC2 instance fetches Identity Document from EC2 Metadata Service. AWS also
provides the PKCS#7 signature of the data, and the public keys for signature
verification.
- 2) The AWS EC2 instance makes a request to Vault with the PKCS#7 signature
(contains the identity document).
- 3) Vault verifies the signature on the PKCS#7 document, certifying against AWS.
- 4) If successful, Vault returns the initial Vault token to the EC2 instance.
This token is mapped to any configured policies for the associated role.

---
name: comparing-aws-auth-methods-summary
# Authenticating to Vault using AWS to Confirm Identity (Summary)

Based on how you attempt to authenticate, Vault will determine if you are
attempting to use the `iam` or `ec2` type. As you learned, each has a different
authentication workflow, and each can solve different use cases.

*Note:* _The `ec2` method was implemented before the primitives to implement the
iam method were supported by AWS. The `iam` method is the recommended approach
as it is more flexible and aligns with best practices to perform access
control and authentication. _

---
name: lab-aws-auth-method-for-vault
# 🔬 Lab - AWS Auth Method for Vault
### [Instruqt - AWS Auth Method for Vault](https://instruqt.com/hashicorp/tracks/vault-aws-auth-method)

#### Lab Summary
- Enable the AWS Auth Method
- Configure the AWS IAM Auth Method
- Log In Using the AWS IAM Auth Method
- Log Out of Vault
