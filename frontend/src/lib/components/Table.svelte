<script>
    import EmptyCard from "./EmptyCard.svelte";

    let selectedID = $state(null);
</script>

<div class="p-4 rounded-xl bg-white/5 backdrop-blur-sm border border-white/10 shadow-2xl">
    <div class="flex flex-row justify-between flex-wrap gap-12">
      <div class="flex gap-12">
        <div class="flex flex-col items-center">
          <div class="text-white text-xl mb-2">Your cards</div>
          <div class="flex gap-2">
            <EmptyCard id="playerCard1" isSelected={selectedID == "playerCard1"} onSelect={() => selectedID = "playerCard1"}/>
            <EmptyCard id="playerCard2" isSelected={selectedID == "playerCard2"} onSelect={() => selectedID = "playerCard2"}/>
          </div>
        </div>
        <div class="flex flex-col items-center">
          <div class="text-white text-xl mb-2">Opponent cards</div>
          <div class="flex gap-2">
            <EmptyCard id="opponentCard1" isSelected={selectedID == "opponentCard1"} onSelect={() => selectedID = "opponentCard1"}/>
            <EmptyCard id="opponentCard2" isSelected={selectedID == "opponentCard2"} onSelect={() => selectedID = "opponentCard2"}/>
          </div>
        </div>
      </div>
      
      <div class="flex flex-col items-center">
        <div class="text-white text-xl mb-2">Board</div>
        <div class="flex gap-2">
            {#each Array(5) as _, index}
            <EmptyCard id={`boardCard${index}`} + i isSelected={selectedID == `boardCard${index}`} onSelect={() => selectedID = `boardCard${index}`}/>
            {/each}
        </div>
      </div>
    </div>
  </div>