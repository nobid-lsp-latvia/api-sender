# Sending messages as email or sms

## Request

```bash
POST {host}/1.0/send
```

### Authorization

`bearer` token `"grant_type": "client_credentials"` un  `"scope": "sender"`

### Body

JSON

```json
{
  "to": [
    {
      "email": "string",
      "phoneNumber": "string"
    }
  ],
  "from": {
    "name": "string",
    "email": "string",
    "phoneNumber": "string"
  },
  "subject": "string",
  "content": [
    {
      "type": "string",
      "value": "string"
    }
  ]
}
```

#### Description of request variables

| Variable | Type | Description | Related
|-|-|-|-
| `to` | *array* | Information about recipient | `email`, `sms`
| `to.email` | *string* | email of the recipient | `email`
| `to.phoneNumber` | *string* | phone number of the recipient (shall include country code) | `sms`
| `from` | *array* | **Optional, if not included, values are taken from environment.** <br/> Information about sender | `email`, `sms`
| `from.name` | *string* | Sender name. If not set, taken from `SENDER_MAIL_NAME` variable | `email`
| `from.email` | *string* | Sender email. If not set, taken from `SENDER_MAIL` variable | `email`
| `from.phoneNumber` | *string* | `Sender` value (number or name) agreed with ESTERIA. If not set, taken from `SENDER_PHONE_NAME` variable | `sms`
| `subject` | *string* | Email subject name | `email`
| `content` | *array* | Content of the message | `email`, `sms`
| `content.type` | *string* | `text/plain` - for usage in `email` & `sms` <br/> `text/html` - for usage in `email` | `email`, `sms`
| `content.value` | *string* | Message to be sent | `email`, `sms`

### Examples

#### Send email

```json
{
  "to": [
    {
      "email": "john.doe@mail.com"
    }
  ],
  "from": {
    "name": "no-reply",
    "email": "no-reply@email.com"
  },
  "subject": "Subject text",
  "content": [
    {
      "type": "text/plain",
      "value": "<p>Test  message</p>"
    }
  ]
}
```

#### Send email without `from` array

```json
{
  "to": [
    {
      "email": "john.doe@mail.com"
    }
  ],
  "subject": "Subject text",
  "content": [
    {
      "type": "text/plain",
      "value": "<p>Test  message</p>"
    }
  ]
}
```

#### Send sms

```json
{
  "to": [
    {
      "phoneNumber": "+37112345678"
    }
  ],
  "from": {
    "phoneNumber": "identificator"
  },
  "content": [
    {
      "type": "text/plain",
      "value": "<p>Test  message</p>"
    }
  ]
}
```

#### Send sms without `from` array

```json
{
  "to": [
    {
      "phoneNumber": "+37112345678"
    }
  ],
  "content": [
    {
      "type": "text/plain",
      "value": "<p>Test  message</p>"
    }
  ]
}
```

## Response

### Body of the response

JSON

```json
{
    "messages": [
        {
            "trackingId": "string"
        }
    ]
}
```

### Description of response variables

| Variable | Type | Description
|-|-|-
| `messages` | *array* | Array with tracking Id's
| `trackingId` | *string* | tracking Id ot the message, can be used to receive statuss of specific message

### Example of response body

```json
{
    "messages": [
        {
            "trackingId": "01JJ96ABZNQF1K33WPK8SS4RJ1"
        }
    ]
}
```

## Corner cases

When you try to send a message and receive HTTP status code `200` with JSON in body:

```json
{
    "trackingId": "",
    "status": ""
}
```

You are using `GET` and shall switch method to `POST`.
