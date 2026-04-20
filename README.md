# 🐹 Cards Service - Microservicio en Go

Microservicio desarrollado en Go que permite gestionar tarjetas (cards) mediante una API REST.
Implementa una arquitectura limpia (Clean Architecture) con separación de responsabilidades en dominio, casos de uso, infraestructura y capa HTTP.

---

## 🚀 Tecnologías

* Go (Golang)
* Gin (HTTP Router)
* PostgreSQL
* JWT (middleware básico de autenticación)

---

## 📁 Estructura del proyecto

```
cards-service/
├── cmd/api/main.go                # Punto de entrada
├── internal/
│   ├── domain/                   # Entidades
│   ├── usecase/                  # Lógica de negocio
│   ├── infrastructure/           # Implementaciones (DB)
│   └── delivery/http/            # Handlers HTTP
├── pkg/middleware/               # Middlewares (JWT)
├── go.mod
```

---

## 🧠 Arquitectura

El proyecto sigue una estructura basada en **Clean Architecture**:

* **domain** → Entidades del negocio (Card)
* **usecase** → Casos de uso (List, Get, Create)
* **infrastructure** → Implementación de acceso a datos (PostgreSQL)
* **delivery/http** → Exposición de endpoints REST
* **middleware** → Autenticación JWT

---

## ⚙️ Configuración

### 1. Clonar el repositorio

```bash
git clone <repo-url>
cd cards-service
```

---

### 2. Configurar base de datos

Ejemplo de conexión:

```bash
postgres://user:pass@localhost:5432/cardsdb?sslmode=disable
```

Crear tabla:

```sql
CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    number VARCHAR(50),
    name VARCHAR(100),
    limit NUMERIC
);
```

---

### 3. Instalar dependencias

```bash
go mod tidy
```

---

### 4. Ejecutar el servicio

```bash
go run cmd/api/main.go
```

Servidor disponible en:

```
http://localhost:8080
```

---

## 🔐 Autenticación

Este servicio usa un middleware JWT básico.

Header requerido:

```
Authorization: Bearer valid-token
```

> ⚠️ Nota: La validación es simplificada. En producción se debe validar firma, expiración y claims.

---

## 🌐 Endpoints

### 📄 Listar tarjetas

```
GET /cards
```

---

### 🔍 Obtener tarjeta por ID

```
GET /cards/:id
```

---

### ➕ Crear tarjeta

```
POST /cards
```

Body ejemplo:

```json
{
  "number": "123456789",
  "name": "Visa Gold",
  "limit": 5000
}
```

---

## 🧪 Ejemplo con curl

```bash
curl -X GET http://localhost:8080/cards \
  -H "Authorization: Bearer valid-token"
```

---

## 🧩 Flujo de la aplicación

1. Request HTTP entra por Gin
2. Middleware valida JWT
3. Handler recibe la petición
4. UseCase ejecuta la lógica de negocio
5. Repository consulta la base de datos
6. Respuesta se retorna en JSON

---

## ⚡ Buenas prácticas aplicadas

* Separación de capas
* Inyección de dependencias manual
* Uso de interfaces para desacoplar repositorios
* Manejo básico de errores
* Arquitectura escalable

---

## 🚧 Mejoras futuras

* Validación real de JWT
* Manejo centralizado de errores
* Logging estructurado (zap/logrus)
* Configuración por variables de entorno
* Dockerización
* Tests unitarios y de integración
* Migraciones de base de datos

---

## 👨‍💻 Autor

Proyecto base para aprendizaje y construcción de microservicios en Go.

---
