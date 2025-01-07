# Glossary

## Setting

Configurable element that can be modified without requiring a new deployment.

## Setting Types

Settings can have one of the four following types:

- "booleanish" representing on/off switches (also known as feature flag)
- "double" representing any decimal nbnumber
- "integer" representing any whole number
- "string" representing text

## Config

Collection of related settings, similar to an online configuration file. They
can represent a frontend, a backend, a mobile application or anything else
within a project.

## Environment

Represents different deployment stages (production, staging, development) where
the same setting can have different values.

## Project

A groups of related configurations, environments, teams and, users. Typically used
to represent one application or service.

## Organisation

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
