# Weaveate Lab

A lab project for me to play with [weaviate](https://weaviate.io/)

## Usage

- Pre-Requisites
  - A container compose method e.g. `podman compose` or `docker compose`
  - [aardvark-dns](https://github.com/containers/aardvark-dns)
- Start container & compile generator
  ```sh
  make
  ```
- Run generator
  ```sh
  ./generator/genertor
  ```

## Scripts

- Install [weaviate-cli](https://github.com/weaviate/weaviate-cli) in virtual env
  ```sh
  [[ ! -d ".venv" ]] && python3 -m venv .venv
  [[ -z "${VIRTUAL_ENV}" ]] && source .venv/bin/activate
  pip install weaviate-cli
  ```
- Get list of collections
  ```sh
  ./scripts/get_collections
  ```
- Get list of records
  ```sh
  ./scripts/get_records
  ```
