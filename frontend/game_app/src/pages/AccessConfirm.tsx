import { useLocation } from "react-router-dom";

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
        <div className="d-flex justify-content-md-center align-items-center vh-100">
            <div className='d-flex flex-column gap-5 justify-content-md-center'>
                <h1 className="mx-auto">Confirme o seu e-mail!</h1>
                <h2>Você está a um passo de ter a melhor experiência de aprendizagem!</h2>
                <h3>{obfuscatedEmail}</h3>
            </div>
        </div>
    );
};

/*============== EXPORT ==============*/

export default AccessConfirmPage;