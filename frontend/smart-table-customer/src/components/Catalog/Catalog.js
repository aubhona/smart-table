import React, { useEffect, useRef, useState, useCallback } from "react";
import { useOrder } from "../OrderContext/OrderContext";
import { SERVER_URL } from "../../config";
import { handleMultipartResponse } from "../hooks/multipartUtils";
import { useNavigate } from "react-router-dom";
import LoadingScreen from "../LoadingScreen/LoadingScreen";
import { getAuthHeaders } from '../hooks/authHeaders';
import "./Catalog.css";

function Catalog() {
  const { customer_uuid, order_uuid, room_code, setRoomCode, jwt_token } = useOrder();
  const [categories, setCategories] = useState([]);
  const [dishes, setDishes] = useState([]);
  const [images, setImages] = useState({});
  const [counts, setCounts] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const catalogLoaded = useRef(false); 
  const navigate = useNavigate();

  const fetchCartInfo = useCallback(async () => {
    if (!customer_uuid || !order_uuid || !jwt_token) {
      console.error('Недостаточно данных для запроса корзины');
      return;
    }
    
    try {
      const res = await fetch(`${SERVER_URL}/customer/v1/order/catalog/updated-info`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "ngrok-skip-browser-warning": "true",
          ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
        }
      });

      if (!res.ok) throw new Error("Failed to fetch cart info");

      const data = await res.json();
      let menuUpdated = data.menu_updated_info || data.items || [];
      const byId = {};
      menuUpdated.forEach(item => {
        byId[item.id || item.menu_dish_uuid] = item.count;
      });
      setCounts(byId);
    } catch (err) {
      console.error("Error fetching cart info:", err);
    }
  }, [customer_uuid, jwt_token, order_uuid]);

  useEffect(() => {
    if (!customer_uuid || !order_uuid) return;

    fetchCartInfo();

    window.addEventListener("focus", fetchCartInfo);
    return () => window.removeEventListener("focus", fetchCartInfo);
  }, [customer_uuid, order_uuid, jwt_token, fetchCartInfo]);

  useEffect(() => {
    if (!customer_uuid || !order_uuid || !jwt_token) return;

    setLoading(true);

    catalogLoaded.current = false;
    (async () => {
      try {
        const res = await fetch(`${SERVER_URL}/customer/v1/order/catalog`, {
          method: "GET",
          headers: {
            Accept: "multipart/mixed, application/json",
            "Content-Type": "application/json",
            "ngrok-skip-browser-warning": "true",
            ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
          }
        });

        const contentType = res.headers.get('content-type') || '';

        if (contentType.includes('application/json')) {
          const data = await res.json();
          if (data.go_tip_screen) {
            navigate("/tip");
            return;
          }
        } else {
          const { list, categories, imagesMap, room_code: rcode, go_tip_screen  } = await handleMultipartResponse(res, "menu");
          setDishes(list.sort((a, b) => a.name.localeCompare(b.name, "ru")));
          setCategories(categories.sort((a, b) => a.localeCompare(b, "ru")));
          setImages(imagesMap);
          setLoading(false);
          if (rcode) setRoomCode(rcode);
          if (go_tip_screen) {
            navigate("/tip");
            return;
          }
        }
      } catch (e) {
        setError("Ошибка загрузки каталога: " + e.message);
        setLoading(false);
      }
    })();
  }, [customer_uuid, order_uuid, jwt_token, navigate, setRoomCode]);

  const updateQuantity = async (dishId, delta) => {
    if (!customer_uuid || !order_uuid || !jwt_token) {
      console.error('Недостаточно данных для обновления количества');
      return;
    }
    
    try {
      setCounts(prev => ({
        ...prev,
        [dishId]: Math.max((prev[dishId] || 0) + delta, 0)
      }));

      const res = await fetch(`${SERVER_URL}/customer/v1/order/items/draft/count/edit`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
        },
        body: JSON.stringify({
          menu_dish_uuid: dishId,
          count: delta
        })
      });

      if (!res.ok) {
        setCounts(prev => ({
          ...prev,
          [dishId]: (prev[dishId] || 0) - delta
        }));
        throw new Error("Failed to update quantity");
      }

      await fetchCartInfo();
    } catch (err) {
      console.error("Error updating quantity:", err);
    }
  };

  const handleItemClick = (id) => {
    navigate(`/catalog/item/${id}`);
  };

  const totalPrice = dishes.reduce(
    (sum, dish) =>
      sum + (Number(counts[dish.id] || 0) * Number(dish.price || 0)),
    0
  );

  const handleScroll = (cat) => {
    const element = document.getElementById(cat);
    if (element) {
      const offset = 210;
      const elementPosition = element.getBoundingClientRect().top + window.scrollY;
      window.scrollTo({
        top: elementPosition - offset,
        behavior: "smooth"
      });
    }
  };

  const handleGoToCart = () => navigate("/cart");
  const handleGoToUsers = () => navigate("/catalog/users-list");
  const handleCheckout = () => navigate("/checkout");

  if (loading) return <LoadingScreen message="Загрузка..." />;
  if (error)
    return (
      <div style={{ color: "red", padding: "10%", textAlign: "center" }}>{error}</div>
    );

  return (
    <div className="app-container">
      <div className="catalog-container">
        <div className="top-buttons">
          <button className="small-button" onClick={handleGoToUsers}>
            Код комнаты: {room_code}
          </button>
          <button className="small-button" onClick={handleCheckout}>
            Статус заказа
          </button>
        </div>

        <div className="category-tabs">
          <div className="category-scroll">
            {categories.map((cat) => (
              <button key={cat} onClick={() => handleScroll(cat)} className="category-link">
                {cat}
              </button>
            ))}
          </div>
        </div>

        {categories.map((cat) => (
          <div key={cat} id={cat} className="category-section">
            <h2 className="category-title">{cat}</h2>
            <div className="menu-grid">
              {dishes
                .filter((dish) => dish.category === cat)
                .map((dish) => (
                  <div
                    key={dish.id}
                    className="menu-item"
                    onClick={() => handleItemClick(dish.id)}
                  >
                    <div className="dish-img">
                      {images[dish.id] ? (
                        <img src={images[dish.id]} alt={dish.name} />
                      ) : (
                        <span>Нет фото</span>
                      )}
                    </div>
                    <div className="dish-info">
                      <p className="dish-price">
                        <strong>{dish.price} ₽</strong>
                      </p>
                      <p className="dish-name">{dish.name}</p>
                      <p className="dish-weight">{dish.weight} г</p>
                      <p className="dish-calories">{dish.calories} ккал</p>
                    </div>
                    <div className="quantity-controls">
                      {counts[dish.id] && counts[dish.id] > 0 ? (
                        <>
                          <button
                            onClick={(e) => {
                              e.stopPropagation();
                              updateQuantity(dish.id, -1);
                            }}
                          >
                            -
                          </button>
                          <span>
                            <strong>{counts[dish.id]}</strong>
                          </span>
                          <button
                            onClick={(e) => {
                              e.stopPropagation();
                              updateQuantity(dish.id, 1);
                            }}
                          >
                            +
                          </button>
                        </>
                      ) : (
                        <button
                          className="add-button"
                          onClick={(e) => {
                            e.stopPropagation();
                            updateQuantity(dish.id, 1);
                          }}
                        >
                          +
                        </button>
                      )}
                    </div>
                  </div>
                ))}
            </div>
          </div>
        ))}
        <div className="scroll-padding"></div>
        <div className="total-price">
          <p>
            Итого: <strong>{totalPrice} ₽</strong>
          </p>
          <button
            className="checkout-button category-link"
            onClick={handleGoToCart}
            disabled={Object.values(counts).reduce((sum, v) => sum + v, 0) === 0}
          >
            Далее
          </button>
        </div>
      </div>
    </div>
  );
}

export default Catalog;
