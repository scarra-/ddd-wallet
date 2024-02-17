### Wallet app written in Golang

#### Installation

1. Clone repository locally
2. Create `.env` file (copy from `.env.example`)
3. Run `make dcd`. Command will setup DB container.
4. Create database by running `make dcd-init`. _(This is hack as it should work by default with `/docker-entrypoint-initdb.d/init.sql` but for some reason it doesn't.)_
5. Run `make dcu`. This command will run full `docker-compose.yml` and start Go wallet service.
6. Run `make migrate` in order to run DB migrations.
7. Add `127.0.0.1 wallet.test` to `/etc/hosts` file. 

**Comment**: _Ideally steps 3-4 should be merged with step 5. But due to the Mysql not picking up startup scripts - running them manually for now._

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
- `application` - Application services. Basically they orchestrate execution of certain process like funding wallet. (They do not contain business logic).
- `domain` - Heart of the application. Contains domain models (`wallet`, `transaction`), domain logic / domain services.
- `infrastructure` - framework code, DB respositories etc..

#### Business logic design

Business logic is structured using [Domain Driven Desing](https://www.amazon.com/Implementing-Domain-Driven-Design-Vaughn-Vernon/dp/0321834577).
`Wallet.go` domain model stores all invariants. This comes handy when you want learn about business rules (domain layer code could be read like a book).

#### API
Import Postman collection `./postman.json`
