
### 平台配置

#### 2. 平台信息[plat_base_info]

| 字段名       | 类型          | 默认值   | 为空 |                   约束                   | 描述                |
| ----------- | ------------ | :-----: | :--: | :--------------------------------------:| :------------------ |
| id     | number(10)   |         |  否  |              PK,CRL,SEQ,DI               | 平台编号            |
| plat_name   | varchar2(32) |         |  否  |                 CRUQL,DN                 | 平台名称            |
| node_name   | varchar2(32) |         |  否  |                UNQ,CRUQL                 | 平台节点名称        |
| cluster_id  | number(10)   |         |  否  |                  CRUQL                   | 集群编号            |
| remark      | varchar2(32) |         |  是  |                  CRUQL                   | 备注                |
| status      | number(1)    |    0    |  否  | CRUQL,DP,DD(base_dictionary_info,status) | 状态(0 启用 1 禁用) |
| create_time | date         | sysdate |  否  |                   CRQL                   | 创建时间            |
| update_time | date         | sysdate |  是  |                   RQUL                   | 更新时间            |
| operator    | number(10)   |         |  否  |                  CRUQL                   | 操作人              |

#### 3. 系统信息[plat_system_info]

| 字段名        | 类型          | 默认值   | 为空 |                 约束                  | 描述                   |
| ------------ | ------------ | :-----: | :--: | :-----------------------------------: | :---------------------|
| system_id    | number(10)   |         |  否  |             PK,CRL,SEQ,DI             | 系统编号               |
| system_name  | varchar2(32) |         |  否  |               CQURL,DN                | 系统名称               |
| plat_id      | number(10)   |         |  否  |     UNQ:1,CRQL,DD(plat_base_info)     | 平台编号               |
| program_name | varchar2(32) |   '-'   |  否  |              UNQ:1,CQURL              | 系统程序名字           |
| node_name    | varchar2(32) |   '-'   |  否  |              UNQ:1,CQURL              | 注册中心服务名节点名字    |
| remark       | varchar2(32) |         |  是  |                 CURL                  | 备注                   |
| status       | number(1)    |    0    |  否  | CQURL,DD(base_dictionary_info,status) | 状态(0 启用 1 禁用)    |
| create_time  | date         | sysdate |  否  |                 RL,DT                 | 创建时间               |
| update_time  | date         | sysdate |  是  |                URL,DT                 | 更新时间               |
| operator     | number(10)   |         |  否  |                 CURL                  | 操作人                 |

### 15. 数据字典[base_dictionary_info]

| 字段名      | 类型          | 默认值  | 为空 |    约束      | 描述                |
| ---------- | ------------ | :----: | :--: | :---------: | :------------------|
| did        | number(10)   |        |  否  | PK,CRL,SEQ  | 编号                |
| name       | varchar2(32) |        |  否  |   CURL,DN   | 名称                |
| value      | varchar2(32) |        |  否  | CURL,UNQ,DI | 值                  |
| type       | varchar2(32) |        |  否  | CRL,UNQ,DP  | 分类                |
| sort_id    | number(10)   |   0    |  否  |    CURL     | 排序编号            |
| group_code | varchar2(32) |  '\*'  |  否  |    CURL     | 分组编号            |
| status     | number(1)    |   0    |  否  |    CURL     | 状态(0 启用 1 禁用) |
