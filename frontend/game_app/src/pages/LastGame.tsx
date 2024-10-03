import Button from "react-bootstrap/Button";
import { Link } from "react-router-dom";
import logo from "./../image/ww-home-logo.png";
import returnIcon from "./../image/seta-esquerda.png";

/*========== MAIN COMPONENT ==========*/

const LastGamePage = () => {
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
					style={{ height: "90vh" }}
				>
					<img src={logo} alt="logo wise wave" className="mx-auto" />
					<h2 className="fs-3 text-light">Parabéns!</h2>
					<h3 className="fs-5 text-light">
						Você chegou na ultima fase!
					</h3>
					<h4 className="fs-6 text-light">
						Dirija-se até a banca para completar o game.
					</h4>
				</div>
			</div>
		</div>
	);
};

/*============== EXPORT ==============*/

export default LastGamePage;
