package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"app/model"
)

func PostFullName(c echo.Context) error {
	exampleRequest := new(model.ExampleRequest)
	if err := c.Bind(exampleRequest); err != nil {
		return err
	}

	greeting := exampleRequest.FirstName + " " + exampleRequest.LastName

	var resp bytes.Buffer
	var b = []byte(
		fmt.Sprintf(`{
			"first_name": %q,
			"last_name": %q,
			"msg": "Hello %s"
		}`, exampleRequest.FirstName, exampleRequest.LastName, greeting),
	)
	err := json.Indent(&resp, b, "", " ")
	if (err != nil) {
		return err
	}

	return c.JSONBlob(
		http.StatusOK,
		[]byte(
			fmt.Sprintf("%s", resp.Bytes()),
		),
	)
}