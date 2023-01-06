<template>
    <div>
        <h1>ここはページ2です</h1>
        <h3>user:{{ name }}</h3>
        <router-link to="/page1">ページ1へ</router-link>
    </div>
</template>
<script>
    export default {
        name: 'page2',
        data () {
            return {
                name: '',
            }
        },
        mounted: async function() {
            await this.$store.dispatch('setLoginUserInformation');
            if (!this.$store.state.login_user_information.email) {
                this.$router.push('/login');
            }
            this.name = this.$store.state.login_user_information.name;
        }
    };
</script>
