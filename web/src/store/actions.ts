import { ActionContext, ActionTree } from "vuex";
import { State } from "./state";
import { Mutations, MutationsType } from "./mutations";
import { useFetch } from "@/utils/fetch";
import { ApiResponse } from "@/models/api";
import { AdminLogin } from "@/models/admin";

export enum ActionTypes {
  adminAutoLogin = "ADMIN_AUTO_LOGIN",
  adminLogin = "ADMIN_LOGIN",
  getProjects = "GET_PROJECTS",
  getClients = "GET_CLIENTS",
}

type ActionAugments = Omit<ActionContext<State, State>, "commit"> & {
  commit<K extends keyof Mutations>(
    key: K,
    payload: Parameters<Mutations[K]>[1]
  ): ReturnType<Mutations[K]>;
};

export type Actions = {
  [ActionTypes.adminAutoLogin](context: ActionAugments): Promise<void>;
  [ActionTypes.adminLogin](
    context: ActionAugments,
    data: AdminLogin
  ): Promise<ApiResponse<void>>;
  [ActionTypes.getProjects](
    context: ActionAugments
  ): Promise<ApiResponse<void>>;
  [ActionTypes.getClients](
    context: ActionAugments,
    id: string
  ): Promise<ApiResponse<void>>;
};

// to stimulate http request time
const sleep = (ms: number) => new Promise((res) => setTimeout(res, ms));

export const actions: ActionTree<State, State> & Actions = {
  [ActionTypes.adminAutoLogin]({ commit }) {
    const fetch = useFetch();
    return new Promise<void>((res) => {
      fetch
        .get("/api/admin/me")
        .then(() => {
          commit(MutationsType.setAdmin, true);
          res();
        })
        .catch(() => {
          res();
        });
    });
  },
  [ActionTypes.adminLogin]({ commit }, data) {
    const fetch = useFetch();
    return new Promise((res, rej) => {
      fetch
        .post("/api/admin/login", data)
        .then(() => {
          commit(MutationsType.setAdmin, true);
          res({});
        })
        .catch((err) => {
          console.log(err.response);
          res({ error: err.response.data });
        });
    });
  },
  [ActionTypes.getProjects]({ commit }) {
    const fetch = useFetch();
    return new Promise((res) => {
      fetch
        .get("/api/projects")
        .then((response) => response.data)
        .then((response) => {
          commit(MutationsType.setPosts, response);
          res({});
        })
        .catch((err) => {
          res({ error: err.response.data });
        });
    });
  },
  [ActionTypes.getClients]({ commit }, id) {
    const fetch = useFetch();
    return new Promise((res) => {
      fetch
        .get(`/api/${id}/clients?show_projects=false`)
        .then((response) => response.data)
        .then((response) => {
          commit(MutationsType.setClients, response);
          res({});
        })
        .catch((err) => {
          res({ error: err.response.data });
        });
    });
  },
};
