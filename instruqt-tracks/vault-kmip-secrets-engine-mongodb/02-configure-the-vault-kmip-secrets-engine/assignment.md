---
slug: configure-the-vault-kmip-secrets-engine
id: wykwv2xfndvy
type: challenge
title: Configure the Vault KMIP Secrets Engine
teaser: In this challenge you will configure Vault's KMIP Secret Engine for external
  KMIP object management.
notes:
- type: text
  contents: |-
    The KMIP secrets engine allows Vault to act as a Key Management Interoperability Protocol (KMIP) server provider and handle the lifecycle of its KMIP managed objects.

    KMIP is a standardized protocol that allows services and applications to perform cryptographic operations without having to manage cryptographic material, otherwise known as managed objects, by delegating its storage and lifecycle to a key management server.

    In this section we will configure a "scope" for managing this application's objects. Within the scope, we designate roles that define access control around allowed KMIP operations.

    Finally, we will create the certificate and the key for MongoDB to authenticate to Vault's KMIP listener.
tabs:
- title: Terminal
  type: terminal
  hostname: base
- title: MongoDB
  type: terminal
  hostname: base
- title: Vault Server
  type: terminal
  hostname: base
- title: Vault UI
  type: service
  hostname: base
  path: /ui
  port: 8200
difficulty: basic
timelimit: 600
---

IMPORTANT: First, switch to the "Vault Server" tab.

Run the following command. This will start the vault process (it will not exit)

```bash,run
VAULT_UI=true vault server -dev-root-token-id=root -dev -log-level=trace
```

IMPORANT: Switch back to the "Terminal" tab.

Now Login to Vault. We've set the dev mode root token to "root".

```bash,run
export VAULT_ADDR="http://127.0.0.1:8200"
echo "export VAULT_ADDR=$VAULT_ADDR" >> /root/.bashrc
vault status
vault login root
```

Enable the KMIP Secrets Engine.

```bash,run
vault secrets enable kmip
```

Configure the KMIP Listener (5696 is the standard default port). You can also set key types and lengths.

```bash,run
vault write kmip/config listen_addrs=0.0.0.0:5696 \
  tls_ca_key_type="rsa" \
  tls_ca_key_bits=2048
```

Next, save the KMIP CA cert that we will pass to MongoDB. These Leaf/CA certs and keys allow MongoDB to authenticate to Vault.

```bash,run
vault read -format json kmip/ca | jq -r .data.ca_pem > ca.pem
```

Then, we create a scope for the HashiCup app's managed objects. Scopes partition KMIP managed objects into multiple named buckets.
Roles are managed within buckets and can be assigned various permitted KMIP operations.
We will also create a "payments" role that specifices the allowed KMIP operations that MongoDB can perform.

```bash,run
vault write -f kmip/scope/hashicups
vault write kmip/scope/hashicups/role/payments operation_all=true
```

Now, create the leaf cert and private key. Then save them as a client.pem
This cert and key will be used by Mongo do authenticate to Vault.

```bash,run
vault write -format=json \
  kmip/scope/hashicups/role/payments/credential/generate \
  format=pem > credential.json
jq -r .data.certificate < credential.json > cert.pem
jq -r .data.private_key < credential.json > key.pem
cat cert.pem key.pem > client.pem
```

With the Vault coniguration all set, we can now encrypt MongoDB.
