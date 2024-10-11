import confirm from "../image/circulo-confirm.png";
import Button from "react-bootstrap/Button";
import { Link } from "react-router-dom";

/*========== MAIN COMPONENT ==========*/

const MagicLinkValidateErrorBoundary = () => {
	return (
		<div className="d-flex justify-content-center align-items-center vh-100 bg-light">
			<div className="position-relative" style={{ width: "100%", height: "100%" }}>
				<div className="d-flex flex-column gap-5 justify-content-center align-items-center w-100 p-5">
					<div
						className="d-flex flex-column gap-2 justify-content-center align-items-center w-100 mt-costum-var2">
						<img src={confirm} alt="Imagem de confirmação"
							 style={{ width: "150px", height: "150px" }} />
						<h3 className="text-blue-900  text-center fs-4">Ocorreu um erro ao confirmar a sua conta!</h3>

						<Link to={"/"} replace={true}>
							<Button variant="primary" type="submit" size="lg" className="w-100 bg-bt-blue-500">
								Gerar novo link de acesso
							</Button>
						</Link>
					</div>
				</div>
			</div>
		</div>

	);
};

/*============== EXPORT ==============*/

export default MagicLinkValidateErrorBoundary;
