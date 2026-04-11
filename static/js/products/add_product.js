"use strict"

function showAddForm() {
    const clone = addProductFormTemplate.content.cloneNode(true);
    const container = clone.querySelector('#add-product-form-container');
    const form = clone.querySelector('#add-product-form');

    showAddFormButton.style.display = "none";
    addProductWrapper.appendChild(clone);

    form.addEventListener("submit", addProduct);

    document.getElementById("cancel-add-button").addEventListener("click", () => {
        container.remove();
        showAddFormButton.style.display = "";
    });
}

async function addProduct(event) {
    event.preventDefault();

    const addProductButton = document.getElementById("add-product-button");
    const token = localStorage.getItem('access_token');

    showErrors(addProductButton, []);

    try {
        if (!token) {
            throw {status: 401};
        }

        const form = document.getElementById("add-product-form");
        const formData = new FormData(form);
        const data = Object.fromEntries(formData.entries());

        const numericFields = ['base_weight', 'base_portion', 'calories', 'proteins', 'fats', 'carbohydrates'];
        numericFields.forEach(field => {
                data[field] = parseFloat(data[field]) || 0;
        });

        const response = await fetch('/api/products', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        const product = await response.json();

        if (!response.ok) {
            throw {
                status: response.status,
                message: product.errors
            };
        }

        addProductToList(listContainer, product, productTemplate);

    } catch (error) {
        if (error.status == 401) {
            error.message = ["Пользователь не авторизован"];
            localStorage.removeItem("access_token");
            renderUserLoggedOut();
        }
        if (!error.message) {
            error.message = ["Неизвестная ошибка"];
        }
        showErrors(addProductButton, error.message);
    }
}
