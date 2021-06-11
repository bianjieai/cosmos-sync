# irita-sync
A server that synchronize irita chain data into a database

# SetUp

# Build And Run

- Build: `make all`
- Run: `make run`
- Cross compilation: `make build-linux`

## Env Variables

### Db config

- DB_ADDR: `required` `string` db addr（example: `127.0.0.1:27017, 127.0.0.2:27017, ...`）
- DB_USER: `required` `string` db user（example: `user`）
- DB_PASSWD: `required` `string` db password（example: `DB_PASSWD`）
- DB_DATABASE：`required` `string` database name（example：`DB_DATABASE`）

### Server config

- SER_BC_FULL_NODES: `required` `string`  full node uri（example: `tcp://127.0.0.1:26657, tcp://127.0.0.2:26657, ...`）
- WORKER_NUM_EXECUTE_TASK: `required` `string` the maximum time (in seconds) that synchronization TX threads are allowed to be out of work（example: `30`）
- WORKER_MAX_SLEEP_TIME: `required` `string` num of worker to create tasks(unit: seconds)（example: `90`）
- BLOCK_NUM_PER_WORKER_HANDLE: `required` `string`  number of blocks per sync TX task（example: `50`）

- CHAIN_BLOCK_INTERVAL: `option` `string` block interval; default `5` (example: `5`)
- BEHIND_BLOCK_NUM: `option` `string` wait block num to handle when retry failed; default `0` (example: `0`)
- PROMETHOUS_PORT: `option` `string` promethous metrics server port

Note: 
> synchronizes irita data from specify block height(such as:17908 current time:1576208532)
  At first,stop the irita-sync and create the task. 
  Run:
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
