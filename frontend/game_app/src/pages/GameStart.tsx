// import Button from "react-bootstrap/Button";
import { Link } from "react-router-dom";
import logo from "./../image/ww-home-logo.png";
import returnIcon from "./../image/seta-esquerda.png";
import { Button, Row, Col, Image, Card } from "react-bootstrap";
// import QRCode from "react-qr-code";


/*========== MAIN COMPONENT ==========*/

const GameStartPage = () => {
	return (
		<div className="d-flex justify-content-center align-items-center vh-100 bg-blue-900">
			<div
				className="flex-column gap-5 text-center "
				style={{ width: "80%", height: "100%" }}
			>
				<Link to={"/access"}>
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
						className="mb-3"
						style={{ width: "auto" }}
					/>

					<div className="d-flex flex-column align-items-center justify-content-center ">
						<Card className="text-center p-4  " style={{ maxWidth: "400px", borderRadius: "15px"}}>
							<Card.Body>

								<Card.Title className="mb-4">
									Escanei o nosso QR code e mergulhe nos desafios.
								</Card.Title>
								{/* <QRCode value="link-ou-texto-para-o-qr-code" size={128} /> */}
								<h5 className="mt-10">Escaneie para jogar!</h5>
								<Button className="mt-3" variant="primary">
									Escanear QR Code
								</Button>
							</Card.Body>
						</Card>
					</div>
				</div>
			</div>
		</div>
	);
};

/*============== EXPORT ==============*/
""
export default GameStartPage;
