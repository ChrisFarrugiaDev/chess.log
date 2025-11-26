<template>
  <div>
    <span>{{ props.gameId }}</span>|
    <span>{{ props.collectionId }}</span>


  </div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import MyChessBoard from '@/components/MyChessBoard.vue';





const props = defineProps<{
    gameId: string,
    collectionId:string
}>();


</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
// Placeholder comment to ensure global styles are imported correctly
</style>