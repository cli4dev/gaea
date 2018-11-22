package sql
//GetSysSupplierInfo 查询单条数据供货商信息
const GetSysSupplierInfo = `select sp_id,sp_name,balance,conpon_prd_mq,status 
from sys_supplier_info where sp_id=@sp_id`

//QuerySysSupplierInfo 查询供货商信息列表数据
const QuerySysSupplierInfo = `select sp_id,sp_name,balance,conpon_prd_mq,status 
from sys_supplier_info where &status`
