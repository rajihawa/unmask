import {
  createRouter,
  createWebHistory,
  NavigationGuardWithThis,
  RouteRecordRaw,
} from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import AdminLogin from "../views/AdminLogin.vue";
import { useStore } from "@/store";
import { watch } from "@vue/runtime-core";
import { ActionTypes } from "@/store/actions";

export enum RouteNames {
  home = "Home",
  adminLogin = "AdminLogin",
  oauthLogin = "OauthLogin",
}

const requireAuth:
  | NavigationGuardWithThis<undefined>
  | NavigationGuardWithThis<undefined>[]
  | undefined = (to, from, next) => {
  const store = useStore();
  console.log(store.state.admin);
  if (store.getters.isAuthenticated) {
    next();
    return;
  }
  store.dispatch(ActionTypes.adminAutoLogin).then(() => {
    if (!store.getters.isAuthenticated) {
      next({ name: RouteNames.adminLogin });
    } else {
      next();
    }
  });
};

const routes: Array<RouteRecordRaw> = [
  {
    path: "/oauth/login/:clientID",
    name: RouteNames.oauthLogin,
    component: Login,
  },
  {
    path: "/login",
    name: RouteNames.adminLogin,
    component: AdminLogin,
  },
  {
    path: "/:params*",
    name: RouteNames.home,
    component: Home,
    beforeEnter: requireAuth,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
