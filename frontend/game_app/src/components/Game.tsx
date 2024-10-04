import {Container, Sprite, Stage} from "@pixi/react";

import {Dispatch, SetStateAction, useState} from "react";
import {Button, Row} from "react-bootstrap";

import useWindowDimensions from "../hooks/windowDimensions";


/*=============== TYPE ===============*/

export enum TileType {
	GROUND = "assets/ground.png",
	WALL = "assets/wall.png",
	WATER = "assets/water.png",
}

export enum ActionType {
	WIN,
	LOSE,
	NONE
}

type Row<T> = [T, T, T, T, T, T, T, T];
type Matrix8x8<T> = [Row<T>, Row<T>, Row<T>, Row<T>, Row<T>, Row<T>, Row<T>, Row<T>];
export type TileMap = Matrix8x8<TileType>;
export type ActionsMap = Matrix8x8<ActionType>;

/*============ INTERFACES ============*/

interface PlayerPosition {
	x: number;
	y: number;
}

interface GameProps {
	levelTitle: string;
	maps: {
		gameTileMap: TileMap;
		actionsMap: ActionsMap;
	}
	startPosition: PlayerPosition;
	totalActions: number;
}

interface GameHeaderProps {
	levelTitle: string;
}

interface GameStageProps {
	gameTileMap: TileMap;
	actionsMap: ActionsMap;
	gameWidth: number;
	gameHeight: number;
	tileHeight: number;
	tileWidth: number;
	characterPosition: PlayerPosition
}

interface GameActionsProps {
	remainingActions: number;
	setRemainingActions: Dispatch<SetStateAction<number>>
	totalActions: number;
	tileHeight: number;
	tileWidth: number;
	characterPosition: PlayerPosition;
	setCharacterPosition: Dispatch<SetStateAction<PlayerPosition>>,
	handleRestart: () => void
}

/*========== SUB COMPONENTS ==========*/

const GameHeader = ({levelTitle}: GameHeaderProps) => {
	return <div className="flex p-3 bg-blue-900 text-white text-center fs-2-custom game-header-custom">
		{levelTitle}
	</div>
}

const GameStage = ({
					   gameTileMap,
					   actionsMap,
					   characterPosition,
					   gameHeight,
					   gameWidth,
					   tileWidth,
					   tileHeight
				   }: GameStageProps) => {

	return (
		<Stage height={gameHeight} width={gameWidth} options={{backgroundColor: 0x145EA8, resolution: 1}}
			   className="p-0">
			<Container zIndex={2}>
				{gameTileMap.map((row, rowIndex) => {
					return row.map((tile, tileIndex) => {
						const x = tileIndex * tileWidth;
						const y = rowIndex * tileHeight;

						return <Sprite key={`${rowIndex}-${tileIndex}`} x={x} y={y} image={tile} width={tileWidth}
									   height={tileHeight}/>
					});
				})}

			</Container>

			<Container zIndex={1}>
				{actionsMap.map((row, rowIndex) => {
					return row.map((tile, tileIndex) => {
						const x = tileIndex * tileWidth;
						const y = rowIndex * tileHeight;

						if (tile === ActionType.NONE) return null;

						return <Sprite key={`${rowIndex}-${tileIndex}`} x={x} y={y} image={'assets/chest.png'}
									   width={tileWidth} height={tileHeight}/>
					});
				})}

			</Container>

			<Sprite x={characterPosition.x} y={characterPosition.y} image={"assets/character.png"} width={tileWidth}
					height={tileHeight} zIndex={2}/>
		</Stage>
	);
}

const GameActions = ({
						 remainingActions,
						 setRemainingActions,
						 totalActions,
						 characterPosition,
						 setCharacterPosition,
						 tileWidth,
						 tileHeight,
						 handleRestart
					 }: GameActionsProps) => {
	const buttonsDisabled = remainingActions === 0;
	const canRestart = !(remainingActions < totalActions);

	const canMoveLeft = (characterPosition.x > 0);
	const canMoveRight = (characterPosition.x < 7 * tileWidth);
	const canMoveTop = (characterPosition.y > 0);
	const canMoveBottom = (characterPosition.y < 7 * tileHeight);

	const subtractAction = () => setRemainingActions((r) => r > 0 ? r - 1 : r);

	const move = (x: number, y: number) => {
		setCharacterPosition((p) => ({
			x: p.x + (tileWidth * x),
			y: p.y + (tileHeight * y)
		}));

		subtractAction();
	}

	const handleMoveLeft = () => move(-1, 0);
	const handleMoveRight = () => move(1, 0);
	const handleMoveTop = () => move(0, -1);
	const handleMoveBottom = () => move(0, 1);

	return (
		<div className="d-flex flex-column actions-custom">
			<div className="d-flex m-0 p-2 bg-blue-900 text-white justify-content-between fs-2-custom">
				Ações: {remainingActions} / {totalActions}

				<Button disabled={canRestart} variant="danger" onClick={handleRestart}>Reiniciar</Button>
			</div>

			<div className="d-flex flex-column align-items-center justify-content-center actions-custom">
				<Button disabled={buttonsDisabled || !canMoveTop} className="mb-2 actions-btn-custom"
						onClick={handleMoveTop}>↑</Button>
				<div className="d-flex actions-side-buttons-custom">
					<Button disabled={buttonsDisabled || !canMoveLeft} className="actions-btn-custom"
							onClick={handleMoveLeft}>←</Button>
					<Button disabled={buttonsDisabled || !canMoveRight} className="actions-btn-custom"
							onClick={handleMoveRight}>→</Button>
				</div>
				<Button disabled={buttonsDisabled || !canMoveBottom} className="mt-2 actions-btn-custom"
						onClick={handleMoveBottom}>↓</Button>
			</div>
		</div>
	)
}

/*========== MAIN COMPONENT ==========*/

const GameComponent = ({levelTitle, totalActions, maps, startPosition}: GameProps) => {
	const {height} = useWindowDimensions();

	const gameHeight = (height / 2) - 68;
	const gameWidth = gameHeight;

	const gameTileHeight = gameHeight / 8;
	const gameTileWidth = gameTileHeight;

	const [remainingActions, setRemainingActions] = useState<number>(totalActions);
	const [characterPosition, setCharacterPosition] = useState<PlayerPosition>(startPosition)

	const handleRestart = () => {
		setCharacterPosition(startPosition);
		setRemainingActions(totalActions);
	}

	return (
		<div className="d-flex flex-column vh-100">
			<GameHeader levelTitle={levelTitle}/>

			<div className="m-0 p-0 d-flex bg-bt-gry-100 game-stage-custom p-0 m-0 justify-content-center">
				<GameStage gameTileMap={maps.gameTileMap} actionsMap={maps.actionsMap}
						   tileHeight={gameTileHeight} tileWidth={gameTileWidth}
						   characterPosition={characterPosition} gameWidth={gameWidth} gameHeight={gameHeight}/>
			</div>

			<GameActions remainingActions={remainingActions} setRemainingActions={setRemainingActions}
						 totalActions={totalActions}
						 tileHeight={gameTileHeight} tileWidth={gameTileWidth}
						 characterPosition={characterPosition} setCharacterPosition={setCharacterPosition}
						 handleRestart={handleRestart}/>
		</div>
	);
};

/*============== EXPORT ==============*/

export default GameComponent;
