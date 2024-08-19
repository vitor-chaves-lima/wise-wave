import { useState, FormEvent } from "react";
import { Link, useSubmit, useNavigate } from "react-router-dom";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

/*========== MAIN COMPONENT ==========*/

const AccessPage = () => {
    let submit = useSubmit();
    let navigate = useNavigate();

    const [email, setEmail] = useState<string>("");

    const onEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.target.value);

    const onFormSubmit = (e: FormEvent) => {
        e.preventDefault();

        submit({
            email
        }, {
            method: "POST",
            action: "/access"
        });

        navigate("/access-confirm", {
            state: { email },
            replace: true,
        });
    }

    return (
        <div className="d-flex justify-content-md-center align-items-center vh-100">
            <div className='d-flex flex-column gap-5 justify-content-md-center'>
                <Link to={"/"} ><Button variant="primary" className="w-100">Voltar</Button></Link>

                <Form onSubmit={onFormSubmit}>
                    <h1 className="mx-auto">Acesse</h1>
                    <p className="mx-auto">Utilizando seu e-mail</p>

                    <Form.Group className="mb-3" controlId="email">
                        <Form.Control type="email" placeholder="Digite seu e-mail" onChange={onEmailChange} value={email} />
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