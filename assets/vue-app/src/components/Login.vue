<template>
    <div class="login">
        <h2>Sign in</h2>
        <div class="container" style="width:500px">
        <div class="row align-items-center">
            <div class="col-md-4">アカウント名</div>
            <div class="col-md-3">
            <input v-model="authId" type="text" placeholder="Username">
            </div>
        </div>
        <div class="row align-items-center">
            <div class="col-md-4">パスワード</div>
            <div class="col-md-3">
            <input v-model="authPass" type="password" placeholder="Password">
            </div>
        </div>
        <div class="row align-items-center">
            <div class="col-md-12">
            {{ msg }}
            </div>
        </div>
        <div class="row align-items-center">
            <div class="col-md-12">
            <button class="btn btn-info btn-block login" v-on:click="post">ログイン</button>
            </div>
        </div>
        </div>
    </div>
</template>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .login {
        margin-top: 20px;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
    }
    input {
        margin: 10px 0;
        padding: 10px;
    }
</style>
<script>
    export default {
        name: 'login',
        data () {
            return {
                showPassword : false,
                msg : 'userIDとpasswordを入力して下さい',
                authId : '',
                authPass : ''
            }
        },
        mounted: async function() {
            await this.$store.dispatch('setLoginUserInformation');
            if (this.$store.state.login_user_information.email) {
                this.$router.push('/page1');
            }
        },
        methods: {
            async post() {
                const data = { id : this.authId, pass : this.authPass };
                const response = await this.$store.dispatch('loginPost', data);

                if(response !== null && response.data.is_ok){
                    await this.$store.dispatch('setLoginUserInformation');
                    this.$router.push('/page1');
                }
            },
        }
    };
</script>