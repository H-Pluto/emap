<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <van-button type="primary">
    <RouterLink to="/">首页</RouterLink>
  </van-button>
  <van-button type="primary">
    <RouterLink to="/map">地图</RouterLink>
  </van-button>

  <RouterView />
</template>

<style scoped></style>
