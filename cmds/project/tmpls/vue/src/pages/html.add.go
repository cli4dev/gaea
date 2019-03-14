package pages

// HTMLAddTpl .
const HTMLAddTpl = `
{{$select := "select" -}}
{{$textarea := "textarea" -}}
{{$empty := "" -}}
{{$dt := "date-picker" -}}
<template>
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
</template>

<script>
    export default {
        name: "{{.name|lName}}.add",
        data() {
            return {
                addData: {},
                dialogAddVisible:false,
				rules: {                    //数据验证规则
					{{range $i,$c:=.createcolumns -}}
					{{$c.name}}: [
					  { required: true, message: "请输入{{$c.descsimple}}", trigger: "blur" }
					],
					{{end -}}
                  },
                {{range $i,$c:=.querycolumns -}}
                {{range $k,$v := $c.source -}}
                {{$k}}:[],
                {{end -}}
                {{- end -}}
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
        methods: {
            resetForm(formName) {
                this.dialogAddVisible = false;
                this.$refs[formName].resetFields();
            },
            show(){
                this.dialogAddVisible = true;
            },
            add(formName) {
                console.log(this.addData);
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
                        console.log("error submit!!");
                        return false;
                    }
                });

            },
        }
    }
</script>

<style scoped>

</style>
`
