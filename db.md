# 投票活动原型系统

## 活动类

### 1. 活动表[atv_activity_info]

| 字段名         | 类型           | 默认值 | 为空 |   约束    | 描述                                |
| -------------- | -------------- | :----: | :--: | :-------: | :---------------------------------- |
| atv_id         | number(10)     |        |  否  | PK,IS,SEQ,DI | 活动编号                            |
| appid          | varchar2(64)   |        |  否  |    IS     | appid                               |
| title          | varchar2(128)  |        |  否  |   IUQS,DN | 活动标题                            |
| tag            | varchar2(32)   |        |  否  |    IUS    | 活动标识                            |
| url            | varchar2(128)  |        |  否  |    IS     | 活动链接                            |
| expire_start   | date           |        |  否  |    IUS    | 开始时间                            |
| expire_end     | date           |        |  否  |    IUS    | 结束时间                            |
| member_count   | number(10)     |        |  否  |    IS     | 参赛人数                            |
| voted_count    | number(10)     |        |  否  |    IS     | 投票总数                            |
| involved_count | number(10)     |        |  否  |    IS     | 投票人数                            |
| visit_count    | number(10)     |        |  否  |    IS     | 访问总数                            |
| gift_count     | number(10)     |        |  否  |    IS     | 礼物总数                            |
| create_time    | date           |        |  否  |    IQS    | 创建时间                            |
| update_time    | date           |        |  否  |    IS     | 更新时间                            |
| voting_rule    | number(10)     |        |  否  |    IUS    | 投票规则(1:每日可投　 2:只能投一票) |
| operator       | varchar2(32)   |        |  否  |    IS     | 操作人                              |
| status         | number(2)      |   0    |  否  |    IUS    | 状态(1：禁用 0：启用)             |
| music          | varchar2(256)  |        |  是  |    IUS    | 背景音乐                            |
| rule           | varchar2(2048) |        |  是  |    IUS    | 活动规则                            |
| intro          | nclob          |        |  是  |    IUS    | 活动介绍                            |
| prize          | nclob          |        |  是  |    IUS    | 奖品明细                            |

### 2. 活动宣传页[atv_banner_page]

| 字段名       | 类型           | 默认值  | 为空 |   约束                       | 描述                         |
| ----------- | ------------- | :----: | :--: | :-------------------------: | :---------------------------|
| atv_page_id | number(10)    |        |  否  | PK,IS,SEQ                   | 页面编号                     |
| atv_id      | number(10)    |        |  否  | IQS,DD(atv_activity_info)   | 活动编号                     |
| page_index  | number(2)     |        |  否  |    IUS                      | 页索引                       |
| bg_img_url  | varchar2(128) |        |  是  |    IUS                      | 背景图片                     |
| content     | nclob         |        |  是  |    IUS                      | 页面内容                     |
| is_show     | number(1)     |        |  否  |    IUS                      | 是否显示 (0:显示　 1:不显示)   |

### 3. 活动礼物信息[atv_gift_info]

| 字段名       | 类型         | 默认值 | 为空 |   约束    | 描述                |
| ------------ | ------------ | :----: | :--: | :-------: | :------------------ |
| atv_gift_id  | number(10)   |        |  否  | PK,IS,SEQ | 礼物编号            |
| atv_id       | number(10)   |        |  否  |    IQS    | 活动编号            |
| vts_prod_id  | number(10)   |        |  否  |    IQS    | 商品编号            |
| ticket_count | number(10)   |        |  否  |    IUQ    | 可抵票数            |
| price        | number(10)   |        |  否  |    IUS    | 售价                |
| show_name    | varchar2(64) |        |  否  |    IUS    | 显示名称            |
| status       | number(1)    |        |  否  |    IUS    | 状态(0 上架 1 下架) |

### 4. 活动参赛者信息[atv_member_info]

| 字段名           | 类型         | 默认值 | 为空 |   约束    | 描述                |
| ---------------- | ------------ | :----: | :--: | :-------: | :------------------ |
| atv_member_id    | number(10)   |        |  否  | PK,IS,SEQ | 参赛编号            |
| atv_id           | number(10)   |        |  否  |    IQS    | 活动编号            |
| uid              | number(10)   |        |  否  |    IS     | 用户编号            |
| member_no        | number(10)   |        |  否  |    IS     | 参赛序号            |
| member_name      | varchar2(64) |        |  否  |    IS     | 用户姓名            |
| head_img         | varchar2(64) |        |  否  |    IS     | 头像                |
| mobile           | varchar2(32) |        |  是  |    IS     | 手机号              |
| voted_count      | number(10)   |        |  否  |    IS     | 投票总数            |
| involved_count   | number(10)   |        |  否  |    IS     | 投票人数            |
| gift_count       | number(10)   |        |  否  |    IS     | 礼物总数            |
| gift_voted_count | number(10)   |        |  否  |    IS     | 礼物票数            |
| visit_count      | number(10)   |        |  否  |    IS     | 访问总数            |
| intro            | nclob        |        |  是  |    IS     | 简介                |
| status           | number(1)    |        |  否  |    IS     | 状态(0 正常 1 禁用) |

### 5. 参赛者投票明细[atv_member_voting]

| 字段名        | 类型         | 默认值 | 为空 |   约束    | 描述     |
| ------------- | ------------ | :----: | :--: | :-------: | :------- |
| atv_voting_id | number(10)   |        |  否  | PK,IS,SEQ | 投票明细 |
| atv_id        | number(10)   |        |  否  |    IQS    | 活动编号 |
| atv_member_id | number(10)   |        |  否  |    IS     | 参赛编号 |
| voting_uid    | number(10)   |        |  否  |    IS     | 投票用户 |
| voting_ip     | varchar2(32) |        |  否  |    IS     | 投票 IP  |
| voted_count   | number(10)   |   1    |  否  |    IS     | 投票数量 |
| create_time   | date         |        |  否  |    IQS    | 投票时间 |

### 6. 参赛者礼物明细[atv_member_gift]

| 字段名           | 类型         | 默认值 | 为空 |   约束    | 描述         |
| ---------------- | ------------ | :----: | :--: | :-------: | :----------- |
| atv_mgift_id     | number(10)   |        |  否  | PK,IS,SEQ | 礼物赠送编号 |
| atv_id           | number(10)   |        |  否  |    IQS    | 活动编号     |
| atv_member_id    | number(10)   |        |  否  |    IQS    | 商品编号     |
| atv_gift_id      | number(10)   |        |  否  |    IQS    | 活动礼物编号 |
| uid              | number(10)   |        |  否  |    IS     | 参赛编号     |
| gift_count       | number(10)   |        |  否  |    IS     | 礼物数量     |
| gift_voted_count | number(10)   |        |  否  |    IS     | 礼物票数     |
| giver_name       | varchar2(32) |        |  是  |    IS     | 赠送人姓名   |
| comment          | varchar2(32) |        |  是  |    IS     | 赠送人留言   |

## 交易类

### 1. 交易订单[trd_order_main]

| 字段名              | 类型           | 默认值  | 为空 |   约束    | 描述                                                             |
| ------------------- | -------------- | :-----: | :--: | :-------: | :--------------------------------------------------------------- |
| order_no            | number(20)     |         |  否  | PK,IS,SEQ | 主订单号                                                         |
| appid               | varchar2(64)   |         |  否  |    IQS    | appid                                                            |
| atv_id              | number(10)     |         |  否  |    IQS    | 活动编号                                                         |
| atv_member_id       | number(10)     |         |  否  |    IS     | 参赛者编号                                                       |
| atv_gift_id         | number(10)     |         |  否  |    IS     | 礼物编号                                                         |
| vts_prod_id         | number(10)     |         |  否  |    IQS    | 商品编号                                                         |
| buyer_uid           | number(10)     |         |  否  |    IS     | 用户编号                                                         |
| num                 | number(10)     |   １    |  否  |    IS     | 数量                                                             |
| amount              | number(20,5)   |         |  否  |    IS     | 支付金额                                                         |
| cost                | number(20,5)   |    0    |  否  |    IS     | 成本金额                                                         |
| profit              | number(20,5)   |    0    |  否  |    IS     | 总利润                                                           |
| status              | number(2)      |         |  否  |    IS     | 订单状态(10.待支付 20.已支付 90.交易关闭　 91.用户取消 0 成功) |
| wx_request_msg      | varchar2(64)   |         |  是  |     I     | 微信请求结果                                                     |
| wx_notify_msg       | varchar2(64)   |         |  是  |     I     | 微信通知结果                                                     |
| create_time         | date           | sysdate |  否  |    IS     | 创建时间                                                         |
| payment_start_time  | date           |         |  是  |     I     | 支付开始时间                                                     |
| payment_finish_time | date           |         |  是  |    IS     | 支付完成时间                                                     |
| order_finish_time   | date           |         |  是  |    IS     | 订单完成时间                                                     |
| buyer_name          | varchar2(32)   |         |  是  |     I     | 购买人姓名                                                       |
| buyer_comment       | varchar2(32)   |         |  是  |     I     | 购买人留言                                                       |
| snapshot            | varchar2(2000) |         |  否  |     I     | 交易快照                                                         |

### 基础信息类

### 1. 用户基本信息[us_user_info]

| 字段名      | 类型          | 默认值  | 为空 |   约束    | 描述                             |
| ----------- | ------------- | :-----: | :--: | :-------: | :------------------------------- |
| uid         | number(10)    |         |  否  | PK,IS,SEQ | 主键                             |
| appid       | varchar2(64)  |         |  否  |  IS,UNQ   | appid                            |
| openid      | varchar2(32)  |         |  否  |  IS,UNQ   | openid                           |
| nick_name   | varchar2(32)  |         |  是  |    IQS    | 昵称                             |
| head_url    | varchar2(128) |         |  是  |    IS     | 头像地址                         |
| phone       | varchar2(16)  |         |  是  |    IS     | 联系电话                         |
| e_mail      | varchar2(64)  |         |  是  |    IS     | 邮箱地址                         |
| status      | number(1)     |         |  否  |    IS     | 状态(0：启用 1.禁用 2.锁定)    |
| create_time | date          | sysdate |  否  |    IS     | 创建时间                         |
| update_time | date          | sysdate |  否  |    IS     | 更新时间                         |
| is_fans     | number(1)     |    1    |  否  |    IS     | 是否是粉丝(0：是 1：否)        |
| source      | number(1)     |    0    |  否  |    IS     | 用户来源(0.微信关注 1. 二维码) |
| referee     | varchar2(32)  |         |  是  |    IS     | 推荐人                           |

### 2. 产品信息[vts_product_info]

| 字段名         | 类型           | 默认值 | 为空 |   约束    | 描述                |
| -------------- | -------------- | :----: | :--: | :-------: | :------------------ |
| vts_prod_id    | number(10)     |        |  否  | PK,IS,SEQ | 商品编号            |
| product_name   | varchar2(64)   |        |  否  |    IUS    | 产品名称            |
| original_price | number(10,5)   |        |  否  |    IUS    | 原价                |
| cost_price     | number(10,5)   |        |  否  |    IUS    | 成本价              |
| image_url      | varchar2(256)  |        |  否  |    IUS    | 图片地址            |
| description    | varchar2(1024) |        |  否  |    IUS    | 使用说明            |
| stock_count    | number(10)     |   0    |  否  |    IUS    | 库存数量            |
| status         | number(1)      |        |  否  |    IUS    | 状态(0 上架 1 下架) |

### 3. 公众号信息[wechat_app_info]

| 字段名            | 类型         | 默认值  | 为空 | 约束  | 描述                   |
| ----------------- | ------------ | :-----: | :--: | :---: | :--------------------- |
| appid             | varchar2(64) |         |  否  | PK,IS | appid                  |
| secret            | varchar2(32) |         |  否  |  IUS  | secret                 |
| token             | varchar2(32) |         |  否  |  IUS  | token                  |
| aes_key           | varchar2(32) |         |  否  |  IUS  | aes key                |
| update_time       | date         | sysdate |  否  |  IS   | 最后更新时间           |
| mchid             | varchar2(64) |         |  否  |  IUS  | 支付服务商编号         |
| pay_key           | varchar2(64) |         |  否  |  IUS  | 支付服务商 key         |
| sub_appid         | varchar2(64) |         |  否  |  IUS  | 子商户 app id          |
| sub_mchid         | varchar2(64) |         |  否  |  IUS  | 子商户编户号           |
| wechat_host       | varchar2(64) |         |  是  |  IUS  | 微信授权域名           |
| wcode_template_id | varchar2(64) |         |  是  |  IUS  | 微信验证码模板消息编号 |

### 4. 数据字典表[vts_dics_info]

| 字段名     | 类型         | 默认值 | 为空 |   约束    | 描述                    |
| ---------- | ------------ | :----: | :--: | :-------: | :----------------- |
| id         | number(10)   |        |  否  | PK,IS,SEQ | 编号                |
| name       | varchar2(32) |        |  否  |    IUS    | 名称                |
| value      | varchar2(32) |        |  否  |  IUS,UNQ  | 值                  |
| type       | varchar2(32) |        |  否  |  IS,UNQ   | 分类                |
| sort_id    | number(10)   |   0    |  否  |    IS     | 排序编号             |
| group_code | varchar2(32) |        |  否  |    IS     | 分组编号             |
| status     | number(2)    |   0    |  否  |    IS     | 状态 0 启用 1 禁用    |
