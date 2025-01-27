<script>
import AllCardBox from "$lib/components/AllCardBox.svelte";
import Card from '$lib/components/Card.svelte';
import EmptyCard from "$lib/components/EmptyCard.svelte";

const cards = ['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'];
const suits = ['S', 'H', 'D', 'C'];

const getCards = (suit) => {
    return cards.map((card) => ({
        src: `/cards/${card}${suit}.svg`,
        alt: `${card} of ${suit}`
    }));
};

const diamondCards = getCards('D');
const clubCards = getCards('C');
const heartCards = getCards('H');
const spadeCards = getCards('S');
const allCards = [diamondCards, clubCards, heartCards, spadeCards];
</script>

<div class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 flex flex-col items-center p-4">
    <div class="w-full max-w-7xl mx-auto text-center mb-4">
        <h1 class="text-5xl font-bold bg-gradient-to-r from-purple-400 to-pink-400 text-transparent bg-clip-text mb-4">Texas Holdem Odds Calculator</h1>
        <div class="h-1 w-64 bg-purple-500 mx-auto rounded-full"></div>
    </div>

    <div class="w-full max-w-7xl mx-auto p-6">
        <div class="p-8 rounded-xl bg-white/5 backdrop-blur-sm border border-white/10 shadow-2xl">
        
        </div>
    </div>
    
    <div class="space-y-2">
        {#each allCards as cardSetColor}
            <AllCardBox>
                {#each cardSetColor as card}
                    <Card src={card.src} alt={card.alt} />
                {/each}
            </AllCardBox>
        {/each}
    </div>
</div>