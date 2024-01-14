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
      /service_name
      ```

- **Proto Directory:**
    - Inside each service directory, you'll find a `proto` directory containing the Protocol Buffer definitions (`.proto` files).

      ```
      /service_name/proto
      ```

- **PB Directory:**
    - Additionally, there's a `pb` directory within each service directory, housing the Golang client-generated code based on the Protocol Buffer definitions.

      ```
      /service_name/pb
      ```
