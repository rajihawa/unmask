<template>
  <div class="container mx-auto h-screen">
    <form
      @submit.prevent="login"
      class="flex flex-col h-full max-w-md mx-auto justify-center content-center text-center"
    >
      <h1 class="text-2xl">{{ loading ? "Loading" : projectName }}</h1>
      <input
        v-model="username"
        class="m-2 p-2 ring-2"
        placeholder="username or email"
      />
      <input
        v-model="password"
        class="m-2 p-2 ring-2"
        placeholder="password"
        type="password"
      />
      <input class="m-2 p-2" type="submit" />
      <div class="text-2xl" v-if="loading">Loading</div>
      <div class="text-2xl">{{ message }}</div>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from "vue";
import axios from "axios";
import { useRoute } from "vue-router";

export default defineComponent({
  name: "OauthLogin",
  setup() {
    const route = useRoute();
    const loading = ref(true);
    const projectName = ref("");
    const username = ref("");
    const password = ref("");
    const message = ref("");
    const callback = ref("");
    onMounted(() => {
      console.log("asfdfds");
      axios
        .get("/api/oauth/" + route.params["clientID"] + "/client", {})
        .then((v) => v.data)
        .then((res) => {
          loading.value = false;
          projectName.value = res.name;
          callback.value = res.callback_url;
          console.log(route.params);
          console.log(res);
        });
    });

    const login = () => {
      loading.value = true;
      message.value = "";
      axios
        .post("/api/oauth/connect", {
          client_id: route.params["clientID"],
          username: username.value,
          password: password.value,
          grant_type: "password",
          scope: "api",
        })
        .then((res) => {
          loading.value = false;
          return res.data;
        })
        .then((data) => {
          if (data.token) {
            window.location.href = callback.value + "?token=" + data.token;
          }
        })
        .catch((err) => {
          loading.value = false;
          message.value = err.response.data.message;
        });
    };

    return {
      loading,
      projectName,
      username,
      password,
      login,
      message,
    };
  },
});
</script>
