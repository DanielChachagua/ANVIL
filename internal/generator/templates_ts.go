package generator

var tsTemplates = map[string]string{
	"models":       tsModelTemplate,
	"schemas":      tsSchemaTemplate,
	"ports":        tsPortTemplate,
	"repositories": tsRepositoryTemplate,
	"services":     tsServiceTemplate,
	"controllers":  tsControllerTemplate,
	"routes":       tsRouteTemplate,
}

const tsModelTemplate = `export interface {{.Entity}} {
  // TODO: Add fields
}
`

const tsSchemaTemplate = `export const {{.Entity}}Schema = {
  // TODO: Add schema definition (e.g., Zod, Mongoose)
};
`

const tsPortTemplate = `import { {{.Entity}} } from "{{.Module}}/{{.ModelsPath}}";

export interface {{.Entity}}Repository {
  // TODO: Add repository methods
}

export interface {{.Entity}}Service {
  // TODO: Add service methods
}
`

const tsRepositoryTemplate = `import { {{.Entity}}, {{.Entity}}Repository } from "{{.Module}}/{{.PortsPath}}";

export class {{.Entity}}RepositoryImpl implements {{.Entity}}Repository {
  // TODO: Add DB connections or clients
}
`

const tsServiceTemplate = `import { {{.Entity}}Repository, {{.Entity}}Service } from "{{.Module}}/{{.PortsPath}}";

export class {{.Entity}}ServiceImpl implements {{.Entity}}Service {
  constructor(private repo: {{.Entity}}Repository) {}
  // TODO: Add service methods
}
`

const tsControllerTemplate = `import { Request, Response } from "express";
import { {{.Entity}}Service } from "{{.Module}}/{{.PortsPath}}";

export class {{.Entity}}Controller {
  constructor(private service: {{.Entity}}Service) {}

  create = async (req: Request, res: Response) => {
    // TODO: implement logic
    res.status(201).json({ message: "Create {{.Entity}}" });
  };
}
`

const tsRouteTemplate = `import { Router } from "express";
import { {{.Entity}}Controller } from "{{.Module}}/{{.ControllersPath}}";

export function setup{{.Entity}}Routes(controller: {{.Entity}}Controller): Router {
  const router = Router();
  router.post("/", controller.create);
  // TODO: Add more routes
  return router;
}
`
