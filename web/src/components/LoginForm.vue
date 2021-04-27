<template>
  <form
    @submit.prevent="submit"
    class="flex flex-col justify-center content-center"
  >
    <p class="text-4xl text-center text-black mb-3">{{ title }}</p>
    <input
      class="m-2 p-1 px-2 text-xl ring-2 ring-black"
      v-model="username"
      type="username"
      placeholder="Username"
      id="username"
    />
    <input
      class="m-2 p-1 px-2 text-xl ring-2 ring-black"
      v-model="password"
      type="password"
      placeholder="Password"
      id="password"
    />
    <input
      :disabled="loading"
      class="m-2 p-1 text-xl bg-blue-500 text-white"
      type="submit"
      :value="loading ? 'Loading...' : 'Submit'"
    />
    <p class="text-red-500 text-center pb-1">{{ message }}</p>
  </form>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
  setup(_, { emit }) {
    const username = ref("");
    const password = ref("");

    const submit = () => {
      emit("submit", {
        username: username.value,
        password: password.value,
      });
    };

    return {
      username,
      password,
      submit,
    };
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    loading: {
      type: Boolean,
      default: false,
    },
    message: {
      type: String,
      default: "",
    },
  },
});
</script>
