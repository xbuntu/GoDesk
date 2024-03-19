<template>
  <div class="table">
    <el-form :model="queryParams" ref="queryForm" size="small" :inline="true" label-width="68px">
      <el-form-item label="岗位名称" prop="postName">
        <el-input
            v-model="queryParams.postName"
            placeholder="请输入岗位名称"
            clearable
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10">
      <el-col :span="1.5">
        <el-button
            type="primary"
            plain
            icon="el-icon-plus"
            size="mini"
            @click="handleAdd">新增
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
            type="warning"
            plain
            icon="el-icon-download"
            size="mini"
            @click="handleExport">导出
        </el-button>
      </el-col>
    </el-row>

    <el-table height="calc(100vh - 180px)" v-loading="loading" :data="postList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center"/>
      <el-table-column label="岗位名称" align="center" prop="postName"/>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
              size="mini"
              type="text"
              icon="el-icon-edit"
              @click="handleUpdate(scope.row)">修改
          </el-button>
          <el-button
              size="mini"
              type="text"
              icon="el-icon-delete"
              @click="handleDelete(scope.row)">删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
        style="margin-top: 20px"
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="1"
        :page-sizes="[10, 20, 30, 50]"
        :page-size="10"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total">
    </el-pagination>

    <!-- 添加或修改岗位对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="500px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="岗位名称" prop="postName">
          <el-input v-model="form.postName" placeholder="请输入岗位名称"/>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" placeholder="请输入内容"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'TableView',
  data() {
    return {
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        postName: undefined
      },
      // 选中数组
      ids: [],
      total:0,
      // 遮罩层
      loading: true,
      // 岗位表格数据
      postList: [],
      title: "添加",
      open: false,
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        postName: [
          {required: true, message: "岗位名称不能为空", trigger: "blur"}
        ],
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询岗位列表 */
    getList() {
      this.loading = true;
      this.postList = [
          {postId: 1, postName: "岗位1"},
          {postId: 2, postName: "岗位2"},
          {postId: 2, postName: "岗位2"},
          {postId: 2, postName: "岗位2"},
          {postId: 2, postName: "岗位2"},
          {postId: 2, postName: "岗位2"},
          {postId: 2, postName: "岗位2"},
          {postId: 2, postName: "岗位2"},
          {postId: 2, postName: "岗位2"},
          {postId: 2, postName: "岗位2"},
          {postId: 2, postName: "岗位2"},
      ];
      this.total = 10;
      this.loading = false;
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageNum = 1;
      this.getList();
    },

    /** 新增按钮操作 */
    handleAdd() {
      this.open = true;
      this.title = "添加岗位";
    },
    /** 导出按钮操作 */
    handleExport() {

    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.open = true;
      this.title = "修改岗位";
      const postId = row.postId || this.ids
      // 修改岗位
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const postIds = row.postId || this.ids;
      this.$confirm('是否确认删除岗位编号为"' + postIds + '"的数据项？').then(function () {
        // 删除成功
      }).then(() => {

      }).catch(() => {
      });
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.postId)
      this.single = selection.length != 1
      this.multiple = !selection.length
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate(valid => {
        if (valid) {

        }
      });
    },
    // 取消按钮
    cancel() {
      this.open = false;
    },
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`);
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`);
    }
  }
}
</script>

<style lang="scss" scoped>
.table {
  padding: 20px;
}
</style>
