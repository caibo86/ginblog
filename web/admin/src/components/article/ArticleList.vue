<template>
    <div>
        <h3>文章列表页面</h3>
        <a-card>
            <a-row :gutter="20" style="margin-bottom: 10px">
                <a-col :span="8">
                    <a-input-search
                        v-model="queryParam.title"
                        placeholder="输入文章名查找"
                        enter-button
                        @search="indexArticle"
                        allowClear
                    ></a-input-search>
                </a-col>

                <a-col :span="4">
                    <a-button type="primary" @click="$router.push('/admin/addarticle')">新增</a-button>
                </a-col>
                <a-col :span="6" offset="4">
                    <a-select placeholder="请选择分类" style="width: 200px" @change="categoryChange">
                        <a-select-option v-for="item in categorylist" :key="item.ID" :value="item.ID">{{
                            item.name
                        }}</a-select-option>
                    </a-select>
                </a-col>
            </a-row>

            <a-table
                rowKey="ID"
                :columns="columns"
                :pagination="pagination"
                :dataSource="articlelist"
                bordered
                @change="handleTableChange"
            >
                <span class="articleImage" slot="image" slot-scope="image">
                    <img :src="image" />
                </span>
                <template slot="action" slot-scope="id">
                    <div class="actionSlot">
                        <a-button
                            size="small"
                            type="primary"
                            icon="edit"
                            style="margin-right: 15px"
                            @click="$router.push(`/admin/addarticle/${id}`)"
                            >编辑</a-button
                        >
                        <a-button
                            size="small"
                            style="margin-right: 15px"
                            icon="delete"
                            type="danger"
                            @click="deleteArticle(id)"
                            >删除</a-button
                        >
                    </div>
                </template>
            </a-table>
        </a-card>
    </div>
</template>

<script>
const columns = [
    {
        title: 'ID',
        dataIndex: 'ID',
        width: '5%',
        key: 'id',
        align: 'center'
    },
    {
        title: '分类',
        dataIndex: 'Category.name',
        width: '7%',
        key: 'name',
        align: 'center'
    },
    {
        title: '标题',
        dataIndex: 'title',
        width: '13%',
        key: 'title',
        align: 'center'
    },
    {
        title: '描述',
        dataIndex: 'desc',
        width: '20%',
        key: 'desc',
        align: 'center'
    },
    {
        title: '缩略图',
        dataIndex: 'image',
        width: '20%',
        key: 'image',
        align: 'center',
        scopedSlots: { customRender: 'image' }
    },
    {
        title: '操作',
        width: '15%',
        key: 'action',
        dataIndex: 'ID',
        scopedSlots: { customRender: 'action' },
        align: 'center'
    }
]
export default {
    data() {
        return {
            pagination: {
                pageSizeOptions: ['5', '10', '20'],
                total: 0,
                showSizeChanger: true,
                showTotal: (total) => `共${total}条`,
                current: 1,
                pageSize: 5
            },
            queryParam: {
                title: '',
                category_id: '',
                pageSize: 5,
                page: 1
            },
            columns,
            articlelist: [],
            categorylist: []
        }
    },
    created() {
        this.indexArticle()
        this.indexCategory()
    },
    methods: {
        // 获取分类列表
        async indexCategory() {
            const { data: res } = await this.$http.get('categories', {
                params: {
                    pageSize: -1,
                    page: 1
                }
            })
            if (res.status !== 0) return this.$message.error(res.message)
            this.categorylist = res.data.categories
        },
        // 获取文章列表
        async indexArticle() {
            const { data: res } = await this.$http.get('articles', {
                params: {
                    title: this.queryParam.title,
                    category_id: this.queryParam.category_id,
                    pageSize: this.queryParam.pageSize,
                    page: this.queryParam.page
                }
            })
            if (res.status !== 0) return this.$message.error(res.message)
            this.articlelist = res.data.articles
            this.pagination.total = res.data.total
            if (
                this.pagination.current > 1 &&
                (this.pagination.current - 1) * this.pagination.pageSize >= this.pagination.total
            ) {
                this.pagination.current -= 1
                this.queryParam.page = this.pagination.current
                await this.indexArticle()
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
            this.indexArticle()
        },
        // 删除文章
        deleteArticle(id) {
            this.$confirm({
                title: '提示：请再次确认',
                content: '确定要删除该文章吗？一旦删除,无法恢复',
                onOk: async () => {
                    const { data: res } = await this.$http.delete(`articles/${id}`)
                    if (res.status !== 0) {
                        return this.$message.error(res.message)
                    }
                    this.$message.success('删除成功')
                    await this.indexArticle()
                },
                onCancel: () => {
                    this.$message.info('已取消删除')
                }
            })
        },
        categoryChange(value) {
            this.indexArticleByCategory(value)
        },
        async indexArticleByCategory(id) {
            const { data: res } = await this.$http.get('articles', {
                params: {
                    category_id: id,
                    pageSize: this.queryParam.pageSize,
                    page: this.queryParam.page
                }
            })
            if (res.status !== 0) return this.$message.error(res.message)
            this.articlelist = res.data.articles
            this.pagination.total = res.data.total
        }
    }
}
</script>

<style scoped>
.actionSlot {
    display: flex;
    justify-content: center;
}

.articleImage {
    width: 100%;
    height: 100%;
}

.articleImage img {
    width: 100px;
    height: 80px;
}
</style>
