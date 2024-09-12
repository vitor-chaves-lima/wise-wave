import { Stage } from "@pixi/react";
import { Button, Col, Container, Row, Stack } from "react-bootstrap";
import useWindowDimensions from "../hooks/windowDimensions";

/*============ INTERFACES ============*/

interface GameHeaderProps {
    levelTitle: string;
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

const GameStage = () => {
    const { height } = useWindowDimensions();

    const gameHeight = (height / 2) - 68;
    const gameWidth = gameHeight;

    return (
        <Row className="game-stage-custom p-0 m-0 justify-content-center">
            <Stage height={gameHeight} width={gameWidth} options={{ background: 0x145EA8, resolution: 1 }} className="p-0" >
            </Stage>
        </Row>
    );
}

const GameActions = ({ remainingActions, totalActions }: GameActionsProps) => {
    return (
        <div className="d-flex flex-column actions-custom">
            <div className="d-flex m-0 p-2 bg-blue-900 text-white justify-content-between fs-2-custom">
                Ações: {remainingActions} / {totalActions}

                <Button disabled={true} variant="danger">Reiniciar</Button>
            </div>

            <div className="d-flex flex-column align-items-center justify-content-center actions-custom">
                <Button className="mb-2 actions-btn-custom">↑</Button>
                <div className="d-flex actions-side-buttons-custom">
                    <Button className="actions-btn-custom">←</Button>
                    <Button className="actions-btn-custom">→</Button>
                </div>
                <Button className="mt-2 actions-btn-custom">↓</Button>
            </div>
        </div>
    )
}

/*========== MAIN COMPONENT ==========*/

const GameComponent = () => {
    return (
        <div className="d-flex flex-column vh-100">
            <GameHeader levelTitle="Fase 1" />

            <div className="m-0 p-0 bg-bt-gry-100">
                <GameStage />
            </div>

            <GameActions remainingActions={3} totalActions={3} />
        </div>
    );
};

/*============== EXPORT ==============*/

export default GameComponent;