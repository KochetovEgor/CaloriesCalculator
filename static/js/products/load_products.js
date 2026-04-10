"use strict"

async function loadProducts() {
    showErrors(loadProductsButton, []);

    const token = localStorage.getItem('access_token');

    try {
        if (!token) {
            throw {status: 401};
        }

        const response = await fetch('/api/products', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        });

        const products = await response.json();

        if (!response.ok) {
            throw {
                status: response.status,
                message: products.errors
            };
        }

        listContainer.innerHTML = '';

        products.forEach(product => addProductToList(listContainer, product, productTemplate));

    } catch (error) {
        if (error.status == 401) {
            error.message = ["Пользователь не авторизован"];
        }
        if (!error.message) {
            error.message = ["Неизвестная ошибка"];
        }
        showErrors(loadProductsButton, error.message);
    }
}
