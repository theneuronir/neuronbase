import PocketBase from "pocketbase";

declare module "#app" {
  interface NuxtApp {
    $pb: PocketBase;
  }
}

export default defineNuxtPlugin((nuxtApp) => {
  const config = useRuntimeConfig()
  const pb = new PocketBase(config.public.PB_BACKEND_URL);
  return {
    provide: {
      pb,
    },
  };
});
