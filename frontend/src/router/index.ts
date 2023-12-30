import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        component: () => import("../views/Home.vue"),
    },
    {
        path: "/about",
        component: () => import("../views/About.vue"),
    },
    {
        path: "/oauth2callback",
        component: () => import("../views/Oauth2callback.vue"),
    },
    {
        path: "/login",
        component: () => import("../views/GoogleLogin.vue"),
    }
];

export default createRouter({
    history: createWebHistory(),
    routes,
});