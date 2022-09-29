# cosmos-sync

A server that synchronize cosmos chain data into a database

# SetUp

# Build And Run

- Build: `make all`
- Run: `make run`
- Cross compilation: `make build-linux`

## Env Variables

- CONFIG_FILE_PATH: `option` `string` config file path

## Config Descriptiom

### [config.toml](https://github.com/bianjieai/cosmos-sync/blob/irishub/1.1.0/config/config.toml)

```text
[database]
node_uri= "mongodb://root:123456@10.1.4.157:27017/?connect=direct&authSource=temp"
database= "bifrost-sync"

[server]
node_urls="tcp://192.168.150.40:26657"
worker_num_create_task=1
worker_num_execute_task=30
worker_max_sleep_time=120
block_num_per_worker_handle=100
max_connection_num=100
init_connection_num=5
bech32_acc_prefix=""
chain_id=""
chain_block_interval=5
behind_block_num=0
promethous_port=9090
support_modules="" 
deny_modules=""
support_types=""
ignore_ibc_header=false
```

### Db config

- node_uri: `required` `string` db connection uri（example: `mongodb://root:123456@10.1.4.157:27017/?connect=direct&authSource=temp, ...`）
- database：`required` `string` database name（example：`DB_DATABASE`）

### Server config

- node_urls: `required` `string`  full node uri（example: `tcp://127.0.0.1:26657, tcp://127.0.0.2:26657, ...`）
- worker_num_create_task: `required` `int` the maximum time (in seconds) that create task are allowed （default: `1`
  example: `1`）
- worker_num_execute_task: `required` `int` the maximum time (in seconds) that synchronization TX threads are allowed
  to be out of work（example: `30`）
- worker_max_sleep_time: `required` `int` num of worker to create tasks(unit: seconds)（example: `90`）
- block_num_per_worker_handle: `required` `int`  number of blocks per sync TX task（example: `50`）

- max_connection_num: `required` `int` client pool config total max connection
- init_connection_num: `required` `int` client pool config idle connection

- bech32_acc_prefix: `option` `string` block chain address prefix（default(cosmoshub): `` example(irishub): `iaa`）
- chain_block_interval: `option` `int` block interval; default `5` (example: `5`)
- behind_block_num: `option` `int` wait block num to handle when retry failed; default `0` (example: `0`)
- promethous_port: `option` `int` promethous metrics server port
- support_modules: `option` `string` setting only support module tx sync,default
  support [all module](https://github.com/bianjieai/cosmos-sync/blob/irishub/1.1.0/libs/msgparser/types.go) (default: ``
  example: `bank,nft`)
- deny_modules: `option` `string` disable support module tx sync
- support_types: `option` `string` setting only support msgType tx sync,default support all types(default: ``
  example: `transfer,recv_packet`)
- ignore_ibc_header: `option` `boolean` setting update_client header info for tx collection ,default not ignore ibc header info(default: `false`
    example: `false`)
- chain_id: `option` `string` setting collection name by chain_id
- use_node_urls: `option` `boolean` use setting full node uri set:`true` or get node uri from github.com set:`false`  (default: `false`)

Note:
> synchronizes cosmos data from specify block height(such as:17908 current time:1576208532)
> At first,stop the cosmos-sync and create the task. Run:

  ```
  db.sync_task.insert({
      'start_height':NumberLong(17908),
      'end_height':NumberLong(0),
      'current_height':NumberLong(0),
      'status':'unhandled',
      'worker_id' : '',
       'worker_logs' : [],
      'last_update_time' : NumberLong(0)
  })
  ```
