"use strict"

async function registerUser(event) {
    event.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    const confirmPassword = document.getElementById("confirm_password").value;

    showErrors(registerButton, []);

    if (password !== confirmPassword) {
        showErrors(registerButton, ["passwords must be the same"]);
        return;
    }

    console.log(password);

    const data = {username: username, password: password};

    const response = await fetch('/api/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });

    const result = await response.json();

    if (response.ok) {
        window.location.href = '/login';;
    } else {
        showErrors(registerButton, [result.errors]);
    }
}

const registerForm = document.getElementById("registerForm");
const registerButton = document.getElementById("registerButton");

registerForm.addEventListener('submit', registerUser);