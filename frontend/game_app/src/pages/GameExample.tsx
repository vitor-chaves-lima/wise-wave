import GameComponent, { ActionsMap, ActionType, TileMap, TileType } from '../components/Game';

/*========== MAIN COMPONENT ==========*/

const GameExamplePage = () => {

    const gameTileMap: TileMap = [
        [TileType.GROUND, TileType.GROUND, TileType.WATER, TileType.WATER, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.GROUND, TileType.WATER, TileType.WATER, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.GROUND, TileType.WATER, TileType.WATER, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.GROUND, TileType.WATER, TileType.WATER, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.GROUND, TileType.WATER, TileType.WATER, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
    ]

    const actionsMap: ActionsMap = [
        [ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE],
        [ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE],
        [ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE],
        [ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE],
        [ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE],
        [ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE],
        [ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.WIN, ActionType.NONE],
        [ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE, ActionType.NONE],
    ]

    return (
        <div>
            <GameComponent levelTitle='Teste' totalActions={4} maps={{
                gameTileMap,
                actionsMap,
            }} startPosition={{
				x: 0,
				y: 0,
			}} />
        </div>

    );
};

/*============== EXPORT ==============*/

export default GameExamplePage;
