---
slug: test-mongodb-kmip-encrypted
id: iyede3ffguja
type: challenge
title: Test MongoDB Encryption via Vault KMIP
teaser: In this challenge you will see the impact of a breach with encrypted database
  files
notes:
- type: text
  contents: |-
    With Vault's KMIP Secret Engine all set up, we can now start MongoDB with KMIP encryption enabled.

    We will role play as the same adversary as before, but show that any exfiltrated data would prove useless to the attacker as it is now encrypted ciphertext via Vault!
tabs:
- title: Terminal
  type: terminal
  hostname: base
- title: Vault Server
  type: terminal
  hostname: base
- title: MongoDB
  type: terminal
  hostname: base
- title: Vault UI
  type: service
  hostname: base
  path: /ui
  port: 8200
difficulty: basic
timelimit: 1200
---

Now, we can start MongoDB with Encryption leveraging Vault as the KMIP Key Management Server.

IMPORANT: switch to the MongoDB tab.

Start MongoD server (with encryption enabled). The process will not exit.

```bash,run
mongod --dbpath /var/lib/mongodb \
  --logpath /var/log/mongodb/mongo.log \
  --enableEncryption \
  --kmipServerName localhost \
  --kmipPort 5696 \
  --kmipServerCAFile ca.pem \
  --kmipClientCertificateFile client.pem
```

IMPORTANT: switch back to the Terminal tab

You can verify that MongoDB was able to connect to Vault's KMIP Secret engine with the following command

```bash,run
cat /var/log/mongodb/mongo.log  | grep KMIP | jq
```

The output should look like this:

```nocopy
{
  "t": {
    "$date": "2021-04-21T16:07:30.855+00:00"
  },
  "s": "I",
  "c": "STORAGE",
  "id": 24199,
  "ctx": "initandlisten",
  "msg": "Created KMIP key",
  "attr": {
    "keyId": "3ggasHBokpcWjwau4En8uGj6XO091QXL"
  }
}
```

Next, login

```bash,run
mongo
```

Now insert the same record as before

```bash,run
db.examples.insertOne(
  {
    name: "sue",
    age: 26
  }
)
```

exit mongo

```bash,run
exit
```

Now cat out the mongodb collection file, there should be several. It will likely be the "7th" file.
NOTE: It will take a minute or two for the full contents to be written to disk. (There will be multiple lines of data)

```bash,run
cat /var/lib/mongodb/collection-7*
```

Now the contents of the file are encrypted! You should not be able to see any object data or metadata in plaintext.
A critical principal in implmenting Zero Trust is to always assume a breach.
With the implmentation of Vault's KMIP secret engine, we've ensured that our customer data is secure even if your adversaries gain access to physical database hosts.
