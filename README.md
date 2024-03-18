# routine


бэк
++вход по логину паролю. токен в куки.
---
**/routine/{id}**
```json

[
    {
        "name":  "",
        "case_id":  1,
        "duration":  1,
        "sound":  "path to mp3",
        "index": 1
    },
    {
        "name":  "",
        "case_id":  3,
        "duration":  1,
        "sound":  "path to mp3",
        "index": 2
    }
]

```
получает всю рутину

---

- **post** /add/{id}

добавляет дело в список дел
```json
{
    "name":  "",
    "case_id":  1,
    "duration":  1,
    "sound":  "path to mp3", 
    "index": 1
}
```


- **update** /update/{id}
  редактирует  дело в список дел

```json
{
  "name": "",
  "case_id": 1,
  "duration": 1,
  "sound": "path to mp3",
  "index": 1
}
```

- **get** /sound/{path to mp3}

```json
{
  "file": "mp3"    
}
```

- **post** /delete/{id}/{case_id}
```json

```

---

База данных


|пользователи|
|-|
|**id**|
|name|
|token|



|Дела|
|-|
|*owner_id* // foregin key пользователи.id|
|**case_id** // key|
|name // text for label|
|duration // second|
|sound //path to mp3|
|index // number in order|


---


фронт:
- окно

- добавка дела
  - настройка:
    ```
    название
    время для дела,
    звук начала и конца( стандартные звуки, добавка звуков до 15 секунд.)
    index 
    ```
- удаление дела
- перемещение дел (поменять местами)

- во время выполнения показывать шкалу выполнения
- Снизу закрепленный элемент запуска и остановки
- запуск и остановка, проригрывание звука