import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      redirect: "/apps",
    },
    {
      path: "/login",
      component: () => import("./pages/Login.vue"),
    },
    {
      path: "/apps",
      component: () => import("./pages/AppsIndex.vue"),
      children: [
        {
          path: "",
          component: () => import("./pages/AppsMainList.vue"),
        },
        {
          path: ":appId",
          component: () => import("./pages/AppDetails.vue"),
        },
      ],
    },
    {
      path: "/plugins",
      component: () => import("./pages/PluginsIndex.vue"),
    },
    {
      path: "/nodes",
      component: () => import("./pages/NodesIndex.vue"),
    },
  ],
});
