document.getElementById('login-form').addEventListener('submit', function (e) {
    e.preventDefault(); // Prevent the form from reloading the page

    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;
    let errorMessage = document.getElementById('error-message');

    fetch('http://127.0.0.1:8080/index', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: `username=${username}&password=${password}`
    })
    .then(response => {
        if (response.ok) {
            window.location.href = "http://127.0.0.1:8080/shop.html"; // Redirect to shop after login
        } else {
            return response.text().then(text => {
                errorMessage.textContent = text || 'Login failed. Please try again.';
            });
        }
    })
    .catch(error => {
        errorMessage.textContent = 'An error occurred. Please try again.';
        console.error('Error:', error);
    });
});
