<script>
    import CardBox from "$lib/components/CardBox.svelte";
    import Card from '$lib/components/Card.svelte';

    const cards = ['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'];
    const suits = ['S', 'H', 'D', 'C'];

    const getCards = (suit) => {
        const result = cards.map((card) => `/cards/${card}${suit}.svg`);
        return result;
    };

    const diamondCards = getCards('D');
    const clubCards = getCards('C');
    const heartCards = getCards('H');
    const spadeCards = getCards('S');

    diamondCards.pop();

    const allCards = [diamondCards, clubCards, heartCards, spadeCards];
    
</script>



<div class="min-h-screen bg-[#3f3f3f] flex flex-col items-center p-4">
    <div class="cards-container max-w-[1400px]">
        {#each allCards as cardSetColor}
        <CardBox>
            {#each cardSetColor as card}
            <Card src={card}></Card>
            {/each}
        </CardBox>
        {/each}
    </div>
</div>