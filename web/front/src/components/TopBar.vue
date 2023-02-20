<template>
    <div>
        <v-app-bar app color="indigo darken-2">
            <v-avatar class="mx-12" size="40" color="grey"></v-avatar>
            <v-container class="py-0 fill-height justify-center">
                <v-btn text color="white" @click="$router.push('/')">首页</v-btn>
                <v-btn v-for="item in categorylist" :key="item.ID" text color="white">{{ item.name }}</v-btn>
            </v-container>
            <v-spacer> </v-spacer>
            <v-responsive max-width="260" color="white">
                <v-text-field dense flat hide-details solo-inverted rounded dark></v-text-field>
            </v-responsive>
        </v-app-bar>
    </div>
</template>

<script>
export default {
    data() {
        return {
            categorylist: []
        }
    },
    created() {
        this.indexCategory()
    },
    methods: {
        // 获取分类
        async indexCategory() {
            const { data: res } = await this.$http.get('categories')
            // if (res.status !== 0) return this.$message.error(res.message)
            this.categorylist = res.data.categories
        }
    }
}
</script>

<style></style>
