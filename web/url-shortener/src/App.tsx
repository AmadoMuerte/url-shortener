import {useCookies} from "react-cookie";
import AdminPanel from "./pages/AdminPanel/AdminPanel.tsx";
import Login from "./pages/Login/Login.tsx";
import {useEffect, useState} from "react";
import {checkAuth} from "./lib/auth/auth.ts";


function App() {
    const [cookies] = useCookies(['login', 'password']);
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const authenticate = async () => {
            if (cookies.login && cookies.password) {
                try {
                    const data = await checkAuth({ login: cookies.login, password: cookies.password });
                    if (data.status === 'OK') {
                        setIsAuthenticated(true);
                    } else {
                        setIsAuthenticated(false);
                    }
                } catch (error) {
                    console.error('Ошибка аутентификации');
                }
            }
            setIsLoading(false);
        };

        if (!isAuthenticated) {
            authenticate();
        }

    }, [cookies, isAuthenticated]);

    if (isLoading) {
        return <div>Загрузка...</div>;
    }

    return isAuthenticated ? <AdminPanel /> : <Login />;
}


export default App
