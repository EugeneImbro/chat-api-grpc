# chat-backend

пилим бекенд для тупого чатика:
1) есть юзер (id, nickname)
2) есть room (id, name, users)
3) есть сообщение(id, from, room_id, text, created_at)

будем итеративно прогать все, поэтому я ожидаю от тебя:
1) grpc сервер с методом GetUsers
2) postgres база (только с юзерами)
3) докерфайл
4) миграция в базу (только с юзерами)

не важно, могу предложить в целом самый популярный https://github.com/golang-migrate/migrate

для работы с базой sqlx

работа с юзерами пока только ридонли

типа GetUsers метод

драйвер для postgres можешь брать любой libpq или pgx

для тестирования https://github.com/stretchr/testify(testify/require и testify/assert), для тестов базы + https://github.com/testcontainers/testcontainers-go
