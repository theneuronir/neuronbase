<script setup>
definePageMeta({
  middleware: "auth",
});

const itemsPerPage = ref(15),
  search = ref(''),
  serverItems = ref([]),
  loading = ref(true),
  totalItems = ref(0)

const capitalize = (string) => {
  return string.charAt(0).toUpperCase() + string.slice(1);
}

const { $pb } = useNuxtApp();

const collections = await $pb.collections.getFullList(200, {
  sort: "+name",
});

</script>

<template>
  <section>
    <v-container>
      <v-card
        title="Collections"
        prepend-icon="mdi-database"
      >
        <v-table>
          <thead>
            <tr>
              <th class="text-left">
                Name
              </th>
              <th class="text-right">
                Created
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="item in collections"
              :key="item.name"
              @click="useRouter().push(`/cms/${item.id}`)"
            >
              <td>{{ capitalize(item.name) }}</td>
              <td class="text-right">{{ item.created }}</td>
            </tr>
          </tbody>
        </v-table>
      </v-card>
    </v-container>
  </section>
</template>

<style>
table tr {
  cursor: pointer;
}
</style>
