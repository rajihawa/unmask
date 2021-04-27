<template>
  <div
    @click="click"
    class="p-2 rounded-md bg-gray-500 text-white text-center shadow-xl cursor-pointer hover:bg-gray-600 duration-200 hover:shadow-2xl"
  >
    <p class="text-2xl">{{ project.name }}</p>
    <p class="text-sm">{{ project.description }}</p>
  </div>
</template>

<script lang="ts">
import { Project } from "@/models/project";
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    project: {
      type: Object as () => Project,
      required: true,
    },
  },
  setup(props, { emit }) {
    const click = () => {
      emit("click", props.project.id);
    };

    return {
      click,
    };
  },
});
</script>
