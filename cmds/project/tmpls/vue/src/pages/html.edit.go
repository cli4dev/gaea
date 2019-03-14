package pages

// HTMLEditTpl .
const HTMLEditTpl = `
{{$select := "select" -}}
{{$textarea := "textarea" -}}
{{$empty := "" -}}
{{$dt := "date-picker" -}}
<template>
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
</template>

<script>
    export default {
		name: "{{.name|lName}}.edit",
        data() {
            return {
                dialogFormVisible: false,    //编辑表单显示隐藏
                editData: {},                //编辑数据对象
            }
        },
        props: {
            refresh: {
                type: Function,
                default: () => {
                },
            }
        },
        methods: {

            closed() {
                console.log("colsed");
                this.refresh()
            },
            show() {
                this.dialogFormVisible = true;
            },
            edit() {
                console.log(this.editData);
                this.$put("{{.path}}", this.editData)
                    .then(res => {
                        this.$message({
                            type: "success",
                            message: "修改成功!"
                        });
                        this.dialogFormVisible = false;
                        this.refresh()
                    })
                    .catch(err => {
                        this.$message({
                            type: "error",
                            message: err.response.data
                        });
                    })
            },
        }
    }
</script>

<style scoped>

</style>
`
