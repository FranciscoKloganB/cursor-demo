# Rollout Rocket

## Glossary

### Glossary: Setting

Configurable element that can be modified without requiring a new deployment.

### Glossary: Setting Types

Settings can have one of the four following types:

- "booleanish" representing on/off switches (also known as feature flag)
- "double" representing any decimal nbnumber
- "integer" representing any whole number
- "string" representing text

### Glossary: Config

Collection of related settings, similar to an online configuration file. They
can represent a frontend, a backend, a mobile application or anything else
within a project.

### Glossary: Environment

Represents different deployment stages (production, staging, development) where
the same setting can have different values.

### Glossary: Project

A groups of related configurations, environments, teams and, users. Typically used
to represent one application or service.

### Glossary: Organisation

Top-level container for projects, teams and users. Also contains shared
information (e.g., billing) and preferences (e.g., authentication rules).

## Entity relationship and hierarchy

- A user has at least one organisation (i.e., it's own personal organisation).
- A user can belong to additional organisations that he creates in addition to the original one or gets invited to.
- An organisation may have multiple users and has at least one user at all times (e.g., the owner - which may or may not be the creator).
- An organisation may have multiple projects and as at least one project (e.g., the default project - which can be renamed).
- A project may have multiple environments and has at least two environments (e.g., called staging and production by default).
- A project may have multiple configurations and has at least one configuration.
- A setting always refers to a configuration and exists across all environments associated with the project.

## Prerequisites

- ASDF
- Docker
- Encore CLI
- Revive
- Taskfile

A set of commands to install the expected tools is provided below. It assumes
you are on MacOS, have access to Homebrew and utilize ZSH. Your mileage will
vary if you are on a different Operating System or use a differen command line
interpreter implementation (e.g., Bash, Elvish, Fish).

  ```bash
  brew install coreutils curl git asdf docker docker-compose encoredev/tap/encore go-task mockery revive sqlc ;

  echo -e "\n. $(brew --prefix asdf)/libexec/asdf.sh" >> ${ZDOTDIR:-~}/.zshrc ;

  asdf plugin add golang https://github.com/asdf-community/asdf-golang.git ;
  asdf install golang ;
  ```

## Running the project

You can launch the entire company stack using docker containers with a single
command:

```bash
task serve
```

While `task serve` (alias to `encore run`) process is running, open
[http://localhost:9400/](http://localhost:9400/) to access Encore's
[developer dashboard](https://encore.dev/docs/go/observability/dev-dash).

Here you can see traces for all requests that you made while using the frontend,
see your architecture diagram, and view API documentation in the Service Catalog.

## Running other development and CI tasks

To list all available project tasks and their descriptions run:

```bash
task --list
```

## Encore

### Deploying to self-hosted infrastructure

See the [self-hosting instructions](https://encore.dev/docs/go/self-host/docker-build)
for how to use `encore build docker` to create a Docker image and configure it.

### Deploying with Encore Cloud Platform

Deploy your application to a free staging environment in Encore's development cloud using `git push encore`:

```bash
git add -A .
git commit -m 'Commit message'
git push encore
```

You can also open your app in the [Cloud Dashboard](https://app.encore.dev) to
integrate with GitHub, or connect your AWS/GCP account, enabling Encore to
automatically handle cloud deployments for you.

#### Linking the application to GitHub

Follow these steps to link your app to GitHub:

1. Create a GitHub repo, commit and push the app.
2. Open your app in the [Cloud Dashboard](https://app.encore.dev).
3. Go to **Settings ➔ GitHub** and click on **Link app to GitHub** to link your
app to GitHub and select the repo you just created.
4. To configure Encore to automatically trigger deploys when you push to a
specific branch name, go to the **Overview** page for your intended environment.
Click on **Settings** and then in the section **Branch Push** configure the
**Branch name** and hit **Save**.
5. Commit and push a change to GitHub to trigger a deploy.

[Learn more in the docs](https://encore.dev/docs/platform/integrations/github)
