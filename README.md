# My Golang App – Gin + Docker

Bu proje, [Gin](https://github.com/gin-gonic/gin) web framework'ü kullanılarak geliştirilmiş basit bir Go web uygulamasıdır. Uygulama, Docker içinde çalışmakta ve aşağıdaki özellikleri sunmaktadır:

- `/api` endpoint'i üzerinden `{"message": "Hello World"}` dönen bir JSON API
- `/` rotasında statik `index.html` dosyasını gösteren bir web sayfası
- Multi-stage Dockerfile yapısı ile minimal imaj

---

## 📁 Proje Yapısı

```
my-golang-app/
├── Dockerfile
├── main.go
├── go.mod
├── go.sum
└── static/
    └── index.html
```

---

## 🚀 Başlatma

### 1. Docker ile Build Et

```bash
docker build -t my-golang-app .
```

### 2. Container'ı Çalıştır

```bash
docker run -p 8080:8080 my-golang-app
```

---

## 🌐 Uygulamayı Kullanma

### API Endpoint

```http
GET http://localhost:8080/api
```

Yanıt:

```json
{
  "message": "Hello World"
}
```

### Statik Web Sayfası

Tarayıcıdan şu adrese git:

```
http://localhost:8080/
```

---

## 🛠 Kullanılan Teknolojiler

- Go (Golang)
- Gin Web Framework
- Docker (multi-stage build)
- HTML/CSS/JS (statik içerik)

---

## 📝 Lisans

Bu proje MIT lisansı ile lisanslanmıştır.