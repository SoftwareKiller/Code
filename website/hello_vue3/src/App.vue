<template>
    <!--html-->
    <div class="app">
        <h1>Hello Vue</h1>
        <Person/>
    </div>
</template>

<script lang="ts">
  import Person from "./components/Person.vue"
  export default{
    name: "App", // 组件名
    components: {Person}
  }
</script>

<style>
 /* 样式 */
 .app {
    background-color: #ddd;
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
 }
</style>