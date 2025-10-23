<script setup lang="ts">
import {computed} from 'vue'
import {useAdminRoutes} from '@/composables/useAdminRoutes'
import CreateNewShootWidget from "@/components/admin/shoots/CreateNewShootWidget.vue";
import SelectShootWidget from "@/components/admin/shoots/SelectShootWidget.vue";
import DeleteShootButton from "@/components/admin/shoots/DeleteShootButton.vue";
import DatabaseUploadWidget from "@/components/admin/database/DatabaseUploadWidget.vue";
import DatabaseDownloadWidget from "@/components/admin/database/DatabaseDownloadWidget.vue";

const {entity, action, selector, setEntity, setAction, setSelector} = useAdminRoutes()

const actions = computed(() => {
  switch (entity.value) {
    case 'shoots':
      return [
        {
          value: 'create',
          label: 'Create New'
        },
        {
          value: 'update',
          label: 'Update'
        },
        {
          value: 'delete',
          label: 'Delete'
        }
      ]
    case 'photos':
      return [
        {
          value: 'favorite',
          label: 'Favorite'
        },
        {
          value: 'delete',
          label: 'Delete'
        },
        {
          value: 'categorize',
          label: 'Set Categories'
        }
      ]
    case 'database':
      return [
        {
          value: 'export',
          label: 'Export'
        },
        {
          value: 'import',
          label: 'Import'
        }
      ]
    default:
      return []
  }
})
</script>

<template>
  <div class="w-full mt-2 ml-2">
    <div class="w-full 2xl:w-5/8 pr-16 2xl:pr-0 flex flex-row justify-end text-white gap-2">
      <select
          :value="entity"
          class="select select-ghost select-md focus:outline-none focus:ring-0 !text-white flex-1 max-w-1/3"
          :class="{
          '!border-primary': entity,
          '!border-neutral': !entity
        }"
          @change="setEntity(($event.target as HTMLSelectElement).value)">
        <option value="" disabled>Select an Entity</option>
        <option value="shoots">Shoots</option>
        <option value="photos">Photos</option>
        <option value="database">Database</option>
      </select>

      <span v-if="entity" class="text-white text-3xl whitespace-nowrap">-></span>
      <select
          v-if="entity"
          :value="action"
          class="select select-ghost select-md focus:outline-none focus:ring-0 !text-white flex-1 max-w-1/3"
          :class="{
            '!border-primary': action,
            '!border-neutral': !action
          }"
          @change="setAction(($event.target as HTMLSelectElement).value)">
        <option value="" disabled>Select an Action</option>
        <option v-for="a in actions" :value="a.value">{{ a.label }}</option>
      </select>


      <span v-if="action" class="text-white text-3xl whitespace-nowrap">-></span>
      <span v-if="entity === 'shoots' && action"
            class="flex-1 max-w-1/3 flex flex-row justify-start items-center flex-nowrap">
        <CreateNewShootWidget v-if="action === 'create'"/>
        <SelectShootWidget v-else/>
        <DeleteShootButton v-if="action == 'delete'" class="ml-4"/>
      </span>

      <span v-if="entity === 'database' && action"
            class="flex-1 max-w-1/3 flex flex-row justify-start items-center flex-nowrap">
        <DatabaseUploadWidget v-if="action === 'import'"/>
        <DatabaseDownloadWidget v-else-if="action === 'export'"/>
      </span>
    </div>
  </div>
</template>

<style scoped>

</style>