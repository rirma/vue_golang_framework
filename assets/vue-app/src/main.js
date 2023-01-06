import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios'
import store from './store'
import router from './router'
import 'bootstrap/dist/css/bootstrap.min.css'

const app = createApp(App)

// axiosを使用できるように定義
app.config.globalProperties.$axios = axios.create({ withCredentials: true })

app.use(store)
app.use(router)
app.mount('#app')