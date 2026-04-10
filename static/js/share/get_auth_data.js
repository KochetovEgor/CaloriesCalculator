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