[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/kyamalabs/proto/main.svg)](https://results.pre-commit.ci/latest/github/kyamalabs/proto/main)

Kyama Games - Proto
=================

## Purpose

This is a centralized repository designed for sharing Protocol Buffer definitions across various Kyama Games services. This repository facilitates seamless communication and data exchange between different services by providing a standardized set of data structures.

## Directory Structure

The directory structure follows a consistent pattern for each service:

- **Service Name Directory:**
    - Contains the specific service's directory.

      ```
      /proto/service_name
      ```

- **Proto Directory:**
    - Inside each service directory, you'll find a `proto` directory containing the Protocol Buffer definitions (`.proto` files).

      ```
      /proto/service_name/proto
      ```

- **PB Directory:**
    - Additionally, there's a `pb` directory within each service directory, housing the Golang client-generated code based on the Protocol Buffer definitions.

      ```
      /proto/service_name/pb
      ```

## Service Integration

To integrate your service, we recommend making use of [kyamalabs/action-open-pull-request-across-repositories](https://github.com/kyamalabs/action-open-pull-request-across-repositories) gitHub action as this repository does not accept direct proto related pull requests.

You could use the workflow below as a starting point to integrate your service to this repository:
```yaml
name: Update Shared Proto Definitions

on:
  push:
    branches:
      - main
  workflow_dispatch:

jobs:
  update_shared_proto:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: Update Shared Proto Definitions
        uses: kyamalabs/action-open-pull-request-across-repositories@v0.1.23
        with:
          destination_repository: 'kyamalabs/proto'
          source_folders: 'api/'
          destination_folders: 'proto/auth/'
          destination_head_branch: 'auth-proto-update'
          pr_title: 'feat: update auth service proto'
          commit_message: 'Update auth service proto'
          destination_base_branch: 'main'
        env:
          GITHUB_TOKEN: ${{ secrets.PROTO_UPDATE_TOKEN }}
          API_TOKEN_GITHUB: ${{ secrets.PROTO_UPDATE_TOKEN }}
```
