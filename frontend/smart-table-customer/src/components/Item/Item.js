import React, { useEffect, useState } from "react";
import { useParams, useNavigate, useLocation } from "react-router-dom";
import { useOrder } from "../OrderContext/OrderContext";
import { SERVER_URL } from "../../config";
import { handleMultipartResponse } from "../hooks/multipartUtils";
import LoadingScreen from "../LoadingScreen/LoadingScreen";
import { getAuthHeaders } from '../hooks/authHeaders';
import "./Item.css";

function RubleIcon({style}) {
  return (
    <svg
      style={{width:'1em',height:'1em',verticalAlign:'middle',...style}}
      viewBox="0 0 16 16"
      fill="currentColor"
    >
      <path d="M4 2.5A.5.5 0 0 1 4.5 2h5a3.5 3.5 0 0 1 0 7H5v2h4.5a.5.5 0 0 1 0 1H5v1.5a.5.5 0 0 1-1 0V12H3.5a.5.5 0 0 1 0-1H4v-1H3.5a.5.5 0 0 1 0-1H4v-7zm1 1v5h4.5a2.5 2.5 0 0 0 0-5H5z"/>
    </svg>
  );
}

function Item() {
  const location = useLocation();
  const initialCount = location.state?.count || 1;
  const initialComment = location.state?.comment || "";
  const originalComment = location.state?.comment ?? "";

  const { id } = useParams();
  const { customer_uuid, order_uuid, jwt_token } = useOrder();
  const navigate = useNavigate();

  const [dish, setDish] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [quantity, setQuantity] = useState(initialCount);
  const [comment, setComment] = useState(initialComment);
  const [success, setSuccess] = useState(false);
  
  const isEdit = location.state && location.state.count;

  useEffect(() => {
    if (!customer_uuid || !order_uuid || !id) return;
    setLoading(true);

    fetch(`${SERVER_URL}/customer/v1/order/item/info/state`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "ngrok-skip-browser-warning": "true",
        ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
      },
      body: JSON.stringify({ dish_uuid: id }),
    })
      .then(async (res) => {
        if (!res.ok) throw new Error("Не удалось загрузить блюдо");
        const data = await res.json();
        setDish({ ...data, img: null });
        setLoading(false);

        fetch(`${SERVER_URL}/customer/v1/order/item/state`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "ngrok-skip-browser-warning": "true",
            ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
            "Accept": "multipart/mixed, application/json",
          },
          body: JSON.stringify({ dish_uuid: id }),
        })
          .then(async (res) => {
            if (!res.ok) return;
            const { list, imagesMap } = await handleMultipartResponse(res, "item");
            const dishData = Array.isArray(list) ? list[0] : list;
            let img = null;
            if(imagesMap && imagesMap[id]) {
              img = imagesMap[id];
            }
            if(!img && imagesMap) {
              const key = Object.keys(imagesMap);
              if(key.length === 1) {
                img = imagesMap[key[0]];
              }
            }
            setDish((prev) => ({ ...prev, ...dishData, img: img || dishData?.img }));
          });
      })
      .catch((e) => {
        setError(e.message || "Ошибка при загрузке блюда");
        setLoading(false);
      });
  }, [customer_uuid, order_uuid, id, initialComment, jwt_token]);

  const handleAddOrSave = async () => {
    try {
      if (isEdit) {
        const diff = quantity - initialCount;
        let delta;
        if (quantity < initialCount) {
          delta = -Math.abs(diff);
        } else {
          delta = diff;
        }
        if (comment !== originalComment) {
          await fetch(`${SERVER_URL}/customer/v1/order/items/draft/count/edit`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              "ngrok-skip-browser-warning": "true",
              "Customer-UUID": customer_uuid,
              "Order-UUID": order_uuid,
              "JWT-Token": jwt_token,
            },
            body: JSON.stringify({
              menu_dish_uuid: id,
              count: 0,
              comment: originalComment || undefined,
            }),
          });
          await fetch(`${SERVER_URL}/customer/v1/order/items/draft/count/edit`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              "ngrok-skip-browser-warning": "true",
              "Customer-UUID": customer_uuid,
              "Order-UUID": order_uuid,
              "JWT-Token": jwt_token,
            },
            body: JSON.stringify({
              menu_dish_uuid: id,
              count: quantity,
              comment: comment || undefined,
            }),
          });
        } else if (delta !== 0){
          await fetch(`${SERVER_URL}/customer/v1/order/items/draft/count/edit`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              "ngrok-skip-browser-warning": "true",
              "Customer-UUID": customer_uuid,
              "Order-UUID": order_uuid,
              "JWT-Token": jwt_token,
            },
            body: JSON.stringify({
              menu_dish_uuid: id,
              count: delta,
              comment: comment || undefined,
            }),
          });
        }
      } else {
        await fetch(`${SERVER_URL}/customer/v1/order/items/draft/count/edit`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "ngrok-skip-browser-warning": "true",
            "Customer-UUID": customer_uuid,
            "Order-UUID": order_uuid,
            "JWT-Token": jwt_token,
          },
          body: JSON.stringify({
            menu_dish_uuid: id,
            count: quantity,
            comment: comment || undefined,
          }),
        });
      }

      setSuccess(true);
      setTimeout(() => {
        navigate(-1);
      }, 500);
    } catch (e) {
      setError(e.message || "Ошибка сохранения/добавления блюда");
    }
  };

  const updateQuantity = (change) => {
    setQuantity((prev) => Math.max(1, prev + change));
    setSuccess(false);
  };

  if (loading) return <LoadingScreen message="Загрузка блюда..." />;
  if (error)
    return (
      <div className="item-container">
        <div className="top-bar">
          <button className="top-button" onClick={() => navigate(-1)}>Назад</button>
        </div>
        <div className="item-error">{error}</div>
      </div>
    );
  if (!dish) return null;

  return (
    <div className="item-container">
      <div className="top-bar">
        <button className="top-button" onClick={() => navigate(-1)}>Назад</button>
      </div>

      <div className="dish-image-item">
        {dish.img ? (
          <img src={dish.img} alt={dish.name} />
        ) : (
          <div className="shimmer shimmer-rect" style={{width:'100%',height:'100%'}} />
        )}
      </div>

      <div className="dish-info-item">
        <p className="description-item">{dish.description || "Описание не указано"}</p>
        <textarea
          placeholder="Комментарий к заказу"
          value={comment}
          onChange={(e) => setComment(e.target.value)}
        />
      </div>

      <div className="item-footer">
        <div className="item-summary">
          <div className="dish-main-info">
            <div className="dish-name-item">{dish.name}</div>
            <div className="dish-metrics">
              <div className="weight-calories-item">{dish.weight} г, {dish.calories} ккал</div>
            </div>
          </div>
          <div className="dish-price-item">{dish.price}<RubleIcon /></div>
        </div>
        <div className="item-actions">
          <div className="quantity-controls-item">
            <button onClick={() => updateQuantity(-1)}>-</button>
            <span><strong>{quantity}</strong></span>
            <button onClick={() => updateQuantity(1)}>+</button>
          </div>
          <button
            className="add-button"
            onClick={handleAddOrSave}
            disabled={quantity <= 0}
          >
            {isEdit ? "Сохранить" : "Добавить"}
          </button>
        </div>
        {success && (
          <div className="success-msg">
            Блюдо {isEdit ? "обновлено" : "добавлено"} в корзину!
          </div>
        )}
      </div>
    </div>
  );
}

export default Item;
