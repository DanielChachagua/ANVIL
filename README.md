# ANVIL CLI

Anvil is a powerful **multi-language** CLI tool backed by `spf13/cobra` designed to instantly generate boilerplate code for new domain entities following a structured architecture and port-and-adapters methodologies across Go, TypeScript, and Python.

## Global Installation 

To use Anvil in any project, install it globally using Go. Run the following from the root of this project:

```bash
go install ./cmd/anvil
```
*(Make sure `$(go env GOPATH)/bin` is in your system's PATH)*

Once installed globally, you can execute `anvil` from anywhere on your terminal.

## Initializing Application Configurations

Anvil uses an `anvil.json` file inside the root of your target project to determine which folders map to what layers (e.g., `ports`, `repositories`, `controllers`), which language is being scaffolded, and the preferred extensions.

To generate this baseline configuration file, simply go to your application logic map and run:

```bash
anvil init .
```
This generates the **Domain-Oriented Hexagonal Architecture** defaults in native `Go` and automatically scaffolds the base project directories.

### Languages & Framework Templates

You can use the `--lang` (or `-l`) flag to automatically construct configurations for different languages. Anvil ships with pre-built boilerplate syntax targeted specifically toward the native paradigms of each ecosystem:

- `-l go` (Go Fiber) - default
- `-l typescript` (TypeScript Express.js)
- `-l python` (Python FastAPI)

### Architectural Templates

You can use the `--arch` (or `-a`) flag to automatically construct routing configurations for different paradigms:

- `-a hexagonal` (Default - Ports & Adapters via `internal/{{module}}/...`)
- `-a clean` (Clean Architecture with explicitly separated Domains, Applications and Infrastructure)
- `-a n-tier` (Traditional monolithic layers: services, models, repositories, controllers)

For example, initializing a **Python Clean Architecture** API:
```bash
anvil init . -l python -a clean
```
Generates an `anvil.json` mapped to Domain Driven Design for `.py` files automatically formatting your `fastapi` integrations.

*Note: Omitting a path by setting its value to `"-"` or `""` will completely skip generating that logic layer altogether.*

## Scaffolding an Entity

Once initialized, generate your module scaffolding via:

```bash
anvil generate User
```

This dynamically parses the layers based on your defined language. It smartly calculates import paths (reading `go.mod`, `package.json`, or applying dot-notation for Python), and creates all the ready-to-go HTTP wrappers and Models!
