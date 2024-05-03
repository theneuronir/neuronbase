<script setup>
const links = [
  { label: "CMS", url: "/", icon: "mdi-database" },
  { label: "Designer", url: "/designer", icon: "mdi-brush" },
  { label: "Settings", url: "/settings", icon: "mdi-cog" },
];

const logout = () => {
  const { $pb } = useNuxtApp();
  $pb.authStore.clear();
  useRouter().push('/login')
};
</script>

<template>
  <v-app id="layout-default">
    <v-app-bar class="px-3" flat density="compact">
      <span>
        theNEURON
        <v-chip color="primary">BASE</v-chip>
      </span>

      <v-spacer />

      <v-tabs class="mr-4 mr-md-0" color="primary">
        <v-tab
          v-for="link in links"
          :key="link"
          :to="link.url"
          :text="link.label"
          :prepend-icon="link.icon"
        ></v-tab>
      </v-tabs>

      <v-spacer />

      <v-menu>
        <template v-slot:activator="{ props }">
          <v-avatar color="primary" size="32" v-bind:="props" />
        </template>

        <v-list>
          <v-list-item>
            <v-list-item-title>Manage Admins</v-list-item-title>
          </v-list-item>
          <v-list-item @click="logout">
            <v-list-item-title>Logout</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>

    <v-main class="bg-grey-lighten-3">
      <slot />
    </v-main>
  </v-app>
</template>
