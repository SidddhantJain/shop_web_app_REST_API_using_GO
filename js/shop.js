// Example product data (12 sample products)
const products = [
    { id: 1, name: "Product 1", description: "Description for product 1", price: 29.99, image_url: "/images/product1.jpg" },
    { id: 2, name: "Product 2", description: "Description for product 2", price: 49.99, image_url: "/images/product2.jpg" },
    { id: 3, name: "Product 3", description: "Description for product 3", price: 19.99, image_url: "/images/product3.jpg" },
    { id: 4, name: "Product 4", description: "Description for product 4", price: 99.99, image_url: "/images/product4.jpg" },
    { id: 5, name: "Product 5", description: "Description for product 5", price: 24.99, image_url: "/images/product5.jpg" },
    { id: 6, name: "Product 6", description: "Description for product 6", price: 59.99, image_url: "/images/product6.jpg" },
    { id: 7, name: "Product 7", description: "Description for product 7", price: 79.99, image_url: "/images/product7.jpg" },
    { id: 8, name: "Product 8", description: "Description for product 8", price: 14.99, image_url: "/images/product8.jpg" },
    { id: 9, name: "Product 9", description: "Description for product 9", price: 39.99, image_url: "/images/product9.jpg" },
    { id: 10, name: "Product 10", description: "Description for product 10", price: 89.99, image_url: "/images/product10.jpg" },
    { id: 11, name: "Product 11", description: "Description for product 11", price: 22.99, image_url: "/images/product11.jpg" },
    { id: 12, name: "Product 12", description: "Description for product 12", price: 32.99, image_url: "/images/product12.jpg" }
];

// Function to dynamically create the product cards
function loadProducts() {
    const productList = document.getElementById('product-list');
    
    products.forEach(product => {
        const productCard = document.createElement('div');
        productCard.classList.add('col-md-3', 'product-card');

        productCard.innerHTML = `
            <img src="${product.image_url}" alt="${product.name}" class="product-img">
            <div class="product-details">
                <h5 class="product-title">${product.name}</h5>
                <p>${product.description}</p>
                <p class="product-price">$${product.price.toFixed(2)}</p>
                <button class="add-to-cart" data-id="${product.id}">Add to Cart</button>
            </div>
        `;

        productList.appendChild(productCard);
    });
}

// Function to handle the add-to-cart action
function addToCart(productId) {
    // Normally, you'd send an API request to the backend to update the cart
    // Here, we're just simulating that functionality
    alert(`Product ${productId} added to cart!`);
}

// Event listeners for the Add to Cart buttons
document.addEventListener('DOMContentLoaded', function() {
    loadProducts();

    document.getElementById('product-list').addEventListener('click', function(event) {
        if (event.target.classList.contains('add-to-cart')) {
            const productId = event.target.getAttribute('data-id');
            addToCart(productId);
        }
    });
});
