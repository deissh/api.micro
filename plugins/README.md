# Plugins

Содержит все плагины и дополнительные модули которые работают с API и обеспечивают второстепенный функционал.
автоматическое добавление нового аниме из сторонних сервисов (MoonWalk, HDGO и т.п.).

## Создание плагина/расширения

Авторизация проходит от имени автора который запустил, то есть расширение будет работать с помощью токена автора и его
разрешений.

(In future) Так же возможно подписка на глобальные события внутри API.

## Список плагинов/расширений

- MoonWalk Watcher - обновляет в базе аниме по его `kinopoisk_id` или `world_art_id` из MoonWalk. Так же возможна загрузка и заполнение базы по умолчанию.
- ...