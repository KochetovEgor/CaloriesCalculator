"use strict"

async function loadRations() {
    showErrors(loadRationsButton, []);

    const token = localStorage.getItem('access_token');

    try {
        if (!token) {
            throw {status: 401};
        }

        const response = await fetch('/api/rations', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        });

        const rations = await response.json();

        if (!response.ok) {
            throw {
                status: response.status,
                message: rations.errors
            };
        }

        listContainer.innerHTML = '';

        rations.forEach(ration => addRationToList(listContainer, ration, rationTemplate));

    } catch (error) {
        if (error.status == 401) {
            error.message = ["Пользователь не авторизован"];
            localStorage.removeItem("access_token");
            renderUserLoggedOut();
        }
        if (!error.message) {
            error.message = ["Неизвестная ошибка"];
        }
        showErrors(loadRationsButton, error.message);
    }
}
