# Advanced Data Protection - Instructor Guide

### Overview

This guide will prepare you to deliver a half-day [Avanced Data Protection focused Vault workshop](https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/adp). This workshop content is suitable for HashiCorp community members, prospects and customers. The workshop is a combination of lecture slides and hands-on Instruqt labs that introduce new users to some of Vault's enterprise features. This workshop focuses on open-source & enterprise features and is targeted for basic to intermediate users. The workshop may be presented in-person, over the web, or as a self-guided tutorial.

The workshop alternates between lectures with accompanying slides and hands-on lab exercises. New concepts that are introduced in the slides are reinforced in the labs. Participants will learn both the theory and practice of Vault. As an instructor you should be well familiar with the slide deck and the Instruqt tracks used in the workshop. Go through the course and make sure you understand all of the major concepts and lab exercises.

### Prerequisites
Prerequisites are minimal. All that is required to participate in the workshop is a web browser and Internet access. No software needs to be downloaded or installed. Self-contained lab environments run on the [Instruqt](https://play.instruqt.com/hashicorp) platform, and markdown-based slide decks are published as Github Pages websites.

The Instruqt tracks include terminal tabs that can be used to execute Vault CLI commands. They also include the Vault UI.

All instructors and TAs from HashiCorp should be sure to register themselves with Instruqt and then post a message in our Slack channel, #proj-instruqt, asking to be added to the HashiCorp organization within Instruqt. This is important even if the tracks are public since only members of the HashiCorp organization can see the useful "Skip to Challenge" button on challenges of tracks within this organization.

### Scheduling your workshop
Please add all workshops, both public and private, to the shared instruqt-workshops Google calendar as follows:

1. Create a new event on the instruqt-workshops calendar and set the name to the name of your workshop which could match a name being used by Field Marketing if it is public or the name of a specific customer and a product if it is private.
2. Set the begin and end times of the event to the begin and end times of your workshop.
3. Include the following information in the description:
    1. The name of the host (probably yourself) after "Host:"
    2. The names of presenters after "Presenters:"
    3. A list of tracks that your workshop will use after "Tracks:", listing the URL of each track on its own line.

Before saving the event, be sure to set the calendar as "instruqt-workshops" instead of your own personal calendar.

### Email Invitation
Here is some boilerplate text you can use or customize when inviting or announcing your workshop:

```
Vault Advanced Data Protection
A hands-on technical workshop

Learn how to leverage Vault for encrypting sensitive customer data. Vault is an API-driven, cloud-agnostic secrets management & encryption solution. Vault is easy for beginners and powerful for experts. Join us for a half-day training that will get you up and running quickly with Vault.

Topics covered in the workshop include:

* Vault Overview
* Interacting with Vault
* Vault Architecture
* Vault Secret Engines and Authentication Methods
* Vault Policies
* Dynamic Database Secrets
* Encryption as a Service
* Format Preserving Encryption
* Tokenization
* Filesystem & Database Encryption

The only prerequisites for this workshop are a web browser and willingness to learn.
```

### Markdown Slide Deck
The slide deck for this training is published here:

#### https://hashicorp.github.io/field-workshops-vault/slides/multi-cloud/adp

#### Navigation
Use the arrow keys (up/down or left/right) to navigate back and forth between slides.

Press the `P` key to enter presenter mode and reveal the speaker notes.

Press the `C` key to pop open an external window that will stay synced with your speaker notes window. This is useful for keeping notes on your laptop while showing the presentation on the projector.

#### RemarkJS
The slide deck for this training is written completely in [Markdown](https://guides.github.com/features/mastering-markdown/) using the [RemarkJS](https://remarkjs.com/#1) framework. This allows us to publish slide decks from a source code repository. The slide deck is easy to maintain, lightweight, and can be accessed from anywhere. Updates and changes to the deck can be made quickly and efficiently by simply editing the markdown source files. If you find any mistakes or issues with the slides please add them to our issue tracker:

https://github.com/hashicorp/field-workshops-vault/issues

### Hands-on Labs
At certain points in the slide deck there are links to the lab exercises. [Instruqt](https://instruqt.com/hashicorp) is our lab platform. While all of the Vault workshop tracks (except for https://play.instruqt.com/hashicorp/tracks/vault-advanced-data-protection-with-transform) are actually public, the slides use links to permanent Instruqt invitations that require users to register with Instruqt. We do this so that HashiCorp Field Marketing can collect the emails of users who start tracks during public or private workshops or on their own after following the links from the slides.

Participants can register with Instruqt [here](https://play.instruqt.com/signup). Users only need to provide an email and a password. They can then login via email, Google, GitHub, and Twitter.=

Students may have questions during the labs. When presenting a workshop be sure to give enough time for all your participants to go through the labs. Remember that this is probably their first time using a tool like Vault.

As mentioned in the Prerequisites section, all HashiCorp instructors and TAs for workshops should register with Instruqt and then post a message in our Slack channel, #proj-instruqt, asking to be added to the HashiCorp organization in Instruqt.

Members of the HashiCorp organization in Instruqt can hover over any challenge in any track in that organization and see a "Skip to Challenge" button. After starting a track, you can use these buttons to run the track's setup and solve scripts up to the challenge you want to skip to. When skipping to a challenge, always be sure to click the "Normal Skip" button too.

https://play.instruqt.com/hashicorp/tracks/vault-basics

Go through this track start to finish and make sure you understand all the challenges. This track covers the Vault CLI and UI, running a dev Vault server and writing a secret to it, using the Vault HTTP API, running a production Vault server, using the KVv2 secrets engine and the Userpass auth method, and Vault policies.

https://play.instruqt.com/hashicorp/tracks/vault-dynamic-database-credentials

Go through this track start to finish and make sure you understand all the challenges. This track covers Vault's database secrets engine.

Go through this track start to finish and make sure you understand all the challenges. This track covers Vault's encryption as a service capabilities as implemented in its Transit and Transform Secret Engines.

The following track is optional and focuses on Tokenization.

https://play.instruqt.com/hashicorp/tracks/vault-advanced-data-protection-with-tokenization

Go through this track start to finish and make sure you understand all teh challenges. This track covers Vault's KMIP secret engine.

https://play.instruqt.com/hashicorp/tracks/vault-kmip-secrets-engine-mongodb


### Configuring the Instruqt Pools
We recommend that you configure Instruqt pools for each Instruqt track used in this workshop 1 hour before you plan to use the track. Please see this Confluence [doc](https://hashicorp.atlassian.net/wiki/spaces/SE/pages/511574174/Instruqt+and+Remark+Contributor+Guide#InstruqtandRemarkContributorGuide-ConfiguringInstruqtPools) for instructions.

### Timing
The following schedule assumes you have a group of participants who are brand new to Vault. You should budget about four hours for this workshop including a short break. (But in virtual workshops, participants are expected to take breaks during the labs.)

If you are presenting to an intermediate or advanced group of Vault users, feel free to skip the first lab.

0:00 - 0:10 - Wait for attendees, make introductions<br>
0:10 - 0:20 - Chapter 1: Vault Overview<br>
0:20 - 0:30 - Chapter 2: Basic Operations<br>
0:30 - 0:45 - Chapter 3: Vault Auth Methods & Static Secrets<br>
0:45 - 1:15 - Lab: Vault Basics (including break time)<br>
1:15 - 1:25 - Chapter 4: Dynamic Secrets<br>
1:25 - 2:00 - Lab: Dynamic MySQL Credentials (5 minute break at end)<br>
2:00 - 2:10 - Chapter 5: Transform and Transit Secret Engines<br>
2:10 - 2:40 - Lab: Transform and Transit Secret Engines<br>
2:40 - 3:00 - Chapter 6: The Vault KMIP Secret Engine<br>
3:00 - 3:30 - Lab: MongoDB Encryption via Vault's KMIP Secret Engine<br>
3:30 - 3:45 - (Optional)Chapter 7: Consul Service Mesh<br>
3:45 - 4:00 - (Optional)Chapter 8: Secure Access Management with Boundary<br>
4:00 - 4:15 - Q&A and Wrap-up<br>
