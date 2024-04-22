document.addEventListener('DOMContentLoaded', function () {
    const quantityButtons = document.querySelectorAll('.quantity-button');
    const totalPriceElement = document.getElementById('total-price');
    let totalPrice = parseFloat("{{ total_price }}");

    quantityButtons.forEach(button => {
        button.addEventListener('click', function () {
            const action = this.dataset.action;
            const productId = this.dataset.id;
            const quantityElement = this.parentNode.querySelector('.quantity');
            let quantity = parseInt(quantityElement.textContent);

            if (action === 'increase') {
                quantity++;
                totalPrice += parseFloat("{{ product.price }}");
            } else if (action === 'decrease' && quantity > 1) {
                quantity--;
                totalPrice -= parseFloat("{{ product.price }}");
            }

            quantityElement.textContent = quantity;
            totalPriceElement.textContent = "&euro; " + totalPrice.toFixed(2);
        });
    });
});

