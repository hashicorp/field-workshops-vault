---
slug: use-kv-secrets-engine
id: gdjubuyolr6a
type: challenge
title: Use the KV V2 Secrets Engine
teaser: Enable and use an instance of the KV v2 Secrets engine.
notes:
- type: text
  contents: |-
    In this challenge, you will mount an instance of Vault's KV v2 secrets engine on the Vault production server you started in the previous challenge.

    You will also write a secret, change its value in the Vault UI, and then see that you can still inspect its original value since the KV v2 secrets engine maintains multiple versions of secrets.
- type: text
  contents: |-
    See https://www.vaultproject.io/docs/secrets/ to learn about Vault secrets engines.

    See https://www.vaultproject.io/docs/secrets/kv/ to learn about Vault's KV secrets engine. You will use version 2 of this engine which can retain multiple versions of any secret.

    See https://www.vaultproject.io/docs/commands/secrets/enable/ for information about the `vault secrets enable` command.
tabs:
- title: Vault CLI
  type: terminal
  hostname: vault-server
- title: Vault UI
  type: service
  hostname: vault-server
  port: 8200
difficulty: basic
timelimit: 900
---
Now that you have a Vault production server running (in the background), you can mount a secrets engine and write secrets to it.

The setup script for this challenge has exported your root token for you from your ".profile" file.

First, you can mount an instance of the KV v2 secrets engine on the default path, "kv":
```
vault secrets enable -version=2 kv
```

Next, write a secret to your new secrets engine:
```
vault kv put kv/a-secret value=1234
```
Feel free to specify your own secret number instead of 1234.

Login to the Vault UI with your root token and verify that the secret "a-secret" under the "kv" secrets engine has the value you provided by clicking its eye icon.

Change the value in the UI by clicking the "Create new version +" button, typing a new value in the field with the dots, and clicking the "Save" button.  Click the eye icon for the secret to validate the change.

Still in the Vault UI, select the History menu for the secret, select "Version 1", and click the eye icon again to see the original value.