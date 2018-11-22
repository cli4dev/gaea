
package supplier
import (
	"fmt"
	"github.com/micro-plat/gaea/modules/const/sql"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

//CreateSysSupplierInfo 创建供货商信息 
type CreateSysSupplierInfo struct {		
	//SpName 供货商名称
	SpName string `json:"sp_name" form:"sp_name" m2s:"sp_name" valid:"required"`
	//Balance 账户余额
	Balance float64 `json:"balance" form:"balance" m2s:"balance" valid:"required"`
	//ConponPrdMq 制券队列
	ConponPrdMq string `json:"conpon_prd_mq" form:"conpon_prd_mq" m2s:"conpon_prd_mq" valid:"required"`
		
}

//UpdateSysSupplierInfo 添加供货商信息 
type UpdateSysSupplierInfo struct {	
	//SpId 供货商编号
	SpId int64 `json:"sp_id" form:"sp_id" m2s:"sp_id" valid:"required"`
	//SpName 供货商名称
	SpName string `json:"sp_name" form:"sp_name" m2s:"sp_name" valid:"required"`
	//Status 状态(0 上架 1 下架)
	Status int `json:"status" form:"status" m2s:"status" valid:"required"`
		
}

//QuerySysSupplierInfo 查询供货商信息 
type QuerySysSupplierInfo struct {		
	//Status 状态(0 上架 1 下架)
	Status int `json:"status" form:"status" m2s:"status" valid:"required"`
		
}
//IDbSysSupplierInfo  供货商信息接口
type IDbSysSupplierInfo interface {
	//Create 创建
	Create(input *CreateSysSupplierInfo) (error)
	//Get 单条查询
	Get(id int64)(db.QueryRow,error)
	//Query 列表查询
	Query(input *QuerySysSupplierInfo) (db.QueryRows,error)
	//Update 更新
	Update(input *UpdateSysSupplierInfo) (err error)
	//Delete 删除
	Delete(id int64) (err error)
}
//DbSysSupplierInfo 供货商信息对象
type DbSysSupplierInfo struct {
	c component.IContainer
}
//NewDbSysSupplierInfo 创建供货商信息对象
func NewDbSysSupplierInfo(c component.IContainer) *DbSysSupplierInfo {
	return &DbSysSupplierInfo{
		c: c,
	}
}

//Get 查询单条数据供货商信息
func(d *DbSysSupplierInfo) Get(id  int64) (db.QueryRow,error){

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetSysSupplierInfo, map[string]interface{}{
		"sp_id":id,
	})
	if err != nil {
		return nil, fmt.Errorf("获取供货商信息数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取供货商信息列表
func(d *DbSysSupplierInfo) Query(input *QuerySysSupplierInfo) (db.QueryRows,error){

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.QuerySysSupplierInfo, map[string]interface{}{
		"status":input.Status,
		})
	if err != nil {
		return nil, fmt.Errorf("获取供货商信息数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, nil
}
