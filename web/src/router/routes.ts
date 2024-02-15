import { createRouter, createWebHashHistory } from "vue-router";

const routes = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "Home",
      component: () => import("../views/Home.vue"),
    },
  ],
});

export default routes;
