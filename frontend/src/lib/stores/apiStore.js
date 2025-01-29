import { writable } from 'svelte/store';

export const probabilityResults = writable({
  winProbability: 0,
  loseProbability: 0,
  tieProbability: 0,
  iterations: 0,
});

const getRankValue = (cardId) => {
  const rank = cardId.charAt(0);
  const rankMap = {
    T: 10,
    J: 11,
    Q: 12,
    K: 13,
    A: 14,
  };
  return rankMap[rank] || parseInt(rank);
};

const getSuitValue = (cardId) => {
  const suit = cardId.charAt(1);
  const suitMap = {
    C: 0,
    D: 1,
    H: 2,
    S: 3,
  };
  return suitMap[suit];
};

const convertCardToApiFormat = (card) => ({
  Rank: getRankValue(card.id),
  Suit: getSuitValue(card.id),
});

export const calculateProbabilities = async (playerCards, opponentCards, communityCards) => {
  const validPlayerCards = playerCards.filter(Boolean);
  if (validPlayerCards.length !== 2) return;

  const requestData = {
    playerCards: validPlayerCards.map(convertCardToApiFormat),
    opponentCards: opponentCards.filter(Boolean).map(convertCardToApiFormat),
    communityCards: communityCards.filter(Boolean).map(convertCardToApiFormat),
    numIterations: 100000,
    numConcurrent: 8,
  };

  try {
    const response = await fetch('http://localhost:8080/api/simulation', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(requestData),
    });

    const data = await response.json();
    probabilityResults.set({
      winProbability: data.winProbability * 100,
      loseProbability: data.loseProbability * 100,
      tieProbability: data.tieProbability * 100,
      iterations: data.iterations,
    });
  } catch (error) {
    console.error('Error calculating probabilities:', error);
  }
};

export const resetResults = () => {
  probabilityResults.set({
    winProbability: 0,
    loseProbability: 0,
    tieProbability: 0,
    iterations: 0,
  });
};
