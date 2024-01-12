---
slug: vault-policies
id: dsoaf3px71xr
type: challenge
title: Use Vault Policies
teaser: Use Vault policies to grant different users access to different secrets.
notes:
- type: text
  contents: |-
    In this challenge, you'll define Vault policies to give two different users access to different secrets within Vault.

    See https://www.vaultproject.io/docs/concepts/policies/ to learn about Vault policies.
tabs:
- title: Vault Policies
  type: code
  hostname: vault-server
  path: /vault/policies/
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
User Management and Policy Assignment
===
Now that you have the Userpass authentication method configured, you can add policies to give different users access to different secrets.

The setup script for this challenge has exported your root token for you from your ".profile" file.

You already created a username for yourself to use with the Userpass auth method. Now, create a second user with the same command you used before, selecting a different username and password:<br>
```
vault write auth/userpass/users/<name> password=<pwd>
```

Policies Configuration
===
Next, edit the two policies on the "Vault Policies" tab, user-1-policy.hcl and user-2-policy.hcl. In user-1-policy.hcl, set all occurences of <user\> to your own username. In user-2-policy.hcl, set all occurences of <user\> to the username of the user you just added.

Save both files by clicking the disk icon above the files. (Do that once for each file.)

Assigning Policies
===
Next, you are going to add the policies to the Vault server:<br>
`vault policy write <user_1> /vault/policies/user-1-policy.hcl`<br>
`vault policy write <user_2> /vault/policies/user-2-policy.hcl`<br>
replacing <user_1\> and <user_2\> with the usernames you are using.

Now, you can assign the new policies to the users by updating the policies assigned to the users:<br>
`vault write auth/userpass/users/<user_1>/policies policies=<user_1>`<br>
`vault write auth/userpass/users/<user_2>/policies policies=<user_2>`<br>
again replacing <user_1\> and <user_2\> with the usernames you are using.

User Access Testing
===
Now, let's see what happens when you log into the Vault UI as the two different users.

Log in as yourself (the first user you created). Remember to specify the login method as "Username".

Click on the kv secrets engine. You should see the secret you previously created called "a-secret".  But if you select it, you'll see a "Not Authorized" message.

Click on the "kv" breadcrumb to go back to the previous screen.

Click the "Create secret +" button, enter <user\>/age for the path, "age" for the key in the "Version data" section of the screen, and your age for the value associated with that key. Be sure to replace <user\> with the username of the logged in user. Then click the "Save" button. You should be allowed to do this.

Logout and log back in as the second user. Try to access the first user's secret.  You should not be able to.

User-Specific Secret Creation and CLI Interactions
===
Now, while still logged in as the second user, repeat the above steps to create a secret, specifying the second user's name for <user\> in the path.  You should be allowed to do this.

If you would like to see what happens when using the Vault CLI, you can run `unset VAULT_TOKEN`, login as each user with `vault login -method=userpass username=<user> password=<pwd>`, and then try commands like the following:<br>
`vault kv get kv/<user>/age`<br>
`vault kv put kv/<user>/weight weight=150`<br>

These will succeed when <user\> matches the username of the logged in user and fail when it does not.

The bottom line for this challenge is that Vault protects each user's secrets from other users.

<strong>Congratulations on finishing the entire track!</strong>