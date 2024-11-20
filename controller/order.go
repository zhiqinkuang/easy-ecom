package controller

import "github.com/gin-gonic/gin"

type OrderController struct {
}
type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (O OrderController) GetList(c *gin.Context) {
	search := &Search{}
	err := c.BindJSON(&search)
	if err == nil {
		SuccessJson(c, 200, search.Cid, search.Name, 1)
		return
	}
	ErrorJson(c, 400, err)

}
