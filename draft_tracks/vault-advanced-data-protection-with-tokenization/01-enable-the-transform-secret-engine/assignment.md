---
slug: enable-the-transform-secret-engine
type: challenge
title: Enable the Transform (Tokenization) Secrets Engine
teaser: In this first challenge, we will enable and configure Vault's Transform Secrets Engine with data tokenization.
notes:
- type: text
  contents: |-
    Secrets engines are Vault plugins that store, generate, tokenize, or encrypt data.
    The Transform Secrets Engine functions as Vault's Encryption-as-a-Service and Tokenization solution.
    The Secrets Engine supports tokenization, format preserving encryption, and data masking transformations.
    To learn more, see https://www.vaultproject.io/docs/secrets/transform/tokenization.

    This lab usually takes 3-5 minutes to spin up.

    ![Tokenization](https://github.com/hashicorp/field-workshops-vault/raw/main/instruqt-tracks/vault-advanced-data-protection-with-tokenization/assets/images/tokenization_overview.png)
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
  The Transform Secrets Engine allows Vault to tokenize data before storing to external storage.

  In this track, you will use the Transform Secrets Engine with a Golang web app that talks to a PostgreSQL server.
  Both of these run on Kubernetes and are deployed via helm. There is also a Vault server running on Kubernetes whose root token is set to `root`.

  Vault can be viewed in the "Vault-ui" tab.
  The Golang web app can be viewed in the "Web App" tab.
  Our application code is shown in the "Golang App Code" tabs. These will be reviewed in a later challenge.

  All secrets engines must be enabled before they can be used. Check which secrets engines are currently enabled on the "Vault CLI" tab.
  ```
  vault secrets list
  ```
  Note that the Transform Secrets Engine is not enabled. Please enable it.
  ```
  vault secrets enable transform
  ```
  Other lines of busines could create their own instances of the Transform engine.