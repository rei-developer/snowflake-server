# Login and Service Server

This repository contains the code for both the login server and the service server.

## Login Server

The login server is a Nest.js application that provides authentication functionality to the service server.

### Installation

```bash
$ yarn install
```

### Running the app

```bash
# development
$ yarn run start

# watch mode
$ yarn run start:dev

# production mode
$ yarn run start:prod
```

### Test

```bash
# unit tests
$ yarn run test

# e2e tests
$ yarn run test:e2e

# test coverage
$ yarn run test:cov
```

---

## Service Server

The service server is a Golang application that provides various services such as maps and chat functionality.

This guide will walk you through the process of setting up a GoLang server in GoLand.

### Prerequisites

Before you begin, you will need to have GoLand-EAP installed on your machine.

### Configuring GoLand

#### Setting the Project GOPATH

1. Open GoLand and go to **Settings > Go > GOPATH**.
2. Under **Project GOPATH**, set the `src` directory of your project as a project source directory by clicking the `+`
   button and selecting the directory.
3. Check the box next to **Use GOPATH that's defined in system environment**.
4. Check the box next to **Index entire GOPATH**.

#### Enabling Go Modules Integration

1. Open GoLand and go to **Settings > Go > Go Modules**.
2. Check the box next to **Enable Go modules integration**.

#### Edit Configurations

1. Open the Edit Configurations dialog.
2. Add a Go Build file.
3. Select "Package" for the kind option.
4. Enter "github.com/snowflake-server/main" in the Package path field.
5. Check the "Run after build" option.
6. Specify the "snowflake-server" directory as the working directory.
7. Specify "snowflake-server" as the module.

Congratulations! You have successfully configured GoLand for GoLang server development.
