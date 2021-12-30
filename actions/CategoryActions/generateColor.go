package CategoryActions

import (
	"errors"
	"gofinance/initializers"
	"gofinance/library"
	"gofinance/models"
	"gorm.io/gorm"
	"log"
)

func GenerateUniqueColor() string {
	var category models.Category

	color := library.GetRandomHexColor()

	for {
		err := initializers.DB.Where("color = ?", color).First(category).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				break
			}

			log.Fatalf("error checking if row exists '%s' %v", color, err)
		}
		color = library.GetRandomHexColor()
	}

	return color
}
