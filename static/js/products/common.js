"use strict"

function addProductToList(productList, product, template) {
    const clone = template.content.cloneNode(true);

    clone.querySelector('.product-name').textContent = product.name;

    clone.querySelector('.base-weight-info .val-base-weight').textContent = product.base_weight;
    clone.querySelector('.base-portion-info .val-base-portion').textContent = product.base_portion;
    clone.querySelector('.calories-info .val-calories').textContent = product.calories;
    
    clone.querySelector('.val-proteins').textContent = product.proteins;
    clone.querySelector('.val-fats').textContent = product.fats;
    clone.querySelector('.val-carbohydrates').textContent = product.carbohydrates;

    productList.appendChild(clone);
}

const listContainer = document.getElementById('products-list');
const productTemplate = document.getElementById('product-template');

const loadProductsButton = document.getElementById("load-products-button");

const addProductFormTemplate = document.getElementById("add-product-form-template");
const showAddFormButton = document.getElementById("show-add-form-button");
const addProductWrapper = document.getElementById("add-product-wrapper");

document.addEventListener('DOMContentLoaded', loadProducts);

listContainer.addEventListener("click", (event) => {
    const target = event.target;

    if (target.classList.contains('delete-product-button')) {
        deleteProduct(event);
    }

    if (target.classList.contains('edit-product-button')) {
        showEditForm(event);
    }
});

loadProductsButton.addEventListener("click", loadProducts);

showAddFormButton.addEventListener("click", showAddForm);
