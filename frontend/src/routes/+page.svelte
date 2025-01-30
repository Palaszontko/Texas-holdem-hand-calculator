<script>
    import AllCardBox from "$lib/components/AllCardBox.svelte";
    import Card from '$lib/components/Card.svelte';
    import ChancesBox from "$lib/components/ChancesBox.svelte";
    import Table from "$lib/components/Table.svelte";

    const cards = ['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'];
    const suits = ['S', 'H', 'D', 'C'];

    const getCards = (suit) => {
        return cards.map((card) => ({
            src: `/cards/${card}${suit}.svg`,
            alt: `${card} of ${suit}`,
            id: `${card}${suit}`
        }));
    };

    const diamondCards = getCards('D');
    const clubCards = getCards('C');
    const heartCards = getCards('H');
    const spadeCards = getCards('S');
    const allCards = [diamondCards, clubCards, heartCards, spadeCards];
</script>

<div class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 flex flex-col items-center p-2 overflow-x-hidden">
    <div class="w-full max-w-7xl mx-auto text-center mb-4 px-4">
       <h1 class="text-4xl md:text-5xl font-bold bg-gradient-to-r from-purple-400 to-pink-400 text-transparent bg-clip-text mb-2">
         Texas Holdem Odds Calculator
       </h1>
       <div class="h-1 w-32 bg-purple-500 mx-auto rounded-full"></div>
    </div>

    <div class="w-full flex flex-col lg:flex-row justify-center items-start gap-4 px-2 mb-4">
        <div class="w-full lg:w-auto">
            <Table/>
        </div>
        <div class="w-full lg:w-64">
            <ChancesBox/>
        </div>
    </div>
    
    <div class="w-full max-w-[95vw] mx-auto space-y-2 p-2 overflow-x-auto">
    {#each allCards as cardSetColor}
        <div class="flex justify-center min-w-min">
            <AllCardBox>
                {#each cardSetColor as card}
                    <Card src={card.src} alt={card.alt} id={card.id}/>
                {/each}
            </AllCardBox>
        </div>
    {/each}
</div>
</div>