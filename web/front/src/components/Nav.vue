<template>
    <v-card class="mx-auto" max-width="300">
        <v-img :src="profile.image">
            <v-card-title>
                <v-col align="center">
                    <v-avatar size="130" color="grey">
                        <img :src="profile.avatar" alt />
                    </v-avatar>
                    <div class="ma-4 white--text">{{ profile.name }}</div>
                </v-col>
            </v-card-title>
            <v-divider></v-divider>
        </v-img>
        <v-card-title> About Me: </v-card-title>
        <v-card-text>{{ profile.desc }}</v-card-text>
        <v-divider color="indigo"></v-divider>
        <v-list nav dense>
            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon color="blue darken-2"> {{ 'mdi-qqchat' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey--text">{{ profile.qq_chat }}</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon color="green darken-2"> {{ 'mdi-wechat' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey--text">{{ profile.wechat }}</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon color="orange darken-2"> {{ 'mdi-sina-weibo' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey--text">{{ profile.weibo }}</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon color="blue darken-2"> {{ 'mdi-youtube' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey--text">{{ profile.bili }}</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon color="indigo darken-2"> {{ 'mdi-mail' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey--text">{{ profile.email }}</v-list-item-content>
            </v-list-item>
        </v-list>
    </v-card>
</template>

<script>
export default {
    data() {
        return {
            profile: {
                ID: 1
            }
        }
    },
    created() {
        this.showProfile(1)
    },
    methods: {
        // 获取个人设置信息
        async showProfile(id) {
            const { data: res } = await this.$http.get(`profiles/${id}`)
            if (res.status !== 0) {
                return this.$message.error(res.message)
            }
            this.profile = res.data.profile
        }
    }
}
</script>

<style></style>
