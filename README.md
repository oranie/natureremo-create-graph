# natureremo-create-graph

Nature Remo APIを叩いてデフォルトでは見れない温度、湿度、明るさをグラフ化する

## 事前準備

API Token発行

https://home.nature.global/

## Data Modeling

sample Jsonデータ
http://swagger.nature.global/#/default/get_1_devices

```Json
[
  {
    "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    "name": "string",
    "temperature_offset": 0,
    "humidity_offset": 0,
    "created_at": "2020-01-20T01:43:56.774Z",
    "updated_at": "2020-01-20T01:43:56.774Z",
    "firmware_version": "string",
    "mac_address": "string",
    "serial_number": "string",
    "newest_events": {
      "te": {
        "val": 0,
        "created_at": "2020-01-20T01:43:56.774Z"
      },
      "hu": {
        "val": 0,
        "created_at": "2020-01-20T01:43:56.774Z"
      },
      "il": {
        "val": 0,
        "created_at": "2020-01-20T01:43:56.774Z"
      },
      "mo": {
        "val": 0,
        "created_at": "2020-01-20T01:43:56.774Z"
      }
    }
  }
]
```

### DynamoDB データモデリング

|DataType(PK)|(SK)|Value|
|---|---|---|
|DeviceId_Type|unixtime|float|
|3fa85f64-5717-4562-b3fc-2c963f66afa6|updated_at|device_name and other info|
|3fa85f64-5717-4562-b3fc-2c963f66afa6_Hu|1579487090|53.2|
|3fa85f64-5717-4562-b3fc-2c963f66afa6_Te|1579487090|25.1|
|3fa85f64-5717-4562-b3fc-2c963f66afa6_Il|1579487090|200.0|






