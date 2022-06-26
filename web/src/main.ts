import App from "./App.vue";
import { createApp, provide, h } from "vue";
import { DefaultApolloClient } from "@vue/apollo-composable";
import { apolloClient } from "./utils/apollo-client";
import Vue3Lottie from "vue3-lottie";
import "vue-global-api";
import "./index.css";
import { router } from "./router";

createApp(App)
  .provide(DefaultApolloClient, apolloClient)
  .use(Vue3Lottie)
  .use(router)
  .mount("#app");
