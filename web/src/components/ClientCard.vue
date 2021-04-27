<template>
  <div
    @click="click"
    class="p-2 rounded-md bg-gray-500 text-white text-center shadow-xl cursor-pointer hover:bg-gray-600 duration-200 hover:shadow-2xl"
  >
    <p class="text-2xl">{{ client.name }}</p>
    <p class="text-sm">{{ client.description }}</p>
  </div>
</template>

<script lang="ts">
import { Client } from "@/models/client";
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    client: {
      type: Object as () => Client,
      required: true,
    },
  },
  setup(props, { emit }) {
    const click = () => {
      emit("click", props.client.id);
    };

    return {
      click,
    };
  },
});
</script>
