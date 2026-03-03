# ANVIL CLI

Anvil is a powerful Go-based CLI tool backed by `spf13/cobra` designed to instantly generate boilerplate code for new domain entities following a structured architecture and port-and-adapters methodologies.

## Global Installation 

To use Anvil in any project, install it globally using Go. Run the following from the root of this project:

```bash
go install ./cmd/anvil
sudo ln -sf /home/daniel/go/bin/anvil /usr/local/bin/anvil
```
*(Make sure `$(go env GOPATH)/bin` is in your system's PATH)*

Once installed globally, you can execute `anvil` from anywhere on your terminal.

## Initializing Application Configurations

Anvil uses an `anvil.json` file inside the root of your target project to determine which folders map to what layers (e.g., `ports`, `repositories`, `controllers`). 

To generate this baseline configuration file, simply go to your application logic map and run:

```bash
anvil init .
```
This generates the **Domain-Oriented Hexagonal Architecture** defaults and automatically scaffolds the base project directories (`cmd/`, `internal/platform/`).

### Architectural Templates

You can use the `--arch` (or `-a`) flag to automatically construct configurations for different paradigms:

- `-a hexagonal` (Default - Ports & Adapters via `internal/{{module}}/...`)
- `-a clean` (Clean Architecture with explicitly separated Domains, Applications and Infrastructure)
- `-a n-tier` (Traditional monolithic layers: services, models, repositories, controllers)

For example, the **Hexagonal** setup scaffolds the configuration using Go templates to route to dynamic directories:
```json
{
  "paths": {
    "models": "internal/{{.EntityLower}}/domain/models.go",
    "schemas": "internal/{{.EntityLower}}/infrastructure/repository/schemas.go",
    "ports": "internal/{{.EntityLower}}/domain/ports.go",
    "services": "internal/{{.EntityLower}}/application/services.go",
    "repositories": "internal/{{.EntityLower}}/infrastructure/repository/repository.go",
    "controllers": "internal/{{.EntityLower}}/infrastructure/handler/http/controller.go",
    "routes": "internal/{{.EntityLower}}/infrastructure/handler/http/routes.go"
  }
}
```

*Note: Omitting a path by setting its value to `"-"` or `""` will completely skip generating that logic layer altogether.*

## Scaffolding an Entity

Once initialized, generate your module scaffolding via:

```bash
anvil generate User
```

This dynamically parses the layers based on your defined configurations. Using `{{.EntityLower}}`, Anvil places `models.go` in `internal/user/domain/models.go`, cleanly separating domains per module! 

It smartly calculates Go import paths based on your `go.mod`, parsing models, routes, controllers, and services with Fiber HTTP wrappers ready-to-go!
