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

Participating actor instances: **Tester:User && System**

Flow of events:
1. Tester starts system without IP
2. System starts up on it's own IP

##### Scenario 2
Scenario Name: System startup - UC2.

Participating actor instances: **Tester:User && System**

Flow of events:
1. Tester starts system with IP
2. System connects to given IP
3A. System fails to connect to IP and closes down
3B. System succesfully connects to IP
4. System gets updated information from the system it connected too
5. System begins it's function.

##### Scenario 3
Scenario Name: System update - UC3.

Participating actor instances: **System && Hive**

Flow of events:
1. x mins passed on local system
2. system calls all system from the hive it knows for updated list of systems in hive
3. system compares all given lists and update it's list with the majority in mind

##### Scenario 4
Scenario Name: System startup - UC4.

Participating actor instances: **System x && System Y**

Flow of events:
1. System X gets a update call from system Y
2. System X sends it's list of known systems to system Y

##### Scenario 5
Scenario Name: Mine Hashvalue - UC5.

Participating actor instances: **System && Hive**

Flow of events:
1. System chooses to change the data in it's blockchain
2. System mines for the new hash value that satify the blockchains requirements
3. System sends out the newly mined data to all other known systems in hive
4A. Hive returns illegal and the systems drops the new blockchain
4B. Hive returns verified and the system saves the new blockchain in it's database
5. the system sends out a signal to all other system that the blockchain was added.

##### Scenario 6
Scenario Name: Verify Hashvalue - UC6.

Participating actor instances: **System && Hive**

Flow of events:
1. System recieves a blockchain to verify from another system in the hive  
2. System mines the hash value with the data given and checks the hashvalue result  
3A. System finds error in the verification and sends out a signal to all other Systems  
-A. System gets enough errors from the hive that verifies the data is false  
-AA. System drops the Data  
-B. System gets enough signals from the hive that verifies the data is correct  
-BB. Systems adds the Data  
3B. System Verifies the content and send out a signal to all other Systems  
-A. System gets enough signals from the hive that verifies the data is correct  
-AA. Systems adds the Data  
-B. System gets enough errors from the hive that verifies the data is false  
-BB. System drops the Data  

##### Scenario 7
Scenario Name: Instigate Change - UC7.

Participating actor instances: **Tester:User && System**

Flow of events:
1. Tester makes a Post call though Postman to a System.
2. System Stops it's own Scenario 5
3. System start a new Scenario 5 with the Post call data

##### Scenario 8
Scenario Name: Get status - UC6.

Participating actor instances: **Tester:User && System**

Flow of events:
1. Tester makes a Get call though Postman to a System.
2. System returns a json formulated response containing it's database content.

#### b. Use case model

###### UC1
Use case name: First system startup.

Participating actors: User -> System.

Flow of events.

1. User starts System through console command without IP.
2. System begins it's function.

Pre-conditions: There are no existing systems running.

Post-conditions:
- A: The system is running and listening on it's own IPv4.

----------------------------
###### UC2
Use case name: System startup.

Participating actors: User -> System.

Flow of events:
1. User starts System through console command with IP.
2. System connects to given IP.
3. System gets updated information from the system it connected too.
4. System begins it's function.

Precondition: Atleast one System is running.

Post-conditions:
- A: System succesfully connects to the Hive.

Exit condition:
- A: System fails to connects and closes down.

--------------------------------
###### UC3
Use case name: System update.

Participating actors: System -> Hive.

Flow of events.

1. x mins passed on local system
2. system calls all system from the hive it knows for updated list of systems in hive
3. system compares all given lists and update it's list with the majority in mind

Pre-conditions: The Hive contains atleast 1 other System.

Post-conditions:
- A: The system gets it's list of IPs updated
- B: The system gets signals and begins listening on it's own IPv4

----------------------------
###### UC4
Use case name: First system startup.

Participating actors: System X && System Y.

Flow of events.

1. System X gets a update call from system Y
2. System X sends it's list of known systems to system Y

Pre-conditions: The Hive contains atleast 1 other System.

Post-conditions:
- A: System Y gets it's list of IPs updated by comparing System Xs and all others in the Hive by UC3.

----------------------------
###### UC5
Use case name: Mine Hashvalue.

Participating actors: System -> Hive

Flow of events:
1. System chooses to change the data in it's blockchain.
2. System mines for the new hash value that satify the blockchains requirements.
3. System sends out the newly mined data to all other known systems in hive.
4. Hive returns verified and the system saves the new blockchain in it's database.
5. System sends out a signal to all other system that the blockchain was added.

Pre-conditions: The Hive contains equal or more than 2 other Systems.

Post-conditions: 
- A: The Hive collectively adds the new blockchain to thier databases.
- B. Hive returns illegal and the systems drops the new blockchain

--------------------------------
###### UC6
Use case name: Verify Hashvalue.

Participating actors: System -> Hive -> System.

Flow of events:
1. System recieves a blockchain to verify from another system in the hive  
2. System mines the hash value with the data given and checks the hashvalue result  
3A. System finds error in the verification and sends out a signal to all other Systems  
-A. System gets enough errors from the hive that verifies the data is false  
-AA. System drops the Data  
-B. System gets enough signals from the hive that verifies the data is correct  
-BB. Systems adds the Data  
3B. System Verifies the content and send out a signal to all other Systems  
-A. System gets enough signals from the hive that verifies the data is correct  
-AA. Systems adds the Data  
-B. System gets enough errors from the hive that verifies the data is false  
-BB. System drops the Data  

Pre-condition: The Hive contains equal or more than 2 other Systems.

Post-condition:
- A: The Hive collectively adds the data to thier database.
- B: The Hive collectively ignores the data.

-------------------------------
###### UC7
Use cast name: Instigate Change.

Participating actors: User -> System.

Flow of events:
1. Tester makes a Post call though Postman to a System.
2. System Stops it's own Scenario 5
3. System start a new Scenario 5 with the Post call data

Pre-condition: the selected System is running

Post-condition:
- A: The system engages UC5 but with the intentional data instead of the random data.

-----------------------------------------
###### UC8
Use case name: Get status.

Participating actors: User -> System.

Flow of events:
1. Tester makes a Get call though Postman to a System.
2. System returns a json formulated response containing it's database content.

Pre-condition: System is running and contains data in it's database

Post-condition:
- A: The system responds with data
-----------------------------------------


- [Usecases](Requirements_Analysis_Appendix.md)

4: Glossary
-------------------

[1]:https://github.com/datsoftlyngby/soft2017fall-lsd-teaching-material/blob/master/assignments/01-HN%20Clone%20Task%20Description.ipynb
