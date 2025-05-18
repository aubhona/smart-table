// import React, { useEffect, useState, useRef } from "react";
// import { useOrder } from "../OrderContext/OrderContext";
// import { SERVER_URL } from "../../config";
// import { handleMultipartResponse } from "../hooks/multipartUtils";
// import { useNavigate } from "react-router-dom";
// import LoadingScreen from "../LoadingScreen/LoadingScreen";
// import "./Catalog.css";

// function Catalog() {
//   const { customer_uuid, order_uuid, room_code, setRoomCode } = useOrder();
//   const [categories, setCategories] = useState([]);
//   const [dishes, setDishes] = useState([]);
//   const [images, setImages] = useState({});
//   const [counts, setCounts] = useState({});
//   const [loading, setLoading] = useState(true);
//   const [error, setError] = useState("");
//   const catalogLoaded = useRef(false);  
//   const navigate = useNavigate();

//   const fetchCatalog = async (showLoading = true) => {
//     if (showLoading) setLoading(true);

//     try {
//       const res = await fetch(`${SERVER_URL}/customer/v1/order/catalog`, {
//         method: "GET",
//         headers: {
//           Accept: "multipart/mixed, application/json",
//           "Content-Type": "application/json",
//           "ngrok-skip-browser-warning": "true",
//           "Customer-UUID": customer_uuid,
//           "Order-UUID": order_uuid,
//           "JWT-Token": "bla-bla-bla"
//         }
//       });

//       if (!res.ok) throw new Error("Failed to fetch catalog");
//       const { list, categories, imagesMap, room_code: rcode, counts: serverCounts } = await handleMultipartResponse(res, "menu");

//       setDishes(list.sort((a, b) => a.name.localeCompare(b.name, "ru")));
//       setCategories(categories.sort((a, b) => a.localeCompare(b, "ru")));
//       setImages(imagesMap);
//       if (rcode) setRoomCode(rcode);
//       if (serverCounts) setCounts(serverCounts);
//       if (showLoading) setLoading(false);
//       catalogLoaded.current = true;
//     } catch (e) {
//       setError("Ошибка загрузки каталога: " + e.message);
//       if (showLoading) setLoading(false);
//     }
//   };

//   const fetchCartInfo = async () => {
//     try {
//       const res = await fetch(`${SERVER_URL}/customer/v1/order/catalog/updated-info`, {
//         method: "GET",
//         headers: {
//           "Content-Type": "application/json",
//           "ngrok-skip-browser-warning": "true",
//           "Customer-UUID": customer_uuid,
//           "Order-UUID": order_uuid,
//           "JWT-Token": "bla-bla-bla"
//         }
//       });

//       if (!res.ok) throw new Error("Failed to fetch cart info");
//       const data = await res.json();
//       let menuUpdated = data.menu_updated_info || data.items || [];
//       const byId = {};
//       menuUpdated.forEach(item => {
//         byId[item.id || item.menu_dish_uuid] = item.count;
//       });
//       setCounts(byId);
//     } catch (err) {
//       console.error("Error fetching cart info:", err);
//     }
//   };

//   useEffect(() => {
//     if (!customer_uuid || !order_uuid) {
//       console.error("Missing customer_uuid or order_uuid:", { customer_uuid, order_uuid });
//       return;
//     }
//     if (!catalogLoaded.current) {
//       fetchCatalog();
//     } else {
//       fetchCartInfo();
//     }
//   }, [customer_uuid, order_uuid]);

//   const updateQuantity = async (dishId, delta) => {
//     try {
//       setCounts(prev => ({
//         ...prev,
//         [dishId]: Math.max((prev[dishId] || 0) + delta, 0)
//       }));

//       const res = await fetch(`${SERVER_URL}/customer/v1/order/items/draft/count/edit`, {
//         method: "POST",
//         headers: {
//           "Content-Type": "application/json",
//           "Customer-UUID": customer_uuid,
//           "Order-UUID": order_uuid,
//           "JWT-Token": "bla-bla-bla"
//         },
//         body: JSON.stringify({
//           menu_dish_uuid: dishId,
//           count: delta
//         })
//       });

//       if (!res.ok) {
//         setCounts(prev => ({
//           ...prev,
//           [dishId]: (prev[dishId] || 0) - delta
//         }));
//         throw new Error("Failed to update quantity");
//       }
//       await fetchCartInfo();
//     } catch (err) {
//       console.error("Error updating quantity:", err);
//     }
//   };

//   const handleItemClick = (id) => {
//     navigate(`/catalog/item/${id}`);
//   };

//   const totalPrice = dishes.reduce(
//     (sum, dish) =>
//       sum + (Number(counts[dish.id] || 0) * Number(dish.price || 0)),
//     0
//   );

//   const handleScroll = (cat) => {
//     const element = document.getElementById(cat);
//     if (element) {
//       const offset = 210;
//       const elementPosition = element.getBoundingClientRect().top + window.scrollY;
//       window.scrollTo({
//         top: elementPosition - offset,
//         behavior: "smooth"
//       });
//     }
//   };

//   const handleGoToCart = () => navigate("/cart");
//   const handleGoToUsers = () => navigate("/catalog/users-list");
//   const handleCheckout = () => navigate("/checkout");

//   if (loading) return <LoadingScreen message="Загрузка..." />;
//   if (error)
//     return (
//       <div style={{ color: "red", padding: "10%", textAlign: "center" }}>{error}</div>
//     );

//   return (
//     <div className="app-container">
//       <div className="catalog-container">
//         <div className="top-buttons">
//           <button className="small-button" onClick={handleGoToUsers}>
//             Код комнаты: {room_code}
//           </button>
//           <button className="small-button" onClick={handleCheckout}>
//             Статус заказа
//           </button>
//         </div>

//         <div className="category-tabs">
//           <div className="category-scroll">
//             {categories.map((cat) => (
//               <button key={cat} onClick={() => handleScroll(cat)} className="category-link">
//                 {cat}
//               </button>
//             ))}
//           </div>
//         </div>

//         {categories.map((cat) => (
//           <div key={cat} id={cat} className="category-section">
//             <h2 className="category-title">{cat}</h2>
//             <div className="menu-grid">
//               {dishes
//                 .filter((dish) => dish.category === cat)
//                 .map((dish) => (
//                   <div
//                     key={dish.id}
//                     className="menu-item"
//                     onClick={() => handleItemClick(dish.id)}
//                   >
//                     <div className="dish-img">
//                       {images[dish.id] ? (
//                         <img src={images[dish.id]} alt={dish.name} />
//                       ) : (
//                         <span>Нет фото</span>
//                       )}
//                     </div>
//                     <div className="dish-info">
//                       <p className="dish-price">
//                         <strong>{dish.price} ₽</strong>
//                       </p>
//                       <p className="dish-name">{dish.name}</p>
//                       <p className="dish-weight">{dish.weight} г</p>
//                       <p className="dish-calories">{dish.calories} ккал</p>
//                     </div>
//                     <div className="quantity-controls">
//                       {counts[dish.id] && counts[dish.id] > 0 ? (
//                         <>
//                           <button
//                             onClick={(e) => {
//                               e.stopPropagation();
//                               updateQuantity(dish.id, -1);
//                             }}
//                           >
//                             -
//                           </button>
//                           <span>
//                             <strong>{counts[dish.id]}</strong>
//                           </span>
//                           <button
//                             onClick={(e) => {
//                               e.stopPropagation();
//                               updateQuantity(dish.id, 1);
//                             }}
//                           >
//                             +
//                           </button>
//                         </>
//                       ) : (
//                         <button
//                           className="add-button"
//                           onClick={(e) => {
//                             e.stopPropagation();
//                             updateQuantity(dish.id, 1);
//                           }}
//                         >
//                           +
//                         </button>
//                       )}
//                     </div>
//                   </div>
//                 ))}
//             </div>
//           </div>
//         ))}
//         <div className="scroll-padding"></div>
//         <div className="total-price">
//           <p>
//             Итого: <strong>{totalPrice} ₽</strong>
//           </p>
//           <button className="checkout-button" onClick={handleGoToCart}>
//             Далее
//           </button>
//         </div>
//       </div>
//     </div>
//   );
// }

// export default Catalog;
import React, { useEffect, useRef, useState } from "react";
import { useOrder } from "../OrderContext/OrderContext";
import { SERVER_URL } from "../../config";
import { handleMultipartResponse } from "../hooks/multipartUtils";
import { useNavigate } from "react-router-dom";
import LoadingScreen from "../LoadingScreen/LoadingScreen";
import "./Catalog.css";

function Catalog() {
  const { customer_uuid, order_uuid, room_code, setRoomCode } = useOrder();
  const [categories, setCategories] = useState([]);
  const [dishes, setDishes] = useState([]);
  const [images, setImages] = useState({});
  const [counts, setCounts] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const catalogLoaded = useRef(false); // Флаг для контроля первого захода
  const navigate = useNavigate();

  // Первый раз (при новом заказе) — полный fetch каталога
  useEffect(() => {
    if (!customer_uuid || !order_uuid) return;

    // При смене заказа — сбрасываем состояние
    catalogLoaded.current = false;
    setLoading(true);
    (async () => {
      try {
        const res = await fetch(`${SERVER_URL}/customer/v1/order/catalog`, {
          method: "GET",
          headers: {
            Accept: "multipart/mixed, application/json",
            "Content-Type": "application/json",
            "ngrok-skip-browser-warning": "true",
            "Customer-UUID": customer_uuid,
            "Order-UUID": order_uuid,
            "JWT-Token": "bla-bla-bla"
          }
        });

        if (!res.ok) throw new Error("Failed to fetch catalog");
        const { list, categories, imagesMap, room_code: rcode, counts: serverCounts } = await handleMultipartResponse(res, "menu");

        setDishes(list.sort((a, b) => a.name.localeCompare(b.name, "ru")));
        setCategories(categories.sort((a, b) => a.localeCompare(b, "ru")));
        setImages(imagesMap);
        if (rcode) setRoomCode(rcode);
        setCounts(serverCounts || {});
        setLoading(false);
        catalogLoaded.current = true;
      } catch (e) {
        setError("Ошибка загрузки каталога: " + e.message);
        setLoading(false);
      }
    })();
    // eslint-disable-next-line
  }, [customer_uuid, order_uuid]);

  // Каждый раз при МОНТИРОВАНИИ (или возврате на каталог) — обновляем только counts
  useEffect(() => {
    if (!catalogLoaded.current) return;
    if (!customer_uuid || !order_uuid) return;
    // Обновляем только количество блюд
    const fetchCartInfo = async () => {
      try {
        const res = await fetch(`${SERVER_URL}/customer/v1/order/catalog/updated-info`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            "ngrok-skip-browser-warning": "true",
            "Customer-UUID": customer_uuid,
            "Order-UUID": order_uuid,
            "JWT-Token": "bla-bla-bla"
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
    };

    // Обновляем при каждом заходе на каталог
    fetchCartInfo();

    // ХАК: обновляем при возврате с других страниц
    window.addEventListener("focus", fetchCartInfo);
    return () => window.removeEventListener("focus", fetchCartInfo);
    // eslint-disable-next-line
  }, [catalogLoaded.current, customer_uuid, order_uuid]);

  // При любом изменении количества — тоже обновляем только counts
  const updateQuantity = async (dishId, delta) => {
    try {
      setCounts(prev => ({
        ...prev,
        [dishId]: Math.max((prev[dishId] || 0) + delta, 0)
      }));

      const res = await fetch(`${SERVER_URL}/customer/v1/order/items/draft/count/edit`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Customer-UUID": customer_uuid,
          "Order-UUID": order_uuid,
          "JWT-Token": "bla-bla-bla"
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
      // После плюса/минуса обновляем только counts
      // можно await-ить, можно не await-ить — не критично
      // Если хочешь моментальное обновление — просто вызови fetchCartInfo();
      // но если "скачет" UI, можно debounce или показывать loader на кнопке
      const fetchCartInfo = async () => {
        try {
          const res = await fetch(`${SERVER_URL}/customer/v1/order/catalog/updated-info`, {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
              "ngrok-skip-browser-warning": "true",
              "Customer-UUID": customer_uuid,
              "Order-UUID": order_uuid,
              "JWT-Token": "bla-bla-bla"
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
      };
      fetchCartInfo();
    } catch (err) {
      console.error("Error updating quantity:", err);
    }
  };

  // --- остальной код рендера (см. твой же шаблон) ---
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
          <button className="checkout-button" onClick={handleGoToCart}>
            Далее
          </button>
        </div>
      </div>
    </div>
  );
}

export default Catalog;
