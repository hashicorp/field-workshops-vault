---
slug: use-userpass-auth
id: o6svotlntaex
type: challenge
title: Use the Userpass Auth Method
teaser: Enable and use the userpass authentication method.
notes:
- type: text
  contents: |-
    Now that you have a running production Vault server with a KV v2 secrets engine mounted, it's time to learn how to authenticate users.

    In this challenge, you'll mount and use the Userpass auth method which allows users to authenticate to Vault with usernames and passwords managed by Vault.

    You'll also start to learn about Vault policies that restrict which secrets different users and applications can read and write. Note that Vault policies are "deny by default"; this means a token can only read or write secrets if explicitly given permission to do so by one of the policies attached to it.
- type: text
  contents: |-
    See https://www.vaultproject.io/docs/auth/ to learn about Vault auth methods.

    See https://www.vaultproject.io/docs/auth/userpass/ to learn about Vault's Userpass auth method.
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
Now that you have a running production Vault server with a KV v2 secrets engine mounted, it's time to learn how to authenticate users.

The setup script for this challenge has exported your root token for you from your ".profile" file.

First, enable the Userpass auth method:
```
vault auth enable userpass
```

Next, add yourself as a Vault user without any policies:<br>
`vault write auth/userpass/users/<name> password=<pwd>`<br>
Be sure to specify an actual username for <name\> and a password for <pwd\> without the angle brackets.

Now, you can sign into the Vault UI by selecting the Username method and providing your username and password.

You can also login with the Vault CLI:<br>
```
vault login -method=userpass username=<name> password=<pwd>
```

Both of these login methods give you a Vault token with Vault's default policy that grants some very limited capabilities. A yellow warning message tells us that we currently have the VAULT_TOKEN environment variable set and that we should either unset it or set it to the new token. Let's unset it:
```
unset VAULT_TOKEN
```

To confirm that your new token is being used, run this command:
```
vault token lookup
```
You will see that the display_name of the current token is "userpass-<name\>" where <name\> is your username and that the only policy listed for the token is the "default" policy.

Try to read the secret you wrote to the KV v2 secrets engine in the last challenge:
```
vault kv get kv/a-secret
```
You will get an error message because your token is not authorized to read any secrets yet.  That is because Vault policies are "deny by default", meaning that a token can only read or write a secret if it is explicitly given permission to do so by one of its policies.

In the next challenge, you'll add a policy to your username that will allow you to read and write some secrets.