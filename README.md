# irita-sync
A server that synchronize irita chain data into a database

# SetUp

# Build And Run

- Build: `make all`
- Run: `make run`
- Cross compilation: `make build-linux`

## Env Variables

### Db config

- DB_ADDR: `required` `string` 数据库连接地址（example: `127.0.0.1:27017, 127.0.0.2:27017, ...`）
- DB_USER: `required` `string` 数据库连接用户（example: `user`）
- DB_PASSWD: `required` `string` 数据库密码（example: `DB_PASSWD`）
- DB_DATABASE：`required` `string` 数据库名称（example：`DB_DATABASE`）

### Server config

- SER_BC_FULL_NODE: `required` `string`  全节点地址（example: `tcp://127.0.0.1:26657, tcp://127.0.0.2:26657, ...`）
- WORKER_NUM_EXECUTE_TASK: `required` `string` 执行任务执行的线程数（example: `30`）
- WORKER_MAX_SLEEP_TIME: `required` `string` 允许线程最大的休眠时间(单位:秒)（example: `90`）
- BLOCK_NUM_PER_WORKER_HANDLE: `required` `string`  每个任务包含的区块数（example: `50`）
- NETWORK: `option` `string` 网络类型（example: `testnet,mainnet`）


Note: 
> synchronizes irishub data from specify block height(such as:17908 current time:1576208532)
  At first,stop the irita-sync and create the task. 
  Run:
  ```
  db.sync_iris_task.insert({
      'start_height':NumberLong(17908),
      'end_height':NumberLong(0),
      'current_height':NumberLong(0),
      'status':'unhandled',
  })
