---
slug: test-the-web-app
type: challenge
title: Test the web app
teaser: Determine how our application works and what data to tokenize
notes:
- type: text
  contents: In this challenge, we will test our current web app to determine what customer PII data needs to be tokenized for compliance.
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
The web app can be viewed on the "Web App" tab. You may want to click the icon in the upper right corner of the tabs area to hide the assignment window so that you can see the entire UI of the web app. Click it again to display the assignment window.
The web app has two sections:
1. Records: This displays what a real user would see after any data has been decoded (de-tokenized).
1. Database View: This displays raw records from the database. If any items are tokenized in the database, they will not be decoded. (i.e. you will NOT see plaintext, this is our goal at the end of the lab!)
In its initial state, no records have been tokenized by Vault. So, you will essentially see the same data in both views. (The headers and order of the columns are different.) Go ahead and confirm that.
The web app was already started by the track's first challenge.

Click the "Add Record" button and add a new record with some fake data. Then check that the new record is not tokenized in the Database View.

In order to meet PCI compliance, our organization needs to ensure that all customer SSNs are tokenized. In the following steps, we will configure Vault and show how to modify our code.
