<template>
    <div>
        <h1>ここはページ1です</h1>
        <h3>user:{{ email }}</h3>
        <router-link to="/page2">ページ2へ</router-link>
        <button class="btn btn-info btn-block logout" v-on:click="post">ログアウト</button>
    </div>
</template>
<script>
    export default {
        name: 'page1',
        data () {
            return {
                email : '',
            }
        },
        mounted: async function() {
            await this.$store.dispatch('setLoginUserInformation');
            if (!this.$store.state.login_user_information.email) {
                this.$router.push('/login');
            }
            this.email = this.$store.state.login_user_information.email;
        },
        methods: {
            async post() {
                await this.$store.dispatch('logoutPost');
                this.$router.push('/login');
            },
        }
    };
</script>