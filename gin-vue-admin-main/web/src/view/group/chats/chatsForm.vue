
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createChats,
  updateChats,
  findChats
} from '@/api/group/chats'

defineOptions({
    name: 'ChatsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findChats({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
