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
   (во втором случае параметры для подключения к БД будут взяты из .env)
```
POSTGRES_USER=swoyo
POSTGRES_PASSWORD=swoyo
POSTGRES_DB=swoyo
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
```
3) Если нужно протестировать как сервис работает с postgresql, 
а она не установлена/лень настраивать окружение, то можно
воспользоваться: "docker-compose up --build -d" (порт так же будет 8080)

```bash
docker-compose up --build -d
```
Дождаться запуска <br/>
Без Docker
![RUN_WITHOUT_DB](https://github.com/aleksandrzhigulin/SWOYO/assets/66275482/65e13f0b-bf34-4de9-b2d4-b85a4858e216)
![RUN_WITH_FLAG](https://github.com/aleksandrzhigulin/SWOYO/assets/66275482/ccaf8830-e41a-42ab-b2ed-9ee5a8c4519f)

С Docker
![RUN_DOCKER](https://github.com/aleksandrzhigulin/SWOYO/assets/66275482/dc8dad9b-67e3-4c2b-85dd-428176330521)
Стек: Go (go-chi, database/sql), PostgreSQL, Docker, приложение работает на порте **8080**, postgresql на порте **5432**
# Описание API
**GET http://localhost:8080/{url}** - *принимает сокращённый url и возвращает полный* <br/> <br/>
**POST http://localhost:8080/** - *создаёт новый сокращённый url и возвращает его* <br/>
Тело запроса:
```
http://cjdr17afeihmk.biz/123/kdni9/z9d112423421
```
## Примеры работы API
**POST http://localhost:8080/**
![POST_NEW_URL](https://github.com/aleksandrzhigulin/SWOYO/assets/66275482/3dec5379-48ac-4be0-baa9-d4dcb63e61a7)
<br/> <br/>
**GET http://localhost:8080/{url}**
![GET_200](https://github.com/aleksandrzhigulin/SWOYO/assets/66275482/b29b4e56-4a2d-4d2c-8c2a-b2b5e1680bea)
![GET_404](https://github.com/aleksandrzhigulin/SWOYO/assets/66275482/b83f77a8-f86a-4566-a1ad-e814d7cdcc92)
