package generator

const modelTemplate = `package {{.PkgName}}

// {{.Entity}} represents a domain model.
type {{.Entity}} struct {
	// TODO: Add fields
}
`

const schemaTemplate = `package {{.PkgName}}

// {{.Entity}}Schema represents the data structure for persistence or API representation.
type {{.Entity}}Schema struct {
	// TODO: Add fields, json/bson/gorm tags
}
`

const portTemplate = `package {{.PkgName}}

import "{{.Module}}/{{.ModelsPath}}"

type {{.Entity}}Repository interface {
	// TODO: Add repository methods
}

type {{.Entity}}Service interface {
	// TODO: Add service methods
}
`

const repositoryTemplate = `package {{.PkgName}}

import (
	"{{.Module}}/{{.PortsPath}}"
)

type {{.EntityLower}}Repository struct {
	// TODO: Add DB connections or clients
}

func New{{.Entity}}Repository() {{.PortsPath | base}}.{{.Entity}}Repository {
	return &{{.EntityLower}}Repository{}
}
`

const serviceTemplate = `package {{.PkgName}}

import (
	"{{.Module}}/{{.PortsPath}}"
)

type {{.EntityLower}}Service struct {
	repo {{.PortsPath | base}}.{{.Entity}}Repository
}

func New{{.Entity}}Service(repo {{.PortsPath | base}}.{{.Entity}}Repository) {{.PortsPath | base}}.{{.Entity}}Service {
	return &{{.EntityLower}}Service{
		repo: repo,
	}
}
`

const controllerTemplate = `package {{.PkgName}}

import (
	"github.com/gofiber/fiber/v2"

	"{{.Module}}/{{.PortsPath}}"
)

type {{.Entity}}Controller struct {
	service {{.PortsPath | base}}.{{.Entity}}Service
}

func New{{.Entity}}Controller(service {{.PortsPath | base}}.{{.Entity}}Service) *{{.Entity}}Controller {
	return &{{.Entity}}Controller{
		service: service,
	}
}

func (c *{{.Entity}}Controller) Create(ctx *fiber.Ctx) error {
	// TODO: implement logic
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create {{.Entity}}",
	})
}
`

const routeTemplate = `package {{.PkgName}}

import (
	"github.com/gofiber/fiber/v2"

	"{{.Module}}/{{.ControllersPath}}"
)

func Setup{{.Entity}}Routes(router fiber.Router, controller *{{.ControllersPath | base}}.{{.Entity}}Controller) {
	group := router.Group("/{{.EntityLower}}")
	
	group.Post("/", controller.Create)
	// TODO: Add more routes
}
`
