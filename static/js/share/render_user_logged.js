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

    container.innerHTML = "";

    const userWrapper = document.createElement("div");
    userWrapper.className = "user-profile-wrapper";

    const usernameSpan = document.createElement("span");
    usernameSpan.className = "user-name-label";
    usernameSpan.textContent = username;

    const logoutLink = document.createElement("a");
    logoutLink.className = "nav-item logout-button";
    logoutLink.textContent = "Выйти";
    logoutLink.setAttribute("href", "/login");

    logoutLink.addEventListener("click", function() {
        localStorage.removeItem("access_token");
    });

    userWrapper.appendChild(usernameSpan);
    userWrapper.appendChild(logoutLink);
    container.appendChild(userWrapper);
}

function renderUserLoggedOut() {
    const container = document.getElementById("auth-container");

    container.innerHTML = "";

    const userWrapper = document.createElement("div");
    userWrapper.className = "user-profile-wrapper";

    const logoutLink = document.createElement("a");
    logoutLink.className = "nav-item logout-button";
    logoutLink.textContent = "Войти";
    logoutLink.setAttribute("href", "/login");

    userWrapper.appendChild(logoutLink);
    container.appendChild(userWrapper);
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
