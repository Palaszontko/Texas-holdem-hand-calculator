<script>
    import AllCardBox from "$lib/components/AllCardBox.svelte";
    import Card from '$lib/components/Card.svelte';
    import ChancesBox from "$lib/components/ChancesBox.svelte";
    import EmptyCard from "$lib/components/EmptyCard.svelte";
    import Table from "$lib/components/Table.svelte";

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

<div class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 flex flex-col items-center p-2">
    <div class="w-full max-w-7xl mx-auto text-center mb-4">
       <h1 class="text-5xl font-bold bg-gradient-to-r from-purple-400 to-pink-400 text-transparent bg-clip-text mb-2">Texas Holdem Odds Calculator</h1>
       <div class="h-1 w- bg-purple-500 mx-auto rounded-full"></div>
    </div>

    <div class="table-cards w-full max-w-fit mx-auto flex flex-row flex-wrap gap-2">
        <Table/>
        <ChancesBox/>
    </div>
    
    <div class="space-y-2 p-2">
        {#each allCards as cardSetColor}
            <AllCardBox>
                {#each cardSetColor as card}
                    <Card src={card.src} alt={card.alt} />
                {/each}
            </AllCardBox>
        {/each}
    </div>
</div>