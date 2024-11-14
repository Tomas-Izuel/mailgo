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

## Casos de Uso

### Crear tipo de notificación

- El administrador crea un tipo de notificación con un nombre y un template de mail.

### Crear un template de mail

- El administrador crea un template de mail con un nombre y un contenido HTML.

### Enviar notificación

- El servicio recibe un evento y un usuario.
- El servicio valida la información del usuario y obtiene el mail del 
  servicio de autenticación.
- El servicio busca el tipo de notificación asociado al evento.
- El servicio busca el template de mail asociado al tipo de notificación.
- El servicio envía el mail al usuario con el contenido del template.

## Modelo de datos


### Tipo de notificación

```
ID: ObjectId
Name: string
TemplateId: ObjectId
```

### Template de mail

```
ID: ObjectId
Subject: string
BodyHTML: string
BodyText: string
```

### Notificación

```
ID: ObjectId
TypeId: ObjectId
Recipient: string
RelatedId: string
CreatedAt: time.Time
Details: map[string]interface{}
```

## Interfaz REST

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
    "templateId": "string"
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
    "templateId": "string"
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
  "templateId": "string"
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
    "templateId": "string"
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

- **Feedback negativo** - `negative-feedback` - Se envía cuando un usuario 
  deja un feedback negativo en la plataforma.
- **Confirmación de compra** - `order-confirmation` - Se envía cuando un 
  usuario realiza una compra en la plataforma.
- **Carrito abandonado** - `abandoned-cart` - Se envía cuando un usuario 
  abandona un carrito con productos en la plataforma.