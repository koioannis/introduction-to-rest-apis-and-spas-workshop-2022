<template>
  <div class="h-screen flex justify-center items-center flex-col space-y-6 text-xs w-4/6 lg:w-3/6 xl:w-2/6 m-auto">
    <h2 class=" text-left text-3xl font-semibold w-full">Notes</h2>

    <section class="h-4/6 flex items-center justify-between flex-col w-full">
      <div class="text-left h-4/6 overflow-scroll w-full">
        <ul class="space-y-4">
          <li v-for="note in notes" :key="note.id" class="p-4 border rounded-lg flex justify-between">
            <p class="w-5/6">
              <b>{{note.id}}</b>: 
              <input
                :value='note.text'
                @change='(event) => note.text = event.target.value'
                class="outline-none w-5/6" />
            </p>

            <div class="flex items-center space-x-2">
              <button @click='updateNote(note.id)'>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-green-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
              </button>
              <button class="font-bold text-red-500 text-sm" @click='deleteNote(note.id)'>x</button>
            </div>
          </li>
          <p v-if='notes.length === 0' class="opacity-60">There are no notes found</p>
        </ul>
      </div>

      <div class="flex flex-col space-y-2 h-2/6 w-full">
        <textarea
          class="h-24 border border-gray-300 rounded-lg p-3 resize-none outline-none"
          placeholder="Write your note"
          v-model='text' />
        <div class="flex space-x-3">
          <button 
            class="px-6 py-2 bg-green-500 rounded-lg text-white shadow-sm hover:bg-green-400"
            @click='createNote'>
              Enter
          </button>
          <button
            class="px-6 py-2 bg-blue-500 rounded-lg text-white shadow-sm hover:bg-blue-400"
            @click='getNotes'>
              Refresh
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { ref } from 'vue'
import axios from 'axios'

export default {
  name: 'Notes',
  setup() {
    axios.defaults.baseURL = 'http://localhost:3000';

    const notes = ref([]);
    const text = ref('');

    const getNotes = () => {
      axios.get('/notes')
        .then((response) => notes.value = response.data);
    }

    const createNote = () => {
      axios.post('/notes', {
        text: text.value
      })
        .then((response) => notes.value.push(response.data));
      text.value = '';
    }

    const deleteNote = (id) => {
      axios.delete(`/notes/${id}`)
        .then(() => notes.value = notes.value.filter((note) => note.id !== id));
    }

    const updateNote = (id) => {
      axios.put(`/notes/${id}`, {
        text: notes.value.filter((note) => note.id === id)[0]?.text
      });
    }

    return {
      notes,
      getNotes,
      createNote,
      deleteNote,
      updateNote,
      text,
    }
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
}
</style>
