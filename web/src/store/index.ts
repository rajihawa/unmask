import { InjectionKey } from "@vue/runtime-core";
import {
  createStore,
  Store as VuexStore,
  createLogger,
  CommitOptions,
  DispatchOptions,
} from "vuex";
import { Actions, actions } from "./actions";
import { Getters, getters } from "./getters";
import { Mutations, mutations } from "./mutations";
import { State, state } from "./state";

const store = createStore<State>({
  state,
  mutations,
  actions,
  getters,
  plugins: process.env.NODE_ENV === "development" ? [createLogger()] : [],
});

export function useStore() {
  return store as Store;
}

export type Store = Omit<
  VuexStore<State>,
  "getters" | "commit" | "dispatch"
> & {
  commit<K extends keyof Mutations, P extends Parameters<Mutations[K]>[1]>(
    key: K,
    payload: P,
    options?: CommitOptions
  ): ReturnType<Mutations[K]>;
} & {
  dispatch<K extends keyof Actions>(
    key: K,
    payload?: Parameters<Actions[K]>[1],
    options?: DispatchOptions
  ): ReturnType<Actions[K]>;
} & {
  getters: {
    [K in keyof Getters]: ReturnType<Getters[K]>;
  };
};

// define injection key
export const key: InjectionKey<Store> = Symbol();

export default store;
