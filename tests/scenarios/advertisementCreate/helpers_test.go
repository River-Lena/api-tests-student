package advertisementCreate

import (
	"strconv"

	"api-tests-template/tests/testdata"
)

// validAdvertisementFields возвращает новый набор валидных полей объявления.
func validAdvertisementFields(title, description string) map[string]string {
	return map[string]string{
		"title":       title,
		"description": description,
		"price":       strconv.Itoa(testdata.AdvertisementPrice),
		"quantity":    strconv.Itoa(testdata.AdvertisementQuantity),
	}
}

// advertisementPhotoPath возвращает путь к фото для тестов создания.
func advertisementPhotoPath() string {
	return testdata.AdvertisementPhotoPath()
}
