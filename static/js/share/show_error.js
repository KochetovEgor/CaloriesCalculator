"use strict"

// showErrors renders 'error-template' with given errors over given inputElement.
function showErrors(inputElement, errors) {
    const card = inputElement.parentElement;

    const existingContainer = card.parentElement.querySelector('.field-errors-container');
    
    if (existingContainer) {
        existingContainer.remove();
    }

    if (!errors || errors.length === 0) return;

    const template = document.getElementById('error-template');
    const clone = template.content.cloneNode(true);
    const list = clone.querySelector('.error-list');

    errors.forEach(text => {
        const li = document.createElement('li');
        li.textContent = text;
        list.appendChild(li);
    });

    card.after(clone);
}
