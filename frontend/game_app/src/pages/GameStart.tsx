// import Button from "react-bootstrap/Button";
import {Link, useNavigate} from "react-router-dom";
import logo from "./../image/ww-home-logo.png";
import returnIcon from "./../image/seta-esquerda.png";
import { Button, Card } from "react-bootstrap";
import { Scanner, IDetectedBarcode } from '@yudiel/react-qr-scanner';

/*========== MAIN COMPONENT ==========*/

const GameStartPage = () => {
	const navigate = useNavigate();

	const handleScan = (detectedCodes: IDetectedBarcode[]) => {
	 	if(detectedCodes[0].rawValue !== "3077d193ec31e39f04c67efe50e7d7b4"){
			alert("Código inválido");
		 	return;
		}

		navigate("/game", {replace: true})
	}

	const handleError = () => {
		alert("Não foi possível ler o QR code")
	}

	return (
		<div className="d-flex justify-content-center align-items-center vh-100 bg-blue-900">
			<div
				className="flex-column gap-5 text-center "
				style={{ width: "80%", height: "100%" }}
			>
				<Link to={"/"}>
					<Button
						variant="link"
						className="position-absolute start-0 mt-costum-return ms-3"
					>
						<img src={returnIcon} alt="Ícone de retorno" />
					</Button>
				</Link>
				<div
					className="d-flex flex-column justify-content-center align-items-center w-100"
					style={{ height: "80vh" }}
				>
					<img
						src={logo}
						className="mb-5"
						style={{ width: "auto" }}
					 alt={"Logo"}/>

					<div className="d-flex flex-column align-items-center justify-content-center ">
						<Card className="text-center p-4 " style={{ maxWidth: "400px", borderRadius: "15px"}}>
							<Card.Body>

								<Card.Title className="mb-4">
									Escaneie o nosso QR code e mergulhe nos desafios.
								</Card.Title>

								<Scanner onScan={handleScan} onError={handleError} classNames={{
									container: "qr-code-container"
								}} allowMultiple={false} formats={["qr_code"]} />

								{/*<h5 className="mt-10">Escaneie para jogar!</h5>*/}
								{/*<Button className="mt-3" variant="primary">*/}
								{/*	Escanear QR Code*/}
								{/*</Button>*/}
							</Card.Body>
						</Card>
					</div>
				</div>
			</div>
		</div>
	);
};

/*============== EXPORT ==============*/
export default GameStartPage;
