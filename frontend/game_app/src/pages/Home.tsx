import { Link } from "react-router-dom";
import Button from 'react-bootstrap/Button';

/*========== MAIN COMPONENT ==========*/

const HomePage = () => {

    return (
        <div className="d-flex justify-content-md-center align-items-center vh-100">
            <div className='d-flex flex-column gap-5 justify-content-md-center'>
                <h1 className="mx-auto">Seja bem-vindo!</h1>

                <Link to={"access"} ><Button variant="primary" className="w-100">Iniciar</Button></Link>
            </div>
        </div>
    );
};

/*============== EXPORT ==============*/

export default HomePage;