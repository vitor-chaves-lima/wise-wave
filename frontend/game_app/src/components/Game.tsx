import { Container, Sprite, Stage } from "@pixi/react";

import { useState } from "react";
import { Button, Row } from "react-bootstrap";

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

interface GameProps {
    levelTitle: string;
    maps: {
        gameTileMap: TileMap;
        actionsMap: ActionsMap;
    }
    totalActions: number;
}

interface GameHeaderProps {
    levelTitle: string;
}

interface GameStageProps {
    gameTileMap: TileMap;
    actionsMap: ActionsMap;
}

interface GameActionsProps {
    remainingActions: number;
    totalActions: number;
}

/*========== SUB COMPONENTS ==========*/

const GameHeader = ({ levelTitle }: GameHeaderProps) => {
    return <div className="flex p-3 bg-blue-900 text-white text-center fs-2-custom game-header-custom">
        {levelTitle}
    </div>
}

const GameStage = ({ gameTileMap, actionsMap }: GameStageProps) => {
    const { height } = useWindowDimensions();

    const gameHeight = (height / 2) - 68;
    const gameWidth = gameHeight;

    const gameTileHeight = gameHeight / 8;
    const gameTileWidth = gameTileHeight;

    return (
        <Stage height={gameHeight} width={gameWidth} options={{ backgroundColor: 0x145EA8, resolution: 1 }} className="p-0">
            <Container zIndex={2}>
                {gameTileMap.map((row, rowIndex) => {
                    return row.map((tile, tileIndex) => {
                        const x = tileIndex * gameTileWidth;
                        const y = rowIndex * gameTileHeight;

                        return <Sprite key={`${rowIndex}-${tileIndex}`} x={x} y={y} image={tile} width={gameTileWidth} height={gameTileHeight} />
                    });
                })}

            </Container>

            <Container zIndex={1}>
                {actionsMap.map((row, rowIndex) => {
                    return row.map((tile, tileIndex) => {
                        const x = tileIndex * gameTileWidth;
                        const y = rowIndex * gameTileHeight;

                        if (tile === ActionType.NONE) return null;

                        return <Sprite key={`${rowIndex}-${tileIndex}`} x={x} y={y} image={'assets/chest.png'} width={gameTileWidth} height={gameTileHeight} />
                    });
                })}

            </Container>
        </Stage>
    );
}

const GameActions = ({ remainingActions, totalActions }: GameActionsProps) => {
    const buttonsDisabled = remainingActions === 0;
    const canRestart = !(remainingActions < totalActions);

    return (
        <div className="d-flex flex-column actions-custom">
            <div className="d-flex m-0 p-2 bg-blue-900 text-white justify-content-between fs-2-custom">
                Ações: {remainingActions} / {totalActions}

                <Button disabled={canRestart} variant="danger">Reiniciar</Button>
            </div>

            <div className="d-flex flex-column align-items-center justify-content-center actions-custom">
                <Button disabled={buttonsDisabled} className="mb-2 actions-btn-custom">↑</Button>
                <div className="d-flex actions-side-buttons-custom">
                    <Button disabled={buttonsDisabled} className="actions-btn-custom">←</Button>
                    <Button disabled={buttonsDisabled} className="actions-btn-custom">→</Button>
                </div>
                <Button disabled={buttonsDisabled} className="mt-2 actions-btn-custom">↓</Button>
            </div>
        </div>
    )
}

/*========== MAIN COMPONENT ==========*/

const GameComponent = ({ levelTitle, totalActions, maps }: GameProps) => {
    const [remainingActions, _] = useState<number>(totalActions);

    return (
        <div className="d-flex flex-column vh-100">
            <GameHeader levelTitle={levelTitle} />

            <div className="m-0 p-0 d-flex bg-bt-gry-100 game-stage-custom p-0 m-0 justify-content-center">
                <GameStage gameTileMap={maps.gameTileMap} actionsMap={maps.actionsMap} />
            </div>

            <GameActions remainingActions={remainingActions} totalActions={totalActions} />
        </div>
    );
};

/*============== EXPORT ==============*/

export default GameComponent;
