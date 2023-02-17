<template>
    <dir>
        <Editor :init="init"></Editor>
    </dir>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'
import tinymce from 'tinymce/tinymce'
import './icons/default/icons.min.js'
import './themes/silver/theme.min.js'
import './langs/zh_CN'

export default {
    components: { Editor },
    data() {
        return {
            init: {
                language: 'zh_CN'
            }
        }
    }
}
</script>

<style>
@import url('./skins/ui/oxide/skin.min.css');
</style>
