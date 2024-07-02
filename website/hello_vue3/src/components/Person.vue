<template>
    <div class="person">
        <h2>姓名： {{ name }}</h2>
        <h2>年龄： {{ age }}</h2>
        <button @click="showTel">查看联系方式</button>
        <button @click="changeName">修改名字</button>
        <button @click="changeAge">修改年龄</button>
    </div>
</template>

<script lang="ts" setup name="Person333">
  // 数据
  let name = "luca"  // 这种写法定义的变量非响应式，响应式即为数据被修改后，页面上不会实时刷新
  let age = 18
  let tel = "138"

  function changeName() {
    name = "alex"
    console.log(name)  // 打印log可以看到，name变量确实被更改了
  }

  function changeAge() {
    age += 1
  }

  function showTel() {
    alert(tel)
  }
</script>

<style scoped>
  .person {
    background-color: skyblue;
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
  }
  button {
    margin: 5px;
  }
</style>