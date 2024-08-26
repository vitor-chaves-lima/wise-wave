import { useState, FormEvent } from "react";
import { Link, useSubmit, useNavigate } from "react-router-dom";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import returnIcon from './../image/seta-esquerda.png';

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
        <div className="d-flex justify-content-center align-items-center vh-100 bg-blue-900">
        <div className='position-relative' style={{ width: '100%', height: '100%' }}>
          <Link to={"/"}>
            <Button variant="link" className="position-absolute start-0 mt-costum-return ms-3">
              <img src={returnIcon} alt="Ícone de retorno" />
            </Button>
          </Link>
      
          <div className="d-flex flex-column gap-5 justify-content-md-center align-items-center" style={{ width: '80%', margin: '0 auto' }}>
            <Form onSubmit={onFormSubmit} className="w-100 mt-costum">
              <h1 className="mx-auto text-light mt-4">Acesse</h1>
              <p className="mx-auto text-light fs-5">Utilizando seu e-mail</p>
      
              <Form.Group className="mb-3" controlId="email">
                <Form.Label htmlFor="inputEmail" className="text-light ">Email</Form.Label>
                <Form.Control
                  type="email" placeholder="Digite seu e-mail" onChange={onEmailChange} value={email} id="inputEmail"
                  size="lg"
                  
                />
              </Form.Group>
      
              <Button variant="primary" type="submit" size="lg" className="w-100 bg-bt-blue-500">
                Acessar
              </Button>
            </Form>
          </div>
        </div>
      </div>
      



    )
}

/*============== EXPORT ==============*/

export default AccessPage