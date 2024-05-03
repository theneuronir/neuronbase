<template>
  <v-app id="inspire">
    <v-app-bar class="px-3" flat density="compact">
      <span>
        theNEURON
        <v-chip color="primary">BASE</v-chip>
      </span>

      <v-spacer />
    </v-app-bar>

    <v-main class="bg-grey-lighten-3">
      <slot />
    </v-main>
  </v-app>
</template>
