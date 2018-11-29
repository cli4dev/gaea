package tmpls

//HTMLTpl vue 页面模板
const HTMLTpl = `
<template>
  <div class="panel panel-default">
    <div class="panel-body">
      <el-form ref="form"  :inline="true" class="form-inline pull-left add-qx-bottom">
		    {{range $i,$c:=.querycolumns}}
        <div class="form-group">
          <el-input clearable  placeholder="请输入{{$c.descsimple}}"></el-input>
        </div>
        {{end}}
        <div class="form-group">
          <el-button type="primary" size="small">查询</el-button>
        </div>
        <div class="form-group">
          <el-button type="success" size="small" @click="addShow">添加</el-button>
        </div>
      </el-form>

      <!-- Add Form -->
      <el-dialog title="添加{{.desc}}" width="30%" :visible.sync="dialogAddVisible">
        <el-form :model="addData" :rules="rules" ref="addForm">
          {{range $i,$c:=.createcolumns}}
          <el-form-item label="{{$c.descsimple}}" prop="{{$c.name}}">
            <el-input v-model="addData.{{$c.name}}"  autocomplete="off"></el-input>
          </el-form-item>
          {{end}}
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button size="small" @click="resetForm('addForm')">取 消</el-button>
          <el-button size="small" type="success" @click="add('addForm')">确 定</el-button>
        </div>
      </el-dialog>
      <!--Add Form -->
    </div>
    <el-scrollbar style="height:100%">
      <el-table :data="tableData" border style="width: 100%">
        {{range $i,$c:=.selectcolumns}}
        <el-table-column prop="{{$c.name}}" label="{{$c.descsimple}}" ></el-table-column>
        {{end}}
        <el-table-column  label="操作">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="editShow(scope.row)">编辑</el-button>
            <el-button type="text" size="small" @click="del(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- Form -->
      <el-dialog title="编辑{{.desc}}" width="30%" @closed="closed" :visible.sync="dialogFormVisible">
        <el-form :model="editData">

          {{range $i,$c:=.updatecolumns}}
          <el-form-item label="{{$c.descsimple}}">
            <el-input v-model="editData.{{$c.name}}"  autocomplete="off"></el-input>
          </el-form-item>
          {{end}}

        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button size="small" @click="dialogFormVisible = false">取 消</el-button>
          <el-button type="success" size="small" @click="edit">确 定</el-button>
        </div>
      </el-dialog>
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
    dialogAddVisible:false,
    dialogFormVisible:false,
		pageSizes: [10, 20, 50, 100],
		params:{pi:1,ps:10},
    totalcount: 0,
    editData:{},
    addData:{},
    rules: {
      {{range $i,$c:=.createcolumns -}}
      {{$c.name|js}}: [
        { required: true, message: "请输入{{$c.descsimple}}", trigger: "blur" }
      ],
      {{end -}}
    },
		tableData: [{
      {{range $i,$c:=.selectcolumns -}}
      {{$c.name|js}}: "{{$c.name}}1",
      {{end -}}
		},{
      {{range $i,$c:=.selectcolumns -}}
      {{$c.name|js}}: "{{$c.name}}2",
      {{end -}}
		},{
      {{range $i,$c:=.selectcolumns -}}
      {{$c.name|js}}: "{{$c.name}}3",
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
    resetForm(formName) {
      this.dialogAddVisible = false
      this.$refs[formName].resetFields();
    },
    closed(){
      console.log("colsed")
    },
    addShow(){
      this.addData = {}
      this.dialogAddVisible = true;
    },
    add(formName){
      console.log(this.addData)
      this.$refs[formName].validate((valid) => {
        if (valid) {
          alert('submit!');
          this.dialogAddVisible = false;
        } else {
          console.log('error submit!!');
          return false;
        }
      });
      
    },
    editShow(val){
      console.log(val);
      this.editData = val
      this.dialogFormVisible = true;
    },
    edit(){
      console.log(this.editData);
      this.dialogFormVisible = false;
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
