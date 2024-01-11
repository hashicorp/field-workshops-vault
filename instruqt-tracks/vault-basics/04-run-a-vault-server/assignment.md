---
slug: run-a-vault-server
id: ztqwxqndnnfz
type: challenge
title: Run a Production Server
teaser: Configure, run, initialize, and unseal a production mode Vault server.
notes:
- type: text
  contents: |-
    In this challenge, you will run a Vault server in "production" mode, initialize it, and unseal it. The server will get its configuration from a file that you will edit before starting it.

    See https://www.vaultproject.io/docs/configuration/ for information on configuring a Vault server.

    See https://www.vaultproject.io/docs/commands/operator/init/ for information on initializing a Vault server.

    See https://www.vaultproject.io/docs/concepts/seal/ for information on the sealing and unsealing of Vault servers.
- type: text
  contents: |-
    This lab will guide you through initializing and unsealing your Vault server with the Vault CLI.

    But, it is also possible to do both of these in the
    Vault UI.
tabs:
- title: Vault Configuration
  type: code
  hostname: vault-server
  path: /vault/config/vault-config.hcl
- title: Vault CLI
  type: terminal
  hostname: vault-server
- title: Vault Server
  type: terminal
  hostname: vault-server
- title: Vault UI
  type: service
  hostname: vault-server
  port: 8200
difficulty: basic
timelimit: 900
---
Let's run a production Vault sever.

A production Vault server gets its configuration from a configuration file.  Open the vault-config.hcl configuration file on the "Vault Configuration" tab to see what your server will use.

On the "Vault Server" tab, run a Vault server that uses the configuration file:
```
vault server -config=/vault/config/vault-config.hcl
```

Since that tab is now running the Vault server, you'll run the rest of the CLI commands on the "Vault CLI" tab.

Initialize the new server, indicating that you want to use one unseal key:
```
vault operator init -key-shares=1 -key-threshold=1
```

This gives you back an unseal key and an initial root token. <strong>Please save these for further use within this track.</strong>

In order to use most Vault commands, you need to set the "VAULT_TOKEN" environment variable, using the initial root token that the "init" command returned:<br>
`export VAULT_TOKEN=<root_token>`<br>
being sure to use your own root token instead of <root_token\>.

Please also add this to your ".profile" file with this command:
```
echo "export VAULT_TOKEN=$VAULT_TOKEN" >> /root/.profile
```

You next need to unseal your Vault server, providing the unseal key that the "init" command returned:
```
vault operator unseal
```
This will return the status of the server which should show that "Initialized" is "true" and that "Sealed` is "false".

To check the status of your Vault server at any time, you can run the `vault status` command. If it shows that "Sealed" is "true", re-run the `vault operator unseal` command.

Finally, log into the Vault UI with your root token. If you have problems, double-check that you ran all of the above commands.

<strong>Don't forget to save your root token and unseal key!</strong>