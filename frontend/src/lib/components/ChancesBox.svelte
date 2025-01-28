<script>
    let outcomes = [
      { label: 'WIN', chance: 99.99, color: 'text-lime-400', bgHighlight: 'bg-lime-400/10' },
      { label: 'LOSE', chance: 10.11, color: 'text-rose-400', bgHighlight: 'bg-rose-400/10' },
      { label: 'DRAW', chance: 10.00, color: 'text-gray-400', bgHighlight: 'bg-gray-400/10' }
    ];
  
    const maxChance = Math.max(...outcomes.map(o => o.chance));
  </script>
  
  <div class="p-4 rounded-xl bg-white/5 backdrop-blur-sm border border-white/10 shadow-2xl">
    <div class="flex flex-col items-center gap-2">
      <div class="text-white text-xl mb-2">Chances</div>
      {#each outcomes as {label, chance, color, bgHighlight}}
        <div class="rounded-xl p-4 w-full border-white/10 border shadow-2xl {chance === maxChance ? bgHighlight + ' animate-pulse' : 'bg-white/5'}">
          <div class="{color} font-bold">
            {label}: {chance}%
          </div>
        </div>
      {/each}
    </div>
  </div>