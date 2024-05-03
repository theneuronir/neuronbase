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

const id = useRoute().params.slug.toString()

const collection = await $pb.collections.getOne(id);

let headers = collection.schema.reduce((prev, curr) => prev.concat([{
  title: capitalize(curr.name),
  align: 'start',
  key: curr.name,
}]), [])

if (collection.type == 'auth') {
  headers.push({
    title: 'Email',
    align: 'start',
    key: 'email',
  })
  headers.push({
    title: 'Username',
    align: 'start',
    key: 'username',
  })
}

headers.push({
  title: 'Created',
  align: 'end',
  key: 'created',
})

headers.push({
  title: 'Updated',
  align: 'end',
  key: 'updated',
})

const loadItems = ({ page, itemsPerPage, sortBy }) => {
  loading.value = true
  $pb.collection(collection.name).getList(page, itemsPerPage).then((res) => {
    serverItems.value = res.items
    totalItems.value = res.totalItems
    loading.value = false
  });
}
</script>

<template>
  <section>
    <v-container>
      <v-card
        :title="capitalize(collection.name)"
        prepend-icon="mdi-database"
      >
        <v-data-table-server
          v-model:items-per-page="itemsPerPage"
          :headers="headers"
          :items="serverItems"
          :items-length="totalItems"
          :loading="loading"
          :search="search"
          item-value="name"
          @update:options="loadItems"
        ></v-data-table-server>
      </v-card>
    </v-container>
  </section>
</template>
