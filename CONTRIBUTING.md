# Contributing to Chainlink Protos

<!-- TOC -->

- [Team Overview](#team-overview)
- [How to Contribute](#how-to-contribute)
  - [Filing a PR on smartcontractkit/spec-generator](#filing-a-pr-on-smartcontractkitspec-generator)
  - [Preparing a release](#preparing-a-release)
  - [Merging Version Packages PR](#merging-version-packages-pr)

<!-- TOC -->

## Team Overview

The Deployment Automation team is responsible for the development and maintenance of this repo. The GitHub
team [@smartcontractkit/deployment-automation](https://github.com/orgs/smartcontractkit/teams/deployment-automation)
are the primary code owners and reviewers for this repo.

## How to Contribute

To contribute, you must:

- Open a pull request (PR) with your changes.
- Request a review from the Deployment Automation
  team ([@smartcontractkit/deployment-automation](https://github.com/orgs/smartcontractkit/teams/deployment-automation))
  to
  ensure adherence to code and design standards.
- Ensure your PR passes all continuous integration checks and adheres to the contribution guidelines specific to each
  repository.

### Filing a PR on smartcontractkit/chainlink-protos

Before creating a PR with your change, you should generate a "changeset" file.

Let's assume that you've made some local changes.
Before filing a PR you need to generate a "changeset" description required for
the automated release process. Follow the steps below:

- Run `pnpm changeset` in the git top level directory.
- Answer remaining questions. At the end, you will have a new
  `.changeset/<random-name>.md` file generated.
- Now you need to commit and push your changes
- Create a Pull request which includes your code change and generated
  "changeset" file.

### Preparing a release

After merging your PR, a changesets CI job will create or update a "Version Packages" PR
like [this one](https://github.com/smartcontractkit/chainlink-protos/pull/31) which contains a release bump.

### Merging Version Packages PR

Now you can Approve/Request approval and Merge the PR from the previous step. After merging, it will kick off the
release workflow and that will release a new version and push tags automatically. You can navigate to
the [tag view](https://github.com/smartcontractkit/chainlink-protos/tags), to check if the latest release is
available.
