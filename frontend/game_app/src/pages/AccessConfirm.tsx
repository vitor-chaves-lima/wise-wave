import { useLocation } from "react-router-dom";
import returnIconBlue from './../image/seta-esquerda-blue-700.png';
import { Link } from "react-router-dom";
import { Button } from "react-bootstrap";
import logo from './../image/ww-home-logo.png';
import confirm from './../image/circulo-confirm.png'

/*========== MAIN COMPONENT ==========*/

const AccessConfirmPage = () => {
    let { state } = useLocation();
    const email = state.email as string;
    const obfuscateEmail = (email: string) => {
        const [username, domain] = email.split("@");
        const obfuscatedUsername = username.slice(0, Math.floor(username.length / 2)) + "*".repeat(username.length - Math.floor(username.length / 2));
        return obfuscatedUsername + "@" + domain;
    };

    const obfuscatedEmail = obfuscateEmail(email);

    return (
        <div className="d-flex justify-content-center align-items-center vh-100 bg-light">
            <div className='position-relative' style={{ width: '100%', height: '100%' }}>
                <Link to={"/access"}>
                    <Button variant="link" className="position-absolute start-0 mt-costum-return ms-3">
                        <img src={returnIconBlue} alt="Ícone de retorno" />
                    </Button>
                </Link>
                <div className='d-flex flex-column gap-5 justify-content-center align-items-center w-100' >
                    <div className="d-flex flex-column gap-2 justify-content-center align-items-center w-100 mt-costum-var2">
                        <img src={confirm} alt="Imagem de confirmação" style={{ width: '150px', height: '150px' }} />
                        <h2 className="text-blue-700 fs-1 fw-bold">Confirme o seu e-mail!</h2>
                        <h3 className="text-blue-900  text-center fs-4">Você está a um passo de ter a melhor experiência de aprendizado!</h3>
                    </div>
                    <div className="d-flex flex-column gap- justify-content-center align-items-center w-100">
                        <h3 className="text-blue-900 fs-5">{obfuscatedEmail}</h3>
                        <img src={logo} alt="Logo Wise-Wave" />
                    </div>
                </div>
            </div>
        </div>


    );
};

/*============== EXPORT ==============*/

export default AccessConfirmPage;