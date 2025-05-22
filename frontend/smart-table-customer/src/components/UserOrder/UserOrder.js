import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { useOrder } from "../OrderContext/OrderContext";
import LoadingScreen from "../LoadingScreen/LoadingScreen";
import { SERVER_URL } from "../../config";
import { getAuthHeaders } from '../../utils/authHeaders';
import "./UserOrder.css";

const UserOrder = () => {
  const navigate = useNavigate();
  const { userLogin } = useParams();
  const { customer_uuid, order_uuid, jwt_token } = useOrder();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [userData, setUserData] = useState(null);
  const [userOrderItems, setUserOrderItems] = useState([]);

  // statusColors и status-circle как в Checkout
  const statusColors = ["white", "blue", "yellow", "green", "cyan"];
  const renderStatus = (status) => (
    <span className={`status-circle ${statusColors[(status || 1) - 1]}`}></span>
  );

  useEffect(() => {
    if (!customer_uuid || !order_uuid || !userLogin) return;
    setLoading(true);

    fetch(`${SERVER_URL}/customer/v1/order/customer/list`, {
      method: "GET",
      headers: {
        "Accept": "multipart/mixed, application/json",
        "ngrok-skip-browser-warning": "true",
        "Content-Type": "application/json",
        ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
      },
    })
      .then(async (res) => {
        if (!res.ok) throw new Error("Не удалось загрузить данные пользователей");
        const data = await res.json();
        const customerList = Array.isArray(data.customer_list) ? data.customer_list : [];
        
        // Находим данные пользователя по логину или uuid
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
            Array.isArray(user.item_list) ? user.item_list.filter(i => i.count > 0) : []
          );
        } else {
          throw new Error("Пользователь не найден");
        }
        setLoading(false);
      })
      .catch((e) => {
        setError(e.message || "Ошибка при загрузке данных пользователя");
        setLoading(false);
      });
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
