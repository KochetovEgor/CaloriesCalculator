"use strict"

function getAuthData() {
    const token = localStorage.getItem('access_token'); 
    if (!token) {
        return null;
    }

    try {
        const base64Url = token.split('.')[1];
        
        const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
        
        const payload = JSON.parse(window.atob(base64));

        const currentTime = Math.floor(Date.now() / 1000);
        
        if (payload.exp && payload.exp < currentTime) {
            return null;
        }

        return payload;

    } catch (error) {
        return null;
    }
}

function renderUserLoggedIn(username) {
    const container = document.getElementById("auth-container");

    while(container.firstChild) {
        container.removeChild(container.firstChild);
    }

    const usernameRender = document.createElement("p");
    usernameRender.textContent = username;

    const logoutLink = document.createElement("a");
    logoutLink.textContent = "Выйти";
    logoutLink.setAttribute("href", "/login");
    logoutLink.addEventListener("click", function(){
        localStorage.removeItem("access_token");
    });

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

function renderUserComponents() {
    let payload = getAuthData();
    if (!payload) {
        localStorage.removeItem("access_token");
        renderUserLoggedOut();
        return;
    }
    renderUserLoggedIn(payload.user_name);
}

document.addEventListener('DOMContentLoaded', renderUserComponents);
