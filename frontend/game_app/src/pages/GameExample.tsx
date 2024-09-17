import GameComponent, { ActionsMap, ActionType, TileMap, TileType } from '../components/Game';

/*========== MAIN COMPONENT ==========*/

const GameExamplePage = () => {

    const gameTileMap: TileMap = [
        [TileType.WALL, TileType.WALL, TileType.WATER, TileType.WATER, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.WALL, TileType.WALL, TileType.WATER, TileType.WATER, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.WALL, TileType.WATER, TileType.WATER, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.WALL, TileType.WATER, TileType.WATER, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.WALL, TileType.WATER, TileType.WATER, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
        [TileType.GROUND, TileType.WALL, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND, TileType.GROUND],
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
                actionsMap
            }} />
        </div>

    );
};

/*============== EXPORT ==============*/

export default GameExamplePage;