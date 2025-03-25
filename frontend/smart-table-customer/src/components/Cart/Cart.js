import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./Cart.css";

const Cart = () => {
  const navigate = useNavigate();
  
  const [cart, setCart] = useState({
    1: { name: "Плов", price: 1200, quantity: 1, calories: 300 },
    2: { name: "Бурурброт", price: 1000, quantity: 2, calories: 250 },
    3: { name: "Шашлыкофф", price: 2000, quantity: 3, calories: 500 },
  });

  const removeFromCart = (id) => {
    const newCart = { ...cart };
    delete newCart[id];
    setCart(newCart);
  };

  const updateQuantity = (id, change) => {
    setCart((prevCart) => {
      const newQuantity = prevCart[id].quantity + change;
      if (newQuantity > 0) {
        return {
          ...prevCart,
          [id]: { ...prevCart[id], quantity: newQuantity },
        };
      }
      return prevCart;
    });
  };

  const clearCart = () => {
    setCart({});
  };

  const handleCheckout = () => {
    // переход к чекауту
  };

  const handleCallWaiter = () => {
    // действия при вызове официанта
  };

  const handleMakeOrder = () => {
    // действия при вызове официанта
  };

  const totalPrice = Object.values(cart).reduce(
    (sum, item) => sum + item.price * item.quantity,
    0
  );

  return (
    <div className="cart-container">
      <div className="cart-header">
        <button className="close-button" onClick={() => navigate(-1)}>назад</button>
        <button className="call-waiter-button" onClick={handleCallWaiter}>
            официант
          </button>
          <button className="checkout-button-cart" onClick={handleCheckout}>
            гоу чекаут
          </button>
          <button className="clear-button" onClick={clearCart}>
          🗑️
            </button>
      </div>

      <div className="cart-items">
        {Object.keys(cart).map((id) => {
          const item = cart[id];
          return (
            <div key={id} className="cart-item">
                <div className="cart-item-img">Фотка</div>
              <div className="cart-item-info">
                <div className="cart-item-name">{item.name}</div>
                <div className="cart-item-price">{item.price} ₽</div>
              </div>
              <div className="cart-item-quantity">
                <button
                  onClick={() => updateQuantity(id, -1)}
                  className="quantity-button"
                >
                  -
                </button>
                <span>{item.quantity}</span>
                <button
                  onClick={() => updateQuantity(id, 1)}
                  className="quantity-button"
                >
                  +
                </button>
              </div>
              <div className="cart-item-total">
                {item.price * item.quantity} ₽
              </div>
              <button
                className="remove-button"
                onClick={() => removeFromCart(id)}
              >
                Remove
              </button>
            </div>
          );
        })}
      </div>

      <div className="cart-total">
        <div>Total: {totalPrice} ₽</div>
        <div className="cart-actions">
          <button className="make-order-button" onClick={handleMakeOrder}>
            Заказать
          </button>
        </div>
      </div>
    </div>
  );
};

export default Cart;
