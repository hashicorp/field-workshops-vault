---
slug: vault-api
id: m8dxmwxagaeb
type: challenge
title: The Vault API
teaser: Use the Vault HTTP API
notes:
- type: text
  contents: |-
    In this challenge, use the Vault HTTP API to check your Vault server's status from the previous challenge.

    Also, employ the same API to retrieve your age stored in the "my-first-secret" secret.

    For additional information on the Vault HTTP API, visit https://www.vaultproject.io/api-docs/.
tabs:
- title: Vault CLI
  type: terminal
  hostname: vault-server
difficulty: basic
timelimit: 600
---
In this challenge, you will use the Vault HTTP API.

On the "Vault CLI" tab, retrive the health of the Vault server by running this command:
```
curl http://localhost:8200/v1/sys/health | jq
```

That should bring back a nicely formatted JSON document indicating that your server is initialized and unsealed.

Now, retrieve your "my-first-secret" secret with this command (on a single line):
```
curl --header "X-Vault-Token: root" http://localhost:8200/v1/secret/data/my-first-secret | jq
```

This should bring back a JSON document showing your age and metadata about the secret including its version.