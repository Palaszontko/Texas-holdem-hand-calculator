<script>
    import { gameState, selectCard } from '$lib/stores/cardStore';
    
    export let src;
    export let alt;
    export let id;

    $: isUsed = Object.values($gameState.positions).some(
        position => position.some(card => card?.id === id)
    );

    const handleSelect = () => {
        if (!isUsed) {
            selectCard({ id, src, alt });
        }
    };

    const handleKeyDown = (event) => {
        if ((event.key === 'Enter' || event.key === ' ') && !isUsed) {
            event.preventDefault();
            handleSelect();
        }
    };
</script>

<div class="group relative m-1">
    <button
        type="button"
        on:click={handleSelect}
        on:keydown={handleKeyDown}
        class="w-24 h-32 relative transform transition-all duration-300 
               {!isUsed ? 'hover:scale-110 hover:-translate-y-1 cursor-pointer' : 'opacity-30 cursor-not-allowed'}
               {$gameState.selectedCard?.id === id ? 'scale-110 -translate-y-1' : ''}"
        aria-label="Select card {alt}"
        disabled={isUsed}
    >
        <img {src} {alt} {id} 
             class="w-full h-full object-contain rounded-lg shadow-xl"/>
        <div class="absolute inset-0 rounded-lg bg-black opacity-0 
             {!isUsed ? 'group-hover:opacity-10' : ''} transition-all duration-300"></div>
    </button>
</div>