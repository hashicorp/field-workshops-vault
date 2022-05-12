---
slug: configure-the-transform-secret-engine
type: challenge
title: Configure the Transform (Tokenization) Secrets Engine
teaser: Specify a role and transformation to Tokenize user SSNs.
notes:
- type: text
  contents: In this challenge, we will configure the Transform Secrets Engine for use with our application.
tabs:
- title: Terminal
  type: terminal
  hostname: kubernetes
- title: Web App
  type: service
  hostname: kubernetes
  path: /
  port: 9090
- title: Vault-ui
  type: service
  hostname: kubernetes
  path: /
  port: 8200
difficulty: basic
timelimit: 10000
---
First, create a role in Vault for our application. Roles are used by Vault ACLs to restrict access to the underlying transformations. (See below for an example).
```
vault write transform/role/vault_go_demo transformations=ssn
```
The transformation from above "ssn" now needs to be configured to leverage tokenization (instead of format preserving encryption).
We also specify a time-to-live "TTL" of the token, and the allowed role from above "vault_go_demo". At the end of the TTL, the tokenized data will be deleted.
More configuration options are listed here: https://www.vaultproject.io/api/secret/transform.
```
vault write transform/transformations/tokenization/ssn \
  allowed_roles=vault_go_demo \
  max_ttl=24h
```
Now, encode a test value.
```
vault write transform/encode/vault_go_demo \
  transformation=ssn \
  value="123-45-6789"
```
A tokenized value should be returned.
The token is stored encrypted along side the provided plaintext within Vault or an external postgres database.
https://learn.hashicorp.com/tutorials/vault/tokenization#setup-external-token-storage

Decode the token
```
vault write transform/decode/vault_go_demo \
  transformation=ssn value=<encoded_value_from_above>
```
Vault ACL policies can restrict which users and applications have access to this Secrets Engine.
They can be read about here: https://www.vaultproject.io/docs/concepts/policies.
As an example, we could give a user or application access to our Secrets Engine role with this policy:
```
path "transform/encode/vault_go_demo" {
  capabilities = [ "update" ]
}
path "transform/decode/vault_go_demo" {
  capabilities = [ "update" ]
}
```
Next, we will review how to modify our golang code to leverage Vault.
