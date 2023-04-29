package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"ticket-backend-gh/contracts/request"
	"ticket-backend-gh/contracts/response"
	"ticket-backend-gh/services/ticket"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (con Controller) MapEndpoints(r *gin.RouterGroup) {
	r.POST("/generateTicket", con.generateTicket)
}

func (con Controller) generateTicket(c *gin.Context) {
	ticket := ticket.TicketServiceImpl{}
	var requestBody request.RequestTicket
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Print(err)
		message := response.Response{
			Status: false,
			Code:   500,
			Data:   response.Data{},
		}
		c.JSON(500, message)
		return
	}

	json.Unmarshal(body, &requestBody)

	ticketResp, err := ticket.GenerateTicket(requestBody)
	if err != nil {
		log.Print(err)
		message := response.Response{
			Status: false,
			Code:   500,
			Data:   response.Data{},
		}
		c.JSON(500, message)
		return
	}

	responseSuccess := response.Response{
		Status: true,
		Code:   201,
		Data:   ticketResp,
	}

	c.JSON(http.StatusCreated, responseSuccess)

}
