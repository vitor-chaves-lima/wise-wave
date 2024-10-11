import confirm from "../image/circulo-confirm.png";
import { Spinner } from "react-bootstrap";
import { Await, Navigate, useLoaderData } from "react-router-dom";
import { Suspense } from "react";

/*========== MAIN COMPONENT ==========*/

const MagicLinkValidatePage = () => {
	const data = useLoaderData();

	return (
		<div className="d-flex justify-content-center align-items-center vh-100 bg-light">
			<div className="position-relative" style={{ width: "100%", height: "100%" }}>
				<div className="d-flex flex-column gap-5 justify-content-center align-items-center w-100 p-5">
					<div
						className="d-flex flex-column gap-2 justify-content-center align-items-center w-100 mt-costum-var2">
						<img src={confirm} alt="Imagem de confirmação"
							 style={{ width: "150px", height: "150px" }} />
						<h3 className="text-blue-900  text-center fs-4">Você está a um passo de ter a melhor
							experiência
							de aprendizado!</h3>
					</div>
					<Suspense fallback={
						<Spinner animation="border" role="status">
							<span className="visually-hidden">Aguarde...</span>
						</Spinner>
					}>
						<Await resolve={data} children={<Navigate to={"/game-start"} />}/>
					</Suspense>
				</div>
			</div>
		</div>

	);
};

/*============== EXPORT ==============*/

export default MagicLinkValidatePage;
