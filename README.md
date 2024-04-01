## Необходимо для запуска:
* Go или Docker Compose*
## Подготовительные действия
Клонирование репозитория
```bash
git clone https://github.com/aleksandrzhigulin/SWOYO.git
```
Переход в папку
```bash
cd SWOYO
```
## Сборка и запуск
Варианты:
1) Запустить готовый экзешник main.exe
2) "go run main.go" или "go run main.go -d"
   (во втором случае username,password,database name будут взяты из окружения + port=5432, host=database)
3) Если нужно протестировать как сервис работает с postgresql, 
а она не установлена/лень настраивать окружение, то можно
воспользоваться: "docker-compose up --build -d" (порт так же будет 8080)

```bash
docker-compose up --build -d
```
Дождаться запуска

Стек: Go (go-chi, database/sql), PostgreSQL, Docker, приложение работает на порте **8080**, postgresql на порте **5432**
# Описание API
**GET http://localhost:8080/{url}** - *возвращает айди пользователя, его имя и баланс* <br/> <br/>
**POST http://localhost:8080/** - *создаёт нового пользователя.* <br/>
Тело запроса:
```
http://cjdr17afeihmk.biz/123/kdni9/z9d112423421
```
## Примеры работы API
**GET http://localhost:8080/{url}**
<br/> <br/>
**POST http://localhost:8080/**
