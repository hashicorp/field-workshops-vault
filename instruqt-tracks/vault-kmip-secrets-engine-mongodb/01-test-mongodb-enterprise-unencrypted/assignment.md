---
slug: test-mongodb-enterprise-unencrypted
id: xtzfvn9exu8x
type: challenge
title: Test MongoDB unecrypted
teaser: In this challenge you will see the impact of a breach with unencrypted database
  files
notes:
- type: text
  contents: |-
    In recent news, it seems as though a new data breach occurs every other week...

    With the advent of events like the SolarWinds hack (a supply chain attack), it is fair to say that we must always assume adversaries are on our network.
    They are likely working to escalate credentials and expand their footprint laterally. As a last line of defense, we must ensure our customer data is encrypted!

    In this challenge, we will role play an attacker who was able to gain remote access to our database machine.
    They do not have priviledged database credentials, only RDP/SSH access to the machine.
    You will simulate exfiltrating data from unencrypted MongoDB storage files.
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
timelimit: 600
---

WARNING: If you see "unauthorized" in the terminal, please be patient for another minute or two until the setup scripts finish.
Click the refresh circular arrow in the top right to bring up the terminal after a bit.

Welcome to the lab. Above you will see several tabs. The Vault Server and Vault UI will be used in the next challenge.
For now, lets focus on the Terminal and MongoDB tabs.

Before we introduce vault for filesystem encryption. Lets test a write to MongoDB.

IMPORANT: switch to the MongoDB tab.

Start MongoD server (without encryption enabled). The process will not exit.

```bash,run
mongod --dbpath /var/lib/mongodb \
  --logpath /var/log/mongodb/mongo.log
```

IMPORTANT: switch back to the Terminal tab.

Next login

```bash,run
mongo
```

Insert an example record

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

As you can see, the contents and metadata are in clear text on disk. Notice the plaintext "namesueage" and the unecrypted metadata in the last collection file.
An adversary in this scenario only needed to gain remote access to the physical database machine in order to exfiltrate critical customer data.
Now lets cleanup mongodb and move to the next challenge.

```bash,run
pkill -9 mongod
rm -rf /var/lib/mongodb/*
rm -rf /var/log/mongodb/*
```
