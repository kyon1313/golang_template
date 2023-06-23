package routes

import (
	"webApp/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	app.Get("/", controller.ViewLogin)
	apiEndpoint := app.Group("/api")
	apiEndpoint.Get("/read_excel", controller.ReadExcelFile)

	// User
	userEndpoint := apiEndpoint.Group("/user")
	userEndpoint.Get("/register", controller.ViewRegistration)
	userEndpoint.Post("/save", controller.RegisterAccount)
	userEndpoint.Post("/verify", controller.VerifyAccount)
	userEndpoint.Get("/setting", controller.ViewUserSetting)
	userEndpoint.Get("/forgot", controller.ViewForgotPassword)
	userEndpoint.Post("/verify_password", controller.VerifyForgotPassword)
	userEndpoint.Get("/logout", controller.UserLogout)
	// Dashboard
	dashboardEndpoint := apiEndpoint.Group("/dashboard")
	dashboardEndpoint.Get("/dashboard_page", controller.ViewDashboard)
	// Upload File
	uploadEndpoint := apiEndpoint.Group("/upload")
	uploadEndpoint.Get("upload_page", controller.ViewUpload)
	uploadEndpoint.Get("/try_upload", controller.ReadExcelFile)
	uploadEndpoint.Get("/upload_image", controller.UploadImage)
	uploadEndpoint.Post("/upload", controller.FileUpload)
	uploadEndpoint.Get("/fetch", controller.FetchByte)
	// Delete File
	deleteEndpoint := apiEndpoint.Group("/delete")
	deleteEndpoint.Get("/delete_page", controller.ViewDelete)
	// Loan
	loanEndpoint := apiEndpoint.Group("/loan")
	loanEndpoint.Get("/report_page", controller.ViewReport)
	loanEndpoint.Get("/audit_page", controller.ViewAuditLog)
	// Client
	clientEndpoint := apiEndpoint.Group("/client")
	clientEndpoint.Get("client_page", controller.ViewClient)
	clientEndpoint.Get("check", controller.ViewCheckClient)

}
