import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useOrder } from "../OrderContext/OrderContext";
import LoadingScreen from "../LoadingScreen/LoadingScreen";
import { SERVER_URL } from "../../config";
import "./Checkout.css";

const Checkout = () => {
  const navigate = useNavigate();
  const { customer_uuid, order_uuid } = useOrder();
  const [users, setUsers] = useState([]);
  const [orderDetails, setOrderDetails] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const statusColors = ["white", "blue", "yellow", "green", "cyan"];
  const renderStatus = (status) => (
    <span
      className={`status-circle ${statusColors[(status || 1) - 1]}`}
      title={`Статус ${status}`}
    ></span>
  );

  useEffect(() => {
    if (!customer_uuid || !order_uuid) return;
    setLoading(true);

    // Получаем список пользователей и их заказы
    fetch(`${SERVER_URL}/customer/v1/order/customer/list`, {
      method: "GET",
      headers: {
        "Accept": "multipart/mixed, application/json",
        "ngrok-skip-browser-warning": "true",
        "Content-Type": "application/json",
        "Customer-UUID": customer_uuid,
        "Order-UUID": order_uuid,
        "JWT-Token": "bla-bla-bla",
      },
    })
      .then(async (res) => {
        setLoading(false);
        if (!res.ok) throw new Error("Не удалось загрузить данные пользователей");
        const data = await res.json();
        setUsers(Array.isArray(data.customer_list) ? data.customer_list : []);
        // Найти свой заказ
        const myOrder = (Array.isArray(data.customer_list) ? data.customer_list : []).find(
          (u) => u.customer_uuid === customer_uuid
        );
        setOrderDetails(Array.isArray(myOrder?.order_details) ? myOrder.order_details : []);
      })
      .catch((e) => {
        setError(e.message || "Ошибка при загрузке пользователей");
        setUsers([]);
        setOrderDetails([]);
        setLoading(false);
      });
  }, [customer_uuid, order_uuid]);

  const handleGoCatalog = () => navigate("/catalog");
  const handleFinish = () => navigate("/tip");
  const handleUserClick = (user) =>
    navigate(`/user-order/${user.login || user.username || user.customer_uuid}`);

  if (loading) return <LoadingScreen message="Загрузка заказа..." />;
  if (error)
    return (
      <div className="checkout-container">
        <div className="checkout-content">
          <div className="checkout-error">{error}</div>
        </div>
      </div>
    );

  const totalPrice = users.reduce((sum, user) => sum + (user.total_sum || 0), 0);

  // Сортировка пользователей по алфавиту
  const sortedUsers = [...users].sort((a, b) => {
    const nameA = (a.username || a.login || a.tg_login || '').toLowerCase();
    const nameB = (b.username || b.login || b.tg_login || '').toLowerCase();
    return nameA.localeCompare(nameB);
  });
  const friends = sortedUsers.filter((u) => u.customer_uuid !== customer_uuid && u.customer_uuid);

  return (
    <div className="checkout-container">
      <div className="checkout-content wide-block">
        <h2>Ваш заказ</h2>
        <div className="order-list">
          {orderDetails.length === 0 && <div className="empty-order">Нет блюд в заказе</div>}
          {orderDetails.map((item, index) => (
            <div key={index} className="order-item">
              <div className="order-item-name">{item.name}</div>
              <div className="order-item-qty">x{item.quantity}</div>
              <div className="order-item-total">{item.total} ₽</div>
              <div className="order-item-status">{renderStatus(item.status)}</div>
            </div>
          ))}
        </div>
      </div>

      <div className="friends-orders-block wide-block">
        <h3>Заказы друзей</h3>
        <div className="user-orders-list">
          {friends.length === 0 && <div className="empty-friends">Нет заказов друзей</div>}
          {friends.map((user, index) => (
            <div
              key={user.customer_uuid || index}
              className="user-order"
              onClick={() => handleUserClick(user)}
            >
              <div className="user-login">
                {user.username || user.login || user.tg_login || `Пользователь #${index + 1}`}
              </div>
              <div className="user-total-price">{user.total_sum || 0} ₽</div>
            </div>
          ))}
        </div>
      </div>

      <div className="checkout-footer">
        <div className="ch-total-price">Итого: {totalPrice} ₽</div>
        <div className="buttons">
          <button className="ch-go-back-button" onClick={handleGoCatalog}>
            Вернуться в каталог
          </button>
          <button className="button" onClick={handleFinish}>
            Завершить заказ
          </button>
        </div>
      </div>
    </div>
  );
};

export default Checkout;
