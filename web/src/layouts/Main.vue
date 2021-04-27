<template>
  <div class="w-screen h-screen max-h-screen overflow-hidden">
    <nav
      class="h-11 bg-gray-600 flex justify-between items-center text-white px-2"
    >
      <button @click="toggleSidebar">
        <MenuIcon class="h-8 w-8" />
      </button>

      <span>logo</span
      ><button>
        <CogIcon class="h-8 w-8" />
      </button>
    </nav>
    <div class="w-full h-full flex">
      <transition name="sidebar">
        <div v-if="sidebar" class="w-52 h-full bg-gray-600"></div>
      </transition>
      <div class="flex-1 block">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { MenuIcon, CogIcon } from "@heroicons/vue/outline";

export default defineComponent({
  name: "Home",
  components: {
    MenuIcon,
    CogIcon,
  },
  setup() {
    const sidebar = ref(false);
    const toggleSidebar = () => {
      sidebar.value = !sidebar.value;
    };
    return {
      sidebar,
      toggleSidebar,
    };
  },
});
</script>

<style scoped>
.sidebar-leave-active {
  transition: all 0.5s ease;
}
.sidebar-enter, .sidebar-leave-to /* .fade-leave-active below version 2.1.8 */ {
  transform: translateX(-13rem);
}

.sidebar-leave {
  transform: translateX(0rem);
}

.sidebar-enter-active {
  animation: slide-in ease 0.5s;
}

@keyframes slide-in {
  0% {
    transform: translateX(-13rem);
  }

  100% {
    transform: translateX(0rem);
  }
}
</style>
