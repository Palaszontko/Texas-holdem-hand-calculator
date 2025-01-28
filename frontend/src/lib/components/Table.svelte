<script>
    import EmptyCard from "./EmptyCard.svelte";
</script>

<div class="p-4 rounded-xl bg-white/5 backdrop-blur-sm border border-white/10 shadow-2xl">
    <div class="flex flex-row justify-between flex-wrap gap-12">
      <div class="flex gap-12">
        <div class="flex flex-col items-center">
          <div class="text-white text-xl mb-2">Your cards</div>
          <div class="flex gap-2">
            <EmptyCard />
            <EmptyCard />
          </div>
        </div>
        <div class="flex flex-col items-center">
          <div class="text-white text-xl mb-2">Opponent cards</div>
          <div class="flex gap-2">
            <EmptyCard />
            <EmptyCard />
          </div>
        </div>
      </div>
      
      <div class="flex flex-col items-center">
        <div class="text-white text-xl mb-2">Board</div>
        <div class="flex gap-2">
            {#each Array(5) as _}
            <EmptyCard/>
            {/each}
        </div>
      </div>
    </div>
  </div>