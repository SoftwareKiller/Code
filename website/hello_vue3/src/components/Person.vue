<template>
    <div class="person">
        <h2>姓名： {{ name }}</h2>
        <h2>年龄： {{ age }}</h2>
        <button @click="showTel">查看联系方式</button>
        <button @click="changeNama">修改名字</button>
        <button @click="changeAge">修改年龄</button>
    </div>
</template>

<script lang="ts">
  export default {
    name:"Person",
    data() {
        return {
            name:"luca",
            age: 18,
            tel: "13886888888"
        }
    },
    methods:{
        showTel() {
            alert(this.tel)
        },
        changeNama() {
            this.name = "alex"
        },
        changeAge() {
            this.age += 1
        }
    }
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