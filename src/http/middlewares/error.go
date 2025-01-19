package middlewares

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func ErrorHanlde(c *fiber.Ctx, err error) error {

	v := reflect.ValueOf(err)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	statusField := v.FieldByName("StatusCode")
	if statusField.IsValid() && statusField.CanInterface() {
		statusCode := statusField.Interface().(int)
		message := v.FieldByName("Message").Interface().(string)
		c.Status(statusCode)
		return c.JSON(map[string]interface{}{
			"message": message,
		})

	}

	c.Status(500)
	return c.JSON(map[string]string{
		"message": err.Error(),
	})
}
