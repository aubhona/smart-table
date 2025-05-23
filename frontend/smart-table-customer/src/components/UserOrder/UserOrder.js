import React, { useEffect, useState, useRef } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { useOrder } from "../OrderContext/OrderContext";
import LoadingScreen from "../LoadingScreen/LoadingScreen";
import { SERVER_URL } from "../../config";
import { getAuthHeaders } from '../hooks/authHeaders';
import "./UserOrder.css";

const UserOrder = () => {
  const navigate = useNavigate();
  const { userLogin } = useParams();
  const { customer_uuid, order_uuid, jwt_token } = useOrder();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [userData, setUserData] = useState(null);
  const [userOrderItems, setUserOrderItems] = useState([]);

  const pollingIntervalRef = useRef(null);

  const statusColorsMap = {
    new: "white",
    accepted: "blue",
    cooking: "yellow",
    cooked: "green",
    served: "cyan",
    payment_waiting: "gray",
    paid: "darkgreen",
    canceled_by_service: "red",
    canceled_by_client: "red",
  };
  const renderStatus = (status) => {
    const color = statusColorsMap[status] || "white";
    return (
      <span
        className={`status-circle ${color}`}
        title={`Статус: ${status}`}
      />
    );
  };

  const fetchUserData = async () => {
    if (!customer_uuid || !order_uuid || !userLogin) return;
    try {
      setError("");

      const res = await fetch(`${SERVER_URL}/customer/v1/order/customer/list`, {
        method: "GET",
        headers: {
          "Accept": "multipart/mixed, application/json",
          "ngrok-skip-browser-warning": "true",
          "Content-Type": "application/json",
          ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
        },
      });

      if (!res.ok) throw new Error("Не удалось загрузить данные пользователей");

      const data = await res.json();
      const customerList = Array.isArray(data.customer_list) ? data.customer_list : [];

      const user = customerList.find(
        (u) =>
          u.login === userLogin ||
          u.username === userLogin ||
          u.tg_login === userLogin ||
          u.customer_uuid === userLogin ||
          u.uuid === userLogin
      );

      if (user) {
        setUserData(user);
        setUserOrderItems(
          Array.isArray(user.item_list) ? user.item_list.filter((i) => i.count > 0) : []
        );
      } else {
        throw new Error("Пользователь не найден");
      }
      setLoading(false);
    } catch (e) {
      setError(e.message || "Ошибка при загрузке данных пользователя");
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUserData();

    pollingIntervalRef.current = setInterval(() => {
      fetchUserData();
    }, 4000);
    return () => {
      clearInterval(pollingIntervalRef.current);
    };
  }, [customer_uuid, order_uuid, jwt_token, userLogin]);

  const handleGoBack = () => navigate(-1);

  if (loading) return <LoadingScreen message="Загрузка заказа пользователя..." />;
  if (error) {
    return (
      <div className="uo-container">
        <div className="uo-header">
          <button className="back-button" onClick={handleGoBack}>
            Назад
          </button>
        </div>
        <div className="uo-content">
          <div className="uo-error">{error}</div>
        </div>
      </div>
    );
  }

  const userName = userData?.username || userData?.login || userData?.tg_login || userLogin;

  return (
    <div className="uo-container">
      <div className="uo-header">
        <h2>Заказ пользователя</h2>
      </div>
      <div className="uo-content">
        <div className="user-name-order">{userName}</div>
        <div className="uo-list">
          {userOrderItems.length === 0 && <div className="empty-order">Нет блюд в заказе</div>}
          {userOrderItems.map((item, index) => (
            <div key={index} className="uo-item">
              <span className="dish-name">{item.name}</span>
              <span className="item-qty">{item.count}</span>
              <span className="item-x">×</span>
              <span className="item-price">{item.price} ₽</span>
              <span className="uo-result-price">{item.result_price || (item.price * item.count)} ₽</span>
              <span className="item-status">{renderStatus(item.status)}</span>
            </div>
          ))}
        </div>
      </div>
      <div className="uo-footer">
        <div className="uo-total-price">Итого: {userData?.total_price || userOrderItems.reduce((sum, item) => sum + (item.result_price || (item.price * item.count)), 0)} ₽</div>
        <button className="go-back-button" onClick={handleGoBack}>
          Назад
        </button>
      </div>
    </div>
  );
};

export default UserOrder;
