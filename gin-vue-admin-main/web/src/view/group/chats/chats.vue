
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="id字段" prop="id" width="120" />

            <el-table-column align="left" label="creatorUserId字段" prop="creatorUserId" width="120" />

            <el-table-column align="left" label="accessHash字段" prop="accessHash" width="120" />

            <el-table-column align="left" label="randomId字段" prop="randomId" width="120" />

            <el-table-column align="left" label="participantCount字段" prop="participantCount" width="120" />

            <el-table-column align="left" label="title字段" prop="title" width="120" />

            <el-table-column align="left" label="about字段" prop="about" width="120" />

            <el-table-column align="left" label="photoId字段" prop="photoId" width="120" />

            <el-table-column align="left" label="defaultBannedRights字段" prop="defaultBannedRights" width="120" />

            <el-table-column align="left" label="migratedToId字段" prop="migratedToId" width="120" />

            <el-table-column align="left" label="migratedToAccessHash字段" prop="migratedToAccessHash" width="120" />

            <el-table-column align="left" label="availableReactionsType字段" prop="availableReactionsType" width="120" />

            <el-table-column align="left" label="availableReactions字段" prop="availableReactions" width="120" />

            <el-table-column align="left" label="deactivated字段" prop="deactivated" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.deactivated) }}</template>
</el-table-column>
            <el-table-column align="left" label="noforwards字段" prop="noforwards" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.noforwards) }}</template>
</el-table-column>
            <el-table-column align="left" label="ttlPeriod字段" prop="ttlPeriod" width="120" />

            <el-table-column align="left" label="version字段" prop="version" width="120" />

            <el-table-column align="left" label="date字段" prop="date" width="120" />

            <el-table-column align="left" label="createdAt字段" prop="createdAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="updatedAt字段" prop="updatedAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateChatsFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="id字段:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
</el-form-item>
            <el-form-item label="creatorUserId字段:" prop="creatorUserId">
    <el-input v-model.number="formData.creatorUserId" :clearable="true" placeholder="请输入creatorUserId字段" />
</el-form-item>
            <el-form-item label="accessHash字段:" prop="accessHash">
    <el-input v-model.number="formData.accessHash" :clearable="true" placeholder="请输入accessHash字段" />
</el-form-item>
            <el-form-item label="randomId字段:" prop="randomId">
    <el-input v-model.number="formData.randomId" :clearable="true" placeholder="请输入randomId字段" />
</el-form-item>
            <el-form-item label="participantCount字段:" prop="participantCount">
    <el-input v-model.number="formData.participantCount" :clearable="true" placeholder="请输入participantCount字段" />
</el-form-item>
            <el-form-item label="title字段:" prop="title">
    <el-input v-model="formData.title" :clearable="true" placeholder="请输入title字段" />
</el-form-item>
            <el-form-item label="about字段:" prop="about">
    <el-input v-model="formData.about" :clearable="true" placeholder="请输入about字段" />
</el-form-item>
            <el-form-item label="photoId字段:" prop="photoId">
    <el-input v-model.number="formData.photoId" :clearable="true" placeholder="请输入photoId字段" />
</el-form-item>
            <el-form-item label="defaultBannedRights字段:" prop="defaultBannedRights">
    <el-input v-model.number="formData.defaultBannedRights" :clearable="true" placeholder="请输入defaultBannedRights字段" />
</el-form-item>
            <el-form-item label="migratedToId字段:" prop="migratedToId">
    <el-input v-model.number="formData.migratedToId" :clearable="true" placeholder="请输入migratedToId字段" />
</el-form-item>
            <el-form-item label="migratedToAccessHash字段:" prop="migratedToAccessHash">
    <el-input v-model.number="formData.migratedToAccessHash" :clearable="true" placeholder="请输入migratedToAccessHash字段" />
</el-form-item>
            <el-form-item label="availableReactionsType字段:" prop="availableReactionsType">
    <el-input v-model.number="formData.availableReactionsType" :clearable="true" placeholder="请输入availableReactionsType字段" />
</el-form-item>
            <el-form-item label="availableReactions字段:" prop="availableReactions">
    <el-input v-model="formData.availableReactions" :clearable="true" placeholder="请输入availableReactions字段" />
</el-form-item>
            <el-form-item label="deactivated字段:" prop="deactivated">
    <el-switch v-model="formData.deactivated" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="noforwards字段:" prop="noforwards">
    <el-switch v-model="formData.noforwards" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="ttlPeriod字段:" prop="ttlPeriod">
    <el-input v-model.number="formData.ttlPeriod" :clearable="true" placeholder="请输入ttlPeriod字段" />
</el-form-item>
            <el-form-item label="version字段:" prop="version">
    <el-input v-model.number="formData.version" :clearable="true" placeholder="请输入version字段" />
</el-form-item>
            <el-form-item label="date字段:" prop="date">
    <el-input v-model.number="formData.date" :clearable="true" placeholder="请输入date字段" />
</el-form-item>
            <el-form-item label="createdAt字段:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="updatedAt字段:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="id字段">
    {{ detailForm.id }}
</el-descriptions-item>
                    <el-descriptions-item label="creatorUserId字段">
    {{ detailForm.creatorUserId }}
</el-descriptions-item>
                    <el-descriptions-item label="accessHash字段">
    {{ detailForm.accessHash }}
</el-descriptions-item>
                    <el-descriptions-item label="randomId字段">
    {{ detailForm.randomId }}
</el-descriptions-item>
                    <el-descriptions-item label="participantCount字段">
    {{ detailForm.participantCount }}
</el-descriptions-item>
                    <el-descriptions-item label="title字段">
    {{ detailForm.title }}
</el-descriptions-item>
                    <el-descriptions-item label="about字段">
    {{ detailForm.about }}
</el-descriptions-item>
                    <el-descriptions-item label="photoId字段">
    {{ detailForm.photoId }}
</el-descriptions-item>
                    <el-descriptions-item label="defaultBannedRights字段">
    {{ detailForm.defaultBannedRights }}
</el-descriptions-item>
                    <el-descriptions-item label="migratedToId字段">
    {{ detailForm.migratedToId }}
</el-descriptions-item>
                    <el-descriptions-item label="migratedToAccessHash字段">
    {{ detailForm.migratedToAccessHash }}
</el-descriptions-item>
                    <el-descriptions-item label="availableReactionsType字段">
    {{ detailForm.availableReactionsType }}
</el-descriptions-item>
                    <el-descriptions-item label="availableReactions字段">
    {{ detailForm.availableReactions }}
</el-descriptions-item>
                    <el-descriptions-item label="deactivated字段">
    {{ detailForm.deactivated }}
</el-descriptions-item>
                    <el-descriptions-item label="noforwards字段">
    {{ detailForm.noforwards }}
</el-descriptions-item>
                    <el-descriptions-item label="ttlPeriod字段">
    {{ detailForm.ttlPeriod }}
</el-descriptions-item>
                    <el-descriptions-item label="version字段">
    {{ detailForm.version }}
</el-descriptions-item>
                    <el-descriptions-item label="date字段">
    {{ detailForm.date }}
</el-descriptions-item>
                    <el-descriptions-item label="createdAt字段">
    {{ detailForm.createdAt }}
</el-descriptions-item>
                    <el-descriptions-item label="updatedAt字段">
    {{ detailForm.updatedAt }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createChats,
  deleteChats,
  deleteChatsByIds,
  updateChats,
  findChats,
  getChatsList
} from '@/api/group/chats'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'Chats'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: undefined,
            creatorUserId: undefined,
            accessHash: undefined,
            randomId: undefined,
            participantCount: undefined,
            title: '',
            about: '',
            photoId: undefined,
            defaultBannedRights: undefined,
            migratedToId: undefined,
            migratedToAccessHash: undefined,
            availableReactionsType: undefined,
            availableReactions: '',
            deactivated: false,
            noforwards: false,
            ttlPeriod: undefined,
            version: undefined,
            date: undefined,
            createdAt: new Date(),
            updatedAt: new Date(),
        })



// 验证规则
const rule = reactive({
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.deactivated === ""){
        searchInfo.value.deactivated=null
    }
    if (searchInfo.value.noforwards === ""){
        searchInfo.value.noforwards=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getChatsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteChatsFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteChatsByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateChatsFunc = async(row) => {
    const res = await findChats({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteChatsFunc = async (row) => {
    const res = await deleteChats({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        id: undefined,
        creatorUserId: undefined,
        accessHash: undefined,
        randomId: undefined,
        participantCount: undefined,
        title: '',
        about: '',
        photoId: undefined,
        defaultBannedRights: undefined,
        migratedToId: undefined,
        migratedToAccessHash: undefined,
        availableReactionsType: undefined,
        availableReactions: '',
        deactivated: false,
        noforwards: false,
        ttlPeriod: undefined,
        version: undefined,
        date: undefined,
        createdAt: new Date(),
        updatedAt: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createChats(formData.value)
                  break
                case 'update':
                  res = await updateChats(formData.value)
                  break
                default:
                  res = await createChats(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findChats({ id: row.id })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
