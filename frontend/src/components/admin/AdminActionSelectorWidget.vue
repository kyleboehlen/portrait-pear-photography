<script setup lang="ts">
import { computed } from 'vue'
import { useAdminRoutes } from '@/composables/useAdminRoutes'

const { entity, action, selector, setEntity, setAction, setSelector } = useAdminRoutes()

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
  <div class="w-full mt-2">
    <div class="w-5/8 flex flex-row justify-end text-white gap-4">
      <select
          :value="entity"
          class="select select-ghost select-md focus:outline-none focus:ring-0 !text-white"
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

      <span v-if="entity" class="text-white text-3xl">-></span>
      <select
          v-if="entity"
          :value="action"
          class="select select-ghost select-md focus:outline-none focus:ring-0 !text-white"
          :class="{
            '!border-primary': action,
            '!border-neutral': !action
          }"
          @change="setAction(($event.target as HTMLSelectElement).value)">
        <option value="" disabled>Select an Action</option>
        <option v-for="a in actions" :value="a.value">{{ a.label }}</option>
      </select>

      <span v-if="action" class="text-white text-3xl">-></span>
    </div>
  </div>
</template>

<style scoped>

</style>