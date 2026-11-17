# API-тесты для проекта https://testboard.avito.com/

## Полезные ссылки

1. [Swagger-спецификация](https://testboard.avito.com/swagger/index.html)
2. [Техническое задание](https://docs.google.com/document/d/1IZ9SxEcePoQT6lyEmukycp1qNok4HZtok73UQnPwKSA/edit?usp=sharing)
3. [Библиотека для написания API-тестов](https://github.com/steinfletcher/apitest)
4. [Мощная библиотека анализа JSON для Go](https://github.com/tidwall/gjson)

## Запуск тестов

1. Установите любое IDE для разработки: [Jetbrains GoLand](https://www.jetbrains.com/ru-ru/go/) (Платное или по студенческой лицензии), [VS Code](https://code.visualstudio.com/) (бесплатное).
2. Установите [Git](https://git-scm.com/install/) на компьютер
3. [Установите](https://go.dev/doc/install) Go версии 1.25.1
4. [Настройте](https://docs.github.com/ru/authentication/connecting-to-github-with-ssh) доступ к своему Github-аккаунту
   по SSH-ключу.
5. Склонируйте себе репозиторий: `git clone <путь к вашему репозиторию>`
6. Подтяните необходимые зависимости через команду в корне проекта:

```bash
   go mod tidy
```

7. Создайте в корне `.env.override` и укажите данные двух разных тестовых аккаунтов TestBoard:

```dotenv
TEST_LOGIN=your_login
TEST_PASSWORD=your_password
TEST_BUYER_LOGIN=buyer_login
TEST_BUYER_PASSWORD=buyer_password
```

`TEST_LOGIN` и `TEST_PASSWORD` относятся к аккаунту продавца, а `TEST_BUYER_LOGIN` и `TEST_BUYER_PASSWORD` — к аккаунту покупателя. Подставьте свои значения. Файл `.env.override` исключён из Git.

Сценарии лабораторной сами создают и удаляют свои объявления. Фото для загрузки уже находится в `tests/testdata/advertisement.png`. У покупателя должно быть свободное место в избранном.

Для полного прогона прежних сценариев заранее подготовьте у продавца хотя бы одно объявление с фотографией. В `tests/scenarios/advertisementSearch/search_test.go` укажите его полное название в `existingAdTitle` и обновите `searchBeginning`, `searchMiddle` и `searchEnd`.

8. Для запуска всех тестов достаточно выполнить команду в корне проекта:

```bash
  go test -v ./...
```

Для запуска только лабораторной работы выполните из корня репозитория:

```shell
go test -count=1 -v ./tests/scenarios/advertisementCreate/ ./tests/scenarios/favorites/
```

9. Отдельный пакет с тестами можно запустить командой:

```bash
  go test -v ./tests/scenarios/myAdvertisement/
```

## Логирование

Для включения логирования request/response в консоль: выставьте переменную окружения `DEBUG=true` в
`.env.override`

## Запуск линтера

1. Установить приложение [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/)
2. Запуск линтера (из корня репозитория):

```bash
  golangci-lint run
```
