<template>
    <div>
        <h3>用户列表页面</h3>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-input-search
                        v-model="queryParam.username"
                        placeholder="输入用户名查找"
                        enter-button
                        @search="indexUser"
                        allowClear
                    ></a-input-search>
                </a-col>

                <a-col :span="4">
                    <a-button type="primary" @click="addUserVisible = true"
                        >新增</a-button
                    >
                </a-col>
            </a-row>

            <a-table
                rowKey="username"
                :columns="columns"
                :pagination="pagination"
                :dataSource="userlist"
                bordered
                @change="handleTableChange"
            >
                <span slot="role" slot-scope="role">{{
                    role == 1 ? '管理员' : '订阅者'
                }}</span>
                <template slot="action" slot-scope="id">
                    <div class="actionSlot">
                        <a-button
                            type="primary"
                            icon="edit"
                            style="margin-right: 15px"
                            @click="editUser(id)"
                            >编辑</a-button
                        >
                        <a-button
                            style="margin-right: 15px"
                            icon="delete"
                            type="danger"
                            @click="deleteUser(id)"
                            >删除</a-button
                        >
                        <a-button
                            icon="info"
                            type="info"
                            @click="deleteUser(id)"
                            >重置密码</a-button
                        >
                    </div>
                </template>
            </a-table>
        </a-card>
        <!-- 新增用户区域 -->
        <a-modal
            closable
            :visible="addUserVisible"
            width="60%"
            title="新增用户"
            @ok="addUserOk"
            @cancel="addUserCancel"
            :destroyOnClose="true"
        >
            <a-form-model
                :model="newUser"
                :rules="addUserRules"
                ref="addUserRef"
            >
                <a-form-model-item has-feedback label="用户名" prop="username">
                    <a-input v-model="newUser.username"></a-input>
                </a-form-model-item>
                <a-form-model-item has-feedback label="密码" prop="password">
                    <a-input-password
                        v-model="newUser.password"
                    ></a-input-password>
                </a-form-model-item>
                <a-form-model-item
                    has-feedback
                    label="确认密码"
                    prop="checkpass"
                >
                    <a-input-password
                        v-model="newUser.checkpass"
                    ></a-input-password
                ></a-form-model-item>
            </a-form-model>
        </a-modal>
        <!-- 编辑用户区域 -->
        <a-modal
            closable
            :visible="editUserVisible"
            width="60%"
            title="编辑用户"
            @ok="editUserOk"
            @cancel="editUserCancel"
        >
            <a-form-model
                :model="userInfo"
                :rules="userRules"
                ref="editUserRef"
            >
                <a-form-model-item has-feedback label="用户名" prop="username">
                    <a-input v-model="userInfo.username"></a-input>
                </a-form-model-item>
                <a-form-model-item label="是否为管理员" prop="role">
                    <!-- <a-select
                        defaultValue="2"
                        style="120px"
                        @change="adminChange"
                    >
                        <a-select-option key="1" value="1">是</a-select-option>
                        <a-select-option key="2" value="2">否</a-select-option>
                    </a-select> -->
                    <a-switch
                        v-model="isAdmin"
                        checked-children="是"
                        un-checked-children="否"
                        @change="adminChange"
                    ></a-switch>
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
        dataIndex: 'ID',
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
                total: 0,
                showSizeChanger: true,
                showTotal: (total) => `共${total}条`,
                current: 1,
                pageSize: 5
            },
            queryParam: {
                username: '',
                pageSize: 5,
                page: 1
            },
            newUser: {
                ID: 0,
                username: '',
                password: '',
                role: 2
            },
            userInfo: {
                ID: 0,
                username: '',
                password: '',
                role: 2
            },
            visible: false,
            addUserVisible: false,
            editUserVisible: false,
            userRules: {
                username: [
                    {
                        validator: (rule, value, callback) => {
                            if (this.userInfo.username === '') {
                                callback(new Error('请输入用户名'))
                            }
                            if (
                                [...this.userInfo.username].length < 4 ||
                                [...this.userInfo.username].length > 12
                            ) {
                                callback(new Error('用户应当在4到12位之间'))
                            } else {
                                callback()
                            }
                        },
                        trigger: 'blur'
                    }
                ],
                password: [
                    {
                        validator: (rule, value, callback) => {
                            if (this.userInfo.password === '') {
                                callback(new Error('请输入密码'))
                            }
                            if (
                                [...this.userInfo.password].length < 6 ||
                                [...this.userInfo.password].length > 20
                            ) {
                                callback(new Error('密码应当在6到20位之间'))
                            } else {
                                callback()
                            }
                        },
                        trigger: 'blur'
                    }
                ],
                checkpass: [
                    {
                        validator: (rule, value, callback) => {
                            if (this.userInfo.checkpass === '') {
                                callback(new Error('请输入密码'))
                            }
                            if (
                                this.userInfo.checkpass !==
                                this.userInfo.password
                            ) {
                                callback(new Error('密码不一致,请重新输入'))
                            } else {
                                callback()
                            }
                        },
                        trigger: 'blur'
                    }
                ]
            },
            addUserRules: {
                username: [
                    {
                        validator: (rule, value, callback) => {
                            if (this.newUser.username === '') {
                                callback(new Error('请输入用户名'))
                            }
                            if (
                                [...this.newUser.username].length < 4 ||
                                [...this.newUser.username].length > 12
                            ) {
                                callback(new Error('用户应当在4到12位之间'))
                            } else {
                                callback()
                            }
                        },
                        trigger: 'blur'
                    }
                ],
                password: [
                    {
                        validator: (rule, value, callback) => {
                            if (this.newUser.password === '') {
                                callback(new Error('请输入密码'))
                            }
                            if (
                                [...this.newUser.password].length < 6 ||
                                [...this.newUser.password].length > 20
                            ) {
                                callback(new Error('密码应当在6到20位之间'))
                            } else {
                                callback()
                            }
                        },
                        trigger: 'blur'
                    }
                ],
                checkpass: [
                    {
                        validator: (rule, value, callback) => {
                            if (this.newUser.checkpass === '') {
                                callback(new Error('请输入密码'))
                            }
                            if (
                                this.newUser.checkpass !== this.newUser.password
                            ) {
                                callback(new Error('密码不一致,请重新输入'))
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
    computed: {
        isAdmin: {
            get: function () {
                if (this.userInfo.role === 1) {
                    return true
                }
                return false
            },
            set: function () {}
        }
    },
    created() {
        this.indexUser()
    },
    methods: {
        async indexUser() {
            const { data: res } = await this.$http.get('users', {
                params: {
                    username: this.queryParam.username,
                    pageSize: this.queryParam.pageSize,
                    page: this.queryParam.page
                }
            })
            if (res.status !== 0) return this.$message.error(res.message)
            this.userlist = res.data.users
            this.pagination.total = res.data.total
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
            this.indexUser()
        },
        // 删除用户
        deleteUser(id) {
            this.$confirm({
                title: '提示：请再次确认',
                content: '确定要删除该用户吗？一旦删除,无法恢复',
                onOk: async () => {
                    const { data: res } = await this.$http.delete(`users/${id}`)
                    if (res.status !== 0) {
                        return this.$message.error(res.message)
                    }
                    this.$message.success('删除成功')
                    this.indexUser()
                },
                onCancel: () => {
                    this.$message.info('已取消删除')
                }
            })
        },
        // 新增用户
        addUserOk() {
            this.$refs.addUserRef.validate(async (valid) => {
                if (!valid) {
                    return this.$message.error('参数不符合要求,请重新输入')
                }
                const { data: res } = await this.$http.post('users', {
                    username: this.newUser.username,
                    password: this.newUser.password,
                    role: this.newUser.role
                })
                if (res.status !== 0) {
                    this.$message.error(res.message)
                } else {
                    this.$message.success('添加用户成功')
                    this.indexUser()
                    this.addUserVisible = false
                }
            })
        },
        addUserCancel() {
            this.$refs.addUserRef.resetFields()
            this.addUserVisible = false
        },
        adminChange(checked) {
            if (checked) {
                this.userInfo.role = 1
            } else {
                this.userInfo.role = 2
            }
        },
        async editUser(id) {
            this.editUserVisible = true
            const { data: res } = await this.$http.get(`users/${id}`)
            if (res.status !== 0) {
                return this.$message.error(res.message)
            }
            this.userInfo = res.data.user
        },
        editUserOk() {
            this.$refs.editUserRef.validate(async (valid) => {
                if (!valid) {
                    return this.$message.error('参数不符合要求,请重新输入')
                }

                const { data: res } = await this.$http.put(
                    `users/${this.userInfo.ID}`,
                    {
                        username: this.userInfo.username,
                        role: this.userInfo.role
                    }
                )
                if (res.status !== 0) {
                    this.$message.error(res.message)
                } else {
                    this.$message.success('更新用户信息成功')
                    this.indexUser()
                    this.editUserVisible = false
                }
            })
        },
        editUserCancel() {
            this.editUserVisible = false
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
