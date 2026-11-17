package testdata

import "path/filepath"

const (
	AdvertisementPrice    = 1557
	AdvertisementQuantity = 2
	AdvertisementPhoto    = "advertisement.png"
)

// AdvertisementPhotoPath возвращает путь к фото из пакетов tests/scenarios.
func AdvertisementPhotoPath() string {
	return filepath.Join("..", "..", "testdata", AdvertisementPhoto)
}
