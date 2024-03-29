# Тестовое задание

Этот репозиторий содержит решения тестовых заданий, выполненных с использованием **Go**, **JS**, **HTML**, **CSS**.

### 1. Backend (GO) - cделать клиента для получения курсов
#### Задачи:

- Сделать клиента для получения курсов криптовалют с [Coingecko API](https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1).
- Добавить возможность получать курс для определенной криптовалюты.
- Курсы обновляются не чаще чем раз в 10 минут.

### 2.Backend (GO) - cделать парсер
#### Задачи:

- Сделать парсер для [списка лучших Instagram аккаунтов в России](https://hypeauditor.com/top-instagram-all-russia/).
- Используем первую страницу результата выборки.
- Собираем данные по всем колонкам (рейтинг, имя, ник и т.д.).
- Результат сохраняется в формате `CSV` ( получаем ~50 строк ).

### 3.Frontend (JS) - fetch 
#### Задачи:

- Через `fetch` получить список валют с [Coingecko API](https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1).
- Создать страницу с таблицей - поля `"id"`, `"symbol"`, `"name"`.
- Для первых 5 валют делаем фон синего цвета (`css background`).
- Для валюты, где `"symbol" = "usdt"`, делаем фон зеленого цвета (`css background`).

### 4.Frontend (CSS) - дизайн веб-страницы
#### Задачи:

- `"Americano"` цвет текста поменять на красный.
- В таблице `"Tea Menu"` добавить колонку `"Manufacturer"`.
- В разделе `"Special Items"` у изображений сделать круглые края.
- Продублировать раздел `"Special Items"` и добавить туда слайдер для изображений.
- Для мобильных устройств (маленький экран) не показывать раздел `"About our cafe"`.
