# github.com/giusepperoro/requestqueue
___
Этот сервис API, состоящий из базы данных PostgresQL, в которую записываются данные изменения баланса. Сам сервис golang API обрабатывает очереди POST запросов на изменение баланса.

Запрос клиента - в отдельной горутине обрабатывается.
Из одной горутины, которая принимает запрос отправляются данные в другую горутину, которая делает саму работу и делает это с помощью каналов - один канал на то, чтобы отправить, другой канал на то, чтобы принять.
___

### Query: Header POST: ```URL/refill ```
### Body:
#### Пополнение:
```
{
    "client_id": 1234,
    "amount": 3000
}
```
___
#### Снятие:
```
{
    "client_id": 1234,
    "amount": -3000
}
```

#### Response on success:
```
{
    "approved": true
}
```

___
# Запуск сервиса локально

#### 1. git clone https://github.com/giusepperoro/requestqueue.git

#### 2. Используйте docker-compose для создания контейнеров с помощью PostgresQL и сервиса API
```
docker compose build
docker compose up
```