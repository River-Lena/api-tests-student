package advertisement

import (
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpPostAdvertisement отправляет форму создания объявления с фотографиями.
func HttpPostAdvertisement(t *testing.T, token string, contentType string, body string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Post(path.AdvertisementPath).
		Header("Content-Type", contentType).
		Body(body).
		Expect(t).
		End().Response
}

// HttpGetAdvertisement получает объявление по идентификатору.
func HttpGetAdvertisement(t *testing.T, token string, id string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Get(path.AdvertisementPath).
		Query("id", id).
		Expect(t).
		End().Response
}

// HttpDeleteAdvertisement удаляет объявление по идентификатору.
func HttpDeleteAdvertisement(t *testing.T, token string, id string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Delete(path.AdvertisementPath).
		Query("id", id).
		Expect(t).
		End().Response
}

// HttpGetAdvertisementPhotos получает список фотографий объявления.
func HttpGetAdvertisementPhotos(t *testing.T, token string, id string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Get(path.AdvertisementsPath + "/" + id + "/photos").
		Expect(t).
		End().Response
}
