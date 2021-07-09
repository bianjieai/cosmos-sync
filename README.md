# irita-sync
A server that synchronize irita chain data into a database

# SetUp

# Build And Run

- Build: `make all`
- Run: `make run`
- Cross compilation: `make build-linux`

## Env Variables

- CONFIG_FILE_PATH: `option` `string` config
  file(https://github.com/bianjieai/irita-sync/blob/opb-bsn/config/config.toml) path

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
