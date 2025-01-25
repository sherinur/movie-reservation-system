document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent form submission

    // Example validation logic
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const errorMessage = document.getElementById('error-message');

    if (email !== 'test@example.com' || password !== 'password123') {
        errorMessage.style.display = 'block';
    } else {
        errorMessage.style.display = 'none';
        alert('Login successful!'); // Example success action
    }
});