<script setup lang="ts">
const { $pb } = useNuxtApp();
const state = reactive({
  valid: false,
  email: "",
  password: "",
});
const login = async () => {
  console.log("login");
  $pb.admins
    .authWithPassword(state.email, state.password)
    .then((res) => {
      useRouter().push("/");
    })
    .catch((err) => {
      useSnack().notify([
        {
          title: err.response.message,
        },
      ]);
    });
};
definePageMeta({
  layout: "blank",
});
</script>

<template>
  <section>
    <v-container>
      <v-row justify="center">
        <v-col md="3">
          <v-form v-model="state.valid" @submit.prevent="login">
            <v-card title="Login" prepend-icon="mdi-login">
              <v-card-text>
                <v-text-field
                  v-model="state.email"
                  label="Email"
                  type="email"
                />
                <v-text-field
                  v-model="state.password"
                  label="Password"
                  type="password"
                />
              </v-card-text>
              <v-card-actions>
                <v-btn
                  type="submit"
                  variant="flat"
                  color="primary"
                  block
                  append-icon="mdi-arrow-right"
                >
                  Login
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-form>
        </v-col>
      </v-row>
    </v-container>
  </section>
</template>
