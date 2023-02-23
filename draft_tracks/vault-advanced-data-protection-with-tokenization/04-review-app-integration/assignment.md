---
slug: review-app-integration
type: challenge
title: Review our application code
teaser: Leverage native libraries to integrate with HashiCorp Vault
notes:
- type: text
  contents: |-
    In this challenge, we will review the application code we used to integrate with the Vault.

    To learn more about Vault native libraries see the following link: https://www.vaultproject.io/api/libraries.
tabs:
- title: Terminal
  type: terminal
  hostname: kubernetes
- title: Golang App Code - init
  type: code
  hostname: kubernetes
  path: /root/app_code/db.go
- title: Golang App Code - tokenization
  type: code
  hostname: kubernetes
  path: /root/app_code/handlers.go
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
NOTE: This challenge is optional for educational purpose only. You may skip below if you choose.

We will be leveraging the official Go Vault client: https://github.com/hashicorp/vault/tree/master/api.

First, open up the code editor tab called "Golang App Code - init".

NOTE: Make sure to click the file name on the left to open up the editor.

In this file, we perform a few steps to login to Vault and pull dynamic database credentials.
For more legacy applications, we also provide the Vault Agent which can template a token or credential to a file.
Vault Agent: https://www.vaultproject.io/docs/agent.

On line 42 we pull the Kubernetes service account JWT token from the pod filesystem. This will be used to login to Vault.
To learn more about Kubernetes authentication in Vault see the following: https://www.vaultproject.io/docs/auth/kubernetes.
```
buf, err := ioutil.ReadFile(tokenPath)
if err != nil {
  log.Fatal(err)
}
jwt := string(buf)
fmt.Printf("K8s Service Account JWT: %v", jwt)
config := map[string]interface{}{
  "jwt":  jwt,
  "role": K8sAuthRole,
}
```
(Line 55) Next, we use our Vault client "Vclient" to write to the Vault login path using the above configuration (JWT and Role).

Once authenticated and authorized, Vault returns a token.
```
//Login
secret, err1 := Vclient.Logical().Write(K8sAuthPath, config)
if err1 != nil {
  log.Fatal(err)
}
Vclient.SetToken(secret.Auth.ClientToken)
```
On line 62, we next pull database credentials from Vault in order to connect to PostgreSQL. In this setup, we are using dynamic secrets.
Dynamic Secrets are just-in-time credentials that are unique and short-lived for every application instance or pod.
Read more about dynamic secrets here: https://www.vaultproject.io/docs/secrets/databases/postgresql.
```
//Pull dynamic database credentials
data, err := Vclient.Logical().Read("database/creds/vault_go_demo")
if err != nil {
  log.Fatal(err)
}
username := data.Data["username"]
password := data.Data["password"]
```

Next, open up the code editor tab called "Golang App Code - tokenization".

NOTE: Make sure to click the file name on the left to open up the editor.

On line 127 in the "AddRecord" function, we added the following line to tokenize the customer record's SSN before writing to Postgres.
```
response, err := config.Vclient.Logical().Write("transform/encode/vault_go_demo", data)
if err != nil {
  log.Fatal(err)
}
encval := response.Data["encoded_value"].(string)
```
Last, we need to detokenize the SSN pulled from the database so we can view as plaintext.
Within the "Records" function (Line 58) we add the following:
```
response, err := config.Vclient.Logical().Write("transform/decode/vault_go_demo", data)
if err != nil {
  log.Fatal(err)
}
decval := response.Data["decoded_value"].(string)
ssn := decval
```
With our code additions made, we can finally launch our updated application and test tokenization.