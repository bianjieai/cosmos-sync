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

### [config.toml](https://github.com/bianjieai/cosmos-sync/blob/cosmos-economy/config/config.toml)

```text
[database]
addrs= "localhost:27018"
user= "iris"
passwd= "irispassword"
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
only_support_module="" 
```

### Db config

- addrs: `required` `string` db addr（example: `127.0.0.1:27017, 127.0.0.2:27017, ...`）
- user: `required` `string` db user（example: `user`）
- passwd: `required` `string` db password（example: `DB_PASSWD`）
- database：`required` `string` database name（example：`DB_DATABASE`）

### Server config

- node_urls: `required` `string`  full node uri（example: `tcp://127.0.0.1:26657, tcp://127.0.0.2:26657, ...`）
- worker_num_create_task: `required` `string` the maximum time (in seconds) that create task are allowed （default: `1`
  example: `1`）
- worker_num_execute_task: `required` `string` the maximum time (in seconds) that synchronization TX threads are allowed
  to be out of work（example: `30`）
- worker_max_sleep_time: `required` `string` num of worker to create tasks(unit: seconds)（example: `90`）
- block_num_per_worker_handle: `required` `string`  number of blocks per sync TX task（example: `50`）

- max_connection_num: `required` `string` client pool config total max connection
- init_connection_num: `required` `string` client pool config idle connection

- bech32_acc_prefix: `option` `string` block chain address prefix（default: `` example: `iaa`）
- chain_block_interval: `option` `string` block interval; default `5` (example: `5`)
- behind_block_num: `option` `string` wait block num to handle when retry failed; default `0` (example: `0`)
- promethous_port: `option` `string` promethous metrics server port
- only_support_module: `option` `string` setting only support module tx sync,default
  support [all module](https://github.com/bianjieai/cosmos-sync/blob/stargate/libs/msgparser/types.go) (default: ``
  example: `bank,nft`)

- chain_id: `option` `string` setting collection name by chain_id

Note:
> synchronizes cosmos data from specify block height(such as:17908 current time:1576208532)
At first,stop the cosmos-sync and create the task. Run:

  ```
  db.sync_task.insert({
      'start_height':NumberLong(17908),
      'end_height':NumberLong(0),
      'current_height':NumberLong(0),
      'status':'unhandled',
      ﻿'worker_id' : '',
       'worker_logs' : [],
      'last_update_time' : NumberLong(0)
  })
  ```
