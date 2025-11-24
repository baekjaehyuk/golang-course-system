package server

import (
	"fmt"
	"golang-course-registration/common/exception"
	"golang-course-registration/config"
	"golang-course-registration/controller/api"
	"golang-course-registration/controller/web"
	"golang-course-registration/infrastructure/database"
	"golang-course-registration/repository"
	"golang-course-registration/service"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Store  *database.SupabaseStore
	echo   *echo.Echo
	config *config.Config
}

type TemplateRenderer struct {
	templates map[string]*template.Template
}

func New(cfg *config.Config) *Server {
	store, _ := database.NewSupabase(cfg.Url, cfg.Key)
	return &Server{
		Store:  store,
		echo:   echo.New(),
		config: cfg,
	}
}

func (s *Server) Init() {
	e := s.echo

	templateCache, err := s.loadTemplates()
	if err != nil {
		panic(err)
	}

	renderer := &TemplateRenderer{
		templates: templateCache,
	}
	e.Renderer = renderer

	lectureRepo := s.InjectLectureRepository()
	enrollmentRepo := s.InjectEnrollmentRepository()
	studentRepo := s.InjectStudentRepository()

	lectureService := s.InjectLectureService(lectureRepo, enrollmentRepo)
	studentService := s.InjectStudentService(studentRepo)
	enrollmentService := s.InjectEnrollmentService(enrollmentRepo, lectureRepo, studentRepo)

	adminController := s.InjectAdminController(lectureService)
	clientController := s.InjectClientController(studentService, lectureService, enrollmentService)
	pageController := s.InjectPageController(lectureService, enrollmentService)

	v1 := e.Group("/api/v1")
	clientGroup := v1.Group("/client")
	clientController.RegisterRoutes(clientGroup)

	adminGroup := v1.Group("/admin")
	adminController.RegisterRoutes(adminGroup)

	pageController.RegisterRoutes(e)
}

func (s *Server) loadTemplates() (map[string]*template.Template, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf(exception.ErrNotFoundDirectory)
	}

	templateDir := filepath.Join(wd, "view", "templates")
	layout := filepath.Join(templateDir, "base.html")

	patterns := []string{
		filepath.Join(templateDir, "*.html"),
		filepath.Join(templateDir, "*", "*.html"),
	}

	var views []string
	for _, pattern := range patterns {
		files, _ := filepath.Glob(pattern)
		views = append(views, files...)
	}

	var partials []string
	for _, dir := range []string{"view/style", "view/script"} {
		partialDir := filepath.Join(wd, dir)
		files, _ := filepath.Glob(filepath.Join(partialDir, "*.html"))
		partials = append(partials, files...)
	}

	cache := make(map[string]*template.Template)
	for _, view := range views {
		if filepath.Base(view) == "base.html" {
			continue
		}

		files := append([]string{layout, view}, partials...)
		tmpl, _ := template.ParseFiles(files...)

		relPath, _ := filepath.Rel(templateDir, view)
		key := filepath.ToSlash(relPath)
		cache[key] = tmpl
	}

	return cache, nil
}

func (s *Server) Start() {
	port := s.config.Port
	s.echo.Logger.Fatal(s.echo.Start(":" + port))
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	tmpl, _ := t.templates[name]
	return tmpl.ExecuteTemplate(w, filepath.Base(name), data)
}

func (s *Server) InjectLectureRepository() repository.LectureRepository {
	return repository.NewLectureRepository(s.Store.Client)
}

func (s *Server) InjectEnrollmentRepository() repository.EnrollmentRepository {
	return repository.NewEnrollmentRepository(s.Store.Client)
}

func (s *Server) InjectStudentRepository() repository.StudentRepository {
	return repository.NewStudentRepository(s.Store.Client)
}

func (s *Server) InjectLectureService(lectureRepo repository.LectureRepository, enrollmentRepo repository.EnrollmentRepository) service.LectureService {
	return service.NewLectureServiceWithEnrollment(lectureRepo, enrollmentRepo)
}

func (s *Server) InjectStudentService(studentRepo repository.StudentRepository) service.StudentService {
	return service.NewStudentService(studentRepo)
}

func (s *Server) InjectEnrollmentService(
	enrollmentRepo repository.EnrollmentRepository,
	lectureRepo repository.LectureRepository,
	studentRepo repository.StudentRepository,
) service.EnrollmentService {
	return service.NewEnrollmentService(enrollmentRepo, lectureRepo, studentRepo)
}

func (s *Server) InjectAdminController(lectureService service.LectureService) *api.AdminController {
	return api.NewAdminController(lectureService)
}

func (s *Server) InjectClientController(
	studentService service.StudentService,
	lectureService service.LectureService,
	enrollmentService service.EnrollmentService,
) *api.ClientController {
	return api.NewClientController(studentService, lectureService, enrollmentService)
}

func (s *Server) InjectPageController(lectureService service.LectureService, enrollmentService service.EnrollmentService) *web.PageController {
	return web.NewPageController(lectureService, enrollmentService)
}
