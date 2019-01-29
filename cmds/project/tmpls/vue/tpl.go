package vue

const Tpl = `<template>
<el-table
	:data="info"
	style="width: 100%" :show-header="false" :default-expand-all="true">
	<el-table-column type="expand">
		<template slot-scope="props">
			<el-form label-position="left" inline class="demo-table-expand">
			{{range $i,$c:=.getcolumns -}}
			<el-form-item label="{{$c.descsimple}}">
				<span v-text="props.row.{{$c.name}}"></span>
			</el-form-item>
			{{- end}}
			</el-form>
		</template>
	</el-table-column>
	<el-table-column
	label="商品 ID"
	>
	</el-table-column>
	<el-table-column
	label="商品名称"
	>
	</el-table-column>
	<el-table-column
	label="描述"
	>
	</el-table-column>
	</el-table>
</template>


<script>
export default {

  data(){
    return {
      info:[],
    }
  },
  mounted() {
    this.Init()
  },
  methods: {
    Init(){
        this.QueryDataList()
    },
    QueryDataList(){
      this.$get(this.$route.query.getpath, 
          this.$route.query
        ).then(res => {
			console.log(res)
			this.info.push(res)
      });
    },
  
    handleClick(tab){
     
    },
   
  },

 
}
</script>

<style>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>
`
