import GameComponent, {ActionsMap, ActionType, TileMap, TileType} from '../components/Game';

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
			<GameComponent levelTitle='Teste' totalActions={10} maps={{
				gameTileMap,
				actionsMap,
			}} startPosition={{
				x: 0,
				y: 0,
			}} question={"Qual é o resultado de 8 + 5?"} alternatives={[
				{
					index: 1,
					text: "10"
				},
				{
					index: 2,
					text: "12"
				},
				{
					index: 3,
					text: "13"
				},
				{
					index: 4,
					text: "15"
				}
			]}/>
		</div>

	);
};

/*============== EXPORT ==============*/

export default GameExamplePage;
