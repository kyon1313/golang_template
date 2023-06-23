package controller

import (
	"log"
	"net/http"
	"strconv"
	"webApp/controller/util"
	"webApp/model"

	"github.com/gofiber/fiber/v2"
)

// CONTROLLER
func RegisterAccount(c *fiber.Ctx) error {

	log.Println("Registration")
	// Get value from input tags
	ic := c.FormValue("instiCode")
	fn := c.FormValue("firstname")
	ln := c.FormValue("lastname")
	un := c.FormValue("username")
	pw := c.FormValue("password")
	cpw := c.FormValue("confirmPassword")

	cic, _ := strconv.Atoi(ic)

	// Initiate Institutions
	instiModel := []model.M_Institution{}
	util.DBConn.Debug().Table("m_institution").Find(&instiModel)

	// Password not match
	if pw != cpw {
		instiModel := []model.M_Institution{}
		util.DBConn.Debug().Table("m_institution").Find(&instiModel)

		return c.Render("registration", fiber.Map{
			"iconDesc":     "error",
			"title":        "USER REGISTRATION",
			"statusCode":   http.StatusUnprocessableEntity,
			"statusDesc":   "Password not match",
			"institutions": instiModel,
			"firstname":    fn,
			"lastname":     ln,
			"username":     un,
			"insticode":    ic,
		})
	}

	// if fields are empty
	if fn == "" || ln == "" || un == "" || pw == "" || ic == "" {

		return c.Render("registration", fiber.Map{
			"iconDesc":     "error",
			"title":        "USER REGISTRATION",
			"statusCode":   http.StatusUnprocessableEntity,
			"statusDesc":   "Registration failed, please fill out the fields",
			"institutions": instiModel,
			"firstname":    fn,
			"lastname":     ln,
			"username":     un,
			"insticode":    ic,
		})

	} else {
		hashPW, _ := util.HashPassword(pw)
		// store data
		userModel := model.TbUsers{
			Firstname: fn,
			Lastname:  ln,
			Username:  un,
			Password:  hashPW,
			InstiCode: cic,
		}

		result := util.DBConn.Debug().Table("tbl_users").Where("username = ?", un).FirstOrCreate(&userModel)

		if result.RowsAffected == 0 {
			return c.Render("registration", fiber.Map{
				"iconDesc":     "error",
				"title":        "USER REGISTRATION",
				"statusCode":   http.StatusUnprocessableEntity,
				"statusDesc":   "Registration failed, user already exist",
				"institutions": instiModel,
				"firstname":    fn,
				"lastname":     ln,
				"username":     un,
				"insticode":    ic,
			})
		}

		return c.Render("login", fiber.Map{
			"iconDesc":   "success",
			"title":      "USER LOGIN",
			"statusCode": http.StatusOK,
			"statusDesc": "Registration successful",
		})
	}
}

func VerifyAccount(c *fiber.Ctx) error {
	userModel := model.TbUsers{}

	un := c.FormValue("username")
	pw := c.FormValue("password")

	log.Println("Verify Account Query:", util.DBConn.Debug().Table("tbl_users").Where("username = ?", un).Find(&userModel))

	result := util.CheckPasswordHash(pw, userModel.Password)

	if result {
		return c.Render("dashboard", fiber.Map{
			"page":       "Dashboard",
			"iconDesc":   "success",
			"title":      "DASHBOARD",
			"statusCode": http.StatusOK,
			"statusDesc": "Account Verified",
			"username":   un,
		})
	}

	return c.Render("login", fiber.Map{
		"iconDesc":   "error",
		"title":      "USER LOGIN",
		"statusCode": http.StatusUnauthorized,
		"statusDesc": "Invalid login account",
		"username":   un,
	})
}

func VerifyForgotPassword(c *fiber.Ctx) error {
	log.Println("Forgot Password")
	// Get value from input tags
	fn := c.FormValue("firstname")
	ln := c.FormValue("lastname")
	un := c.FormValue("username")
	pw := c.FormValue("password")
	cpw := c.FormValue("confirmPassword")

	// Password not match
	if pw != cpw {
		instiModel := []model.M_Institution{}
		util.DBConn.Debug().Table("m_institution").Find(&instiModel)

		return c.Render("forgotpassword", fiber.Map{
			"iconDesc":   "error",
			"title":      "FORGOT PASSWORD",
			"statusCode": http.StatusUnprocessableEntity,
			"statusDesc": "Password not match",
			"firstname":  fn,
			"lastname":   ln,
			"username":   un,
		})

	}
	return nil
}
func UserLogout(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"title":      "USER LOGIN",
		"statusCode": "",
		"iconDesc":   "",
		"statusDesc": "",
	})
}

// VIEWS
func ViewLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"title": "USER LOGIN",
	})
}

func ViewRegistration(c *fiber.Ctx) error {
	instiModel := []model.M_Institution{}

	util.DBConn.Debug().Table("m_institution").Find(&instiModel)

	return c.Render("registration", fiber.Map{
		"title":        "USER REGISTRATION",
		"institutions": instiModel,
	})
}

func ViewDashboard(c *fiber.Ctx) error {
	return c.Render("dashboard", fiber.Map{
		"page":  "Dashboard",
		"title": "DASHBOARD",
	})
}

func ViewUserSetting(c *fiber.Ctx) error {
	return c.Render("usersetting", fiber.Map{
		"page":  "User Setting",
		"title": "USER SETTING",
	})
}

func ViewForgotPassword(c *fiber.Ctx) error {
	return c.Render("forgotpassword", fiber.Map{
		"page":  "Update User Password",
		"title": "FORGOT PASSWORD",
	})
}
