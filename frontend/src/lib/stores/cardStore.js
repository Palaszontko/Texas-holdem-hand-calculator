import { writable, derived } from 'svelte/store';
import { calculateProbabilities, resetResults } from '$lib/stores/apiStore';

const initialState = {
  selectedCard: null,
  positions: {
    playerCards: [null, null],
    opponentCards: [null, null],
    boardCards: [null, null, null, null, null],
  },
};

export const gameState = writable(initialState);

const shouldCalculate = derived(gameState, ($state) => {
  const playerCards = $state.positions.playerCards;
  const opponentCards = $state.positions.opponentCards;
  const communityCards = $state.positions.boardCards;

  if (playerCards[0] && playerCards[1]) {
    return {
      playerCards,
      opponentCards,
      communityCards,
    };
  }
  return null;
});

shouldCalculate.subscribe(($state) => {
  if ($state) {
    calculateProbabilities($state.playerCards, $state.opponentCards, $state.communityCards);
  } else {
    resetResults();
  }
});

const isCardUsed = (state, cardId) => {
  for (const position of Object.values(state.positions)) {
    if (position.some((card) => card?.id === cardId)) {
      return true;
    }
  }
  return false;
};

export const selectCard = (card) => {
  gameState.update((state) => {
    if (isCardUsed(state, card.id)) {
      return state;
    }
    return {
      ...state,
      selectedCard: state.selectedCard?.id === card.id ? null : card,
    };
  });
};

export const placeCard = (position, index) => {
  gameState.update((state) => {
    const newPositions = { ...state.positions };
    const currentCard = newPositions[position][index];

    if (currentCard) {
      newPositions[position] = [...newPositions[position]];
      newPositions[position][index] = null;
      return {
        ...state,
        positions: newPositions,
        selectedCard: null,
      };
    }

    if (state.selectedCard) {
      newPositions[position] = [...newPositions[position]];
      newPositions[position][index] = state.selectedCard;
      return {
        ...state,
        positions: newPositions,
        selectedCard: null,
      };
    }

    return state;
  });
};

export const resetGame = () => {
  gameState.set(initialState);
  resetResults();
};
