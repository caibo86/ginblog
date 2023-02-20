<template>
    <div>
        <a-card>
            <h3>个人设置</h3>
            <a-form-model>
                <a-form-model-item label="作者名称">
                    <a-input style="width: 300px" v-model="profile.name"></a-input>
                </a-form-model-item>

                <a-form-model-item label="个人简介">
                    <a-input type="textarea" v-model="profile.desc"></a-input>
                </a-form-model-item>

                <a-form-model-item label="QQ号码">
                    <a-input style="width: 300px" v-model="profile.qq_chat"></a-input>
                </a-form-model-item>

                <a-form-model-item label="微信">
                    <a-input style="width: 300px" v-model="profile.wechat"></a-input>
                </a-form-model-item>

                <a-form-model-item label="微博">
                    <a-input style="width: 300px" v-model="profile.weibo"></a-input>
                </a-form-model-item>

                <a-form-model-item label="B站地址">
                    <a-input style="width: 300px" v-model="profile.bili"></a-input>
                </a-form-model-item>

                <a-form-model-item label="Email">
                    <a-input style="width: 300px" v-model="profile.email"></a-input>
                </a-form-model-item>

                <a-form-model-item label="头像">
                    <a-upload
                        listType="picture"
                        name="file"
                        :multiple="false"
                        :action="upUrl"
                        :headers="headers"
                        @change="avatarChange"
                    >
                        <a-button><a-icon type="upload"></a-icon>点击上传</a-button>
                        <br />
                        <template v-if="profile.avatar">
                            <img :src="profile.avatar" style="width: 120px; height: 100px" />
                        </template>
                    </a-upload>
                </a-form-model-item>

                <a-form-model-item label="背景图">
                    <a-upload
                        listType="picture"
                        name="file"
                        :multiple="false"
                        :action="upUrl"
                        :headers="headers"
                        @change="imageChange"
                    >
                        <a-button><a-icon type="upload"></a-icon>点击上传</a-button>
                        <br />
                        <template v-if="profile.image">
                            <img :src="profile.image" style="width: 120px; height: 100px" />
                        </template>
                    </a-upload>
                </a-form-model-item>

                <a-form-model-item>
                    <a-button type="primary" style="margin-right: 15px" @click.once="updateProfile()">更新</a-button>
                    <a-button type="info" @click.once="articleCancel">取消</a-button>
                </a-form-model-item>
            </a-form-model>
        </a-card>
    </div>
</template>

<script>
import { Url } from '../../plugin/http'

export default {
    data() {
        return {
            profile: {
                ID: 1,
                name: '',
                desc: '',
                qq_chat: '',
                wechat: '',
                weibo: '',
                bili: '',
                email: '',
                image: '',
                avatar: ''
            },
            upUrl: Url + '/upload',
            headers: {}
        }
    },
    created() {
        this.showProfile(1)
        this.headers = {
            Authorization: `Bearer ${window.sessionStorage.getItem('token')}`
        }
    },
    methods: {
        async showProfile(id) {
            const { data: res } = await this.$http.get(`profiles/${id}`)
            if (res.status !== 0) {
                return this.$message.error(res.message)
            }
            this.profile = res.data.profile
        },
        avatarChange(info) {
            // if (info.file.status !== 'uploading') {
            // }
            if (info.file.status === 'done') {
                this.$message.success('图片上传成功')
                this.profile.avatar = info.file.response.data.url
            } else if (info.file.status === 'error') {
                this.$message.error('图片上传失败')
            }
        },
        imageChange(info) {
            // if (info.file.status !== 'uploading') {
            // }
            if (info.file.status === 'done') {
                this.$message.success('图片上传成功')
                this.profile.image = info.file.response.data.url
            } else if (info.file.status === 'error') {
                this.$message.error('图片上传失败')
            }
        },
        // 更新
        async updateProfile(id) {
            const { data: res } = await this.$http.put(`profiles/${this.profile.ID}`, this.profile)
            if (res.status !== 0) {
                return this.$message.error(res.message)
            }
            this.profile = res.data.profile
            this.$message.success('更新成功')
            this.$router.push('index')
        }
    }
}
</script>

<style></style>
