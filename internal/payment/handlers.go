package payment

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var newPayment Payment

	if err := c.BindJSON(&newPayment); err != nil {
		fmt.Printf("Error binding JSON for creating payment: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPayment.ID = fmt.Sprintf("%d", len(PaymentsData)+1)

	PaymentsData = append(PaymentsData, newPayment)
	c.IndentedJSON(http.StatusCreated, newPayment.ID)
}

func GetPaymentByID(c *gin.Context) {
	id := c.Param("id")
	for _, p := range PaymentsData {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Payment not found"})
}
