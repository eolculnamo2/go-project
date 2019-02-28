const loginButton = document.getElementById('login-button');

function makeFetch() {
    const payload = {
                        user: document.getElementById('user').value,
                        password: document.getElementById('password').value
                    }
    fetch('/login',{
        method: "POST",
        body: JSON.stringify(payload),
        headers: { "Content-Type": "application/json" }
        });
}

loginButton.addEventListener('click', () => makeFetch() );