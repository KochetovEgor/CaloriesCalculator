"use strict"

async function loginUser(event) {
    event.preventDefault();

    showErrors(loginButton, []);

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    const authHeader = "Basic " + btoa(username + ":" + password)

    const response = await fetch("/api/login", {
        method: "POST",
        headers: {
            "Authorization": authHeader
        }
    });

    const result = await response.json();

    if (response.ok) {
        localStorage.setItem("access_token", result.access_token);
        window.location.href = "/";
    } else {
        showErrors(loginButton, result.errors);
    }
}

const loginForm = document.getElementById("loginForm")
const loginButton = document.getElementById("loginButton")

loginForm.addEventListener("submit", loginUser)