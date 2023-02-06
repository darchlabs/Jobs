# Jobs

## Reviewers

- github.com/cagodoy
- github.com/mtavano
- github.com/nicobarra

## Definition

Jobs is a service designed to manage executions over smart contracts methods based on time and on-chain logic, in an automatized way.

### Context

There are multiple options for jobs (or keepers) in order to manage off-chain smart contract interactions nowadays. Some of them are very expensive, or some have a lot of friction when trying to setup and configure it. Also, there is a lack of well-delivered information about the jobs state and the interactions.

### Why Do We Do This?

The purpose of this module is to offer a cheaper, frictionless and autonomous (self-hosted) option for the user in order him to pay the real computing costs of the jobs he is running and not an expensive price derivated from changes in the token price.

In addition, we also want to make it easier to monitor and manage jobs, so it provides a flexible way for creating and managing (updating, stopping, starting or deleting if necessary) jobs and also for accessing to the logs of them.

### Proposed Solution

In the first MVP that will be built by `Darch Labs`, the solution is implemented in the `Golang` programming language.
The code can be divided in different parts managing different tasks that will be connected between them:

- Interfaces definition for implementations, and integrations for them

- An API for the user (from the fronted) to interact with one of the module actions

- Read and Write in a DB the Jobs providers implemented available, the smart contracts being used by the user

- Write in a DB the state of the jobs providers being used, just like using the `synchronizer V2` for getting the smart contracts interactions and writing it in the DB

- A dashboard that will the DB data in the frontend

### Diagrams

- Architecture diagram:

https://app.diagrams.net/#G1PxvFkkQKAgMXKkp0dnIGnzweYV3uBXMT
