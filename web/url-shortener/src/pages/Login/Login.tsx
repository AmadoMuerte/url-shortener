import './Login.css';
import {useState} from 'react';
import { useCookies } from 'react-cookie';
import {checkAuth} from "../../lib/auth/auth.ts";

function Login() {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');
    const [, setCookies] = useCookies(['login', 'password']);
    const [errorMessage, setErrorMessage] = useState('');

    const  handleSubmit  = async (event: { preventDefault: () => void; }) => {
        event.preventDefault();

        if (login && password) {
            try {
                const data = await checkAuth({ login: login, password: password });
                if (data.status === 'OK') {
                    setCookies('login', login, { maxAge: 300600 });
                    setCookies('password', password, { maxAge: 300600 });
                } else {
                    setErrorMessage('Неверный логин или пароль');
                }
            } catch (error) {
                console.error('Ошибка аутентификации:', error);
            }
        }
    };

    return (
        <div className="login-container">
            <form className="login-form" onSubmit={handleSubmit}>
                <h2>Вход в систему</h2>
                <div className="input-group">
                    <label htmlFor="login">Логин:</label>
                    <input
                        type="text"
                        id="login"
                        name="login"
                        value={login}
                        onChange={(e) => setLogin(e.target.value)}
                        required
                    />
                </div>
                <div className="input-group">
                    <label htmlFor="password">Пароль:</label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                <button type="submit" className="login-button">Войти</button>
            </form>
            {errorMessage && <p className="error-message">{errorMessage}</p>}
        </div>
    );
}

export default Login;
