# Get message status by trackingId

## Request

```bash
GET {host}/1.0/{trackingId}
```

`{trackingId}` - `trackingId` value received from send message response

### Authorization

`bearer` token `"grant_type": "client_credentials"` un  `"scope": "sender"`

## Response

```json
{
    "trackingId": "string",
    "status": "string"
}
```

### Description of response variables

| Variable | Type | Description
|-|-|-
| `trackingId` | *string* | tracking Id ot the message
| `status` | *string* | Status of the message. For SMS delivery statuses refer to Esteria documentation

### Example of response body

```json
{
    "trackingId": "01JJ94CH607RK5DV83B3SFTQCV",
    "status": "sent"
}
```
