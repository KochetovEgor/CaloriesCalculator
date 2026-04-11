"use strict"

function showAddForm() {
    const clone = addRationTemplate.content.cloneNode(true);
    const container = clone.querySelector('#add-ration-form-container');
    const form = clone.querySelector('#add-ration-form');
    const productsList = clone.querySelector('#ration-products-list');
    const addRowBtn = clone.querySelector('#add-product-row-btn');

    showAddFormButton.style.display = "none";

    addRationWrapper.appendChild(clone);

    addProductRow(productsList);

    addRowBtn.addEventListener('click', () => {
        addProductRow(productsList);
    });

    form.addEventListener("submit", addRation);

    document.getElementById("cancel-add-ration-btn").addEventListener("click", () => {
        container.remove(); 
        showAddFormButton.style.display = "";
    });
}

function addProductRow(container) {
    const clone = productRowTemplate.content.cloneNode(true);

    const row = clone.querySelector('.product-row');
    const removeBtn = clone.querySelector('.remove-product-row-btn');

    removeBtn.addEventListener('click', () => {
        row.remove();
    });

    container.appendChild(clone);
}

async function addRation(event) {
    event.preventDefault();

    const addRationBtn = document.getElementById("add-ration-btn");
    const token = localStorage.getItem('access_token');

    showErrors(addRationBtn, []);

    try {
        if (!token) {
            throw {status: 401};
        }

        const form = event.target;
        const date = form.querySelector('#ration-date-input').value;
        const productRows = form.querySelectorAll('.product-row');

        const products = [];

        productRows.forEach(row => {
            const name = row.querySelector('.product-search-input').value;
            const weight = parseFloat(row.querySelector('.product-weight-input').value);
            const portion = parseFloat(row.querySelector('.product-portion-input').value);

            if (name) {
                products.push({
                    name: name,
                    weight: parseFloat(weight) || 0,
                    portion: parseFloat(portion) || 0
                });
            }
        });

        const data = {date: date, products: products};

        const response = await fetch('/api/rations', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        const ration = await response.json();

        if (!response.ok) {
            throw {
                status: response.status,
                message: ration.errors
            };
        }

        addRationToList(listContainer, ration, rationTemplate);
        document.getElementById("cancel-add-ration-btn").click();
        showAddFormButton.click();

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