<template>
    <div>
        <Editor :init="init" v-model="content"></Editor>
    </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'
// eslint-disable-next-line
import tinymce from './tinymce.min.js'
import './icons/default/icons.min.js'
import './themes/silver/theme.min.js'
import './langs/zh_CN'

// 注册插件
import './plugins/preview/plugin.min.js'
import './plugins/paste/plugin.min.js'
import './plugins/code/plugin.min.js'
import './plugins/wordcount/plugin.min.js'
import './plugins/image/plugin.min.js'
import './plugins/imagetools/plugin.min.js'

export default {
    components: { Editor },
    props: {
        value: {
            type: String,
            default: ''
        }
    },
    data() {
        return {
            init: {
                language: 'zh_CN',
                height: '500px',
                margin: '0',
                padding: '0',
                branding: false,
                plugins: 'preview paste code wordcount image imagetools',
                toolbar: [
                    'undo redo | formatselect | alignleft aligncenter alignright alignjustify',
                    'preview paste code wordcount | image imagetools'
                ],
                // 上传图片
                images_upload_handler: async (blobInfo, succFun, failFun) => {
                    const formdata = new FormData()
                    formdata.append('file', blobInfo.blob(), blobInfo.name())
                    const { data: res } = await this.$http.post('upload', formdata)
                    if (res.status !== 0) return this.$message.error(res.message)
                    succFun(res.data.url)
                },
                imagetools_cors_hosts: ['*'],
                imagetools_proxy: '*'
            },
            content: this.value
        }
    },
    watch: {
        value(newValue) {
            this.content = newValue
        },
        content(newValue) {
            this.$emit('input', newValue)
        }
    }
}
</script>

<style>
@import url('./skins/ui/oxide/skin.min.css');
</style>
