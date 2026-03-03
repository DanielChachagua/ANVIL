package generator

import (
	"bufio"
	"bytes"
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
	Language  string            `json:"language"`
	Extension string            `json:"extension"`
	Paths     map[string]string `json:"paths"`
}

// getModuleName extracts the module name from the configuration based on semantics
func getModuleName(lang string) string {
	if lang == "go" || lang == "" {
		file, err := os.Open("go.mod")
		if err == nil {
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "module ") {
					return strings.TrimSpace(strings.TrimPrefix(line, "module "))
				}
			}
		}
	} else if lang == "typescript" || lang == "ts" {
		file, err := os.ReadFile("package.json")
		if err == nil {
			var pkg map[string]interface{}
			if json.Unmarshal(file, &pkg) == nil {
				if name, ok := pkg["name"].(string); ok {
					return name
				}
			}
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

	lang := anvilCfg.Language
	if lang == "" {
		lang = "go"
	}

	ext := anvilCfg.Extension
	if ext == "" {
		ext = ".go"
	}

	data := TemplateData{
		Entity:      entityUpper,
		EntityLower: strings.ToLower(entity),
		Module:      getModuleName(lang),
	}

	parsePath := func(p string) string {
		t, err := template.New("path").Parse(p)
		if err != nil {
			return p
		}
		var buf bytes.Buffer
		if err := t.Execute(&buf, data); err != nil {
			return p
		}
		return buf.String()
	}

	resolveImportDir := func(p string) string {
		p = parsePath(p)
		if strings.HasSuffix(p, ext) {
			return filepath.Dir(p)
		}
		return p
	}

	data.ModelsPath = resolveImportDir(targets["models"])
	data.SchemasPath = resolveImportDir(targets["schemas"])
	data.PortsPath = resolveImportDir(targets["ports"])
	data.ReposPath = resolveImportDir(targets["repositories"])
	data.ServicesPath = resolveImportDir(targets["services"])
	data.ControllersPath = resolveImportDir(targets["controllers"])
	data.RoutesPath = resolveImportDir(targets["routes"])

	var templatesMap map[string]string
	switch lang {
	case "typescript", "ts":
		templatesMap = tsTemplates
	case "python", "py":
		templatesMap = pyTemplates
	default:
		templatesMap = goTemplates
	}

	fmt.Printf("Generating scaffolding for module %q...\n", data.Entity)

	funcMap := template.FuncMap{
		"base":    filepath.Base,
		"replace": strings.ReplaceAll,
	}

	for key, folderRaw := range targets {
		if folderRaw == "" || folderRaw == "-" || folderRaw == "false" {
			// Skip this layer if the path is explicitly disabled
			continue
		}

		folder := parsePath(folderRaw)

		tmplStr, ok := templatesMap[key]
		if !ok {
			fmt.Printf("Warning: unknown template section %q, skipping\n", key)
			continue
		}

		var filePath string
		if strings.HasSuffix(folder, ext) {
			filePath = folder
			folder = filepath.Dir(filePath)
		} else {
			filePath = filepath.Join(folder, fmt.Sprintf("%s%s", data.EntityLower, ext))
		}

		if err := os.MkdirAll(folder, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", folder, err)
		}
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
