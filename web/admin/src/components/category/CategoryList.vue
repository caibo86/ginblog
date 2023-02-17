<template>
    <div>
        <h3>分类列表页面</h3>
        <a-card>
            <a-row :gutter="20" style="margin-bottom: 10px">
                <a-col :span="4">
                    <a-button type="primary" @click="addCategoryVisible = true">新增</a-button>
                </a-col>
            </a-row>

            <a-table
                rowKey="ID"
                :columns="columns"
                :pagination="pagination"
                :dataSource="categorylist"
                bordered
                @change="handleTableChange"
            >
                <template slot="action" slot-scope="id">
                    <div class="actionSlot">
                        <a-button type="primary" icon="edit" style="margin-right: 15px" @click="editCategory(id)"
                            >编辑</a-button
                        >
                        <a-button style="margin-right: 15px" icon="delete" type="danger" @click="deleteCategory(id)"
                            >删除</a-button
                        >
                    </div>
                </template>
            </a-table>
        </a-card>
        <!-- 新增分类区域 -->
        <a-modal
            closable
            :visible="addCategoryVisible"
            width="60%"
            title="新增分类"
            @ok="addCategoryOk"
            @cancel="addCategoryCancel"
            :destroyOnClose="true"
        >
            <a-form-model :model="newCategory" :rules="addCategoryRules" ref="addCategoryRef">
                <a-form-model-item has-feedback label="分类名" prop="name">
                    <a-input v-model="newCategory.name"></a-input>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
        <!-- 编辑分类区域 -->
        <a-modal
            closable
            :visible="editCategoryVisible"
            width="60%"
            title="编辑分类"
            @ok="editCategoryOk"
            @cancel="editCategoryCancel"
        >
            <a-form-model :model="category" :rules="categoryRules" ref="editCategoryRef">
                <a-form-model-item has-feedback label="分类名" prop="name">
                    <a-input v-model="category.name"></a-input>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
    </div>
</template>

<script>
const columns = [
    {
        title: 'ID',
        dataIndex: 'ID',
        width: '10%',
        key: 'id',
        align: 'center'
    },
    {
        title: '分类名',
        dataIndex: 'name',
        width: '30%',
        key: 'name',
        align: 'center'
    },
    {
        title: '操作',
        width: '40%',
        key: 'action',
        dataIndex: 'ID',
        scopedSlots: { customRender: 'action' },
        align: 'center'
    }
]
export default {
    data() {
        return {
            categorylist: [],
            columns,
            pagination: {
                pageSizeOptions: ['5', '10', '20'],
                total: 0,
                showSizeChanger: true,
                showTotal: (total) => `共${total}条`,
                current: 1,
                pageSize: 5
            },
            queryParam: {
                pageSize: 5,
                page: 1
            },
            newCategory: {
                ID: 0,
                name: ''
            },
            category: {
                ID: 0,
                name: ''
            },
            addCategoryVisible: false,
            editCategoryVisible: false,
            categoryRules: {
                name: [
                    {
                        validator: (rule, value, callback) => {
                            if (this.category.name === '') {
                                callback(new Error('请输入分类名'))
                            }
                            if ([...this.category.name].length < 4 || [...this.category.name].length > 12) {
                                callback(new Error('分类名应当在4到12位之间'))
                            } else {
                                callback()
                            }
                        },
                        trigger: 'blur'
                    }
                ]
            },
            addCategoryRules: {
                name: [
                    {
                        validator: (rule, value, callback) => {
                            if (this.newCategory.name === '') {
                                callback(new Error('请输入分类名'))
                            }
                            if ([...this.newCategory.name].length < 4 || [...this.newCategory.name].length > 12) {
                                callback(new Error('分类名应当在4到12位之间'))
                            } else {
                                callback()
                            }
                        },
                        trigger: 'blur'
                    }
                ]
            }
        }
    },
    created() {
        this.indexCategory()
    },
    methods: {
        // 获取分类列表
        async indexCategory() {
            const { data: res } = await this.$http.get('categories', {
                params: {
                    pageSize: this.queryParam.pageSize,
                    page: this.queryParam.page
                }
            })
            if (res.status !== 0) return this.$message.error(res.message)
            this.categorylist = res.data.categories
            this.pagination.total = res.data.total
            if (
                this.pagination.current > 1 &&
                (this.pagination.current - 1) * this.pagination.pageSize >= this.pagination.total
            ) {
                this.pagination.current -= 1
                this.queryParam.page = this.pagination.current
                await this.indexCategory()
            }
        },
        // 更改分页
        handleTableChange(pagination, filters, sorter) {
            const pager = { ...this.pagination }
            pager.current = pagination.current
            pager.pageSize = pagination.pageSize
            this.queryParam.pageSize = pagination.pageSize
            this.queryParam.page = pagination.current
            if (pagination.pageSize !== this.pagination.pageSize) {
                this.queryParam.page = 1
                pager.current = 1
            }
            this.pagination = pager
            this.indexCategory()
        },
        // 删除分类
        deleteCategory(id) {
            this.$confirm({
                title: '提示：请再次确认',
                content: '确定要删除该分类吗？一旦删除,无法恢复',
                onOk: async () => {
                    const { data: res } = await this.$http.delete(`categories/${id}`)
                    if (res.status !== 0) {
                        return this.$message.error(res.message)
                    }
                    this.$message.success('删除成功')
                    this.indexCategory()
                },
                onCancel: () => {
                    this.$message.info('已取消删除')
                }
            })
        },
        // 新增分类
        addCategoryOk() {
            this.$refs.addCategoryRef.validate(async (valid) => {
                if (!valid) {
                    return this.$message.error('参数不符合要求,请重新输入')
                }
                const { data: res } = await this.$http.post('categories', {
                    name: this.newCategory.name
                })
                if (res.status !== 0) {
                    this.$message.error(res.message)
                } else {
                    this.addCategoryVisible = false
                    this.$message.success('添加分类成功')
                    await this.indexCategory()
                }
            })
        },
        addCategoryCancel() {
            this.$refs.addCategoryRef.resetFields()
            this.addCategoryVisible = false
        },
        async editCategory(id) {
            this.editCategoryVisible = true
            const { data: res } = await this.$http.get(`categories/${id}`)
            if (res.status !== 0) {
                return this.$message.error(res.message)
            }
            this.category = res.data.category
        },
        editCategoryOk() {
            this.$refs.editCategoryRef.validate(async (valid) => {
                if (!valid) {
                    return this.$message.error('参数不符合要求,请重新输入')
                }

                const { data: res } = await this.$http.put(`categories/${this.category.ID}`, {
                    name: this.category.name
                })
                if (res.status !== 0) {
                    this.$message.error(res.message)
                } else {
                    this.$message.success('更新分类信息成功')
                    this.indexCategory()
                    this.editCategoryVisible = false
                }
            })
        },
        editCategoryCancel() {
            this.editCategoryVisible = false
            this.$message.info('编辑已取消')
        }
    }
}
</script>

<style scoped>
.actionSlot {
    display: flex;
    justify-content: center;
}
</style>
