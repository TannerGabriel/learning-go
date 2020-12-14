<template>
<!-- eslint-disable -->
  <section>
    <link href="https://fonts.googleapis.com/css?family=Press+Start+2P" rel="stylesheet">
    <h1>TODOS</h1>
    <input type="text" id="todo-input" placeholder="Add todo...">  
    <button v-on:click="addTodo">Add todo</button>
    <ul>
      <li v-for="todo in todos" :key="todo.text" class="flex">
        <label>
          <section v-if="todo.status == true">
            <input type="checkbox" v-on:click="changeStatus(todo._id, true)" checked>
          </section>
          <section v-else>
            <input type="checkbox" v-on:click="changeStatus(todo._id, false)">
          </section>

          <span>&nbsp</span>
        </label>
        <span>{{ todo.task }}</span>
        <div class="space"></div>
        <button class="is-error padding" v-on:click="removeTodo(todo._id)">X</button>
      </li>
    </ul>
  </section>
</template>

<script>
import axios from 'axios'
/* eslint-disable */
export default {
  name: 'Main',
  data: function () {
    return {
      todos: [],
      baseURL: `http://${process.env.VUE_APP_API_URL || 'localhost'}:3000/`
    }
  },
  beforeMount(){
    this.getTodos()
  },
  methods: {
    addTodo: async function () {
      const element = document.getElementById("todo-input");
      axios.post(`${this.baseURL}todo`, {
        task: element.value
      })   
      this.todos.push({task: element.value})  
    },
    removeTodo: function(id) {
      axios.delete(`${this.baseURL}todo/deleteTask/${id}`, () => {})   
      const index = this.todos.findIndex(x => x._id === id)
      this.todos.splice(index, 1);
    },
    changeStatus: function(id, status) {
      if(status == true) {
        axios.put(`${this.baseURL}todo/undoTask/${id}`, () => {}) 
      } else {
        axios.put(`${this.baseURL}todo/${id}`, () => {}) 
      }
    },
    getTodos: async function(){
      const todos = await axios.get(`${this.baseURL}todo`)
      todos.data.forEach(element => {
        this.todos.push(element)
      });
    }
  },
}
</script>

<style>
html, body, pre, code, kbd, samp {
  font-family: "Press Start 2P";
}
body {
  background: #20262E;
  padding: 20px;
}
#app {
  background: #fff;
  border-radius: 4px;
  padding: 20px;
  transition: all 0.2s;
}
li {
  margin: 8px 0;
}
del {
  color: rgba(0, 0, 0, 0.3);
}
.padding {
  padding: 1px 7px 2px;
}
.flex {
  display: flex;
}
.space {
  flex-grow: 1;
}
</style>