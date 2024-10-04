import React, { useState } from "react";
import { Button } from "react-bootstrap";
import { Row, Col } from 'react-bootstrap';
import returnIcon from './../image/seta-esquerda.png';
import { Link } from "react-router-dom";
import { CircularProgressbar, buildStyles } from 'react-circular-progressbar';
import 'react-circular-progressbar/dist/styles.css'; 

/*========== MAIN COMPONENT ==========*/

const ScorePage = () => {

    const [score, setScore] = useState(50); 

    return (
        <div className="d-flex justify-content-center align-items-center vh-100 bg-blue-900">
            <div className='position-relative' style={{ width: '100%', height: '100%' }}>
                <Link to={"/"}>
                    <Button variant="link" className="position-absolute start-0 mt-costum-return ms-3">
                        <img src={returnIcon} alt="Ícone de retorno" />
                    </Button>
                </Link>

                <div className="d-flex flex-column justify-content-center align-items-center vh-100" style={{ width: '80%', margin: '0 auto' }}>
                    <div className="mt-custom d-flex flex-column">
                        <h1 className="text-light mt-4 text-center">Parabéns!</h1>
                        <p className="text-light fs-5 text-center mb-5">Pontuação</p>

                        <Row className="d-flex justify-content-center align-items-center text-center ">
                            <Col md={6} className="d-flex justify-content-center align-items-center">
                                <div className="d-flex flex-column align-items-center" style={{ position: 'relative' }}>
                                    <div
                                        className="d-flex flex-column justify-content-center align-items-center border border-5 border-white rounded-circle bg-bt-gry-100"
                                        style={{ width: '250px', height: '250px', backgroundColor: 'transparent', zIndex: 10 }}
                                    >
                                        
                                        <CircularProgressbar
                                            value={score}
                                            text={`${score}`}
                                            styles={buildStyles({
                                                pathColor: `rgba(255, 255, 255, ${score / 100})`,
                                                textColor: '#fff',
                                                trailColor: 'rgba(255, 255, 255, 0.2)',
                                                backgroundColor: 'transparent',
                                            })}
                                        />
                                        
                                    </div>
                                    
                                </div>
                            </Col>
                        </Row>
                    </div>
                </div>
            </div>
        </div>
    );
};

/*============== EXPORT ==============*/

export default ScorePage;
