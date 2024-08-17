import { Link } from "react-router-dom";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

/*========== MAIN COMPONENT ==========*/

const AccessPage = () => {
    return (
        <div className="d-flex justify-content-md-center align-items-center vh-100">
            <div className='d-flex flex-column gap-5 justify-content-md-center'>
                <Link to={"/"} ><Button variant="primary" className="w-100">Voltar</Button></Link>

                <Form>
                    <h1 className="mx-auto">Acesse</h1>
                    <p className="mx-auto">Utilizando seu e-mail</p>

                    <Form.Group className="mb-3" controlId="email">
                        <Form.Control type="email" placeholder="Digite seu e-mail" />
                    </Form.Group>
                    <Button variant="primary" type="submit" className="w-100">
                        Acessar
                    </Button>
                </Form>
            </div>
        </div >
    )
}

/*============== EXPORT ==============*/

export default AccessPage