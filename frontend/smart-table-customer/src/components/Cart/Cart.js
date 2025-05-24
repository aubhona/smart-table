import React, { useEffect, useState, useCallback } from "react";
import { useOrder } from "../OrderContext/OrderContext";
import { SERVER_URL } from "../../config";
import { handleMultipartResponse } from "../hooks/multipartUtils";
import { useNavigate } from "react-router-dom";
import LoadingScreen from "../LoadingScreen/LoadingScreen";
import "./Cart.css";
import { getAuthHeaders } from '../hooks/authHeaders';

function Cart() {
  const { customer_uuid, order_uuid, jwt_token } = useOrder();
  const [cartItems, setCartItems] = useState([]);
  const [images, setImages] = useState({});
  const [loading, setLoading] = useState(true);
  const [commitLoading, setCommitLoading] = useState(false);
  const [error, setError] = useState("");
  const navigate = useNavigate();

  // Быстрая загрузка корзины без картинок
  const fetchCartInfo = useCallback(async () => {
    setLoading(true);
    try {
      const res = await fetch(`${SERVER_URL}/customer/v1/order/cart-info`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "ngrok-skip-browser-warning": "true",
          ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
        },
      });
      if (!res.ok) throw new Error(`Ошибка загрузки корзины: ${res.status}`);
      const data = await res.json();
      setCartItems((data.items || []).sort((a, b) => a.name.localeCompare(b.name, "ru")));
      setLoading(false);
    } catch (e) {
      setError(e.message || "Не удалось загрузить корзину");
      setLoading(false);
    }
  }, [customer_uuid, jwt_token, order_uuid]);

  // Полная загрузка корзины с картинками
  const fetchCartWithImages = useCallback(async () => {
    try {
      const res = await fetch(`${SERVER_URL}/customer/v1/order/cart`, {
        method: "GET",
        headers: {
          Accept: "multipart/mixed, application/json",
          "Content-Type": "application/json",
          "ngrok-skip-browser-warning": "true",
          ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
        },
      });
      if (!res.ok) throw new Error(`Ошибка загрузки корзины: ${res.status}`);
      const { list, imagesMap } = await handleMultipartResponse(res, "dish_list");
      const items = Array.isArray(list)
        ? (list[0]?.items || list).sort((a, b) => a.name.localeCompare(b.name, "ru"))
        : [];
      setImages(imagesMap || {});
      setCartItems(items);
    } catch (e) {
      // Не блокируем интерфейс, просто не обновляем картинки
      console.error("Ошибка загрузки корзины с картинками:", e);
    }
  }, [customer_uuid, jwt_token, order_uuid]);

  useEffect(() => {
    if (!customer_uuid || !order_uuid) {
      setError("Ошибка: отсутствует ID пользователя или заказа");
      setLoading(false);
      return;
    }
    fetchCartInfo();
    fetchCartWithImages();
  }, [customer_uuid, order_uuid, fetchCartInfo, fetchCartWithImages]);

  useEffect(() => {
    if (!loading && cartItems.length === 0) {
      navigate('/catalog');
    }
  }, [cartItems, loading, navigate]);

  const updateQuantity = async (item, newCount) => {
    const currentCount = item.count;
    const delta = newCount - currentCount;
    if (delta === 0) return;

    try {
      await fetch(`${SERVER_URL}/customer/v1/order/items/draft/count/edit`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
        },
        body: JSON.stringify({
          menu_dish_uuid: item.id || item.menu_dish_uuid,
          count: delta,
          comment: item.comment || undefined,
        }),
      });
      fetchCartWithImages();
    } catch (e) {
      setError(e.message || "Не удалось обновить корзину");
    }
  };

  const handleRemove = (item) => {
    updateQuantity(item, 0);
  };

  const handleItemClick = (item) => {
    navigate(`/catalog/item/${item.id || item.menu_dish_uuid}`, {
      state: {
        count: item.count,
        comment: item.comment,
      },
    });
  };

  const totalPrice = cartItems.reduce(
    (sum, item) => sum + Number(item.price) * Number(item.count),
    0
  );

  const handleCheckout = () => {
    navigate("/checkout");
  };

  const handleOrderCommit = async () => {
    setCommitLoading(true);
    setError("");
    try {
      const res = await fetch(`${SERVER_URL}/customer/v1/order/items/commit`, {
        method: "POST",
        headers: {
          ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
        },
      });
      if (res.status === 204) {
        navigate("/checkout");
      } else {
        let err = "Ошибка при оформлении заказа";
        try {
          const data = await res.json();
          if (data?.error) err = data.error;
        } catch {}
        setError(err);
      }
    } catch (e) {
      setError(e.message || "Ошибка оформления заказа");
    }
    setCommitLoading(false);
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
            <div key={item.id || item.menu_dish_uuid} className="cart-item" onClick={() => handleItemClick(item)}>
              <div className="cart-item-img">
                {images[item.id] ? (
                  <img src={images[item.id]} alt={item.name}/>
                ) : (
                  <div className="image-placeholder" />
                )}
              </div>
              <div className="cart-item-info">
                <div className="cart-item-name">{item.name}</div>
              </div>
              <div className="cart-item-price-with-comment">
                <span>{item.price} ₽</span>
                {item.comment && (
                  <span className="comment-icon" title={item.comment}>🗨️</span>
                )}
              </div>
              <div className="cart-item-quantity" onClick={e => e.stopPropagation()}>
                <button onClick={() => updateQuantity(item, item.count - 1)} className="quantity-button">-</button>
                <span>{item.count}</span>
                <button onClick={() => updateQuantity(item, item.count + 1)} className="quantity-button">+</button>
              </div>
              <div className="cart-item-total">{item.price * item.count} ₽</div>
              <button className="remove-button" onClick={e => { e.stopPropagation(); handleRemove(item); }}>
                Удалить
              </button>
            </div>
          ))
        )}
      </div>
      <div className="cart-total">
        <div>Итого: {totalPrice} ₽</div>
        <div className="cart-actions">
          <button className="make-order-button" onClick={handleOrderCommit} disabled={commitLoading}>
            {commitLoading ? "Оформляем..." : "Заказать"}
          </button>
        </div>
      </div>
    </div>
  );
}

export default Cart;