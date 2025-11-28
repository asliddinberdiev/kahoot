<script lang="ts">
  import Button from "./lib/Button.svelte";
  import QuizCard from "./lib/QuizCard.svelte";
  import type { QuizQuestion } from "./model/quiz";
  import { NetService } from "./service/net";

  let quizzes: { _id: string; name: string }[] = [];

  let currentQuestion: QuizQuestion | null = null;

  let netService = new NetService();
  netService.connect();
  netService.onPacket((packet: any) => {
    switch (packet.id) {
      case 2: {
        currentQuestion = packet.question;
        break;
      }
    }
  });

  async function getQuizzes() {
    let response = await fetch("http://localhost:3000/api/quizzes");
    if (!response.ok) {
      alert("Failed");
      return;
    }

    quizzes = await response.json();
  }

  let code = "";
  let msg = "";

  function connect() {
    netService.sendPacket({});
  }

  function hostQuiz(quiz: any) {
    netService.sendPacket({
      id: 1,
      quiz_id: quiz.id,
    });
  }
</script>

<Button on:click={getQuizzes}>Get quizzes</Button>
Message: {msg}

{#each quizzes as quiz}
  <QuizCard onHost={() => hostQuiz(quiz)} {quiz} />
{/each}

<input bind:value={code} class="border" type="text" placeholder="Game Code" />
<Button on:click={connect}>Join game</Button>

{#if currentQuestion != null}
  <h2 class="text-4xl font-bold mt-8">{currentQuestion.name}</h2>
  <div class="flex">
    {#each currentQuestion.choices as choice}
      <dev class="flex-1 bg-blue-400 text-center font-bold text-2xl text-white justify-center items-center p-8">
        {choice.name}
      </dev>
    {/each}
  </div>
{/if}
