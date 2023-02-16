<template>
    <div>
        <h3>用户列表页面</h3>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-input-search
                        placeholder="输入用户名查找"
                        enter-button
                    ></a-input-search>
                </a-col>

                <a-col :span="4">
                    <a-button type="primary">新增</a-button>
                </a-col>
            </a-row>

            <a-table
                rowKey="username"
                :columns="columns"
                :pagination="pagination"
                :dataSource="userlist"
                bordered
            >
                <span slot="role" slot-scope="role">{{
                    role == 1 ? '管理员' : '订阅者'
                }}</span>
                <template slot="action">
                    <div class="actionSlot">
                        <a-button type="primary" style="margin-right: 15px"
                            >编辑</a-button
                        >
                        <a-button type="danger">删除</a-button>
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
        width: '10%',
        key: 'id',
        align: 'center'
    },
    {
        title: '用户名',
        dataIndex: 'username',
        width: '20%',
        key: 'username',
        align: 'center'
    },
    {
        title: '角色',
        dataIndex: 'role',
        width: '20%',
        key: 'role',
        scopedSlots: { customRender: 'role' },
        align: 'center'
    },
    {
        title: '操作',
        width: '30%',
        key: 'action',
        scopedSlots: { customRender: 'action' },
        align: 'center'
    }
]
export default {
    data() {
        return {
            userlist: [],
            columns,
            pagination: {
                pageSizeOptions: ['5', '10', '20'],
                defaultCurrent: 1,
                defaultPageSize: 5,
                total: 0,
                showSizeChanger: true,
                showTotal: (total) => `共${total}条`,
                onChange: (page, perPage) => {
                    this.pagination.defaultCurrent = page
                    this.pagination.defaultPageSize = perPage
                    this.indexUser()
                },
                onShowSizeChange: (page, perPage) => {
                    this.pagination.defaultCurrent = page
                    this.pagination.defaultPageSize = perPage
                    this.indexUser()
                }
            }
        }
    },
    created() {
        this.indexUser()
    },
    methods: {
        async indexUser() {
            const { data: res } = await this.$http.get('users', {
                params: {
                    perPage: this.pagination.defaultPageSize,
                    page: this.pagination.defaultCurrent
                }
            })
            console.log(res)
            if (res.status !== 0) return this.$message.error(res.message)
            this.userlist = res.data.users
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
</style>
