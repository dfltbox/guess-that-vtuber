<script lang="ts">
import { goto } from '$app/navigation';
import { onMount } from 'svelte';

  let difficulty: string = '';
  let selectedorg: string = '';

  function pickdifficulty(selecteddifficulty: string) {
    difficulty = selecteddifficulty;
  }
  
  function pickorg(org: string) {
    selectedorg = org;
  }

  async function start() {
    if (!difficulty) {
      alert("Please pick a difficulty")
      return;
    }
    if (!selectedorg) {
      alert("Please pick an organisation!")
      return;
    }
    await goto(`/guess?difficulty=${encodeURIComponent(difficulty)}&org=${encodeURIComponent(selectedorg)}`)
  }

  onMount(async () => {
    const org = encodeURIComponent('hololive production');
    const response = await fetch(`/api/vtubers/org/${org}`);
  });
</script>

<div>
  <div class="flex flex-col items-center justify-center min-h-screen text-center space-y-4">
    <h1 class="title">Pick a difficulty and organisation!</h1>
    <div class="space-x-3">
      <button on:click={() => pickdifficulty("easy")} class="easy {difficulty === 'easy' ? 'selecteddiff' : ''}">Easy</button>
      <button on:click={() => pickdifficulty("hard")} class="hard {difficulty === 'hard' ? 'selecteddiff' : ''}">Hard</button>
      <button on:click={() => pickdifficulty("extreme")} class="extreme {difficulty === 'extreme' ? 'selecteddiff' : ''}">Extreme</button>
    </div>
    <div class="space-x-3">
      <button on:click={() => pickorg("VShojo")} id="vshojo" class="org {selectedorg === 'VShojo' ? 'selected' : ''}">VShojo</button>
      <button on:click={() => pickorg("hololive production")} id="hololive" class="org {selectedorg === 'hololive production' ? 'selected' : ''}">Hololive</button>
      <button on:click={() => pickorg("V&U")} id="v&u" class="org {selectedorg === 'V&U' ? 'selected' : ''}">V&U</button>
      <button on:click={() => pickorg("Mythic Talent")} id="mythic" class="org {selectedorg === 'Mythic Talent' ? 'selected' : ''}">Mythic Talent</button>
      <button on:click={() => pickorg("Independent")} id="indie" class="org {selectedorg === 'Independent' ? 'selected' : ''}">Independent</button>
    </div>
    <button 
    on:click={start}
    class="start mt-8 {(!difficulty || !selectedorg) ? 'opacity-50 cursor-not-allowed' : ''}"
    disabled={!difficulty || !selectedorg}
  >
    Start Game
  </button>
  </div>
</div>