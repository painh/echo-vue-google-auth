import { createApp } from 'vue'
import './style.css'

// import app.vue
import App from './App.vue'

import store, { key } from "./store";
import router from "./router";

// router oauth2callback 페이지


createApp(App).use(store, key).use(router).mount("#app");