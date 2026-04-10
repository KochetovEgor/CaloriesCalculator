"use strict"

function showEditForm(event) {
    const productItem = event.target.closest('.product-item');
    const template = document.getElementById('edit-product-template');
    const editFormClone = template.content.cloneNode(true);

    const currentData = {
        name: productItem.querySelector('.product-name').textContent,
        weight: productItem.querySelector('.val-base-weight').textContent,
        portion: productItem.querySelector('.val-base-portion').textContent,
        calories: productItem.querySelector('.val-calories').textContent,
        proteins: productItem.querySelector('.val-proteins').textContent,
        fats: productItem.querySelector('.val-fats').textContent,
        carbohydrates: productItem.querySelector('.val-carbohydrates').textContent
    };

    const form = editFormClone.querySelector('.edit-product-form');
    form.querySelector('.edit-name').value = currentData.name;
    form.querySelector('.edit-base-weight').value = currentData.weight;
    form.querySelector('.edit-base-portion').value = currentData.portion;
    form.querySelector('.edit-calories').value = currentData.calories;
    form.querySelector('.edit-proteins').value = currentData.proteins;
    form.querySelector('.edit-fats').value = currentData.fats;
    form.querySelector('.edit-carbohydrates').value = currentData.carbohydrates;

    const originalContent = productItem.innerHTML;

    productItem.innerHTML = '';
    productItem.appendChild(editFormClone);

    productItem.querySelector('.cancel-edit-button').addEventListener("click", function (){
        productItem.innerHTML = originalContent;
    });

    productItem.querySelector('.edit-product-form').addEventListener("submit", changeProduct    );
}

async function changeProduct(event) {
    event.preventDefault();

    const saveEditButton = event.target.querySelector(".save-edit-button");
    showErrors(saveEditButton, []);

    const token = localStorage.getItem('access_token');

    try {
        if (!token) {
            throw {status: 401};
        }

        const formData = new FormData(event.target);
        const data = Object.fromEntries(formData.entries());

        const numericFields = ['base_weight', 'base_portion', 'calories', 'proteins', 'fats', 'carbohydrates']
        numericFields.forEach(field => {
            data[field] = parseFloat(data[field]) || 0;
        });

        const response = await fetch('/api/products', {
            method: 'PUT',
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

        loadProducts();

    } catch (error) {
        if (error.status == 401) {
            error.message = ["Пользователь не авторизован"];
        }
        if (!error.message) {
            error.message = ["Неизвестная ошибка"];
        }
        showErrors(saveEditButton, error.message);
    }
}