<script>
  import { probabilityResults } from '$lib/stores/apiStore';
</script>

<div class="p-4 rounded-xl bg-white/5 backdrop-blur-sm border border-white/10 shadow-2xl">
  <div class="flex flex-col items-center gap-2">
      <div class="text-white text-xl mb-2">Chances</div>
      <div class="rounded-xl p-4 w-full border-white/10 border shadow-2xl bg-lime-400/10">
          <div class="text-lime-400 font-bold">
              WIN: {$probabilityResults.winProbability.toFixed(2)}%
          </div>
      </div>
      <div class="rounded-xl p-4 w-full border-white/10 border shadow-2xl bg-white/5">
          <div class="text-rose-400 font-bold">
              LOSE: {$probabilityResults.loseProbability.toFixed(2)}%
          </div>
      </div>
      <div class="rounded-xl p-4 w-full border-white/10 border shadow-2xl bg-white/5">
          <div class="text-gray-400 font-bold">
              DRAW: {$probabilityResults.tieProbability.toFixed(2)}%
          </div>
      </div>

  </div>
</div>