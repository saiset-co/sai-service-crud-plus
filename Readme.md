# Microservice store dictionary

## Ports

`6080`: **HTTP**  
`6081`: WebSocket   
`6082`: Socket

## How to run

`make build`: rebuild and start service  
`make up`: start service  
`make down`: stop service  
`make logs`: display service logs

## API

### Create dictionary document:

```json lines
{
  "method": "create",
  "data": {
    "dictionary": "$dictionaryName",
    "documents": [
      {
        "sto_id": "$stoId",
        //Any data to save
        "$anyFieldName1": "$anyFieldValue1",
        "$anyFieldName2": "$anyFieldValue2"
      }
    ]
  }
}
```

## Read dictionary:

```json lines
{
  "method": "read",
  "data": {
    "dictionary": "$dictionaryName",
    "select": {
      "sto_id": "$stoId",
      // any fields you want
      "$anyFieldName1": "$anyFieldValue1"
    },
    "options": {
      "count": 1,
      "limit": 1,
      "sort": {
        "$anyFieldName1": -1
      }
    },
    "fields": [
      "$anyFieldName1",
      "$anyFieldName2",
      ...
    ]
  }
}
```

## Update dictionary documents:

```json lines
{
  "method": "update",
  "data": {
    "dictionary": "$dictionaryName",
    "select": {
      "sto_id": "$stoId",
      // any fields you want
      "$anyFieldName1": "$anyFieldValue1"
    },
    "document": {
      "sto_id": "$stoId",
      // any data you want
      "$anyFieldName2": "$anyFieldValue2-1"
    }
  }
}
```

## Delete dictionary documents:

```json lines
{
  "method": "delete",
  "data": {
    "dictionary": "$dictionaryName",
    "select": {
      "sto_id": "$stoId",
      // any fields you want
      "$anyFieldName1": "$anyFieldValue1"
    }
  }
}
```