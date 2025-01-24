<script>
    import CardBox from "$lib/components/CardBox.svelte";
    import Card from '$lib/components/Card.svelte';
    import EmptyCardSlot from '$lib/components/EmptyCardSlot.svelte';

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

    const allCards = [diamondCards, clubCards, heartCards, spadeCards];
    
</script>



<div class="min-h-screen bg-[#3f3f3f] flex flex-col items-center p-4">
    <div class="cardsHolder flex flex-row justify-center m-12">
        <div class="playerCards">
            <h1 class="text-white text-2xl">Player Cards</h1>
            <div class="flex flex-wrap gap-2 justify-center mt-4">
                
                <CardBox>
                    {#each Array(2) as _}
                    <EmptyCardSlot></EmptyCardSlot>
                    {/each}
                </CardBox>
            </div>
        </div>

        <div class="opponentCards">
            <h1 class="text-white text-2xl">Opponent cards</h1>
            <div class="flex flex-wrap gap-2 justify-center mt-4">
                <CardBox>
                    {#each Array(2) as _}
                    <EmptyCardSlot></EmptyCardSlot>
                    {/each}
                </CardBox>
            </div>
        </div>

        <div class="boardCards">
            <h1 class="text-white text-2xl">Board</h1>
            <div class="flex flex-wrap gap-2 justify-center mt-4">
                <CardBox>
                    {#each Array(5) as _}
                    <EmptyCardSlot></EmptyCardSlot>
                    {/each}
                </CardBox>
            </div>
        </div>



    </div>



    <div class="m-20 cards-container max-w-[1400px]">
        {#each allCards as cardSetColor}
        <CardBox>
            {#each cardSetColor as card}
            <Card src={card} playingCards={false}></Card>
            {/each}
        </CardBox>
        {/each}
    </div>
</div>