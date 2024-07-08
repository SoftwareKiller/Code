<template>
    <div class="person">
      <h2>这是一辆大奔，型号为{{ car.No }}，价值{{ car.price }} 万</h2>
      <button @click="refresh">修改骑车价格</button>
      <hr>
      <h2>游戏列表</h2>
      <ul>
        <li v-for="g in games"  :key="g.id">{{ g.name }}</li>
      </ul>
      <button @click="changeFirstName">修改第一个游戏名</button>
    </div>
</template>

<script lang="ts" setup name="person">
  import { reactive } from 'vue';

  let car = reactive({No:"AMG", price:100 })
  let games = reactive([
    {id: "01", name:"旷野之息"},
    {id: "02", name:"王国之泪"},
    {id: "03", name:"黑神话：悟空"}
  ])

  function refresh() {
    car.price += 10
  }

  function changeFirstName() {
    games[0].name = "巫师3"
  }
</script>

<style scoped>
  .person {
    background-color: rgb(117, 142, 152);
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
  }
  button {
    margin: 5px;
  }
  li {
    font-size: 20px;
    margin: 10px;
  }
</style>