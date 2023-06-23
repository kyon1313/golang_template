package controller

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"webApp/controller/util"
	"webApp/model"

	"github.com/gofiber/fiber/v2"
)

// CONTROLLER
func ConvertToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func FileUpload(c *fiber.Ctx) error {

	imgSrc, uploadErr := c.FormFile("imageSource")
	imgSize := imgSrc.Size

	fmt.Printf("IMAGE SIZE: %d", imgSize)
	if uploadErr != nil {
		return c.SendString("Error in FormFile")
	}

	imgPath := fmt.Sprintf("./template/images/uploads/%s", imgSrc.Filename)
	c.SaveFile(imgSrc, imgPath)
	bytes, err := ioutil.ReadFile(imgPath)

	if err != nil {
		log.Fatal(err)
	}

	// var base64Encoding string

	// Determine the content type of the image file
	// mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	// switch mimeType {
	// case "image/jpeg":
	// 	base64Encoding += "data:image/jpeg;base64,"
	// case "image/png":
	// 	base64Encoding += "data:image/png;base64,"
	// }

	// Append the base64 encoded output
	// base64Encoding += ConvertToBase64(bytes)

	fmt.Println("Byte:", bytes)
	imgStruct := model.Uploaded_Images{
		UserID:  1,
		ImgData: bytes,
		ImgType: "mimeType",
	}

	util.DBConn.Debug().Table("uploaded_images").Create(&imgStruct)

	return c.Render("uploadimage", fiber.Map{
		"page":       "Testing Upload Image File",
		"iconDesc":   "success",
		"imageData":  bytes,
		"title":      "UPLOAD IMAGE",
		"statusCode": http.StatusOK,
		"statusDesc": "Registration successful",
	})
}

func FetchByte(c *fiber.Ctx) error {
	imgStruct := model.Uploaded_Images{}
	util.DBConn.Debug().Table("uploaded_images").Select("img_data").First(&imgStruct)
	return c.JSON(imgStruct.ImgData)
}

// VIEWS
func UploadImage(c *fiber.Ctx) error {
	return c.Render("uploadimage", fiber.Map{
		"page":  "Testing Upload Image File",
		"title": "UPLOAD IMAGE",
	})
}
