import { PairsManagerBehavior } from '../../interfaces/Managers';
import { shuffleArray } from '../../lib';
import { logger } from '../Logger';
import { Player } from '../Player';

export class BasePairsManager implements PairsManagerBehavior {
  createPairs(players: Player[]): [Player, Player][] {
    const battlePairs = [];
    const shuffledPlayers = shuffleArray(players);
    logger.debug({
      message: 'shuffledPlayers',
    });
    logger.log(shuffledPlayers.length);
    for (let i = 0; i < shuffledPlayers.length; i += 2) {
      // @ts-ignore
      battlePairs.push([shuffledPlayers[i], shuffledPlayers[i + 1]] as [Player, Player]);
    }

    return battlePairs;
  }
}
