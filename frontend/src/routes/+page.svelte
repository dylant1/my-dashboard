<script lang="ts">
  /** @type {import('./$types').PageData} */
  import type { Assignment, Assignments } from "./+page.server.ts";
  import { browser } from '$app/environment';
  import { Datepicker } from 'svelte-calendar';

  export let data;
  let assignments: Assignments = data.assignments!;

  function getHoursTill(date: string): Number {
    const tmp = Date.parse(date);
    const now = new Date();

    var hours: number = Math.abs(tmp - Date.parse(now.toString())) / 36e5;

    return Math.floor(hours);
  }

  let newItem = "";
  interface IItem {
    text: string;
    status: boolean;
  }
  

  let todoList: IItem[];

  if (!browser) {
    todoList = [];
  }
  else {
    let tmp = window.localStorage.getItem("todo");

if (tmp) {
  todoList = JSON.parse(tmp)
}
else {
  todoList = [];
}
  }
   
  function addToList() {
    todoList = [...todoList, { text: newItem, status: false }];
    if (browser) {
        window.localStorage.setItem("todo", JSON.stringify(todoList))
    }
    newItem = "";
  }

  function removeFromList(index: number) {
    todoList.splice(index, 1);
    todoList = todoList;
    if (browser) {
        window.localStorage.setItem("todo", JSON.stringify(todoList))
    }
  }
</script>

<div class="page-wrapper unit">
  <div
    style="width: 207px; height: 450px; overflow: hidden; border-radius: 5px; margin-top: 90px;"
  >
    <iframe
      width="800"
      height="760"
      title="test"
      style="pointer-events: none; position: relative; left: -299px; top: -157px"
      src="https://www.youtube.com/embed/eRXE8Aebp7s?autoplay=1&mute=1&controls=0"
    />
  </div>

  <div class="wrapper unit">
    <h2>Assignments</h2>
    <div class="assignment-wrapper">
      {#each assignments as assignment}
        <div class="assignment">
          <h3>{assignment.name}</h3>
          <div>Due in {getHoursTill(assignment.due_at)} hours</div>
        </div>
      {/each}
    </div>
  </div>
  <div class="unit">
    <div style="width:360px; height:400px;overflow:hidden;margin-top:80px">
    <iframe title="chess-puzzle" style="width: 400px; height: 560px;position:relative;left:-20px;" src="https://www.chess.com/daily_puzzle" frameborder="0"></iframe></div>

  </div>
  <div class="unit">
    <h2 style="text-align: center">TODO</h2>
    <input bind:value={newItem} type="text" placeholder="" />
    <button on:click={addToList}>+</button>
    {#if todoList}
    {#each todoList as item, index}
    <div class="list-item">
      <input bind:checked={item.status} type="checkbox" />
      <span class:checked={item.status}>{item.text}</span>
      <span on:click={() => removeFromList(index)} role="presentation">❌</span></div>
    {/each}
    {/if}
  </div>

</div>
<div style="position: absolute; top: 0; right: 0;">
    <Datepicker theme={{
      "calendar": {
        "width": "700px",
        "maxWidth": "100vw",
        "legend": {
          "height": "45px"
        },
        "shadow": "0px 10px 26px rgba(0, 0, 0, 0.25)",
        "colors": {
          "text": {
            "primary": "#eee",
            "highlight": "#fff"
          },
          "background": {
            "primary": "#333",
            "highlight": "#5829d6",
            "hover": "#222"
          },
          "border": "#222"
        },
        "font": {
          "regular": "1.5em",
          "large": "37em"
        },
        "grid": {
          "disabledOpacity": ".5",
          "outsiderOpacity": ".7"
        }
      }
    }}/></div>
<style>
    .list-item {
        color: white;
    }
    .unit {
        margin-left: 20px;
        margin-right: 20px;
    }
  .checked {
    text-decoration: line-through;
  }
  .page-wrapper {
    display: flex;
    flex-direction: row;
    justify-content: center;
  }
  .wrapper {
    display: flex;
    justify-content: center;
    flex-direction: column;
    text-align: center;
    width: fit-content;
  }
  h2 {
    color: #f63e46;
  }
  .assignment {
    display: flex;
    background-color: #252c47;
    margin: 10px;
    flex-direction: column;
    padding: 20px;
    color: white;
    border-radius: 10px;
  }
  .assignment-wrapper {
    display: flex;
    flex-direction: column;
    width: fit-content;
    overflow-y: scroll;
    height: 500px;
  }
  h3 {
    padding-top: 0px;
    margin-top: 0px;
  }
  :global(body) {
    background-color: #1b1f38;
    font-family: "Roboto", sans-serif;
  }

  ::-webkit-scrollbar {
    width: 5px;
  }

  ::-webkit-scrollbar-track {
    background: #f1f1f1;
  }

  ::-webkit-scrollbar-thumb {
    background: #f63e46;
  }
</style>
