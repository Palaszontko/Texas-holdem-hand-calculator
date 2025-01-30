<script>
  import { probabilityResults } from '$lib/stores/apiStore';
</script>

<div class="p-4 rounded-xl bg-white/5 backdrop-blur-sm border border-white/10 shadow-2xl w-full">
  <div class="flex flex-col items-center gap-2">
      <div class="text-white text-xl mb-2">Chances</div>
      <div class="rounded-xl p-4 w-full border-white/10 border shadow-2xl  { $probabilityResults.winProbability > $probabilityResults.loseProbability && $probabilityResults.winProbability > $probabilityResults.tieProbability ? 'bg-lime-400/20 animate-pulse' : '' }">
          <div class="text-lime-400 font-bold">
              WIN: {$probabilityResults.winProbability.toFixed(2)}%
          </div>
      </div>
      <div class="rounded-xl p-4 w-full border-white/10 border shadow-2xl  { $probabilityResults.loseProbability > $probabilityResults.winProbability && $probabilityResults.loseProbability > $probabilityResults.tieProbability ? 'bg-rose-400/20 animate-pulse' : '' }">
          <div class="text-rose-400 font-bold">
              LOSE: {$probabilityResults.loseProbability.toFixed(2)}%
          </div>
      </div>
      <div class="rounded-xl p-4 w-full border-white/10 border shadow-2xl  { $probabilityResults.tieProbability > $probabilityResults.winProbability && $probabilityResults.tieProbability > $probabilityResults.loseProbability ? 'bg-gray-400/20 animate-pulse' : '' }">
          <div class="text-gray-400 font-bold">
              DRAW: {$probabilityResults.tieProbability.toFixed(2)}%
          </div>
      </div>
  </div>
</div>