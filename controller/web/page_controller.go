package web

import (
	"golang-course-registration/common/constants"
	"golang-course-registration/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PageController struct {
	lectureService    service.LectureService
	enrollmentService service.EnrollmentService
}

func NewPageController(lectureService service.LectureService, enrollmentService service.EnrollmentService) *PageController {
	return &PageController{
		lectureService:    lectureService,
		enrollmentService: enrollmentService,
	}
}

func (c *PageController) RegisterRoutes(e *echo.Echo) {
	e.GET("/", c.Index)
	e.GET("/admin/dashboard", c.AdminDashboard)
	e.GET("/client/dashboard", c.ClientDashboard)
}

func (c *PageController) Index(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "index.html", map[string]interface{}{})
}

func (c *PageController) AdminDashboard(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "admin.html", map[string]interface{}{})
}

func (c *PageController) ClientDashboard(ctx echo.Context) error {
	studentID, err := strconv.Atoi(ctx.QueryParam("studentId"))
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/?error=invalid_student_id")
	}

	if studentID < constants.StudentIdMin || studentID > constants.StudentIdMax {
		return ctx.Redirect(http.StatusFound, "/?error=invalid_student_id")
	}

	return ctx.Render(http.StatusOK, "client.html", map[string]interface{}{"StudentID": studentID})
}
