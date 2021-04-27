<template>
  <Main>
    <div class="p-2 flex">
      <project-card
        @click.self="onProjectClick"
        v-for="project in projects"
        :project="project"
        :key="project"
      />
    </div>
    <hr />
    <div class="p-2 flex">
      <client-card
        @click.self="() => {}"
        v-for="client in clients"
        :client="client"
        :key="client"
      />
    </div>
  </Main>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref, watch } from "vue";
import Main from "@/layouts/Main.vue";
import { useStore } from "@/store";
import { ActionTypes } from "@/store/actions";
import ProjectCard from "@/components/ProjectCard.vue";
import ClientCard from "@/components/ClientCard.vue";
import { useRoute, useRouter } from "vue-router";

export default defineComponent({
  name: "Home",
  components: {
    Main,
    ProjectCard,
    ClientCard,
  },
  setup() {
    const router = useRouter();
    const route = useRoute();
    const store = useStore();
    const loading = ref(true);
    const message = ref("");
    const projects = computed(() => store.state.projects.all);
    const clients = computed(() => store.state.clients.all);
    const selectedProject = ref(route.params.params[0]);
    onMounted(() => {
      if (selectedProject.value) {
        store.dispatch(ActionTypes.getClients, selectedProject.value);
      }
      store.dispatch(ActionTypes.getProjects).then((res) => {
        loading.value = false;
        if (res.error) {
          message.value = res.error.message;
        }
      });
    });
    watch(
      () => selectedProject.value,
      (newValue) => {
        if (newValue) {
          store.dispatch(ActionTypes.getClients, selectedProject.value);
        }
      }
    );

    const onProjectClick = (id: string) => {
      selectedProject.value = id;
      router.push("/" + id);
    };

    return {
      projects,
      clients,
      loading,
      message,
      onProjectClick,
      selectedProject,
    };
  },
});
</script>
