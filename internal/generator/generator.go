package generator

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type TemplateData struct {
	Entity          string
	EntityLower     string
	Module          string
	PkgName         string
	ModelsPath      string
	SchemasPath     string
	PortsPath       string
	ReposPath       string
	ServicesPath    string
	ControllersPath string
	RoutesPath      string
}

type AnvilConfig struct {
	Paths map[string]string `json:"paths"`
}

// getModuleName extracts the module name from the go.mod file, fallback if not found
func getModuleName() string {
	file, err := os.Open("go.mod")
	if err != nil {
		return "your_project"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	return "your_project"
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func GenerateModule(entity string) error {
	entityUpper := capitalize(entity)

	targets := make(map[string]string)

	configData, err := os.ReadFile("anvil.json")
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("anvil.json not found in the current directory.\nPlease run 'anvil init .' (optionally with --arch) to generate the configuration file first")
		}
		return fmt.Errorf("could not read anvil.json: %w", err)
	}

	var anvilCfg AnvilConfig
	if err := json.Unmarshal(configData, &anvilCfg); err != nil {
		return fmt.Errorf("invalid anvil.json format: %w", err)
	}

	if anvilCfg.Paths != nil {
		for k, v := range anvilCfg.Paths {
			targets[k] = v
		}
	}

	data := TemplateData{
		Entity:          entityUpper,
		EntityLower:     strings.ToLower(entity),
		Module:          getModuleName(),
		ModelsPath:      targets["models"],
		SchemasPath:     targets["schemas"],
		PortsPath:       targets["ports"],
		ReposPath:       targets["repositories"],
		ServicesPath:    targets["services"],
		ControllersPath: targets["controllers"],
		RoutesPath:      targets["routes"],
	}

	templatesMap := map[string]string{
		"models":       modelTemplate,
		"schemas":      schemaTemplate,
		"ports":        portTemplate,
		"repositories": repositoryTemplate,
		"services":     serviceTemplate,
		"controllers":  controllerTemplate,
		"routes":       routeTemplate,
	}

	fmt.Printf("Generating scaffolding for module %q...\n", data.Entity)

	funcMap := template.FuncMap{
		"base": filepath.Base,
	}

	for key, folder := range targets {
		if folder == "" || folder == "-" || folder == "false" {
			// Skip this layer if the path is explicitly disabled
			continue
		}

		tmplStr, ok := templatesMap[key]
		if !ok {
			fmt.Printf("Warning: unknown template section %q, skipping\n", key)
			continue
		}

		if err := os.MkdirAll(folder, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", folder, err)
		}

		filePath := filepath.Join(folder, fmt.Sprintf("%s.go", data.EntityLower))
		if _, err := os.Stat(filePath); err == nil {
			fmt.Printf("File %s already exists, skipping.\n", filePath)
			continue
		}

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", filePath, err)
		}

		data.PkgName = filepath.Base(folder)

		t := template.Must(template.New(key).Funcs(funcMap).Parse(tmplStr))
		if err := t.Execute(file, data); err != nil {
			file.Close()
			return fmt.Errorf("failed to execute template for %s: %w", filePath, err)
		}
		file.Close()

		fmt.Printf("Created: %s\n", filePath)
	}

	fmt.Println("Scaffolding complete!")
	return nil
}
