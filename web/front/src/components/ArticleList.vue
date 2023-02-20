<template>
    <v-col>
        <v-card
            class="ma-3"
            v-for="item in articlelist"
            :key="item.id"
            link
            @click="$router.push(`/detail/${item.ID}`)"
        >
            <v-row no-gutters>
                <v-col class="d-flex justify-center align-center mx-3" cols="1">
                    <v-img max-height="100" max-width="100" :src="item.image"></v-img>
                </v-col>
                <v-col>
                    <v-card-title class="my-2"
                        ><v-chip color="pink" label class="mr-3 white--text">{{ item.Category.name }}</v-chip>
                        {{ item.title }}
                    </v-card-title>
                    <v-card-subtitle v-text="item.desc"></v-card-subtitle>
                    <v-divider></v-divider>
                    <v-card-text>
                        <v-icon class="mr-1" small>{{ 'mdi-calendar-month' }}</v-icon>
                        <span>{{ item.CreatedAt | dateformat('YYYY-MM-DD HH:SS') }}</span>
                    </v-card-text>
                </v-col>
            </v-row>
        </v-card>
        <div class="text-center">
            <v-pagination
                v-model="queryParam.page"
                :length="Math.ceil(this.total / queryParam.pageSize)"
                total-visible="7"
                @input="indexArticle"
            >
            </v-pagination>
        </div>
    </v-col>
</template>

<script>
export default {
    created() {
        this.indexArticle()
    },
    data() {
        return {
            articlelist: [],
            queryParam: {
                pageSize: 5,
                page: 1
            },
            total: 0
        }
    },
    methods: {
        // 获取文章列表
        async indexArticle() {
            const { data: res } = await this.$http.get('articles', {
                params: {
                    page: this.queryParam.page,
                    pageSize: this.queryParam.pageSize
                }
            })
            // if (res.data.status !== 0) return this.$message.error(res.data.$message)
            this.articlelist = res.data.articles
            console.log(this.articlelist)
            this.total = res.data.total
        }
    }
}
</script>

<style></style>
