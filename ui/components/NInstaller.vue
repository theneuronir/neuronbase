<template>
  <v-app>
    <v-main>
      <v-container class="fill-height">
        <v-row class="fill-height" justify="center" align="center">
          <v-col md="4">
            <div class="text-h4 text-center mb-8">theNEURON</div>
            <v-card title="Setup">
              <v-card-text>
                <p>Create your first admin account in order to continue</p>
              </v-card-text>
              <v-card-text>
                <v-form @submit="createAndLogin">
                  <v-text-field v-model="admin.email" label="Email" />
                  <v-text-field
                    v-model="admin.password"
                    label="Password"
                    type="password"
                  />
                  <v-text-field
                    v-model="admin.passwordConfirm"
                    label="Password confirm"
                    type="password"
                  />
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-btn
                  @click.stop="createAndLogin"
                  color="primary"
                  variant="flat"
                  block
                >
                  Create And Login
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
const admin = reactive({
  email: "",
  password: "",
  passwordConfirm: "",
});
const { $pb } = useNuxtApp();
const emit = defineEmits(["install"]);
const createAndLogin = async () => {
  try {
    await $pb.admins.create(admin);
    await $pb.admins.authWithPassword(admin.email, admin.password);
    await emit("install");
  } catch {}
};
</script>
