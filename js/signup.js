document.getElementById('signup-form').addEventListener('submit', function (e) {
    e.preventDefault(); // Prevent form from reloading the page

    let username = document.getElementById('username').value;
    let phone = document.getElementById('phone').value;
    let password = document.getElementById('password').value;
    let confirmPassword = document.getElementById('confirm-password').value;
    let errorMessage = document.getElementById('error-message');

    // Basic validation
    if (username === '' || phone === '' || password === '' || confirmPassword === '') {
        errorMessage.textContent = 'All fields are required.';
    } else if (password.length < 6) {
        errorMessage.textContent = 'Password must be at least 6 characters long.';
    } else if (password !== confirmPassword) {
        errorMessage.textContent = 'Passwords do not match.';
    } else {
        // Send the data to the backend via fetch
        fetch('http://127.0.0.1:8080/signup', {
            method: 'POST',
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
            body: `username=${username}&phone=${phone}&password=${password}`
        })
        .then(response => {
            if (response.ok) {
                window.location.href = "http://127.0.0.1:8080/index"; // Redirect to login after successful signup
            } else {
                return response.text().then(text => {
                    errorMessage.textContent = text || 'Signup failed.';
                });
            }
        })
        .catch(error => {
            errorMessage.textContent = 'An error occurred. Please try again.';
            console.error('Error:', error);
        });
    }
});
