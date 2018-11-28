package tmpls

//HTMLTpl vue 页面模板
const HTMLTpl = `
<template>
  <div class="panel panel-default">
    <div class="panel-body">
      <el-form ref="form"  :inline="true" class="form-inline pull-left add-qx-bottom">
		{{range $i,$c:=.querycolumns}}
        <div class="form-group">
          <el-input clearable placeholder="请输入{{$c.desc}}"></el-input>
        </div>
        {{end}}
        <div class="form-group">
          <el-button type="primary">查询</el-button>
        </div>
      </el-form>
    </div>
    <el-scrollbar style="height:100%">
      <el-table :data="tableData" border style="width: 100%">
        {{range $i,$c:=.selectcolumns}}
        <el-table-column prop="{{$c.name}}" label="{{$c.desc}}" ></el-table-column>
        {{end}}

        <el-table-column  label="操作">
          <template slot-scope="scope">

            <el-button type="text" size="small" @click="edit(scope.row)">编辑</el-button>
            <el-button type="text" size="small" @click="del(scope.row)">删除</el-button>

          </template>
        </el-table-column>

      </el-table>
    </el-scrollbar>
    <div class="page-pagination">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="params.pi"
        :page-size="params.ps"
        :page-sizes="pageSizes"
        layout="total, sizes, prev, pager, next, jumper"
        :total="totalcount">
      </el-pagination>
    </div>

  </div>
</template>

<script>
export default {
	name: "{{.name|cname}}",
	data () {
		return {
		pageSizes: [10, 20, 50, 100],
		params:{pi:1,ps:10},
		totalcount: 0,
		tableData: [{
      {{range $i,$c:=.selectcolumns -}}
      {{$c.name}}: "{{$c.name}}",
      {{end -}}
		}]
		}
	},
	methods:{
		handleSizeChange(val) {
		  this.params.ps = val
		},
    handleCurrentChange(val) {
      this.params.pi = val
    },
    edit(val){
      console.log(val)
    },
    del(val){
      console.log(val)
      this.$confirm("此操作将永久删除该数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        console.log(val);
        this.$message({
          type: "success",
          message: "删除成功!"
        });
      }).catch(() => {
        this.$message({
          type: "info",
          message: "已取消删除"
        });          
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
`
