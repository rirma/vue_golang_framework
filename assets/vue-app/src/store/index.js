import { createStore } from 'vuex'
import axios_source from 'axios'

const axios = axios_source.create({ withCredentials: true })

export default createStore({
    state: {
        login_user_information: {
            is_valid: false,
			email: null,
			name: null,
        }
    },
    mutations: {
        loginUserInformation(state, login_user_information) {
            state.login_user_information = login_user_information
        },
    },
    actions: {
        async setLoginUserInformation(context) {
            if (this.state.login_user_information.email === null) {
                await axios.get('http://localhost:8082/api/v1/auth/user_information')    // ログインユーザの情報取得
                    .then(
                        function (response) {
                            context.commit('loginUserInformation', response.data);
                        }
                    )
                    .catch(
                        function (error) {
                            console.log(error);
                        }
                    );
            }
        },
        async loginPost(context, data) {
            const response = await axios.post('http://localhost:8082/api/v1/login', data)    // バックエンド側にPOST
                .then(
                    function (response) {
                        return response
                    }
                )
                .catch(
                    function (error) {
                        console.log(error);
                        return null
                    }
                );
            return response;
        },
        async logoutPost(context) {
            await axios.post('http://localhost:8082/api/v1/logout')    // バックエンド側にPOST
                .then(
                    function (response) {
                        return response
                    }
                )
                .catch(
                    function (error) {
                        console.log(error);
                        return null
                    }
                );
            context.commit(
                'loginUserInformation',
                {
                    is_valid: false,
                    email: null,
                    name: null,
                }
            );
        }
    }
})