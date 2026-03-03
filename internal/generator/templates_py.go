package generator

var pyTemplates = map[string]string{
	"models":       pyModelTemplate,
	"schemas":      pySchemaTemplate,
	"ports":        pyPortTemplate,
	"repositories": pyRepositoryTemplate,
	"services":     pyServiceTemplate,
	"controllers":  pyControllerTemplate,
	"routes":       pyRouteTemplate,
}

const pyModelTemplate = `from dataclasses import dataclass

@dataclass
class {{.Entity}}:
    pass # TODO: Add fields
`

const pySchemaTemplate = `from pydantic import BaseModel

class {{.Entity}}Schema(BaseModel):
    pass # TODO: Add fields
`

const pyPortTemplate = `from abc import ABC, abstractmethod

class {{.Entity}}Repository(ABC):
    pass # TODO: Add repository methods

class {{.Entity}}Service(ABC):
    pass # TODO: Add service methods
`

const pyRepositoryTemplate = `from {{replace .PortsPath "/" "."}} import {{.Entity}}Repository

class {{.Entity}}RepositoryImpl({{.Entity}}Repository):
    def __init__(self):
        pass # TODO: Add DB connections or clients
`

const pyServiceTemplate = `from {{replace .PortsPath "/" "."}} import {{.Entity}}Repository, {{.Entity}}Service

class {{.Entity}}ServiceImpl({{.Entity}}Service):
    def __init__(self, repo: {{.Entity}}Repository):
        self.repo = repo
`

const pyControllerTemplate = `from fastapi import APIRouter
from {{replace .PortsPath "/" "."}} import {{.Entity}}Service

class {{.Entity}}Controller:
    def __init__(self, service: {{.Entity}}Service):
        self.service = service
        self.router = APIRouter(prefix="/{{.EntityLower}}")
        self._setup_routes()

    def _setup_routes(self):
        @self.router.post("/")
        async def create():
            # TODO: implement logic
            return {"message": "Create {{.Entity}}"}
`

const pyRouteTemplate = `# Routes are handled in the controller using FastAPI router
# or import controller.router here.
`
