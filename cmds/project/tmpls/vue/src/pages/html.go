package pages

//HTMLTpl vue 页面模板
const HTMLTpl = `
{{$select := "select" -}}
{{$textarea := "textarea" -}}
{{$empty := "" -}}
{{$dt := "date-picker" -}}
<template>
  <div class="panel panel-default">
    <div class="panel-body">
      <el-form ref="form"  :inline="true" class="form-inline pull-left">
		    {{range $i,$c:=.querycolumns}}
          {{if eq $c.domType $textarea -}}
          <el-form-item >
            <el-input
            type="textarea"
            :rows="2"
            placeholder="请输入{{$c.descsimple}}"
            v-model="queryData.{{$c.name}}">
            </el-input>
            </el-form-item >
          {{- else -}}
          <el-form-item >
          <el-{{$c.domType}} clearable  v-model="queryData.{{$c.name}}" {{if eq $c.domType $dt -}} value-format="yyyy-MM-dd" {{- end}} placeholder="请输入{{$c.descsimple}}">

          {{if ne $c.domType $select -}}
          {{$c.descsimple}}
          {{- end -}}

          {{if eq $c.domType $select -}}
            <el-option
              v-for="item in {{$c.name}}"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          {{- end}}
          </el-{{$c.domType}}>
          </el-form-item>
          {{- end}}
        
        {{end}}

        <el-form-item >
          <el-button type="primary" @click="query" size="small">查询</el-button>
        </el-form-item >
        <el-form-item >
          <el-button type="success" size="small" @click="addShow">添加</el-button>
        </el-form-item >
      </el-form>

      <!-- Add Form -->
      <el-dialog title="添加{{.desc}}" {{if gt (.createcolumns|len) 4 -}} width="35%" {{else}} width="26%" {{- end}} :visible.sync="dialogAddVisible">
        <el-form :model="addData"  {{if gt (.createcolumns|len) 4 -}}:inline="true"{{- end}} :rules="rules" ref="addForm">
       
        {{range $i,$c:=.createcolumns}}
            {{if eq $c.domType $textarea -}}
            <el-form-item label="{{$c.descsimple}}" prop="{{$c.pname}}">
            <el-input
            type="textarea"
            :rows="2"
            placeholder="请输入{{$c.descsimple}}"
            v-model="addData.{{$c.name}}">
            </el-input>
            </el-form-item>
            {{- else -}}
            <el-form-item label="{{$c.descsimple}}" prop="{{$c.name}}">
            <el-{{$c.domType}} clearable  v-model="addData.{{$c.name}}" {{if eq $c.domType $dt -}} value-format="yyyy-MM-dd HH:mm:ss" {{- end}}   placeholder="请输入{{$c.descsimple}}">
            {{if eq $c.domType $select -}}
              <el-option
                v-for="item in {{$c.name}}"
                :key="item.id"
                :label="item.name"
                :value="item.id">
              </el-option>
            {{- end}}
            </el-{{$c.domType}}>
            </el-form-item>
            {{- end}}
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
            <el-button type="text" size="small" @click="detailShow(scope.row)">详情</el-button>
            <el-button type="text" size="small" @click="del(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- edit Form -->
      <el-dialog title="编辑{{.desc}}" {{if gt (.updatecolumns|len) 4 -}} width="35%" {{else}} width="26%" {{- end}}  @closed="closed" :visible.sync="dialogFormVisible">
        <el-form :model="editData" {{if gt (.updatecolumns|len) 4 -}}:inline="true"{{- end}} >
        
          {{range $i,$c:=.updatecolumns}}
            {{if eq $c.domType $textarea -}}
            <el-form-item label="{{$c.descsimple}}">
            <el-input
            type="textarea"
            :rows="2"
            placeholder="请输入{{$c.descsimple}}"
            v-model="editData.{{$c.name}}">
            </el-input>
            </el-form-item>
            {{- else -}}
           
            <el-form-item label="{{$c.descsimple}}">
            <el-{{$c.domType}} clearable  v-model="editData.{{$c.name}}" {{if eq $c.domType $dt -}} value-format="yyyy-MM-dd HH:mm:ss" {{- end}}   placeholder="请输入{{$c.descsimple}}">
            {{if eq $c.domType $select -}}
              <el-option
                v-for="item in {{$c.name}}"
                :key="item.id"
                :label="item.name"
                :value="item.id">
              </el-option>
            {{- end}}
            </el-{{$c.domType}}>
            </el-form-item>
            {{- end}}
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
    dialogAddVisible:false,     //添加表单显示隐藏
    dialogFormVisible:false,    //编辑表单显示隐藏
		pageSizes: [10, 20, 50, 100], 
		params:{pi:1,ps:10},        //页码，页容量控制
    totalcount: 0,              //数据总条数
    editData:{},                //编辑数据对象
    addData:{},                 //添加数据对象 
    queryData:{},
    {{range $i,$c:=.querycolumns -}}
    {{range $k,$v := $c.source -}}
    {{$k}}:[],
    {{end -}}
    {{- end -}}
    rules: {                    //数据验证规则
      {{range $i,$c:=.createcolumns -}}
      {{$c.name}}: [
        { required: true, message: "请输入{{$c.descsimple}}", trigger: "blur" }
      ],
      {{end -}}
    },
		tableData: [{                //表数据
      {{range $i,$c:=.selectcolumns -}}
      {{$c.name}}: "{{$c.name}}1",
      {{end -}}
		}]
		}
  },
  created(){
    {{range $i,$c:=.querycolumns -}}
    {{range $k,$v:= $c.source}}
      this.$get("{{$v.path}}",{{if ne $v.params "" -}} {{$v.params}} {{- else -}} {} {{- end}})
      .then(res => {
        this.{{$k}}= res
      })
      .catch(err => {
          console.log(err)
      });
    {{end}}
    {{- end}}
  },
  mounted(){
    this.init()
   
  },
	methods:{
    /*
    *初始化操作
    */
    init(){
      this.query()
    },
    detailShow(val){
      val.getpath ="{{.path}}"
      this.$emit('addTab',"详情"+val.{{range $i,$c:=.pk}}{{$c.name}}{{end}},"{{.path}}.view",val);
    },
    /*
    *查询数据并赋值
    */
    query(){
      this.queryData.pi = this.params.pi
      this.queryData.ps = this.params.ps
      this.$get("{{.path}}/query", this.queryData)
          .then(res => {
            this.tableData = res.data
            this.totalcount = res.count
          })
          .catch(err => {
              console.log(err)
          })

    },
    /*
    *改变页容量
    */
		handleSizeChange(val) {
      this.params.ps = val
      this.query()
    },
    /*
    *改变当前页码
    */
    handleCurrentChange(val) {
      this.params.pi = val
      this.query()
    },
    /*
    *重置添加表单
    */
    resetForm(formName) {
      this.dialogAddVisible = false
      this.$refs[formName].resetFields();
    },
    /*
    *编辑表单关闭事件
    */
    closed(){
      console.log("colsed")
      this.query()
    },
    /*
    *添加表单显示
    */
    addShow(){
      this.addData = {}
      this.dialogAddVisible = true;
    },
    /*
    *添加数据提交操作
    */
    add(formName){
      console.log(this.addData)
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$post("{{.path}}", this.addData)
          .then(res => {
            this.dialogAddVisible = false;
            this.query()
          })
          .catch(err => {
              console.log(err)
          })
          
        } else {
          console.log('error submit!!');
          return false;
        }
      });
      
    },
    /**
    *编辑表单显示
    *@val 当前行的数据对象
    */
    editShow(val){
      this.editData = val
      this.dialogFormVisible = true;
    },
    /*
    *编辑数据提交
    */
    edit(){
      console.log(this.editData);
      this.$put("{{.path}}", this.editData)
      .then(res => {
        this.$message({
          type: "success",
          message: "修改成功!"
        });
        this.dialogFormVisible = false;
        this.query()


      })
      .catch(err => {
        this.$message({
          type: "error",
          message: err.response.data
        });
      })
      
    },
    /**
    *删除数据
    *@val 要删除的数据对象
    */
    del(val){
      console.log(val)
      this.$confirm("此操作将永久删除该数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        console.log(val);
        this.$del("{{.path}}", val)
        .then(res => {
          this.$message({
            type: "success",
            message: "删除成功!"
          });
          this.dialogFormVisible = false;
          this.query()
        })
        .catch(err => {
            console.log(err)
        })
       
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
