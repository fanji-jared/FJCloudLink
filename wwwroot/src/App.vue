<template>
  <div>
    <img alt="Vue logo" src="./assets/logo.png" />
    <HelloWorld msg="Welcome to Your Vue.js App" />
    <FileUploadVue action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"></FileUploadVue>
  </div>
</template>

<script>
import FileUploadVue from './components/FileUpload.vue';
import HelloWorld from "./components/HelloWorld.vue";

export default {
  name: "App",
  components: {
    HelloWorld,
    FileUploadVue,
  },
  data() {
    return {
    };
  },
  methods: {
  },
};
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
</style>
