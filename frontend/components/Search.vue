<template>
  <div>
    <el-row class="search_area" :gutter="20">
      <el-col :span="16">
        <el-input @focus="handleFocus" ref="textInput" placeholder="Please input" v-model="input"></el-input>
      </el-col>
      <el-col :span="8">
        <el-button @click="handleClick">検索</el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    name: 'Search',
    data() {
      return {
        input: 'default'
      }
    },
    methods: {
      async handleClick(val) {
        const input = this.$refs.textInput.value
        console.log('this.$refs.textInput', val)
        try {
          const res = await axios.post('/api?search='+ input)
          console.log('res', res)
          return true
        } catch (e) {
          console.error('error', e)
          return false
        }
      },
      handleFocus(val) {
        this.$refs.textInput.select()
      }
    }
  }
</script>

<style scoped>
.search_area {
  padding: 20px;
}
</style>
