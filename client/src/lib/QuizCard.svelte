<script lang="ts">
  import Button from "./Button.svelte";

  export let quiz: { _id: string; name: string };
  export let onHost: () => void;
</script>

<div
  class="flex justify-between items-center bg-white border-2 border-green-500"
>
  <p>{quiz.name}</p>

  <Button on:click={onHost}>Host</Button>
</div>
