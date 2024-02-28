import { createRouter, createWebHashHistory } from "vue-router";
import { getCookie } from "../utils/cookies";

const verifyAuth = (to: any, _from: any, next: any) => {
  const cookie = getCookie("access_token");

  if (
    to.name !== "Login" &&
    (cookie === null || cookie === undefined || cookie === "")
  )
    next({ name: "Login" });
  else next();
};

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/login",
      name: "Login",
      component: () => import("../views/Login.vue"),
    },
    {
      path: "/register",
      name: "Register",
      component: () => import("../views/Register.vue"),
    },
    {
      path: "/landing",
      name: "Landing",
      component: () => import("../views/Landing.vue"),
    },
    {
      path: "/",
      name: "Dashboard",
      component: () => import("../views/Dashboard.vue"),
      beforeEnter: [verifyAuth],
    },
  ],
});

export default router;
