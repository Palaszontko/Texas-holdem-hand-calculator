<script>
    let {id, isSelected, onSelect} = $props();

    function handleClick() {
        onSelect();
        console.log("Card selected: ", id);
    }
</script>

<button type="button" onclick={handleClick} class="group relative w-full" aria-label="Select card slot">
    <div class="absolute -inset-0.5 bg-gradient-to-r from-purple-500/50 to-pink-500/50 rounded-lg blur opacity-0 group-hover:opacity-30 transition duration-300"></div>
    <div class="relative w-32 h-48 bg-white/5 rounded-xl border-2 transition-all duration-300 group-hover:border-purple-500/30 {isSelected ? 'border-purple-300 animate-pulse' : 'border-white/10'}">
        <div class="w-full h-full flex items-center justify-center">
            <div class="text-white/20 text-4xl group-hover:text-white/40 transition-all duration-300">+</div>
        </div>
    </div>
</button>