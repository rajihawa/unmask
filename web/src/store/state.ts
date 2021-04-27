import { Client } from "@/models/client";
import { Project } from "@/models/project";

export type AdminState = {
  logged: boolean;
};

type ProjectsState = {
  all: Project[];
  ids: string[];
};

type ClientsState = {
  all: Client[];
  ids: string[];
};

export type State = {
  admin: AdminState;
  projects: ProjectsState;
  clients: ClientsState;
};

export const state: State = {
  admin: {
    logged: false,
  },
  projects: {
    all: [],
    ids: [],
  },
  clients: {
    all: [],
    ids: [],
  },
};
