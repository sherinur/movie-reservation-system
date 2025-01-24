document.addEventListener('DOMContentLoaded', () =>{
    const submitBtn = document.querySelector('#submit-btn');

    if (submitBtn) {
        submitBtn.addEventListener('click', ValidateRegister);
    }
});

const fullnameRegex = /^[a-zA-Z0-9_-]{1,39}$/;
const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*(),.?"{}<>])[A-Za-z\d!@#$%^&*(),.?"{}<>]{8,}$/;
const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

const fullnameErrorMsg = `Fullname can contain only alphanumeric characters, hyphens (-), 
                          and underscores (_), and must be between 1 and 39 characters in length.`;

const passwordErrorMsg = `Password must be at least 8 characters long, 
                          contain at least one lowercase letter, one uppercase letter,
                          one digit, and one special character (e.g., !@#$%^&*).`;

const confirmPasswordMsg = `Passwords do not match. Please make sure both passwords are identical.`;

const emailErrorMsg = `Invalid email address. Please enter a valid email address.`;

async function ValidateRegister() {
    const fullnameInput = document.querySelector('#fullname');
    const emailInput = document.querySelector('#email');
    const passwordInput = document.querySelector('#password');
    const confirmPasswordInput = document.querySelector('#confirm-password');

    if (!fullnameInput || !emailInput || !passwordInput || !confirmPasswordInput) {
        return
    }

    const fullname = fullnameInput.value;
    const email = emailInput.value;
    const password = passwordInput.value;
    const confirmPassword = confirmPasswordInput.value;

    if (!fullnameRegex.test(fullname)) {
        const fullnameError = document.querySelector('#fullname-error');
        ShowError(fullnameError, fullnameInput, fullnameErrorMsg);
    }

    if (!emailRegex.test(email)) {
        const emailError = document.querySelector('#email-error');
        ShowError(emailError, emailInput, emailErrorMsg);
    }

    if (!passwordRegex.test(password)) {
        const passwordError = document.querySelector('#password-error');
        ShowError(passwordError, passwordInput, passwordErrorMsg);
    }

    if (confirmPassword != password) {
        const confirmPasswordError = document.querySelector('#confirm-password-error');
        ShowError(confirmPasswordError, confirmPasswordInput, confirmPasswordMsg);
    }

    MakeRegisterRequest(fullname, email, password);
}

async function HideError(errorText, input) {
    errorText.style.display = 'none';
    input.classList.remove('error-input');
}

function ShowError(errorMsgElement, input, errorText) {
    input.classList.add('error-input');

    errorMsgElement.textContent = errorText;

    errorMsgElement.style.display = 'block';

    input.addEventListener('click', () => HideError(errorMsgElement, input));
}

// TODO: Rewrite the fetch to the user-service:
// ! Uncaught (in promise) SyntaxError: JSON.parse: unexpected character at line 1 column 1 of the JSON data
//     MakeRegisterRequest http://localhost:4200/register.validation.js:98
//     AsyncFunctionThrow self-hosted:804
//     (Async: async)
//     ValidateRegister http://localhost:4200/register.validation.js:59
//     (Async: EventListener.handleEvent)
//     <anonymous> http://localhost:4200/register.validation.js:5
//     (Async: EventListener.handleEvent)
//     <anonymous> http://localhost:4200/register.validation.js:1


const url = 'http://127.0.0.1:8080/register';

async function MakeRegisterRequest(username, email, password) {
    const data = {
        username: username,
        email: email,
        password: password
    }

    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    });

    if (!response.ok) {
        console.log(`error: ${response.status}`);
    }

    const result = await response.json();
    console.log(`${result}`)
}