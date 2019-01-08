

###  消息模板[base_message_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| mid | int(11)||否| PRI | 消息编号                         |
| message_name | varchar(32)||否|  | 消息名称                         |
| message_tag | varchar(32)||否| UNI | 消息 tag,和 influxdb 表对应                         |
| message_group | int(11)||否|  | 消息分组（运维组、研发组）                         |
| message_type | int(11)||否|  | 消息类型(cup、内存、硬盘、网络、打开文件数....)                         |
| alarm_type | varchar(32)||否|  | 报警类型(默认值)(微信、短信、邮件、语音)多个使用逗号分割                         |
| alarm_value | varchar(32)||是|  | 阀值（默认值）                         |
| status | tinyint(1)|0|否|  | 状态(0 启用 1 禁用)                         |
| create_time | datetime|CURRENT_TIMESTAMP|否|  | 创建时间                         |
| update_time | datetime|CURRENT_TIMESTAMP|是|  | 更新时间                         |
| operator | int(11)||否|  | 操作人                         |

###  [conf_db_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| sql_id | int(11)||否| PRI | 数据库验证编号                         |
| sql_name | varchar(32)||否|  | 数据库验证名称                         |
| sys_id | int(11)||否|  | 系统编号                         |
| cc_id | int(11)||否|  | 配置编号                         |
| sql_str | varchar(1024)||否|  | 执行 sqlc                         |
| operator | varchar(32)||否|  | 最后修改人                         |
| update_time | datetime||是|  | 最后修改时间                         |

###  系统 注册中心 表[conf_register_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| conf_id | bigint(20)||否| PRI | 服务编号                         |
| conf_name | varchar(32)||否|  | 服务名称                         |
| sys_id | bigint(20)||否| MUL | 系统编号                         |
| type | varchar(32)||否|  | 注册中心类型(zookeeper,redis 等...)                         |
| plat_tag | varchar(32)||否|  | 注册中心 平台 tag                         |
| service_type | varchar(32)||否|  | 注册中心 服务类型                         |
| sys_tag | varchar(32)||否|  | 注册中心 系统 tag                         |
| cluster_tag | varchar(32)||否|  | 注册中心 集群 tag                         |
| cc_id | int(11)||否|  | 配置编号                         |
| operator | int(11)||否|  | 最后修改人                         |
| update_time | datetime||否|  | 最后修改时间                         |

###  报警消息[alarm_message_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| id | bigint(20)||否| PRI | 编号                         |
| plat_id | int(11)||否|  | 平台编号                         |
| sys_id | int(11)||是|  | 系统编号                         |
| server_id | int(11)||是|  | 服务器编号                         |
| mid | int(11)||否|  | 消息编号                         |
| map_id | int(11)||否|  | 映射编号                         |
| is_alarm | tinyint(1)||否|  | 报警类型 0 报警 1 恢复                         |
| create_time | datetime|CURRENT_TIMESTAMP|否|  | 创建时间                         |
| content | varchar(128)||否|  | 报警内容                         |

###  系统消息配置表[sys_message_map]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| map_id | int(11)||否| PRI | 映射编号                         |
| plat_id | int(11)||否| MUL | 平台编号                         |
| sys_id | int(11)||否|  | 系统编号                         |
| mid | int(11)||否|  | 消息编号                         |
| config_id | int(11)||否|  | 配置编号                         |
| alarm_value | varchar(32)||否|  | 阀值                         |
| operator | int(11)||否|  | 最后修改人                         |
| update_time | datetime||否|  | 最后修改时间                         |

###  人员消息订阅表[message_subscriblist]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| sub_id | int(11)||否| PRI | 编号                         |
| uid | int(11)||否| MUL | 用户编号                         |
| mid | int(11)||否|  | 消息编号                         |
| alarm_type | varchar(32)||否|  | 报警类型(微信、短信、邮件、语音)多个使用逗号分割                         |
| operator | int(11)||否|  | 最后修改人                         |
| update_time | datetime||否|  | 最后修改时间                         |

###  系统信息[base_sys_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| sys_id | int(11)||否| PRI | 系统编号                         |
| plat_id | int(11)||否| MUL | 平台编号                         |
| sys_tag | varchar(32)||否|  | 系统标志，和 influxdb 系统字段对应                         |
| sys_name | varchar(32)||否|  | 系统名称                         |
| remark | varchar(32)||是|  | 备注                         |
| status | tinyint(1)|0|否|  | 状态(0 启用 1 禁用)                         |
| create_time | datetime|CURRENT_TIMESTAMP|否|  | 创建时间                         |
| update_time | datetime|CURRENT_TIMESTAMP|是|  | 更新时间                         |
| operator | int(11)||否|  | 操作人                         |

###  系统 http 表[conf_http_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| http_id | bigint(20)||否| PRI | 服务 http 编号                         |
| http_name | varchar(32)||否|  | 服务 http 名称                         |
| sys_id | int(11)||否|  | 系统编号                         |
| http_addr | varchar(32)||否| UNI | 访问地址                         |
| operator | int(11)||否|  | 最后修改人                         |
| update_time | datetime||否|  | 最后修改时间                         |

###  数据字典表[base_dictionary_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| id | int(11)||否| PRI | 编号                         |
| name | varchar(32)||否|  | 名称                         |
| value | varchar(32)||否| MUL | 值                         |
| type | varchar(32)||否|  | 分类                         |
| sort_id | int(11)|0|否|  | 排序编号                         |
| group_code | varchar(32)|*|否|  | 分组编号                         |
| status | tinyint(1)|0|否|  | 状态 0 启用 1 禁用                         |

###  系统队列表[conf_queue_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| queue_id | bigint(20)||否| PRI | 队列编号                         |
| queue_name | varchar(32)||否|  | 队列名称                         |
| sys_id | bigint(20)||否| MUL | 系统编号                         |
| queue_tag | varchar(32)||否|  | 队列 tag                         |
| cc_id | int(11)||否|  | 配置编号                         |
| redis_tag | varchar(32)||否|  | redis tag                         |
| operator | int(11)||否|  | 最后修改人                         |
| update_time | datetime||否|  | 最后修改时间                         |

###  系统 tcp 表[conf_tcp_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| tcp_id | bigint(20)||否| PRI | 服务 tcp 编号                         |
| tcp_name | varchar(32)||否|  | 服务 tcp 名称                         |
| sys_id | int(11)||否|  | 系统编号                         |
| tcp_addr | varchar(32)||否| UNI | 访问地址                         |
| operator | int(11)||否|  | 最后修改人                         |
| update_time | datetime||否|  | 最后修改时间                         |

###  系统 http 表[conf_http_info_copy]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| http_id | bigint(20)||否| PRI | 服务 http 编号                         |
| http_name | varchar(32)||否|  | 服务 http 名称                         |
| sys_id | int(11)||否|  | 系统编号                         |
| http_addr | varchar(32)||否| UNI | 访问地址                         |
| operator | int(11)||否|  | 最后修改人                         |
| update_time | datetime||否|  | 最后修改时间                         |

###  报警发送列表[send_message_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| id | bigint(20)||否| PRI | 编号                         |
| msg_id | bigint(20)||否|  | 消息编号                         |
| uid | int(11)||否|  | 人员编号                         |
| create_time | datetime|CURRENT_TIMESTAMP|否|  | 创建时间                         |
| send_mq | varchar(32)||是|  | 发送队列                         |
| alarm_type | varchar(32)||否|  | 报警类型 (微信、短信、邮件、语音)                         |
| send_count | int(11)|0|否|  | 发送次数                         |
| flow_time | datetime||否|  | 超时时间                         |
| status | int(11)|10|否|  | 状态:10 等待发送 20 正在发送 0 发送成功 90 发送失败                         |
| batch_id | bigint(20)|0|否|  | 批次号                         |
| complete_time | datetime||是|  | 完成时间                         |

###  平台人员配置表[plat_member_map]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| map_id | int(11)||否| PRI | 配置编号                         |
| plat_id | int(11)||否| MUL | 平台编号                         |
| uid | int(11)||否|  | 人员编号                         |
| operator | int(11)||否|  | 最后修改人                         |
| update_time | datetime||否|  | 最后修改时间                         |

###  服务器消息配置表[server_message_map]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| map_id | int(11)||否| PRI | 配置编号                         |
| server_id | int(11)||否| MUL | 服务器编号                         |
| mid | int(11)||否|  | 消息编号                         |
| alarm_value | varchar(32)||否|  | 阀值                         |
| operator | int(11)||否|  | 最后修改人                         |
| update_time | datetime||否|  | 最后修改时间                         |

###  人员信息[base_member_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| uid | int(11)||否| PRI | 编号                         |
| user_name | varchar(32)||否|  | 用户名                         |
| role | varchar(32)||否|  | 角色                         |
| message_group | int(11)||否|  | 默认关注的消息分组（0 无、1 运维组、2 研发组）                         |
| email | varchar(32)||是|  | 邮箱                         |
| phone_num | varchar(11)||是|  | 手机号                         |
| open_id | varchar(32)||是|  | open_id                         |
| status | tinyint(1)||否|  | 状态(0 正常 1 禁用)                         |
| sync_time | datetime||否|  | 同步时间                         |

###  系统配置[alarm_type_conf]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| id | bigint(20)||否| PRI | 编号                         |
| alarm_type | varchar(32)||否|  | 报警类型(微信、短信、邮件、语音)                         |
| mq | varchar(32)||否|  | mq                         |
| alarm_name | varchar(32)||否|  | 名称                         |
| conf | varchar(2048)||是|  | 配置信息(json)                         |
| status | tinyint(1)||否|  | 状态(0 启用 1 禁用)                         |
| create_time | datetime||否|  | 创建时间                         |
| update_time | datetime||是|  | 更新时间                         |
| operator | int(11)||否|  | 最后修改人                         |

###  [component_config_list]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| cc_id | int(11)||否| PRI | 编号                         |
| conf_type | varchar(32)||否|  | 类型(influx,zk,redis,db)                         |
| conf_name | varchar(32)||否|  | 名称                         |
| config | varchar(1024)||是|  | 配置                         |

###  平台信息[base_plat_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| plat_id | int(11)||否| PRI | 平台编号                         |
| plat_tag | varchar(32)||否| UNI | 平台 tag,和 influxdb 数据库名称对应                         |
| plat_name | varchar(32)||否|  | 平台名称                         |
| cluster_id | int(11)||否|  | 集群编号                         |
| remark | varchar(32)||是|  | 备注                         |
| status | tinyint(1)|0|否|  | 状态(0 启用 1 禁用)                         |
| create_time | datetime|CURRENT_TIMESTAMP|否|  | 创建时间                         |
| update_time | datetime|CURRENT_TIMESTAMP|是|  | 更新时间                         |
| operator | int(11)||否|  | 操作人                         |

###  [server_seq_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| seq_id | int(11)||否| PRI |                          |
| server_name | varchar(32)||否|  |                          |

###  平台服务器配置表[plat_server_map]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| map_id | bigint(20)||否| PRI | 配置编号                         |
| plat_id | bigint(20)||否| MUL | 平台编号                         |
| server_id | bigint(20)||否|  | 服务器编号                         |
| operator | int(11)||否|  | 最后修改人                         |
| update_time | datetime||否|  | 最后修改时间                         |

###  服务器信息[base_server_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| server_id | int(11)||否| PRI | 服务器编号                         |
| server_name | varchar(32)||否|  | 服务器名称                         |
| server_ip | varchar(32)||否| UNI | 服务器 ip                         |
| remark | varchar(32)||是|  | 备注                         |
| status | tinyint(1)|0|否|  | 状态(0 启用 1 禁用)                         |
| create_time | datetime|CURRENT_TIMESTAMP|否|  | 创建时间                         |
| update_time | datetime|CURRENT_TIMESTAMP|是|  | 更新时间                         |
| operator | int(11)||否|  | 操作人                         |

###  [server_report_status]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| id | int(11)||否| PRI | 编号                         |
| plat_tag | varchar(32)||是|  | 系统tag                         |
| server_ip | varchar(32)||否|  | 服务器ip                         |
| status | int(1)|1|否|  | 绑定状态 0 已绑定 1 未绑定                         |
| update_time | datetime||否|  | 更新时间                         |
| expired_time | datetime||否|  | 过期时间                         |

###  [base_cluster_info]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
| cluster_id | int(11)||否| PRI | 集群编号                         |
| cluster_tag | varchar(32)||否|  | 集群 tag                         |
| cluster_name | varchar(32)||否|  | 集群名称                         |
| cc_id | int(11)||否|  | influxdb 配置                         |
| create_time | datetime|CURRENT_TIMESTAMP|否|  | 创建时间                         |
| update_time | datetime|CURRENT_TIMESTAMP|否|  | 更新时间                         |
| operator | int(11)||否|  | 操作人                         |
