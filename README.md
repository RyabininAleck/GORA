# GORA

## Задача
Реализовать HTTP API фотогалереи
Возможности API:
- Загрузка фото
  - Сохранять данные о новых файлах в БД
  - Генерировать preview
- Просмотр списка фото
- Удаление фото
Дополнительные требования:
- Формат ответа - json
- В качестве БД использовать sqlite
- Время на выполнение: не более 4х часов


## Решение
Было реализовано api

**POST /api/v1/upload:**  предназначен для загрузки изображений. Ожидается файл изображения в теле запроса и возвращает информацию о загруженном изображении. 

**GET /api/v1/photo:** предназначен для получения списка фотографий. Возвращает список всех доступных фотографий.

**GET /api/v1/photo/:id:** предназначен для получения конкретной фотографии по её идентификатору (:id). Возвращает информацию о запрошенной фотографии.

**DELETE /api/v1/photo/:id:** предназначен для удаления фотографии по её идентификатору (:id). Удаляет фотографию из базы данных.

**GET /api/v1/show/preview/:id:**  предназначен для отображения превью запрошенной фотографии по её идентификатору (:id). Возвращает превью фотографии.

**GET /api/v1/show/image/:id:**  предназначен для отображения полноразмерной версии запрошенной фотографии по её идентификатору (:id). Возвращает полноразмерное изображение.



### Обработчик UploadPhotoHandler для загрузки фото:

- Принимает файл из формы запроса.
- Генерирует уникальное имя файла.
- Сохраняет оригинальное фото и его превью.
- Возвращает данные о загруженном фото.

### Обработчик GetPhotoListHandler для получения списка фото:

- Извлекает список фото из базы данных.
- Возвращает список фото в ответе.

### Обработчик GetPhotoHandler для получения конкретного фото:

- Извлекает фото по его идентификатору из базы данных.
- Возвращает фото в ответе.

### Обработчик DelPhotoHandler для удаления фото:

- Получает идентификатор фото из параметра запроса.
- Удаляет фото из базы данных.
- Возвращает идентификатор удаленного фото в ответе.

### Обработчики ShowPhotoHandler и ShowPreviewHandler для отображения фото и их превью:

- Получают идентификатор фото из параметра запроса.
- Извлекают информацию о фото из базы данных.
- Открывают файл фото или превью.
- Возвращают файл фото или превью в ответе.