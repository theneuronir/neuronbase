<template>
  <div class="NApp">
    <n-installer v-if="installer" @install="onInstall" />
    <nuxt-layout v-else>
      <nuxt-page />
    </nuxt-layout>
    <div class="snackbars">
      <v-card v-for="s in snack" :key="s.title" :title="s.title" class="mt-2" />
    </div>
  </div>
</template>

<script setup lang="ts">
const active = ref(true);
const router = useRouter();
const { $pb } = useNuxtApp();
if (!$pb.authStore.isValid) {
  router.push("/login");
}
const { snack } = useSnack();
const route = useRoute();
const installer = computed(() => route.query.installer !== undefined);
const onInstall = () => {
  router.push("/");
};
</script>

<style lang="scss">
.snackbars {
  position: fixed;
  right: 0;
  bottom: 0;
  width: 400px;
  max-height: 100vh;
  overflow: auto;
  padding: 12px;
  display: flex;
  flex-direction: column;
}
</style>
