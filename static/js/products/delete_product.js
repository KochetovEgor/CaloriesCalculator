"use strict"

async function deleteProduct(event) {
    const button = event.target;
    const productItem = button.closest(".product-item");
    const productName = productItem.querySelector(".product-name").textContent;

    showErrors(button, []);

    const token = localStorage.getItem('access_token');

    try {
        if (!token) {
            throw {status: 401};
        }

        const data = {name: productName};

        const response = await fetch("/api/products", {
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

        productItem.remove();

    } catch (error) {
        if (error.status == 401) {
            error.message = ["Пользователь не авторизован"];
        }
        if (!error.message) {
            error.message = ["Неизвестная ошибка"];
        }
        showErrors(button, error.message);
    }
}