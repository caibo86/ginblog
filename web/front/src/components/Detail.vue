<template>
    <div>
        <div class="d-flex justify-center pa-3 text-h4 font-weight-bold">
            {{ article.title }}
        </div>
        <div class="d-flex justify-center">
            <v-icon class="mr-1" small>{{ 'mdi-calendar-month' }}</v-icon>
            <span>{{ article.CreatedAt | dateformat('YYYY-MM-DD HH:SS') }}</span>
        </div>
        <v-divider class="pa-3 ma-3"></v-divider>
        <v-alert class="ma-4" elevation="2" color="indigo" dark border="left" outlined>{{ article.desc }}</v-alert>
        <div class="content ma-5 pa-3 text-justify" v-html="article.content"></div>
    </div>
</template>

<script>
export default {
    props: ['id'],
    data() {
        return {
            article: {}
        }
    },
    created() {
        this.showArticle(this.id)
    },
    methods: {
        // 查询文章
        async showArticle(id) {
            const { data: res } = await this.$http.get(`articles/${id}`)
            //  if (this.status !=== 0) return this.$message.error(res.message)
            this.article = res.data.article
        }
    }
}
</script>

<style scoped>
.content >>> img,
span,
p {
    width: 90%;
}
</style>
