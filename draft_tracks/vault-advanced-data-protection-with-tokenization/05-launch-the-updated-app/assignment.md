---
slug: launch-the-updated-app
type: challenge
title: Launch the updated application
teaser: Deploy the updated app to securely tokenizen customer SSNs (PII data).
notes:
- type: text
  contents: |-
    We will now deploy our updated application to use following tokenization workflow.

    ![Tokenization Workflow](https://github.com/hashicorp/field-workshops-vault/raw/main/instruqt-tracks/vault-advanced-data-protection-with-tokenization/assets/images/tokenization_workflow.png)
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
Our application is ready for prime time!

Delete the old application.
```
kubectl delete -f /root/vault_go_demo_non_tokenization/
```
Deploy the updated application
```
kubectl apply -f /root/vault_go_demo/
kubectl wait --timeout=180s --for=condition=Ready $(kubectl get pods --selector=app=vault-go-demo -o name)
sleep 5s
```
This may take a minute. You can check if the new pods are ready via "kubectl get pods"
```
kubectl get pods
```
Go back to the web app, add another record, and then look at the record in the Database View. Verify that the SSN is now tokenized.

Congratulations!

You have now ensured that customer PII/PAN data is protected and access to the associated plaintext is only available to properly authenticated and authorized applications.

Further reading: Although it was out of scope for this lab, we also integrated our application with dynamic database secrets.
These ensure that credentials are just-in-time, short-lived, and unique.
https://learn.hashicorp.com/tutorials/vault/database-secrets.