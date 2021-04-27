<template>
  <div class="container max-w-md mx-auto h-screen flex flex-col justify-center">
    <LoginForm
      @submit.self="login"
      :loading="loading"
      :message="message"
      class="shadow-2xl"
      title="Admin Login"
    />
  </div>
</template>

<script lang="ts">
import { AdminLogin } from "@/models/admin";
import { useStore } from "@/store";
import { ActionTypes } from "@/store/actions";
import { defineComponent, ref } from "vue";
import { useRouter } from "vue-router";
import LoginForm from "../components/LoginForm.vue";

export default defineComponent({
  name: "AdminLogin",
  setup() {
    const store = useStore();
    const router = useRouter();
    const loading = ref(false);
    const message = ref("");

    const login = (data: AdminLogin) => {
      store.dispatch(ActionTypes.adminLogin, data).then((res) => {
        loading.value = false;
        if (res.error) {
          message.value = res.error.message;
        } else {
          router.push("/");
        }
      });
    };

    return {
      login,
      loading,
      message,
    };
  },
  components: {
    LoginForm,
  },
});
</script>
