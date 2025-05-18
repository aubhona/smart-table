import React, { useEffect, useState } from "react";
import { useOrder } from "../OrderContext/OrderContext";
import { SERVER_URL } from "../../config";
import { handleMultipartResponse } from "../hooks/multipartUtils";
import { useNavigate } from "react-router-dom";
import LoadingScreen from "../LoadingScreen/LoadingScreen";
import "./Cart.css";

function Cart() {
  const { customer_uuid, order_uuid } = useOrder();
  const [cartItems, setCartItems] = useState([]);
  const [images, setImages] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const JWT_TOKEN = "bla-bla-bla"; 

  useEffect(() => {
    if (!customer_uuid || !order_uuid) {
      setError("Ошибка: отсутствует ID пользователя или заказа");
      setLoading(false);
      return;
    }

    setLoading(true);
    fetch(`${SERVER_URL}/customer/v1/order/cart`, {
      method: "GET",
      headers: {
        Accept: "multipart/mixed, application/json",
        "Content-Type": "application/json",
        "ngrok-skip-browser-warning": "true",
        "Customer-UUID": customer_uuid,
        "Order-UUID": order_uuid,
        "JWT-Token": JWT_TOKEN
      },
    })
      .then(async (res) => {
        alert(res);
        if (!res.ok) {
          throw new Error(`Ошибка загрузки корзины: ${res.status}`);
        }
        const { list, imagesMap } = await handleMultipartResponse(res, "menu");
        setCartItems(list); 
        setImages(imagesMap || {});
        setLoading(false);
      })
      .catch((e) => {
        console.error("Ошибка загрузки корзины:", e);
        setError(e.message || "Не удалось загрузить корзину");
        setLoading(false);
      });
  }, [customer_uuid, order_uuid]);

  const updateQuantity = async (dishId, delta) => {
    try {
      const res = await fetch(`${SERVER_URL}/customer/v1/order/items/draft/count/edit`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Customer-UUID": customer_uuid,
          "Order-UUID": order_uuid,
          "JWT-Token": JWT_TOKEN
        },
        body: JSON.stringify({
          menu_dish_uuid: dishId,
          count: delta
        })
      });

      if (!res.ok) {
        throw new Error("Не удалось обновить количество");
      }

      const cartRes = await fetch(`${SERVER_URL}/customer/v1/order/cart`, {
        method: "GET",
        headers: {
          Accept: "multipart/mixed, application/json",
          "ngrok-skip-browser-warning": "true",
          "Customer-UUID": customer_uuid,
          "Order-UUID": order_uuid,
          "JWT-Token": JWT_TOKEN
        },
      });

      if (!cartRes.ok) {
        throw new Error("Не удалось обновить корзину");
      }

      const { list, imagesMap } = await handleMultipartResponse(cartRes, "dish_list");
      setCartItems(list);
      setImages(imagesMap || {});
    } catch (e) {
      console.error("Ошибка обновления количества:", e);
      setError(e.message || "Не удалось обновить корзину");
    }
  };

  const handleRemove = (dishId) => updateQuantity(dishId, -10000);

  const totalPrice = cartItems.reduce(
    (sum, item) => sum + Number(item.price) * Number(item.count),
    0
  );

  const handleCheckout = () => {
    navigate("/checkout");
  };

  if (loading) return <LoadingScreen message="Загрузка корзины..." />;
  if (error) return (
    <div className="cart-error">
      <div className="error-message">{error}</div>
      <button className="back-button" onClick={() => navigate(-1)}>Назад</button>
    </div>
  );

  return (
    <div className="cart-container">
      <div className="cart-header">
        <button className="close-button" onClick={() => navigate(-1)}>Назад</button>
        <button className="checkout-button-cart" onClick={handleCheckout}>Статус заказа</button>
      </div>
      <div className="cart-items">
        {cartItems.length === 0 ? (
          <p>Корзина пуста</p>
        ) : (
          cartItems.map((item) => (
            <div key={item.menu_dish_uuid} className="cart-item">
              <div className="cart-item-img">
                {images[item.menu_dish_uuid] ? (
                  <img src={images[item.menu_dish_uuid]} alt={item.name}/>
                ) : (
                  <span>Нет фото</span>
                )}
              </div>
              <div className="cart-item-info">
                <div className="cart-item-name">{item.name}</div>
                <div className="cart-item-price">{item.price} ₽</div>
              </div>
              <div className="cart-item-quantity">
                <button onClick={() => updateQuantity(item.menu_dish_uuid, -1)} className="quantity-button">-</button>
                <span>{item.count}</span>
                <button onClick={() => updateQuantity(item.menu_dish_uuid, 1)} className="quantity-button">+</button>
              </div>
              <div className="cart-item-total">{item.price * item.count} ₽</div>
              <button className="remove-button" onClick={() => handleRemove(item.menu_dish_uuid)}>
                Удалить
              </button>
            </div>
          ))
        )}
      </div>
      <div className="cart-total">
        <div>Итого: {totalPrice} ₽</div>
        <div className="cart-actions">
          <button className="make-order-button" onClick={handleCheckout}>
            Заказать
          </button>
        </div>
      </div>
    </div>
  );
}

export default Cart;
