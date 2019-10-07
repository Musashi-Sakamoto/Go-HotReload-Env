package handler

import (
  "encoding/json"
  "io/ioutil"
  "net/http"

  "github.com/labstack/echo"
  "app/model"
)

func AboutHandler(c echo.Context) error {
  // Please note the the second parameter "about.html" is the template name and should
  // be equal to one of the keys in the TemplateRegistry array defined in main.go
  tr := &http.Transport{}
  client := &http.Client{Transport: tr}

  resp, err := client.Get(
	  "http://localhost:1323/api/get-full-name?first_name=Musashi&last_name=Sakamoto",
  )

  if err != nil {
	return err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  var exampleResponse model.ExampleResponse
  if err = json.Unmarshal(body, &exampleResponse); err != nil {
	  return err
  }

  return c.Render(http.StatusOK, "about.html", map[string]interface{}{
    "name": "About",
    "msg": exampleResponse.Msg,
  })
}
