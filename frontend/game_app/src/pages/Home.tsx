import Button from 'react-bootstrap/Button';
import { Link } from "react-router-dom";
import logo from './../image/ww-home-logo.png';

/*========== MAIN COMPONENT ==========*/

const HomePage = () => {

    return (
        <div className="d-flex justify-content-center align-items-center vh-100 bg-blue-900">
            <div className="d-flex flex-column gap-5 text-center " style={{ width: '80%' }}>
                <div>
                    <img src={logo} alt="logo wise wave" className="mx-auto" />
                    <h1 className="text-light fs-1-custom text-logo-ww">WiseWave</h1>
                </div> 
                <h2 className="fs-3 text-light">Seja bem-vindo!</h2>
                <Link to={"access"}>
                    <Button variant="primary" size="lg" className="w-100 bg-bt-blue-500">Iniciar</Button>
                </Link>
            </div>
        </div>

    );
};

/*============== EXPORT ==============*/

export default HomePage;