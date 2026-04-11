"use strict"

async function deleteRation(event) {
    const button = event.target;
    const rationItem = button.closest(".ration-item");
    const rationDate = rationItem.querySelector(".ration-date").textContent;

    showErrors(button, []);

    const token = localStorage.getItem('access_token');

    try {
        if (!token) {
            throw {status: 401};
        }

        const data = {date: rationDate};

        const response = await fetch("/api/rations", {
            method: "DELETE",
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (!response.ok) {
            const deleteResp = await response.json();
            throw {
                status: response.status,
                message: deleteResp.errors
            };
        }

        rationItem.remove();

    } catch (error) {
        if (error.status == 401) {
            error.message = ["Пользователь не авторизован"];
            localStorage.removeItem("access_token");
            renderUserLoggedOut();
        }
        if (!error.message) {
            error.message = ["Неизвестная ошибка"];
        }
        showErrors(button, error.message);
    }
}