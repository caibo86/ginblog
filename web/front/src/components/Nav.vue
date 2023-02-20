<template>
    <v-card class="mx-auto" max-width="300">
        <v-img src="../assets/navbg.jpg">
            <v-card-title>
                <v-col align="center">
                    <v-avatar size="130" color="grey">
                        <img src="../assets/avatar.jpeg" alt />
                    </v-avatar>
                    <div class="ma-4 white--text">caibo</div>
                </v-col>
            </v-card-title>
            <v-divider></v-divider>
        </v-img>
        <v-col>
            <div class="ma-3">About Me:</div>
            <div class="ma-3">12321321313</div>
        </v-col>
        <v-divider color="indigo"></v-divider>
        <v-list nav dense>
            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon> {{ 'mdi-qqchat' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey-text">sdsadsa</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon> {{ 'mdi-wechat' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey-text">sdsadsa</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon> {{ 'mdi-sina-weibo' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey-text">sdsadsa</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon> {{ 'mdi-facebook' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey-text">sdsadsa</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon> {{ 'mdi-twitter' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey-text">sdsadsa</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon class="ma-3"
                    ><v-icon> {{ 'mdi-mail' }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content class="grey-text">sdsadsa</v-list-item-content>
            </v-list-item>
        </v-list>
    </v-card>
</template>

<script>
export default {}
</script>

<style></style>
