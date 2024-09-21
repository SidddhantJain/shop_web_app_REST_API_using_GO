window.onload = function() {
    // Fetch user details from the session API and update UI
    fetch('/api/session')
        .then(response => response.json())
        .then(data => {
            if (data.username) {
                document.getElementById('username').textContent = data.username;
            } else {
                window.location.href = "index.html"; // Redirect to login if not authenticated
            }
        });
};

function logout() {
    fetch('/api/logout', { method: 'POST' })
        .then(() => {
            window.location.href = "index.html";
        });
}
