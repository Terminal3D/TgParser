# Telegram Bot для отслеживания цен в интернет-магазинах

## Описание
Этот Telegram бот предназначен для мониторинга цен и доступности товаров в интернет-магазинах. Пользователи могут добавлять товары для отслеживания, удалять их, просматривать список отслеживаемых товаров, подписываться на уведомления об изменениях цен и доступности, а также отписываться от них.
На данный момент реализована работа с сайтом sportsdirect.com.

## Начало работы

Клонирование репозитория
```sh
git clone github.com/Terminal3D/TgParser
```


Настройка переменных окружения
```makefile
BOT_TOKEN=<ваш токен бота Telegram>
DB_URL=<URL подключения к базе данных>
```

Установка зависимостей для бота
```go
go mod tidy
```


## Использование


* /start - начать отслеживание товаров
* /stop - прекратить отслеживание товаров
* /subsribe - подписаться на рассылку обновлений о товарах
* /unsubscribe - отписаться от рассылки обновлений о товарах

Получение списка добавление нового предмета,удаление или получение списка доступных предметов происходит при помощи разных меню, используемых ботом. (internal/tgbot/menus.go)