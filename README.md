# avito-segmentation
Приложение запускается на localhost:8080
с помощью команды docker-compose up -d
файл иницилизации базы данных находится в assets/postgres

Примеры запросов:
Метод создания сегмента. Принимает slug (название) сегмента:
Post:
  curl --location 'localhost:8080/segments/' \
--header 'Content-Type: application/json' \
--data '{
    "segment_name":"AVITO_DISCOUNT_80"
}'

Метод удаления сегмента. Принимает slug (название) сегмента.
Delete:
  curl --location 'localhost:8080/segments/' \
--header 'Content-Type: application/json' \
--data '{
    "segment_name":"AVITO_DISCOUNT_80"
}'

Метод добавления пользователя в сегмент. Принимает список slug (названий) сегментов которые нужно добавить пользователю, список slug (названий) сегментов которые нужно удалить у пользователя, id пользователя.
Post:
 curl --location 'localhost:8000/users/' \
--header 'Content-Type: application/json' \
--data '{
  "segment_for_adding":[
       {
			"slug": "AVITO_DISCOUNT_20",
			"ttl": "2023-08-29 13:59:05"
		},
        {
			"slug": "AVITO_DISCOUNT_30",
			"ttl": "2023-08-30 14:30:05"
		},
		{
			"slug": "AVITO_DISCOUNT_50"
		}
      ],
  "segment_for_deleting":["AVITO_DISCOUNT_30"],
  "user_id":2
}'

Метод получения активных сегментов пользователя. Принимает на вход id пользователя.
Post:
curl --location 'localhost:8080/users/activeSegments' \
--header 'Content-Type: application/json' \
--data '{
  "user_id":2
}'

Получение ссылки на csv-отчет с историей сегментов пользователей за период

post
curl --location 'localhost:8080/users/reports' \
--header 'Content-Type: application/json' \
--data '{
    "period":"2023-09"
}'
Тут мы получим ссылку на файл и делаем get запрос на эту ссылку, чтобы получить данные с файла
