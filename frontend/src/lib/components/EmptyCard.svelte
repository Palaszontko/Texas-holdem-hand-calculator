<script>
    import { gameState, placeCard } from '$lib/stores/cardStore';
    
    export let position;
    export let index;
    export let id;
    export let currentCard = null;

    const handleSelect = () => {
        placeCard(position, index);
    };

    const handleKeyDown = (event) => {
        if (event.key === 'Enter' || event.key === ' ') {
            event.preventDefault();
            handleSelect();
        }
    };

    $: isSelected = $gameState.selectedCard && !currentCard;
</script>

<button 
    class="group relative"
    type="button"
    on:click={handleSelect}
    on:keydown={handleKeyDown}
    {id}
    aria-label={currentCard ? `Remove ${currentCard.alt}` : "Empty card position"}
>
    {#if !currentCard}
        <div class="absolute -inset-0.5 bg-gradient-to-r from-purple-500/50 to-pink-500/50 rounded-lg 
             blur opacity-0 group-hover:opacity-30 transition duration-300
             {isSelected ? 'opacity-30' : ''}">
        </div>
        <div class="relative w-32 h-48 bg-white/5 rounded-lg border-2 border-white/10 
             transition-all duration-300 group-hover:border-purple-500/30 
             {isSelected ? 'border-pink-300 animate-pulse' : ''}">
            <div class="w-full h-full flex items-center justify-center">
                <div class="text-white/20 text-4xl group-hover:text-white/40 transition-all duration-300">+</div>        
            </div>
        </div>
    {:else}
        <div class="relative w-32 h-48 transform transition-all duration-300 group-hover:scale-105">
            <img 
                src={currentCard.src} 
                alt={currentCard.alt} 
                class="w-full h-full object-scale-down rounded-lg"
            />
            <div class="absolute inset-0 flex items-center justify-center rounded-lg transition-all duration-300">
                <div class="absolute inset-0 bg-black opacity-0 group-hover:opacity-40 rounded-lg transition-all duration-300"></div>
            </div>
        </div>
    {/if}
</button>