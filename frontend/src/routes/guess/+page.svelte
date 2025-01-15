<script lang="ts">
    import { page } from '$app/state';
    import { onMount } from 'svelte';

    const difficulty = decodeURIComponent(page.url.searchParams.get('difficulty') || '');
    const org = decodeURIComponent(page.url.searchParams.get('org') || '');
  

    interface Vtuber {
      id: number;
      name: string;
      org: string;
      gender: string;
      country: string;
      language: string;
      datedebuted: string;
      nicknames: string;
      url: string;
    }
  
    let vtubers: Vtuber[] = [];
    let selectedvtuber: Vtuber;
    let loading = true
    let streak = 0
    let title = "Make your guess!"
    let nicknames: string | string[] = []
    let inputValue = '';
    
    function sleep(ms: number) {
      return new Promise(resolve => setTimeout(resolve, ms));
    }

    function handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Enter') {
      if (inputValue.trim()) {
        console.log('Input value:', inputValue);
        guess(inputValue.toLowerCase());
        inputValue = '';
      }
    }
  }

    async function fetchvtubers(organization: string) {
      loading = true
      const response = await fetch(`/api/vtubers/org/${encodeURIComponent(organization)}`);
      vtubers = await response.json();
      selectedvtuber = randvtuber(vtubers)
      loading = false;
      
      nicknames = selectedvtuber.nicknames 
    ? (Array.isArray(selectedvtuber.nicknames) 
        ? selectedvtuber.nicknames.map(n => n.toLowerCase())
        : selectedvtuber.nicknames.split(',').map(n => n.toLowerCase()))
    : [];

      title = "Make your guess!"
    }
    
    async function skipvtuber() {
      streak = 0
      title = "Correct answer was " + selectedvtuber.name
      await sleep(1500)
      await fetchvtubers(org)
    }

    async function guess(guessed: string) {
      if (guessed == selectedvtuber.name.toLowerCase() || nicknames.includes(guessed)) {
        streak++
        fetchvtubers(org);
        console.log(guessed)
      }
      else {
        console.log(guessed)
        streak = 0
        title = "Incorrect"
        await sleep(1000)
        title = "Make your guess!"
      }
    }

    function randvtuber(array: Vtuber[]): Vtuber {
    return array[Math.floor(Math.random() * array.length)];
  }
    console.log(`Difficulty: ${difficulty}, Organization: ${decodeURIComponent(org)}`);

    onMount(() => {
    if (org) {
      fetchvtubers(org);
    }
  });
  </script>

  <div>
    {#if loading}
    <p>Loading...</p>
  {:else if selectedvtuber}
    <h1>{title}</h1>
    <!-- svelte-ignore a11y-missing-attribute -->
    <img src="http://127.0.0.1:8080/images/{selectedvtuber.id}/{difficulty}.jpg" class="w-[500px]">
  {/if}
  <div class="flex w-[500px] text-center mt-3 space-x-1 gap-2">
    <button on:click={async () => await skipvtuber()} class="easy">Skip</button>
    <input placeholder="Make your guess" class="guessing flex-grow" on:keydown={handleKeydown} bind:value={inputValue}>
  </div>
  <div class="flex w-[500px] justify-center w-full mt-4">
    <a class="hard flex-grow" href="/difficulty">Return</a>
  </div>
  <div class="flex justify-center w-full mt-4">
    <p>Streak: {streak}</p>
  </div>
  </div>