import { Client } from "@/models/client";
import { Project } from "@/models/project";
import { MutationTree } from "vuex";
import { State } from "./state";

export enum MutationsType {
  setAdmin = "SET_ADMIN",
  setPosts = "SET_POSTS",
  setClients = "SET_CLIENTS",
}

export type Mutations = {
  [MutationsType.setAdmin](state: State, value: boolean): void;
  [MutationsType.setPosts](state: State, projects: Project[]): void;
  [MutationsType.setClients](state: State, clients: Client[]): void;
};

export const mutations: MutationTree<State> & Mutations = {
  [MutationsType.setAdmin](state, value) {
    state.admin.logged = value;
  },
  [MutationsType.setPosts](state, projects) {
    const ids = projects.map((project) => project.id);
    state.projects = {
      all: projects,
      ids,
    };
  },
  [MutationsType.setClients](state, clients) {
    const ids = clients.map((client) => client.id);
    state.clients = {
      all: clients,
      ids,
    };
  },
};
