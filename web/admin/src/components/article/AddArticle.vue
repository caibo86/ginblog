<template>
    <div>
        <a-card>
            <h3>{{ id ? '编辑文章' : '新增文章' }}</h3>
            <a-form-model :model="article" ref="articleRef" :rules="articleRules" :hideRequiredMark="true">
                <a-form-model-item label="文章标题" prop="title">
                    <a-input style="width: 300px" v-model="article.title"></a-input>
                </a-form-model-item>
                <a-form-model-item label="文章分类" prop="category_id">
                    <a-select style="width: 200px" v-model="article.category_id" placeholder="请选择分类">
                        <a-select-option v-for="item in categorylist" :key="item.ID" :value="item.ID">{{
                            item.name
                        }}</a-select-option>
                    </a-select>
                </a-form-model-item>

                <a-form-model-item label="文章描述" prop="desc"
                    ><a-input type="textarea" v-model="article.desc"></a-input>
                </a-form-model-item>
                <a-form-model-item label="文章缩略图" prop="image">
                    <a-upload
                        listType="picture"
                        name="file"
                        :multiple="false"
                        :action="upUrl"
                        :headers="headers"
                        @change="upChange"
                    >
                        <a-button><a-icon type="upload"></a-icon>点击上传</a-button>
                        <br />
                        <template v-if="id">
                            <img :src="article.image" style="width: 120px; height: 100px" />
                        </template>
                    </a-upload>
                </a-form-model-item>
                <a-form-model-item label="文章内容" prop="content"
                    ><MyEditor v-model="article.content"></MyEditor
                ></a-form-model-item>

                <a-form-model-item>
                    <a-button type="primary" style="margin-right: 15px" @click="articleOK(article.ID)">{{
                        article.ID ? '更新' : '提交'
                    }}</a-button>
                    <a-button type="info" @click="articleCancel">取消</a-button>
                </a-form-model-item>
            </a-form-model>
        </a-card>
    </div>
</template>

<script>
import { Url } from '../../plugin/http'
import MyEditor from '../editor/Index'

export default {
    components: { MyEditor },
    props: ['id'],
    data() {
        return {
            article: {
                ID: 0,
                title: '',
                category_id: undefined,
                desc: '',
                content: '',
                image: ''
            },
            categorylist: [],
            upUrl: Url + '/upload',
            headers: {},
            articleRules: {
                title: [
                    {
                        required: true,
                        message: '请输入文章标题',
                        trigger: 'blur'
                    }
                ],
                category_id: [
                    {
                        required: true,
                        message: '请选择文章分类',
                        trigger: 'blur'
                    }
                ],
                desc: [
                    {
                        required: true,
                        message: '请输入文章描述',
                        trigger: 'blur'
                    },
                    {
                        max: 120,
                        message: '最长120个字符',
                        trigger: 'change'
                    }
                ],
                content: [
                    {
                        required: true,
                        message: '请输入文章内容',
                        trigger: 'blur'
                    }
                ],
                image: [
                    {
                        required: true,
                        message: '请选择文章缩略图',
                        trigger: 'blur'
                    }
                ]
            }
        }
    },
    created() {
        this.indexCategory()
        this.headers = {
            Authorization: `Bearer ${window.sessionStorage.getItem('token')}`
        }
        if (this.id) {
            this.showArticle(this.id)
        }
    },
    methods: {
        // 获取单个文章信息
        async showArticle(id) {
            const { data: res } = await this.$http.get(`articles/${id}`)
            if (res.status !== 0) return this.$message.error(res.message)
            this.article = res.data.article
        },
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
        // 上传图片
        upChange(info) {
            // if (info.file.status !== 'uploading') {
            // }
            if (info.file.status === 'done') {
                this.$message.success('图片上传成功')
                this.article.image = info.file.response.data.url
            } else if (info.file.status === 'error') {
                this.$message.error('图片上传失败')
            }
        },
        articleOK(id) {
            this.$refs.articleRef.validate(async (valid) => {
                if (!valid) {
                    return this.$message.error('参数不符合要求,请重新输入')
                }
                if (id === 0) {
                    const { data: res } = await this.$http.post('articles', this.article)
                    if (res.status !== 0) {
                        return this.$message.error(res.message)
                    }
                    this.$router.push('/admin/articlelist')
                    this.$message.success('添加文章成功')
                } else {
                    const { data: res } = await this.$http.put(`articles/${id}`, this.article)
                    if (res.status !== 0) {
                        return this.$message.error(res.message)
                    }
                    this.$router.push('/admin/articlelist')
                    this.$message.success('更新文章成功')
                }
            })
        },
        articleCancel() {
            this.$refs.articleRef.resetFields()
        }
    }
}
</script>
