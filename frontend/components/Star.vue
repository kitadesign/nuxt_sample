<template>
  <div class="block">
    <h4>Nuxt.js + Element？</h4>
    <el-rate v-model="value" :colors="['#2F3D4D', '#64B486', '#4B7C6E']"></el-rate>
  </div>
</template>

<script>
  export default {
    name: 'Star',
    data () {
      return {
        value: null
      }
    }
  }
</script>

<style scoped>
  .block {
    padding: 30px 20px;
  }
</style>
