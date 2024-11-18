# Arquitectura de microservicios
## Servicio de notificaciones mail

```
Author: Izuel Tomas
Legajo: 47700
Año: 2024
```

## Descripción

El servicio de mensajeria por mail esta encargado de hacer llegar 
notificaciones a los usuarios de la plataforma debido a diferentes eventos.

Para que estas notificaciones sean eficientes en cuanto al contenido y 
presentación de los mensajes, se utilizá tipos de eventos, administrables 
por un administrador, los cuales contienen un template de mail personalizado 
para cada caso.

## Tecnologías
- Golang
- MongoDB
- RabbitMQ

## Casos de Uso

### Consultar notificaciones

- Se obtienen todas las notificaciones de un usuario.
- No se necesita permisos especiales.
- Un usuario solo puede ver sus notificaciones.

### Consultar notificaciones una notificación

- Se obtiene una notificación de un usuario.
- No se necesita permisos especiales.
- Un usuario solo puede ver sus notificaciones.

### Crear tipo de notificación

- Solo los usuarios con permiso de "admin" pueden crear un tipo de notificación.
- Se pueden asignar tantas notificaciones como se desee.
- Se valida que el template de mail exista.

### Modificar tipo de notificación

- Solo los usuarios con permiso de "admin" pueden modificar un tipo de notificación.
- Se pueden modificar los campos de nombre, template de mail y notificaciones.

### Eliminar tipo de notificación

- Solo los usuarios con permiso de "admin" pueden eliminar un tipo de notificación.

### Obtener tipo de notificación

- El administrador obtiene un tipo de notificación.

### Obtener todos los tipos de notificación

- El administrador obtiene todos los tipos de notificación.

### Crear un template de mail

- Solo los usuarios con permiso de "admin" pueden crear un template de mail.
- Se debe ingresar un asunto y un cuerpo de mail.

### Enviar notificación

- Se envía una notificación a un usuario.
- Se obtiene el template de mail correspondiente al tipo de notificación.
- Se reemplazan las variables del template por los valores correspondientes.

## Modelo de datos


### Tipo de notificación

```
ID: ObjectId
Name: string
TemplateId: ObjectId
EventKeys: []string
```

### Template de mail

```
ID: ObjectId
Subject: string
BodyHTML: string
```

### Notificación

```
ID: ObjectId
TypeId: ObjectId
UserId: ObjectId
Recipient: string
RelatedId: string
CreatedAt: time.Time
Details: map[string]interface{}
Mail: {
    Subject: string
    BodyHTML: string
    BodyText: string
}
```

## Interfaz REST

### Notificaciones

*Consultar notificaciones* `GET /notification`

**Headers**

| Cabecera      | Contenido                         |
| ------------- | --------------------------------- |
| Authorization | Bearer xxx (Token en formato JWT) |

**Response**:

- **200 OK**
```json
[
  {
    "id": "string",
    "typeId": "string",
    "recipient": "string",
    "relatedId": "string",
    "createdAt": "string",
    "mail": {
      "subject": "string",
      "bodyHTML": "string",
      "bodyText": "string"
    }
  }
]
```

- **401 UNAUTHORIZED** si el token es inválido
- **500 INTERNAL SERVER ERROR** si hay un error en el servidor

*Consultar una notificación* `GET /notification/{id}`

**Headers**

| Cabecera      | Contenido                         |
| ------------- | --------------------------------- |
| Authorization | Bearer xxx (Token en formato JWT) |

**Response**:

- **200 OK**
```json
{
  "id": "string",
  "typeId": "string",
  "recipient": "string",
  "relatedId": "string",
  "createdAt": "string",
  "mail": {
    "subject": "string",
    "bodyHTML": "string",
    "bodyText": "string"
  }
}
```

- **401 UNAUTHORIZED** si el token es inválido
- **500 INTERNAL SERVER ERROR** si hay un error en el servidor

--- 

### Tipos de notificación

*Creación de tipo de notificación* `POST /notification-type`

**Headers**

| Cabecera      | Contenido                         |
| ------------- | --------------------------------- |
| Authorization | Bearer xxx (Token en formato JWT) |

**Body**

```json
{
    "name": "string",
    "templateId": "string",
    "notifications": ["string"]
}
```

**Response**:

- **200 OK** si la notificación fue creada con éxito
- **400 BAD REQUEST** si hay errores en la solicitud
- **401 UNAUTHORIZED** si el token es inválido
- **500 INTERNAL SERVER ERROR** si hay un error en el servidor

*Obtener todas las notificaciones* `GET /notification-type`

**Headers**

| Cabecera      | Contenido                         |
| ------------- | --------------------------------- |
| Authorization | Bearer xxx (Token en formato JWT) |

**Response**:

- **200 OK**
```json
[
  {
    "id": "string",
    "name": "string",
    "templateId": "string",
    "notifications": ["string"]
  }
]
```

- **401 UNAUTHORIZED** si el token es inválido
- **500 INTERNAL SERVER ERROR** si hay un error en el servidor

*Obtener una notificación* `GET /notification-type/{id}`

**Headers**

| Cabecera      | Contenido                         |
| ------------- | --------------------------------- |
| Authorization | Bearer xxx (Token en formato JWT) |

**Response**:

- **200 OK**
```json
{
  "id": "string",
  "name": "string",
  "templateId": "string",
  "notifications": ["string"]
}
```

- **401 UNAUTHORIZED** si el token es inválido
- **500 INTERNAL SERVER ERROR** si hay un error en el servidor

*Actualizar una notificación* `PUT /notification-type/{id}`

**Headers**

| Cabecera      | Contenido                         |
| ------------- | --------------------------------- |
| Authorization | Bearer xxx (Token en formato JWT) |

**Body**

```json
{
    "name": "string",
    "templateId": "string",
    "notifications": ["string"]
}
```

**Response**:

- **200 OK** si la notificación fue actualizada con éxito
- **400 BAD REQUEST** si hay errores en la solicitud
- **401 UNAUTHORIZED** si el token es inválido
- **500 INTERNAL SERVER ERROR** si hay un error en el servidor

*Eliminar una notificación* `DELETE /notification-type/{id}`

**Headers**

| Cabecera      | Contenido                         |
| ------------- | --------------------------------- |
| Authorization | Bearer xxx (Token en formato JWT) |

**Response**:

- **200 OK** si la notificación fue eliminada con éxito
- **401 UNAUTHORIZED** si el token es inválido
- **500 INTERNAL SERVER ERROR** si hay un error en el servidor

---

### Templates de mail

*Creación de template de mail* `POST /mail-template`

**Headers**

| Cabecera      | Contenido                         |
| ------------- | --------------------------------- |
| Authorization | Bearer xxx (Token en formato JWT) |

**Body**

```json
{
    "subject": "string",
    "bodyHTML": "string"
}
```

**Response**:

- **200 OK** 
```json
{
  "id": "string",
  "subject": "string",
  "bodyHTML": "string"
}
```

- **400 BAD REQUEST** si hay errores en la solicitud
- **401 UNAUTHORIZED** si el token es inválido
- **500 INTERNAL SERVER ERROR** si hay un error en el servidor

---

## Interfaz asíncrona (RabbitMQ)

### Eventos

- **Exchange**: `notification`
- **Tipo**: `direct`
- **routingKey**: `send-notification`

**Mensaje**

```json
{
    "eventKey":  "string",
    "relatedId": "string",
    "userId":   "string",
    "details": [
      {
        "key": "value"
      }
    ]
}
```