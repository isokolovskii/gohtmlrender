package handlers

import (
	"github.com/alexedwards/scs/v2"
	"github.com/ivan3177/gohtmlrender/pkg/models"
	"net/http"
)

// Renderer is an interface that defines a method for rendering templates.
// The RenderTemplate method takes an http.ResponseWriter, the name of the template,
// and a pointer to a models.TemplateData struct that contains the data needed for rendering the template.
// This interface can be used in a repository struct, such as the example below:
//
//	type Repository struct {
//	    render Renderer
//	}
//
// It can also be used when initializing a new Repository struct, as shown in the example below:
//
//	func New(r Renderer) *Repository {
//	    return &Repository{
//	        render: r,
//	    }
//	}
//
// The RenderTemplate method should be implemented by any type that wants to act as a renderer
// and provide the functionality to render templates using the provided data.
type Renderer interface {
	RenderTemplate(w http.ResponseWriter, name string, data *models.TemplateData)
}

// Repository is a type that represents a repository of rendered templates.
// It contains a pointer to a render.Repository, which provides methods for rendering templates.
type Repository struct {
	render         Renderer
	sessionManager *scs.SessionManager
}

var Repo Repository

// New creates a new instance of Repository with the given render object
func New(r Renderer, sessionManage *scs.SessionManager) *Repository {
	Repo = Repository{
		render:         r,
		sessionManager: sessionManage,
	}

	return &Repo
}

// Home is a method of the Repository type that renders the "home.page.tmpl" template.
// It takes in an http.ResponseWriter and an http.Request as parameters.
// The RenderTemplate method is called from the m.render field of the Repository,
// which renders the provided template and writes the output to the http.ResponseWriter.
// Example usage:
//
//	repo := &Repository{
//	  render: render.New(&config.AppConfig{}),
//	}
//	http.HandleFunc("/", repo.Home)
//
// The template cache is used to retrieve the template by name before rendering.
// If the template is not found in the cache, it is created using the createTemplateByName function.
// The executeTemplate function is responsible for executing the template and writing the output to the http.ResponseWriter.
func (m *Repository) Home(w http.ResponseWriter, _ *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "Home page"

	m.render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About method renders the "about.page.tmpl" template and writes it to the http.ResponseWriter.
// It uses the render.RenderTemplate method of the Repository's render field to render the template.
// If the template is not found in the cache, it tries to create one using the createTemplateByName function.
// If the template rendering or writing to the response fails, an error is logged.
func (m *Repository) About(w http.ResponseWriter, _ *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "About"

	m.render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
