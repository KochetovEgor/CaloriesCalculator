"use strict"

function addRationToList(rationList, ration, template) {
    const clone = template.content.cloneNode(true);

    clone.querySelector('.ration-date').textContent = ration.date;

    clone.querySelector('.calories-info .val-calories').textContent = ration.calories;
    
    clone.querySelector('.val-proteins').textContent = ration.proteins;
    clone.querySelector('.val-fats').textContent = ration.fats;
    clone.querySelector('.val-carbohydrates').textContent = ration.carbohydrates;

    rationList.prepend(clone);
}

async function loadProducts() {
    showErrors(showAddFormButton, []);

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

        window.allProducts = products;

        populateFoodDatalist(window.allProducts);

    } catch (error) {
        if (error.status == 401) {
            error.message = ["Пользователь не авторизован"];
            localStorage.removeItem("access_token");
            renderUserLoggedOut();
        }
        if (!error.message) {
            error.message = ["Неизвестная ошибка"];
        }
        showErrors(showAddFormButton, error.message);
    }
}

function populateFoodDatalist(products) {
    const datalist = document.getElementById('products-datalist');
    if (!datalist) return;

    datalist.innerHTML = '';

    products.forEach(product => {
        const option = document.createElement('option');

        option.value = product.name; 
        // Храним ID в data-атрибуте, чтобы потом найти его при сохранении
        option.dataset.id = product.id; 
        
        datalist.appendChild(option);
    });
}

const listContainer = document.getElementById('rations-list');
const rationTemplate = document.getElementById('ration-template');
const addRationTemplate = document.getElementById("add-ration-template");
const productRowTemplate = document.getElementById("product-row-template");

const loadRationsButton = document.getElementById("load-rations-button");
const showAddFormButton = document.getElementById("show-add-form-button");
const addRationWrapper = document.getElementById('add-ration-wrapper');

document.addEventListener('DOMContentLoaded', loadRations);
document.addEventListener('DOMContentLoaded', loadProducts);

listContainer.addEventListener("click", (event) => {
    const target = event.target;

    if (target.classList.contains('delete-ration-button')) {
        deleteRation(event);
    }
});

loadRationsButton.addEventListener("click", loadRations);
showAddFormButton.addEventListener("click", showAddForm);