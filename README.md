# Go Youtrack Rest Client
Клиент для YouTrack

Нейминг структур и методов аналогичен неймингу YouTrack'а.

### Описание
#### Файлы проекта:
- [client.go](client.go) - содержит логику инициализации клиента. Создание нового клиента + options-функции.

- [methods.go](methods.go) - содержит все методы клиента.

- [entities.go](entities.go) - содержит нужные нам entities YouTrack'а. Все ссылки есть в комментариях.

- [common.go](common.go) - содержит общие приватные методы, константы и т.п.

- [requests.go](requests.go) - содержит структуры реквестов, которые мы маршаллим в JSON request body.