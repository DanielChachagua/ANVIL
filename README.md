# ANVIL CLI

Anvil is a powerful Go-based CLI tool backed by `spf13/cobra` designed to instantly generate boilerplate code for new domain entities following a structured architecture and port-and-adapters methodologies.

## Global Installation 

To use Anvil in any project, install it globally using Go. Run the following from the root of this project:

```bash
go install ./cmd/anvil
```
*(Make sure `$(go env GOPATH)/bin` is in your system's PATH)*

Once installed globally, you can execute `anvil` from anywhere on your terminal.

## Initializing Application Configurations

Anvil uses an `anvil.json` file inside the root of your target project to determine which folders map to what layers (e.g., `ports`, `repositories`, `controllers`). 

To generate this baseline configuration file, simply go to your application logic map and run:

```bash
anvil init .
```
This generates the Hexagonal Architecture defaults.

### Architectural Templates

You can use the `--arch` (or `-a`) flag to automatically construct configurations for different paradigms:

- `-a hexagonal` (Default - Ports & Adapters via internal packages)
- `-a clean` (Clean Architecture with explicitly separated Domains, Applications and Infrastructure)
- `-a n-tier` (Traditional monolithic layers: services, models, repositories, controllers)

For example:
```bash
anvil init . --arch clean
```
Generates a Domain Driven Design compatible `anvil.json` map:
```json
{
  "paths": {
    "models": "domain/models",
    "schemas": "infrastructure/schemas",
    "ports": "domain/ports",
    "services": "application/services",
    "repositories": "infrastructure/repositories",
    "controllers": "interfaces/http/controllers",
    "routes": "interfaces/http/routes"
  }
}
```

*Note: Omitting a path by setting its value to `"-"` or `""` will completely skip generating that logic layer altogether.*

## Scaffolding an Entity

Once initialized, generate your module scaffolding via:

```bash
anvil generate User
```

This dynamically maps the layers based on your defined configurations, smartly calculating Go paths based on `go.mod` imports, parsing models, routes, controllers, and services with Fiber HTTP wrappers ready-to-go!
