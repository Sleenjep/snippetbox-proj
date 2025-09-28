# 1 Изначальная структура веб-приложения на Go

```bash
go mod init
```

go.mod

```bash
go env
```

# 2. Основы веб-приложений на Golang

server mux в golang

```go
w http.ResponseWriter, r *http.Request
```

```go
mux := http.NewServeMux()
mux.HandleFunc("/", home)
```

```go
w.Write([]byte("message"))
```

nginx, apache

```bash
go run main.go
```

# 3. Маршрутизация HTTP-запросов используя ServeMux

фиксированные и многоуровневые пути

```go
if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
}
```

```go
http.Handle()
http.HandleFunc()
```

DefaultServeMux

```go
var DefaultServeMux = NewServeMux()
```

проблемы с использованием дефолтного server mux

особенности проверки url-шаблонов в server mux

# 4. Настройка HTTP заголовков веб-приложения

```go
if r.Method != http.MethodPost {
    w.WriteHeader(405)
    // w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("GET-метод запрещён!"))
    return
}
```

```go
w.Header().Set("Allow", http.MethodPost)
```

```go
http.Error(w, "Метод запрещён!", http.StatusMethodNotAllowed)
```

```go
w.Header().Set("Cache-Control", "public, max-age=31536000")
 
w.Header().Add("Cache-Control", "public")
w.Header().Add("Cache-Control", "max-age=31536000")
 
w.Header().Del("Cache-Control")
 
w.Header().Get("Cache-Control")
```

```go
http.DetectContentType()
```

```go
w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{"name":"Alex"}`))
```

```go
textproto.CanonicalMIMEHeaderKey()
```

```go
w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
```

```go
w.Header()["Date"] = nil
```

```go
id, err := strconv.Atoi(r.URL.Query().Get("id"))
if err != nil || id < 1 {
    http.NotFound(w, r)

    return
}
```

# 5. Обработка URL-запросов в Golang

NOTHING


# 6. Организация файлов веб-приложения на Go


```bash
go run ./cmd/web
```

# 7. Шаблонизатор в Golang при создании веб-приложения

```go
ts, err := template.ParseFiles("./ui/html/home-page.html")
```

```go
err = ts.Execute(w, nil)
```

рабочая директория зависит от того, в каком месте запускается бинарник

```html
{{define "base"}}
...
{{define end}}
```

```html
{{template "title" .}}

{{template "main" .}}
```

golang-шаблонизатор

{{template}} и {{block}} … {{end}}.

# 8. Получаем доступ к статическим файлам — CSS и JS

```go
fileServer := http.FileServer(http.Dir("./ui/static"))
```

```html
<link rel="stylesheet" href="/static/css/main.css">
<link rel="shortcut icon" href="/static/img/favicon.ico" type="image/x-icon">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700">
```

```html
<script src="/static/js/main.js" type="text/javascript"></script>
```

```go
fileServer := http.FileServer(http.Dir("./ui/static/"))
mux.Handle("/static/", http.StripPrefix("/static", fileServer))
```

> обработчик статических файлов в Go 

> path.Clean() и атаки по обходу нижних уровней директорий

> настраиваемая файловая система и ее последующая её передача в http.FileServer

```go
fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
mux.Handle("/static", http.NotFoundHandler())
mux.Handle("/static/", http.StripPrefix("/static", fileServer))
```

# 9. Интерфейс http.Handler — Обработчик запросов

```go
mux := http.NewServeMux()
mux.Handle("/", http.HandlerFunc(home))
```

# 10. Настройка веб-приложения из командной строки

```go
addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
flag.Parse()

...

log.Printf("Запуск сервера на %s", *addr)
err := http.ListenAndServe(*addr, mux)
```

> flag.String(), flag.Int(), flag.Bool(), flag.Float64()


```bash
go run ./cmd/web/ -help
```

```go
// export SNIPPETBOX_ADDR=":9999"
addr := os.Getenv("SNIPPETBOX_ADDR")
```

```go
type Config struct {
    Addr      string
    StaticDir string
}

...
cfg := new(Config)
flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
flag.Parse()
```

# 11. Логирование в Golang — Записываем лог в файл

> стандартный логгер в golang

> информационные сообщения и сообщения об ошибках

> многоуровневое логгирование

> os.Stdin, os.Stdout, os.Stderr

```go
infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

...

infoLog.Printf("Запуск сервера на %s", *addr)
errorLog.Fatal(err)
```

```bash
go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
```

```go
srv := &http.Server{
    Addr:     *addr,
    ErrorLog: errorLog,
    Handler:  mux,
}
```

```go
err := srv.ListenAndServe()
```

```go
errorFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
    log.Fatal(err)
}
defer errorFile.Close()
```

```go
mw := io.MultiWriter(os.Stderr, errorFile)
errorLog := log.New(mw, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
```

# 12. Внедрение зависимостей в Golang (Dependency Injection)

```go
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

...

mux.HandleFunc("/", app.home)
mux.HandleFunc("/snippet", app.showSnippet)
mux.HandleFunc("/snippet/create", app.createSnippet)
```

# 13. Создание методов-помощников для обработки ошибок

> debug.Stack()

# 14. Изоляция маршрутизации приложения в отдельный файл

NOTHING

# 15. Установка MySQL для веб-приложения на Golang

NOTHING

# 16. Установка MySQL драйвера для работы в Golang

```bash
go get github.com/jackc/pgx/v5
```

```bash
go get -u github.com/jackc/pgx/v5
```

> go.mod и go.sum

```bash
go mod verify
```

```bash
go mod download
```

```bash
go get -u github.com/jackc/pgx/v5@none
```

```bash
go mod tidy
```

```bash
go mod tidy -v
```

# 17. Создание пула подключений к MySQL в Go

> пул соединений с базой данных

> импорт драйверов для баз данных в golang

```go
import _ "github.com/jackc/pgx/v5/stdlib"
```

# 18. Проектирование модели в Go

> создание моделей БД в Go

# 19. Выполнение SQL запросов в Golang

> плейсхолдеры и SQL-инъекции

```go
DB.Query()
```

```go
DB.QueryRow()
```

```go
DB.Exec()
```


```go
if, err := result.LastInsertId() // для MySQL

err = db.QueryRow(`INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`, "Alice", "alice@example.com").Scan(&id) // для PostgreSQL
```

Запрос с использованием MySQL:
```go
stmt := `
    INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))
`
```

Запрос с использованием PostgreSQL:
```go
stmt := `
    INSERT INTO snippets (title, content, created, expires)
    VALUES ($1, $2, NOW(), NOW() + ($3 || 'days')::interval)
    RETURNING id
`
```

> SQL-инъекции

# 20. Выводим запись из базы данных по её ID из URL

> использование параметра parseTime=true

> функция error.ls()

```go
if errors.Is(err, sql.ErrNoRows) {
    return nil, models.ErrNoRecord
} else {
    return nil, err
}
```

# 21. Вывод последних записей из базы данных

```go
defer rows.Close()
```

# 22. SQL Транзакции через Golang

```go
```

> пакет database/sql

```go
tx, err := m.DB.Begin()

_, err = tx.Exec("INSERT INTO ...")
if err != nil {
    tx.Rollback()
    return err
}

err = tx.Commit()
```

```go
db.SetMaxOpenConns(100)

db.SetMaxIdleConns(5)
```

# 23. Отображение контента из MySQL в HTML-шаблон

> XSS-атаки 

> экранирование данных с использованием  {{}}

> html/template и text/template

```html
<span>{{"<script>alert('xss attack')</script>"}}</span>

<span>&lt;script&gt;alert(&#39;xss attack&#39;)&lt;/script&gt;</span>
```

Вызов одного шаблона из другого c использованием .:
```html
{{template "base" .}}
{{template "main" .}}
{{template "footer" .}}
```

Вызов метода в шаблоне:
```html
<span>{{.Snippet.Created.Weekday}}</span>
```

Вызов метода и передача в него параметров:
```html
<span>{{.Snippet.Created.AddDate 0 6 0}}</span>
```

> tml/template удаляет любые HTML комментарии и любые условные комментарии

