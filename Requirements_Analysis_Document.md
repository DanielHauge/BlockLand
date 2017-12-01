Requirements Analysis Document (RAD)
==================
This is a Requirements Analysis Document by studentgroup: Emmely, Kristian, Daniel.
The analysis is for a Blockchain. Assignments description: [Unknown][Unknown].
This document come with an appendix, which is included in this repository. It can be found here: [Unknown](Requirements_Analysis_Appendix.md). The appendix has all the usecases.(But for hand-in Unknown it is included in this document)
> **Content:**
> -  1: [Introduction](#1-introduction)
>> - A. [Purpose of the system](#a-purpose-of-the-system)
>> - B. [Scope of the system](#b-scope-of-the-system)
>> - C. [Objectives and success criteria of the project](#c-objectives-and-success-criteria-of-the-project)
>> - D. [Definitions, acronyms, and abbreviations](#d-definitions-acronyms-and-abbreviations)
>> - E. [References](#e-references)
>> - F. [Overview](#f-overview)
> - 2: [Current system](#2-current-system)
> - 3: [Proposed system](#3-proposed-system)
>> - A. [Overview](#a-overview)
>> - B. [Functional requirements](#b-functional-requirements)
>> - C. [Nonfunctional requirements](#c-nonfunctional-requirements)
>> - D. [Systemmodels](#d-systemmodels)
> - 4: [Glossary](#4-glossary)

----------


1: Introduction
-------------------

### A. Purpose of the system
____________
The system is made to allow itself to communicate with other copies of itself, while using hash functions with the blockchain model, to make a self regulating autonomous system, while being able to be setup from any location.

### B. Scope of the system
____________
The system needs to be able to autonomously verify each action individual systems does are legal, while also keep track of every system that's currently active, while doing so over the internet or locally, yhe system will also share all data among other systems, while allowing new systems to join the network as fast as possible.
The system will also include a API call function that allows outside users to initiate change or retrieve status of the "hive", through the use of a browser or postman.

### C. Objectives and success criteria of the project
____________
The project must be able to avoid any tempering of the blockchains data while a minimum of <50% systems are functioning correctly, the proecjt needs to allow users to forcefully override and instigate a change to the blockchain for observational and proof of concept purposes.

### D. Definitions, acronyms, and abbreviations
____________
We will use the term "System" whenever we talk of a singular active autonomous unit or node, while we will use the term "Hive" when the sum of all active Systems are in play.
When we use the term Database, we will be refering to the local systems database, as each system will have thier own personal database.
We will use "Product" when we talk of the entire project as a solution.

### E. References
____________
- Unknown

### F. Overview
____________
The product will be a student project focused on making a decentralized product, this will be achieved through the use of P2P networking, the product must also be self regulated through the use of blockchains, in order to make it self verifying each change in the product.

----------


2: Current system
-------------------
The system called Bitcoing, is a only P2P network whoes purpose it to calculate hash values, it rewards each participants for discovering new hash values whenever the first finds a hash value and gets it verified by all other participents through the use of blockchains.


----------


3: Proposed system
-------------------
### A. Overview
____________
The proposed product would be a P2P system with a API and a database, that enables it to communicate with other systems in a collective P2P Hive, this is to allow it to regulate and propose changes to the "hives" database for blockchain purposes.
### B. Functional requirements

##### API system Functional requirements
- The system must be able to Mine a hashvalue
- The system must be able to submit a mined hashvalue
- The system must be able to verify a hash value
- The system must be able to update it's database
- The system must be able to share it's database content (both blockchain and connected systems)
- The system must be able to connect to another system

------------------------------------------
##### Human(API) User Functional requirements
- The User must be able to instigate a new mining operation
- The User must be able to receive update on system status

### C. Nonfunctional requirements
____________
#### a. Usability
- We will not put too much effort into the usability of the system.
#### b. Reliability
- The system cannot lose any information which has been sent from another system.
- the system cannot allow other systems to instigate false hashvalues in the shared blockchain.
#### c. Performance
- The system needs to be able to handle multiple users at one time.
#### d. Supportability
- (extra)The system needs to work on popular browsers such as Google Chrome, Mozilla Firefox, Safari, Internet Explorer 9.
#### e. Implementation
- The system needs to have a API for both the systems interaction with other systems, and for user interaction.
#### f. Interface
- The system needs to interact with a database for storage of information.
- The system needs to interact with a API for the Hive and outside users.
#### g. Packing
- No Packing(physical) requirements.
#### h. Legal
- No legal requirements

### D. Systemmodels
____________
#### a. Scenarios

##### Scenario 1
Scenario Name: First system startup - UC1.

Participating actor instances: **Tester:System**

Flow of events:
1. Tester selects login in the menu
2. HackerNewsClone presents a form for creating a new account
3. Tester completes filling in the form and then submits.
4. (B) Tester is prompted that the user name has been taken.

##### Scenario 2
Scenario Name: Login - UC2.

Participating actor instances: **Tester:HN User & HackerNewsClone:System.**

Flow of events:
1. Tester selects Login
2. HackerNewsClone presents a form for login
3. Tester completes the form by inputting username and password and then submits.
4. (A) HackerNewsClone responds by going back to the previous page Tester was on. New links are added to the menu, links for an introduction to HN, threads and edit profile information.

##### Scenario 3
Scenario Name: Submit Thread - UC4.

Participating actor instances: **Tester:HN User & HackerNewsClone:System.**

Flow of events:
1. Tester selects submit
2. HackerNewsClone responds with a submit form
3. Tester fill in the form giving a Title and URL linking to the news article and submits
4. (A) HackerNewsClone responds with a thread successfully submitted message

##### Scenario 4
Scenario Name: API-Query - UC8.

Participating actor instances: **Helge's Program:Simulator program & HackerNewsClone:System.**

Flow of events:
1. Helge's Program requests the HackerNewsClone's API for the latest ingested thread or comment.
2. (C) HackerNewsClone's API responds with status: System is unreachable or offline.

##### Scenario 5
Scenario Name: Story Discussion - UC2, UC5, UC6.

Participating actor instances: **Anders:HN User, Anna:HN User & HackerNewsClone:System.**

Flow of events:
1. Anders logs into Hacker News on his break at work. He sees Anna's story about a subject he is very passionate about. Since he is already logged in he opens up the story to view the discussion.
2. The system presents Anders with details of the story, view of the story discussion and a form to make a comment.
3. Anders fills in a comment and submits it.
4. Later that evening Anna sits in front of the TV she picks up her smart phone and open Hacker News. She can see that more comments have been added to her story. She clicks on the story to view more details.
5. The system presents Anna with details of the story, and view of the discussion and a form to make a comment. Anna sees Anders comment, she chooses to answer by clicking reply on the comment.
6. Because Anna is not yet logged in, the system presents the Anna with a form to log in or create a new account.
7. Anna fills in her user name and password to log in.
8. The system presents the story details, and a form to reply. 
9. Anna fills in an answer and submits.
10. The system presents the Anna with the page of the story of the discussion she replied to.


#### b. Use case model

###### UC1
Use case name: Create New Account.

Participating actors: HN User - System.

Flow of events.

1. User selects Login in the menu.
2. The system presents a form for creating a new account.
3. The user completes the form by filling in username and password. The user then submits the form.

Pre-conditions: The user is not already logged into an account.

Post-conditions:
- A: The system responds by going back to the previous page the user was on. A link with the user name and number of karma points shows up in the menu. Links with the title welcome and threads also shows up in the menu.
- B: User is prompted that the user name has been taken.
- C: Invalid username or password

----------------------------
###### UC2
Use case name: Login.

Participating actors: HN User -> System.

Flow of events:
1. User selects Login
2. The system presents a form for login.
3. The user completes the form by filling in username and password. The user then submits the form.

Precondition: The user has an account, and is not already logged in to an account.

Exit condition:
- A: The system responds by going back to the previous page the user was on. New links are added to the menu, links for an introduction to HN, threads and edit profile information.
- B: User is prompted that the user name or password was incorrect.

--------------------------------
###### UC3
Use case name: Update User Information.

Participating actors: HN User -> System.

Flow of events:
1. Users select the link to user information.
2. The system responds by presenting a form to the HN User, with the editable fields for the accounts information.  and noneditable information about the HN User.
3. The HN User make changes and submits the form.

Pre-conditions: The users are logged in.

Post-conditions: 
- A: The system responds by presenting the form with the updated information.
- B: The system present to the user that something went wrong.

--------------------------------
###### UC4
Use case name: Submit a thread.

Participating actors: HN User -> System.

Flow of events:
1. The user selects the link to submit in the main menu.
2. The system responds with a submit form.
3. The user fills in the form giving a Title and URL linking to the news article. Leave URL blank to submit a question for discussion. If there is no URL, the text (if any) will appear at the top of the thread. Titles beginning with "Show HN" will appear under show.

Pre-condition: The users are logged in.

Post-condition:
- A: The system responds with a thread successfully submitted message
- B: The system present to the user that something went wrong.

-------------------------------
###### UC5
Use cast name: Comment on a thread.

Participating actors: HN User -> System.

Flow of events:
1. The user selects to comment on a specific a thread in the display of stories.
2. The system responds with a form and details about the thread, and previous comments. The thread information is made up by the title, the number of points, creator of the thread, days since threaded, number of comments. The comments are presented with the username of comment submitter, days since threaded, the comments and a link to reply.
3. The user fills in the comment in the form and submits it.

Pre-condition: The users are logged in.

Post-condition:
- A: The system responds with showing the page of the comments.
- B: The system present to the user that something went wrong.

-----------------------------------------
###### UC6
Use case name: Reply to a thread comment.

Participating actors: HN User -> System.

Flow of events:
1. The user selects the comment section of a thread.
2. The system responds with a form and details about the thread and previous comments. The thread information is made up by the title, the number of points, creator of the thread, days since threaded, number of comments. The comments are presented in a hierarchy 
3. The user selects reply to a comment.with the username of comment submitter, days since threaded, the comments and a link to reply.
4. The system responds with a form the title of the thread and the parent comment to reply to.
5. The HN user writes the reply by filling in the form and submitting. 

Pre-condition: The users are logged in.

Post-condition:
- A: The system responds with showing the page of the comment.
- B: The system present to the user that something went wrong.
-----------------------------------------
###### UC7
Use case name: Create Thread.

Participating actors: Simulator Program -> System.

Flow of events:
1. The user submits a post to the system API. The post consists of a Title, Text or URL.

Pre-condition: none

Post-condition:
- A: The system API responds with a successful response and details about the created post.
- B: The system API responds with error and details about the error.
- C: The system API responds with status: System is unreachable, most likely offline.

--------------------------------
###### UC8
Use case name: Query system.

Participating actors: Simulator Program -> System.

Flow of events:
1. The user requests the systems API for the latest ingested thread or comment.

Precondition: None

Post-condition:
- A: The system API responds with a thread or comment
- B: The system API responds with status: System is upgrading
- C: The system API responds with status: System is unreachable, most likely offline.
------------------------------
###### UC9

Use case name: API Create Comment.

Participating actors: Simulator Program -> System.

Flow of events:

1. The user submits a comment with a designated thread.

Pre-condition: 

1. An existing thread for the comment to be added too.

Post-condition:

- A: The system API responds with a successful response and details about the created post.
- B: The system API responds with error and details about the error.
- C: The system API responds with status: System is unreachable, most likely offline.


- [Usecases](Requirements_Analysis_Appendix.md)

4: Glossary
-------------------

[1]:https://github.com/datsoftlyngby/soft2017fall-lsd-teaching-material/blob/master/assignments/01-HN%20Clone%20Task%20Description.ipynb
