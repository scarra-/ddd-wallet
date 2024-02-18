### Wallet app written in Golang

#### Installation

1. Clone repository locally
2. Create `.env` file (copy from `.env.example`)
3. Run `make dcu`. This command will run full `docker-compose.yml` and start Go wallet service.
4. Run `make migrate` in order to run DB migrations.
5. Add `127.0.0.1 wallet.test` to `/etc/hosts` file. 

#### Functionality

- Create wallet
- Fund wallet
   - Uses pessimistic lock to lock the wallet.
   - Adjusts wallet balance and creates `fund` transaction.
- Spend wallet funds
   - Uses pessimistic lock to lock the wallet.
   - Adjusts wallet balance and creates `spend` transaction.
- Get wallet (shows current balance)

#### Folder structure

- `bin` - contains binary files.
- `build` - folder contains docker setup files.
- `cmd` - application entrypoints.
- `database` - mysql setup.
- `internal` - wallet application files
- `tests` - tests

#### Application design

In order to organize application files I used [Multitier architecture](https://en.wikipedia.org/wiki/Multitier_architecture) with four layers:
- `ui` - "User interface". Stores HTTP, CLI, etc.. handlers
- `application` - Application services. Basically they orchestrate execution of certain process like funding wallet. (They do not contain business logic, just route).
- `domain` - Heart of the application. Contains domain models (`wallet`, `transaction`), domain logic / domain services.
- `infrastructure` - Framework code, DB respositories etc..

#### Business logic design

Business logic is structured using [Domain Driven Desing](https://www.amazon.com/Implementing-Domain-Driven-Design-Vaughn-Vernon/dp/0321834577).
`Wallet.go` domain model stores all invariants. This comes handy when you want learn about business rules (domain layer code could be read like a book).

#### Domain model (core business logic)

`Wallet` contains two main functions:
- `Fund`
   1. Checks if amount is positive.
   2. Checks if operation is not duplicate (by querying transactions with same originId)
   3. Adjusts balance
   4. Creates `fund` transaction.
- `Spend`
   1. Checks if amount is positive.
   2. Checks if wallet has sufficient funds.
   3. Checks if operation is not duplicate (by querying transactions with same originId).
   4. Adjusts balance
   5. Creates `spend` transaction.
 
Transactions act as a log/receipt/idempotency mechanism for a wallet. Even though wallet has `balance` integer field - it in theory can be replayed using transactions as they contain history of wallet state changes.


#### API
Import Postman collection `./postman.json`
