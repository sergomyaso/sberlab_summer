<template>
  <img class="logo" alt="Terraform logo" src="./assets/tr.jpg">
  <UploadModule/>
</template>

<script>
import UploadModule from './components/UploadModule.vue'

export default {
  name: 'App',
  components: {
    UploadModule
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.logo {
  width: 60%;
  height: 60%;
}
</style>
