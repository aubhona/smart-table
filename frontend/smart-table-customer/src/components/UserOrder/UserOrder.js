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
      <div className="user-order-container">
        <div className="user-order-header">
          <button className="back-button" onClick={handleGoBack}>
            Назад
          </button>
        </div>
        <div className="user-order-content">
          <div className="user-order-error">{error}</div>
        </div>
      </div>
    );
  }

  const userName = userData?.username || userData?.login || userData?.tg_login || userLogin;

  return (
    <div className="user-order-container">
      <div className="user-order-header">
        <h2>Заказ пользователя</h2>
      </div>
      <div className="user-order-content">
        <div className="user-name-order">{userName}</div>
        <div className="user-order-list">
          {userOrderItems.length === 0 && <div className="empty-order">Нет блюд в заказе</div>}
          {userOrderItems.map((item, index) => (
            <div key={index} className="user-order-item">
              <div className="item-details">
                <div className="dish-name">{item.name}</div>
                <div className="item-price">{item.price} ₽</div>
              </div>
              <div className="item-right">
                <div className="item-qty">{item.count}</div>
                <div className="result-price">{item.result_price || (item.price * item.count)} ₽</div>
                <div className="item-status">{renderStatus(item.status)}</div>
              </div>
            </div>
          ))}
        </div>
      </div>
      <div className="user-order-footer">
        <div className="user-total-price">Итого: {userData?.total_price || userOrderItems.reduce((sum, item) => sum + (item.result_price || (item.price * item.count)), 0)} ₽</div>
        <button className="go-back-button" onClick={handleGoBack}>
          Назад
        </button>
      </div>
    </div>
  );
};

export default UserOrder;
