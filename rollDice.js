class Player {
    constructor(name, diceCount) {
      this.name = name;
      this.diceCount = diceCount;
      this.points = 0;
    }
  
    rollDice() {
      const diceResults = [];
      for (let i = 0; i < this.diceCount; i++) {
        diceResults.push(Math.floor(Math.random() * 6) + 1);
      }
      return diceResults;
    }
  }
  
  function playGame(N, M) {
    const players = [];
    for (let i = 1; i <= N; i++) {
      players.push(new Player(`Pemain #${i}`, M));
    }
  
    let round = 1;
    while (players.length > 1) {
      console.log(`[
  Pemain = ${N}, Dadu = ${M}
  ==================
  "Round ${round}"`);
  
      const diceResults = [];
      players.forEach(player => {
        if (player.diceCount > 0) {
          const playerRoll = player.rollDice();
          console.log(
            `${player.name} (${player.points}): ${playerRoll.join(', ')}`
          );
          diceResults.push({ player, rolls: playerRoll });
        }
      });
  
      console.log(`Setelah evaluasi:`);
      diceResults.forEach(({ player, rolls }) => {
        rolls.forEach(result => {
          if (result === 6) {
            player.points++;
          } else if (result === 1) {
            const nextPlayerIndex =
              (players.indexOf(player) + 1) % players.length;
            const nextPlayer = players[nextPlayerIndex];
            nextPlayer.diceCount++;
            // Hapus angka 1 dari hasil lemparan
            rolls.splice(rolls.indexOf(1), 1);
          }
        });
  
        const validResults = rolls.filter(result => result >= 1 && result <= 5);
        player.diceCount = validResults.length;
        console.log(
          `${player.name} (${player.points}): ${[
            ...validResults,
            ...Array(player.diceCount - validResults.length).fill('_'),
          ].join(', ')}`
        );
      });
  
      console.log('==================');
      round++;
  
      players
        .filter(player => player.diceCount === 0)
        .forEach(player => {
          console.log(`${player.name} telah selesai bermain.`);
          players.splice(players.indexOf(player), 1);
        });
    }
  
    const winner = players.reduce((prev, current) =>
      prev.points > current.points ? prev : current
    );
    console.log(`${winner.name} adalah pemenang dengan ${winner.points} poin!`);
  }
  
  const N = 2; // masukkan jumlah player disini
  const M = 4; // masukkan jumlah putaran dadu disini
  playGame(N, M);
  