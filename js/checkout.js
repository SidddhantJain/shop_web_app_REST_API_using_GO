document.getElementById('checkout-form').addEventListener('submit', function(event) {
    event.preventDefault();
    
    let formData = new FormData(this);
    let data = {
        address: formData.get('address'),
        credit_card: formData.get('credit_card')
    };

    fetch('/api/checkout', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(result => {
        if (result.success) {
            alert('Payment successful! Your order is on the way.');
            window.location.href = "profile.html";
        } else {
            alert('Payment failed. Please try again.');
        }
    });
});
