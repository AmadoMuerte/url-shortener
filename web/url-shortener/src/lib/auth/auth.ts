interface Auth {
    login: string
    password: string
}

export async function checkAuth(Auth: Auth) {
    const response = await fetch('http://localhost:8082/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(Auth)
    });
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return await response.json();
}