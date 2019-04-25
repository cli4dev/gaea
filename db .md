<!-- 
C: 创建（有默认值的字段比如创建日期，不需要配置C，直接设置默认值）
R: 单条数据读取：比如编辑，预览时列出此字段 
U: 修改字段（确定此字段需要前端传入）
D:不使用
Q: 查询条件 
L：列表里列出的字段
DI: 字典编号
DN: 字典名称
DP：字典参数
DD: 下拉选项
-->
# 投票活动原型系统

## 活动类

### 1. 活动表[atv_activity_info]

| 字段名         | 类型           | 默认值 | 为空  |                约束                | 描述                              |
| -------------- | -------------- | :----: | :---: | :--------------------------------: | :-------------------------------- |
| atv_id         | number(10)     |        |  否   |            PK,CL,SEQ,DI            | 活动编号                          |
| appid          | varchar2(64)   |        |  否   |                 CL                 | appid                             |
| title          | varchar2(128)  |        |  否   |              CUQL,DN               | 活动标题                          |
| tag            | varchar2(32)   |        |  否   |                CUL                 | 活动标识                          |
| url            | varchar2(128)  |        |  否   |                 CL                 | 活动链接                          |
| expire_start   | date           |        |  否   |                CUL                 | 开始时间                          |
| expire_end     | date           |        |  否   |                CUL                 | 结束时间                          |
| member_count   | number(10)     |        |  否   |                 CL                 | 参赛人数                          |
| voted_count    | number(10)     |        |  否   |                 CL                 | 投票总数                          |
| involved_count | number(10)     |        |  否   |                 CL                 | 投票人数                          |
| visit_count    | number(10)     |        |  否   |                 CL                 | 访问总数                          |
| gift_count     | number(10)     |        |  否   |                 CL                 | 礼物总数                          |
| create_time    | date           |        |  否   |                CQL                 | 创建时间                          |
| update_time    | date           |        |  否   |                 CL                 | 更新时间                          |
| voting_rule    | number(10)     |        |  否   | CQUL,DD(vts_dics_info,voting_rule) | 投票规则(1:每日可投,2:只能投一票) |
| operator       | varchar2(32)   |        |  否   |                 CL                 | 操作人                            |
| status         | number(2)      |   0    |  否   |   CQUL,DD(vts_dics_info,status)    | 状态（1：禁用,0：启用）           |
| music          | varchar2(256)  |        |  是   |                CUL                 | 背景音乐                          |
| rule           | varchar2(2048) |        |  是   |                CUL                 | 活动规则                          |
| intro          | varchar2(2048) |        |  是   |                CUL                 | 活动介绍                          |
| prize          | varchar2(2048) |        |  是   |                CUL                 | 奖品明细                          |

### 2. 活动宣传页[atv_banner_page]

| 字段名      | 类型           | 默认值 | 为空  |              约束              | 描述                       |
| ----------- | -------------- | :----: | :---: | :----------------------------: | :------------------------- |
| atv_page_id | number(10)     |        |  否   |           PK,CL,SEQ            | 页面编号                   |
| atv_id      | number(10)     |        |  否   |   CQL,DD(atv_activity_info)    | 活动编号                   |
| page_index  | number(2)      |        |  否   |              CUL               | 页索引                     |
| bg_img_url  | varchar2(128)  |        |  是   |              CUL               | 背景图片                   |
| content     | varchar2(2048) |        |  是   |              CUL               | 页面内容                   |
| is_show     | number(1)      |        |  否   | CQUL,DD(vts_dics_info,is_show) | 是否显示 (0:显示,1:不显示) |

### 3. 活动礼物信息[atv_gift_info]

| 字段名       | 类型         | 默认值 | 为空  |             约束             | 描述                |
| ------------ | ------------ | :----: | :---: | :--------------------------: | :------------------ |
| atv_gift_id  | number(10)   |        |  否   |         PK,CL,SEQ,DI         | 礼物编号            |
| atv_id       | number(10)   |        |  否   |  CQL,DD(atv_activity_info)   | 活动编号            |
| vts_prod_id  | number(10)   |        |  否   |             CQL              | 商品编号            |
| ticket_count | number(10)   |        |  否   |             CUQ              | 可抵票数            |
| price        | number(10)   |        |  否   |             CUL              | 售价                |
| show_name    | varchar2(64) |        |  否   |            CUL,DN            | 显示名称            |
| status       | number(1)    |        |  否   | CUL,DD(vts_dics_info,status) | 状态(0 上架,1 下架) |

### 4. 活动参赛者信息[atv_member_info]

| 字段名           | 类型           | 默认值 | 为空  |             约束              | 描述                |
| ---------------- | -------------- | :----: | :---: | :---------------------------: | :------------------ |
| atv_member_id    | number(10)     |        |  否   |           PK,CL,SEQ           | 参赛编号            |
| atv_id           | number(10)     |        |  否   |   CQL,DD(atv_activity_info)   | 活动编号            |
| u_id             | number(10)     |        |  否   |              CL               | 用户编号            |
| member_no        | number(10)     |        |  否   |              CL               | 参赛序号            |
| member_name      | varchar2(64)   |        |  否   |              CL               | 用户姓名            |
| head_img         | varchar2(64)   |        |  否   |              CL               | 头像                |
| mobile           | varchar2(32)   |        |  是   |              CL               | 手机号              |
| voted_count      | number(10)     |        |  否   |              CL               | 投票总数            |
| involved_count   | number(10)     |        |  否   |              CL               | 投票人数            |
| gift_count       | number(10)     |        |  否   |              CL               | 礼物总数            |
| gift_voted_count | number(10)     |        |  否   |              CL               | 礼物票数            |
| visit_count      | number(10)     |        |  否   |              CL               | 访问总数            |
| intro            | varchar2(2048) |        |  是   |              CL               | 简介                |
| status           | number(1)      |        |  否   | CQUL,DD(vts_dics_info,status) | 状态(0 正常,1 禁用) |

### 5. 参赛者投票明细[atv_member_voting]

| 字段名        | 类型         | 默认值 | 为空  |           约束            | 描述     |
| ------------- | ------------ | :----: | :---: | :-----------------------: | :------- |
| atv_voting_id | number(10)   |        |  否   |         PK,CL,SEQ         | 投票明细 |
| atv_id        | number(10)   |        |  否   | CQL,DD(atv_activity_info) | 活动编号 |
| atv_member_id | number(10)   |        |  否   |            CQL            | 参赛编号 |
| voting_uid    | number(10)   |        |  否   |            CL             | 投票用户 |
| voting_ip     | varchar2(32) |        |  否   |            CL             | 投票 IP  |
| voted_count   | number(10)   |   1    |  否   |            CL             | 投票数量 |
| create_time   | date         |        |  否   |            CQL            | 投票时间 |

### 6. 参赛者礼物明细[atv_member_gift]

| 字段名           | 类型         | 默认值 | 为空  |           约束            | 描述         |
| ---------------- | ------------ | :----: | :---: | :-----------------------: | :----------- |
| atv_mgift_id     | number(10)   |        |  否   |         PK,CL,SEQ         | 礼物赠送编号 |
| atv_id           | number(10)   |        |  否   | CQL,DD(atv_activity_info) | 活动编号     |
| atv_member_id    | number(10)   |        |  否   | CQL,DD(vts_product_info)  | 商品编号     |
| atv_gift_id      | number(10)   |        |  否   |   CQL,DD(atv_gift_info)   | 活动礼物编号 |
| uid              | number(10)   |        |  否   |            CL             | 参赛编号     |
| gift_count       | number(10)   |        |  否   |            CL             | 礼物数量     |
| gift_voted_count | number(10)   |        |  否   |            CL             | 礼物票数     |
| giver_name       | varchar2(32) |        |  是   |            CL             | 赠送人姓名   |
| comment          | varchar2(32) |        |  是   |            CL             | 赠送人留言   |

## 交易类

### 1. 交易订单[trd_order_main]

| 字段名              | 类型           | 默认值  | 为空  |                约束                | 描述                                                             |
| ------------------- | -------------- | :-----: | :---: | :--------------------------------: | :--------------------------------------------------------------- |
| order_no            | number(20)     |         |  否   |             PK,CL,SEQ              | 主订单号                                                         |
| appid               | varchar2(64)   |         |  否   |                CQL                 | appid                                                            |
| atv_id              | number(10)     |         |  否   |     CQL,DD(atv_activity_info)      | 活动编号                                                         |
| atv_member_id       | number(10)     |         |  否   |                 CL                 | 参赛者编号                                                       |
| atv_gift_id         | number(10)     |         |  否   |       CQL,DD(atv_gift_info)        | 礼物编号                                                         |
| vts_prod_id         | number(10)     |         |  否   |      CQL,DD(vts_product_info)      | 商品编号                                                         |
| buyer_uid           | number(10)     |         |  否   |        CQL,DD(us_user_info)        | 用户编号                                                         |
| num                 | number(10)     |    1    |  否   |                 CL                 | 数量                                                             |
| amount              | number(20,5)   |         |  否   |                 CL                 | 支付金额                                                         |
| cost                | number(20,5)   |    0    |  否   |                 CL                 | 成本金额                                                         |
| profit              | number(20,5)   |    0    |  否   |                 CL                 | 总利润                                                           |
| status              | number(2)      |         |  否   | CQL,DD(vts_dics_info,order_status) | 订单状态（10.待支付 20.已支付 90.交易关闭　 91.用户取消 0 成功） |
| wx_request_msg      | varchar2(64)   |         |  是   |                 CL                  | 微信请求结果                                                     |
| wx_notify_msg       | varchar2(64)   |         |  是   |                 CL                  | 微信通知结果                                                     |
| create_time         | date           | sysdate |  否   |                 CL                 | 创建时间                                                         |
| payment_start_time  | date           |         |  是   |                 CL                  | 支付开始时间                                                     |
| payment_finish_time | date           |         |  是   |                 CL                 | 支付完成时间                                                     |
| order_finish_time   | date           |         |  是   |                 CL                 | 订单完成时间                                                     |
| buyer_name          | varchar2(32)   |         |  是   |                 CL                  | 购买人姓名                                                       |
| buyer_comment       | varchar2(32)   |         |  是   |                 CL                  | 购买人留言                                                       |
| snapshot            | varchar2(2048) |         |  否   |                 CL                  | 交易快照                                                         |

### 基础信息类

### 1. 用户基本信息[us_user_info]

| 字段名      | 类型          | 默认值  | 为空  |             约束              | 描述                           |
| ----------- | ------------- | :-----: | :---: | :---------------------------: | :----------------------------- |
| uid         | number(10)    |         |  否   |         PK,CL,SEQ,DI          | 主键                           |
| appid       | varchar2(64)  |         |  否   |            CL,UNQ             | appid                          |
| openid      | varchar2(32)  |         |  否   |            CL,UNQ             | openid                         |
| nick_name   | varchar2(32)  |         |  是   |            CQL,DN             | 昵称                           |
| head_url    | varchar2(128) |         |  是   |              CL               | 头像地址                       |
| phone       | varchar2(16)  |         |  是   |              CL               | 联系电话                       |
| e_mail      | varchar2(64)  |         |  是   |              CL               | 邮箱地址                       |
| status      | number(1)     |         |  否   | CQUL,DD(vts_dics_info,status) | 状态（0：启用 1.禁用 2.锁定）  |
| create_time | date          | sysdate |  否   |              CL               | 创建时间                       |
| update_time | date          | sysdate |  否   |              CL               | 更新时间                       |
| is_fans     | number(1)     |    1    |  否   |              CL               | 是否是粉丝（0：是,1：否）      |
| source      | number(1)     |    0    |  否   |              CL               | 用户来源（0微信关注,1 二维码） |
| referee     | varchar2(32)  |         |  是   |              CL               | 推荐人                         |

### 2. 产品信息[vts_product_info]

| 字段名         | 类型           | 默认值 | 为空  |             约束              | 描述              |
| -------------- | -------------- | :----: | :---: | :---------------------------: | :---------------- |
| vts_prod_id    | number(10)     |        |  否   |         PK,CL,SEQ,DI          | 商品编号          |
| product_name   | varchar2(64)   |        |  否   |            CUL,DN             | 产品名称          |
| original_price | number(10,5)   |        |  否   |              CUL              | 原价              |
| cost_price     | number(10,5)   |        |  否   |              CUL              | 成本价            |
| image_url      | varchar2(256)  |        |  否   |              CUL              | 图片地址          |
| description    | varchar2(1024) |        |  否   |              CUL              | 使用说明          |
| stock_count    | number(10)     |   0    |  否   |              CUL              | 库存数量          |
| status         | number(1)      |        |  否   | CQUL,DD(vts_dics_info,status) | 状态(0上架,1下架) |

### 3. 公众号信息[wechat_app_info]

| 字段名            | 类型         | 默认值  | 为空  | 约束  | 描述                   |
| ----------------- | ------------ | :-----: | :---: | :---: | :--------------------- |
| appid             | varchar2(64) |         |  否   | PK,CL | appid                  |
| secret            | varchar2(32) |         |  否   |  CUL  | secret                 |
| token             | varchar2(32) |         |  否   |  CUL  | token                  |
| aes_key           | varchar2(32) |         |  否   |  CUL  | aes key                |
| update_time       | date         | sysdate |  否   |  CL   | 最后更新时间           |
| mchid             | varchar2(64) |         |  否   |  CUL  | 支付服务商编号         |
| pay_key           | varchar2(64) |         |  否   |  CUL  | 支付服务商key          |
| sub_appid         | varchar2(64) |         |  否   |  CUL  | 子商户app id           |
| sub_mchid         | varchar2(64) |         |  否   |  CUL  | 子商户编户号           |
| wechat_host       | varchar2(64) |         |  是   |  CUL  | 微信授权域名           |
| wcode_template_id | varchar2(64) |         |  是   |  CUL  | 微信验证码模板消息编号 |

### 4. 数据字典表[vts_dics_info]

| 字段名     | 类型         | 默认值 | 为空  |    约束    | 描述                |
| ---------- | ------------ | :----: | :---: | :--------: | :-----------------|
| id         | number(10)   |        |  否   | PK,CL,SEQ  | 编号                |
| name       | varchar2(32) |        |  否   |  CQUL,DN   | 名称                |
| value      | varchar2(32) |        |  否   | CUL,UNQ,DI | 值                  |
| type       | varchar2(32) |        |  否   | CUL,UNQ,DP | 分类                |
| sort_id    | number(10)   |   0    |  否   |     CL     | 排序编号            |
| group_code | varchar2(32) |        |  否   |     CL     | 分组编号            |
| status     | number(2)    |   0    |  否   |     CL     | 状态(0 启用,1 禁用) |
