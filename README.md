# neurocapsule

A service for brewing magical eerie enchanted mysterious otherwordly spellbinding potions.

## 🚀 Перед началом работы

Проект использует линтеры и pre-commit hook'и, поэтому для разработки проекта требуется установить
утилиту [pre-commit](https://pre-commit.com/#install)

Для MacOS через [homebrew](https://brew.sh/):

```bash
brew install pre-commit
```

Или используя утилиту pip:

```bash
pip install pre-commit
```

После установки pre-commit нужно запустить установку hook'ов:

```bash
pre-commit install
```

После установки hook'ов на каждый коммит в репозиторий будет запускаться ряд обязательных проверок.

## 📁 Структура папок проекта

Стандартную структуру Go-приложений в Lamoda можно посмотреть [здесь](https://stash.lamoda.ru/projects/GOTOOLS/repos/scratch/browse/docs/project_structure).

## 🚜 Запуск

### Docker-compose

В целях разработки проект можно запустить используя [docker-compose](https://docs.docker.com/compose/)

```bash
make dev-up-app
```

Если нужно изменить параметры конфигурации, то желательно это делать через файл _deployments/compose.dev.env_.

### Локально

Перед первым запуском необходимо скопировать `deployments/local.dev.env.example` в `deployments/local.dev.env`. Последний уже содержится в `.gitignore`.

Как вариант окружение можно запустить в docker-compose, а приложение локально:

```bash
make dev-local-env
make run-server
```

### Рестарт

Окружение в docker-compose можно перезапустить командой `make dev-restart`.

## 🛠 Сборка

Проект может быть собран как в виде бинарного файла (`make build`), так и в виде Docker образов:

* CI: `make docker-build-ci-image`
* Prod: `make @build`

## 👨🏼‍💻 Разработка

`make setup` - установка всех необходимых зависимостей для проекта (например, golangci-lint)

`make lint` - запуск линтеров локально (должен быть установлен golangci-lint)

`make test` - запуск тестов локально

`make generate` - перегенерация сгенерированного кода в `/app/scratch/generated`

## 🧪 Тестирование

Проект использует стандартные тесты go c фреймворками
[testify](https://github.com/stretchr/testify) и [gonkey](https://github.com/lamoda/gonkey). Тесты могут быть запущены в
docker контейнере в docker-compose окружении:

```bash
make @test
```

Или локально, но при этом используя БД из docker-compose окружения:

```bash
make dev-env
NEUROCAPSULE_CONF=$(PWD)/deployments/local.dev.env make test
```

В процессе прохождения тестов также строится отчет по покрытию кода, который находится в файле coverage.html

## 🔗 Полезные ссылки

* 🏗 Сборка и деплой (Bamboo)
  * [Тесты и линтинг]() - TO BE DONE.
  * [Сборка]() - TO BE DONE.
  * [Деплой]() - TO BE DONE.
  * [Деплой в QA-кластер]() - TO BE DONE.
* [💳 Карточка сервиса (Confluence)]() - TO BE DONE.
* 📈 Метрики (Grafana)
  * [Бизнес-метрики]() - TO BE DONE.
  * [Технический дашборд]() - TO BE DONE.
* 📜 [Логи (Kibana)]() - TO BE DONE.
* 🛡 [Sentry]() - TO BE DONE.

## 📌 Внешние зависимости
Инфраструктурные:- PostgreSQL


