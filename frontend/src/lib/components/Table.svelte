<script>
  import EmptyCard from "./EmptyCard.svelte";
  import { gameState } from '$lib/stores/cardStore';

  const positions = {
      player: { name: 'Your cards', count: 2, id: 'playerCards' },
      opponent: { name: 'Opponent cards', count: 2, id: 'opponentCards' },
      board: { name: 'Board', count: 5, id: 'boardCards' }
  };
</script>

<div class="p-4 rounded-xl bg-white/5 backdrop-blur-sm border border-white/10 shadow-2xl">
  <div class="flex flex-row justify-between flex-wrap gap-12">
      <div class="flex gap-12">
          <div class="flex flex-col items-center">
              <div class="text-white text-xl mb-2">{positions.player.name}</div>
              <div class="flex gap-2">
                  {#each Array(positions.player.count) as _, i}
                      <EmptyCard 
                          position={positions.player.id}
                          index={i}
                          id={`${positions.player.id}${i + 1}`}
                          currentCard={$gameState.positions?.[positions.player.id]?.[i]}
                      />
                  {/each}
              </div>
          </div>

          <div class="flex flex-col items-center">
              <div class="text-white text-xl mb-2">{positions.opponent.name}</div>
              <div class="flex gap-2">
                  {#each Array(positions.opponent.count) as _, i}
                      <EmptyCard 
                          position={positions.opponent.id}
                          index={i}
                          id={`${positions.opponent.id}${i + 1}`}
                          currentCard={$gameState.positions?.[positions.opponent.id]?.[i]}
                      />
                  {/each}
              </div>
          </div>
      </div>
      
      <div class="flex flex-col items-center">
          <div class="text-white text-xl mb-2">{positions.board.name}</div>
          <div class="flex gap-2">
              {#each Array(positions.board.count) as _, i}
                  <EmptyCard 
                      position={positions.board.id}
                      index={i}
                      id={`${positions.board.id}${i + 1}`}
                      currentCard={$gameState.positions?.[positions.board.id]?.[i]}
                  />
              {/each}
          </div>
      </div>
  </div>
</div>