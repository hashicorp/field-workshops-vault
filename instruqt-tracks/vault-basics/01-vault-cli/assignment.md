---
slug: vault-cli
id: naobffcl3cxv
type: challenge
title: The Vault CLI
teaser: Run the Vault Command Line Interface (CLI).
notes:
- type: text
  contents: |-
    Vault Community Edition is a command line application that you can download and run from your laptop or virtual workstation.

    It is written in Go and runs on macOS, Windows, Linux and other operating systems.

    You can always download the latest version of vault here:
    https://www.vaultproject.io/downloads/
- type: text
  contents: |-
    Installing Vault on your laptop or workstation is easy. You simply download the zip file, unpack it, and place it somewhere in your PATH.

    Check out this tutorial for step-by-step instructions:
    https://learn.hashicorp.com/tutorials/vault/getting-started-install

    We've launched Vault in a Docker container in your Instruqt lab environment so that you don't need to download or install it.
tabs:
- title: Vault CLI
  type: terminal
  hostname: vault-server
difficulty: basic
timelimit: 600
---
The Vault Command Line Interface (CLI) allows you to interact with Vault servers.

Let's start with some basic vault commands, running them in the "Vault CLI" tab on the left.

Check the version of Vault running on your machine:
```
vault version
```

See the list of Vault CLI commands:
```
vault
```

Get help for the "vault secrets" command:
```
vault secrets -h
```
Note that you can also use the "-help" and "--help" flags instead of "-h".

Get help for the "vault read" command:
```
vault read -h
```
Scroll back up to get a feeling for how you would read secrets from Vault.