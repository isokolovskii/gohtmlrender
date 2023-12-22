package render

import (
	"bytes"
	"fmt"
	"github.com/ivan3177/gohtmlrender/pkg/config"
	"github.com/ivan3177/gohtmlrender/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Repository represents a data repository that stores and manages templates and configuration.
// templateCache is a map that stores template names as keys and corresponding *template.Template objects as values.
// The templates are used for rendering HTML pages.
// The templateCache is initialized empty and populated when needed.
// config is a pointer to AppConfig that contains application configuration.
// It is used to determine whether or not to use the template cache.
// init is a method of Repository that initializes the template cache based on the configuration.
// If the UseCache field is true, it calls initCache to populate the templateCache.
// initCache is a method of Repository that populates the template cache.
// It reads the template files from the ./templates directory and creates *template.Template objects.
// The created templates are stored in the templateCache map.
// RenderTemplate is a method of Repository that renders a template and writes the generated HTML to the http.ResponseWriter.
// It retrieves the template from the template cache based on the provided template name.
// If the template is not found in the cache, it creates the template using the createTemplateByName function.
// The rendered HTML is written to the http.ResponseWriter, and any error that occurs during rendering is logged.
// templateFromCacheByName is a method of Repository that retrieves a template from the template cache based on its name.
// It checks if the template exists in the cache map.
// If not, it logs a message and calls createTemplateByName to create the template.
// The retrieved template and any error are returned.
// createTemplateCache is a function that creates and returns a map of template names and *template.Template objects.
// It reads the template files from the ./templates directory and creates *template.Template objects.
// The templates are stored in a map with their names as keys.
// executeTemplate is a function that executes a template and writes the rendered HTML to the http.ResponseWriter.
// It uses a buffer to capture the rendered HTML.
// The rendered HTML is then written to the http.ResponseWriter.
// The number of bytes written is logged, and any error that occurs during execution is returned.
// createTemplateByName is a function that creates and returns a *template.Template object.
// It resolves the file name from the template name and calls createTemplateFromFile to create the template.
// The created template and any error are returned.
// createTemplateFromFile is a function that creates and returns a *template.Template object from a file.
// It parses the file using template.New and template.ParseFiles.
// If successful, the template is passed to addLayoutsToTemplate to add layout templates.
// The created template, the base name of the file, and any error are returned.
// resolveFileFromTemplateName is a function that resolves the file path from the template name.
// It returns the resolved file path by formatting it with the "./templates" prefix.
type Repository struct {
	templateCache map[string]*template.Template
	config        *config.AppConfig
}

// New creates a new instance of Repository using the provided AppConfig configuration. It initializes the template cache and returns a pointer to the Repository.
func New(config *config.AppConfig) *Repository {
	repo := Repository{
		templateCache: make(map[string]*template.Template),
		config:        config,
	}

	repo.init()

	return &repo
}

// init initializes the repository by calling the initCache method with the specified UseCache value from the configuration.
func (r *Repository) init() {
	r.initCache(r.config.UseCache)
}

// initCache initializes the template cache in the Repository struct based on the specified useCache parameter.
// If useCache is true, it creates the template cache by calling the createTemplateCache function and assigns it to the Repository's templateCache field.
// If an error occurs while creating the cache, it logs the error message.
func (r *Repository) initCache(useCache bool) {
	if useCache {
		tc, err := createTemplateCache()
		if err != nil {
			log.Printf("Cannot create templates cache: %v", err)
		} else {
			r.templateCache = tc
		}
	}
}

// createTemplateCache creates a cache of templates by reading all "*.page.tmpl" files in the "./templates" directory.
// The function returns a map of template names to *template.Template pointers and an error, if any.
// The function uses the createTemplateFromFile function to read each template file and add it to the cache.
func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		ts, name, err := createTemplateFromFile(page)

		if err != nil {
			return cache, err
		}

		cache[name] = ts
	}

	return cache, nil
}

// createTemplateFromFile parses a template from a file and adds layouts to it if available.
// It takes a string parameter `page` which represents the filename of the template file.
// It returns a pointer to the parsed template (`ts`), the name of the template (`name`), and an error (`err`).
// The name of the template is obtained by removing the directory and file extension from the `page`.
// The template is parsed using the `New` method of the `template` package and the `ParseFiles` function.
// If there are no errors during parsing, the `addLayoutsToTemplate` function is called to add layouts to the template.
// Finally, the parsed template, template name, and any error encountered are returned.
func createTemplateFromFile(page string) (ts *template.Template, name string, err error) {
	name = filepath.Base(page)

	ts, err = template.New(name).ParseFiles(page)

	if err == nil {
		ts, err = addLayoutsToTemplate(ts)
	}

	return
}

// addLayoutsToTemplate parses layout templates and adds them to the existing template set.
// It takes a pointer to the template set and returns a new template set and an error, if any.
// If layout templates are found, they are parsed and added to the template set.
// Otherwise, the original template set is returned.
// The layout templates are expected to have the ".layout.tmpl" extension.
// Usage example:
//
//	ts, err := addLayoutsToTemplate(existingTemplates)
func addLayoutsToTemplate(ts *template.Template) (t *template.Template, err error) {
	layouts, _ := filepath.Glob("./templates/*.layout.tmpl")

	if len(layouts) > 0 {
		t, err = ts.ParseGlob("./templates/*.layout.tmpl")

		return
	}

	return ts, nil
}

// RenderTemplate renders a template by name.
// It retrieves the template from the cache if available, and executes it with the provided http.ResponseWriter.
// If the template is not found in the cache, it creates a new template by the provided name.
// If there is an error executing the template, it logs the error.
func (r *Repository) RenderTemplate(w http.ResponseWriter, name string, td *models.TemplateData) {
	t, err := r.templateFromCacheByName(name)
	if err != nil {
		log.Fatalf("Cannot create template by provided template name %s: %v", name, err)
		return
	}

	err = executeTemplate(w, t, td)
	if err != nil {
		log.Printf("Error executing template %s: %v", name, err)
	}
}

// templateFromCacheByName retrieves a template from the cache based on its name.
// If the template is not found in the cache, it creates the template using the provided name.
// The function returns the template and any error that occurred during the process.
func (r *Repository) templateFromCacheByName(name string) (t *template.Template, err error) {
	var ok bool

	t, ok = r.templateCache[name]

	if !ok {
		log.Printf("Template %s not found in cache", name)
		t, err = createTemplateByName(name)
	}

	return
}

// createTemplateByName creates a template by name.
func createTemplateByName(name string) (t *template.Template, err error) {
	page := resolveFileFromTemplateName(name)
	t, _, err = createTemplateFromFile(page)
	return
}

// resolveFileFromTemplateName takes a template name and returns the file path where the template is located.
// It appends the template name to the "./templates/" directory path.
func resolveFileFromTemplateName(tmpl string) string {
	return fmt.Sprintf("./templates/%s", tmpl)
}

// executeTemplate executes the provided template and writes the rendered output to the http.ResponseWriter.
// It logs the number of bytes written.
func executeTemplate(w http.ResponseWriter, t *template.Template, td *models.TemplateData) (err error) {
	td = addDefaultData(td)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, td)
	if err != nil {
		return
	}

	var bytesWritten int64
	bytesWritten, err = buf.WriteTo(w)
	if err != nil {
		return
	}

	log.Printf("Written %d bytes while writting rendered template", bytesWritten)
	return nil
}

// addDefaultData adds default data to the provided TemplateData object and returns it.
func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
