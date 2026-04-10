"use strict"

function renderUsername() {
    let payload = getAuthData();
    if (!payload) {
        localStorage.removeItem("access_token");
        renderUserLoggedOut();
        return;
    }
    renderUserLoggedIn(payload.user_name);
}

renderUsername();