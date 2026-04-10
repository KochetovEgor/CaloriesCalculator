"use strict"

function renderUserLoggedIn(username) {
    const container = document.getElementById("auth-container");

    while(container.firstChild) {
        container.removeChild(container.firstChild);
    }

    const usernameRender = document.createElement("p");
    usernameRender.textContent = username;

    const logoutLink = document.createElement("a");
    logoutLink.textContent = "Выйти";
    logoutLink.setAttribute("href", "/logout");

    container.appendChild(usernameRender);
    container.appendChild(logoutLink);
}

function renderUserLoggedOut() {
    const container = document.getElementById("auth-container");

    while(container.firstChild) {
        container.removeChild(container.firstChild);
    }

    const loginLink = document.createElement("a");
    loginLink.textContent = "Войти";
    loginLink.setAttribute("href", "/login");

    container.appendChild(loginLink);
}