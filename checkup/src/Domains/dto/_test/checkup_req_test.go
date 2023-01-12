package checkupdto_test

import (
	"testing"

	checkupdto "github.com/bimaagung/laboratcall-microservices/checkup/src/Domains/dto"
	"github.com/stretchr/testify/assert"
)

type Payload struct {
		Id string
		Name string
		Unit string
		Price int
}

func TestCheckup(t *testing.T) {

	t.Run("should throw error when payload did not contain needed property", func(t *testing.T) {
		// Arrange
		payload := &Payload{
			Unit: "mg",
			Price: 10000,
		} 	

		// action
		_, err := checkupdto.NewCheckupReq(payload.Name, payload.Unit, payload.Price)
		

		// Assert
		assert.EqualError(t,err, "CHECKUP.NOT_CONTAIN_NEEDED_PROPERTY")
	})

	t.Run("should checkup object correctly", func(t *testing.T) {
		// Arrange
		payload := &Payload{
			Name: "Hematology",
			Unit: "mg",
			Price: 10000,
		} 	

		// action
		result, err := checkupdto.NewCheckupReq(payload.Name, payload.Unit, payload.Price)

		// Assert
		assert.NoError(t,err)
		assert.Equal(t, result.Name, payload.Name)
		assert.Equal(t, result.Unit, payload.Unit)
		assert.Equal(t, result.Price, payload.Price)
	})
}