---
slug: your-first-secret
id: jns4aecue8e3
type: challenge
title: Your First Secret
teaser: Run a Vault dev server and write your first secret.
notes:
- type: text
  contents: |-
    In this challenge, you will run your first Vault server in "dev" mode and write your first secret to Vault.

    See https://www.vaultproject.io/docs/concepts/dev-server/ for more on Vault's "dev" mode.

    You will use the default KV v2 secrets engine that Vault automatically creates in "dev" mode. It is used to store multiple versions of static secrets.

    See https://www.vaultproject.io/docs/secrets/kv/kv-v2/ for more on the KV v2 secrets engine.
tabs:
- title: Vault CLI
  type: terminal
  hostname: vault-server
- title: Vault Dev Server
  type: terminal
  hostname: vault-server
- title: Vault UI
  type: service
  hostname: vault-server
  port: 8200
difficulty: basic
timelimit: 600
---
There are three tabs to use in this challenge: Vault CLI, Vault Dev Server, and Vault UI.

Vault servers can be run in "dev" and "production" modes. You can see this by getting help for the "vault server" CLI command:

Select the "Vault CLI" tab and run the following:
```
vault server -h
```
Scroll back up to see information about the two modes.

Select the "Vault Dev Server" tab

Let's run Vault in Dev Server mode. The simpest command to do this would be "vault server -dev", but we want to set the initial root token to "root" and make the Vault server bind to all IP addresses. So, please run this instead:
```
vault server -dev -dev-listen-address=0.0.0.0:8200 -dev-root-token-id=root
```

Select the "Vault UI" tab.

Log into the Vault UI with `Method` set to `Token` and `Token` set to `root`.

Select the "Vault CLI" tab.

Write a secret to the KV v2 secrets engine that the Vault dev server automatically mounted:<br>
```
vault kv put secret/my-first-secret age=<age>
```
where <age\> is your age.

Alternatively, you could create this secret in the Vault UI by selecting the "secret/" KV v2 secrets engine on the "Secrets" tab, clicking the "Create secret +" button, specifying "my-first-secret" as the path, "age" as the secret's first key, and your age as the corresponding value, and then finally clicking the "Save" button.

In the Vault UI tab, select the "secret/" KV v2 secrets engine, select the "my-first-secret" secret, and click the eye icon to see your age.

If you lied about your age, you can correct it in the Vault UI by clicking the "Create new version +" button, or with the Vault CLI by repeating the "vault kv put" command with your real age. Don't worry, nobody else can see it!