# Storyboard

The descriptions below are not prompt ready. You should convert them to
step-by-step instructions for best results using Cursor (or any other LLM).

You can see all project commands running `task --list`.

You will need to run at least the steps below. We expect you to take anywhere
between 5m to 10m setting up the environment if you are not familiar with Go
or Encore (assuming you follow the README instructions and the instructions
below).

- Create an encore account at <https://encore.dev>;
- Install the Homebrew dependencies listed on `README.md`;
- Run `encore app init`, select `Go` and insert any name you want (e.g., `cursor-demo`);
- Run `task modules-sync`;
- Run `task generate-jwt-secrets` and set the encore secrets `JwtPublicKey` and `JwtPrivateKey`
  - See <https://encore.dev/docs/go/primitives/secrets>
- Run `task serve`

## Bug - CreateOrganisation and CreateSetting flows do not properly set created_by

This happens they use utils.GetCtxActor. Golang contexts are immutable and
our source code is not receiving ctx that includes our custom set actor. Fixing
auth_user_guard.go to call utils.SetCtxActor would not be enough. We would
need to implement middlewares to achieve the desired outcome.

Alternatevely, we should use the recommended approach for the time being:

- See <https://encore.dev/docs/go/develop/auth#using-auth-data>

## Feature - Implement get setting endpoint

Currently we are able to create a setting, but we should also be able to get
them by ID. Implement this feature in apicore/settings. You will need ask cursor
to add a new SQL query in `apicore/common/datasource/db/queries/settings.sql`.

After that you must run the command SQLC command managed with Taskfile called
`task generate-sql-apicore`.

Instruct add new repository ports at the settings application layer, implement
them in the infrastructure layer and map the SQLC structs to domain entities.

Finally, you must ask Cursor to inject the repositories on settings facade and
define a new endpoint on `apicore.go` and the respective controller on the
settings presentation layer

## Feature - Implement refresh token rotation

Can you figure this out? What if cursor helps you out?
