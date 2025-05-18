import React from "react";
import { useNavigate } from "react-router-dom";
import "./Checkout.css";

const Checkout = () => {
  const navigate = useNavigate();

  const orderDetails = [
    { name: "Бургер", price: 1000, quantity: 2, total: 2000, status: 1 },
    { name: "Плов", price: 1200, quantity: 1, total: 1200, status: 3 },
    { name: "Яблоко", price: 600, quantity: 3, total: 1800, status: 5 },
  ];

  const totalPrice = orderDetails.reduce((sum, item) => sum + item.total, 0);
  const users = [
    { login: "aubhona", totalPrice: 5000 },
    { login: "saundler", totalPrice: 3200 },
  ];

  const handleGoCatalog = () => navigate("/catalog");
  const handleFinish = () => navigate("/tip");
  const handleUserClick = (userLogin) => {
    navigate(`/user-order/${userLogin}`);
  };

  const renderStatus = (status) => {
    const statusColors = ["white", "blue", "yellow", "green", "cyan"];
    return (
      <span
        className={`status-circle ${statusColors[status - 1]}`}
        title={`Статус ${status}`}
      ></span>
    );
  };

  return (
    <div className="checkout-container">
      <div className="checkout-header">
        <button className="back-button" onClick={() => navigate(-1)}>
          Назад
        </button>
      </div>

      <div className="checkout-content">
        <h2>Ваш заказ</h2>
        <div className="order-list">
          {orderDetails.map((item, index) => (
            <div key={index} className="order-item">
              <div className="order-item-name">{item.name}</div>
              <div className="order-item-total">{item.total} ₽</div>
              <div className="order-item-status">
                {renderStatus(item.status)}
              </div>
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

      <div className="user-orders">
        {users.map((user, index) => (
          <div 
          key={index} 
          className="user-order"
          onClick={() => handleUserClick(user.login)}>
            <div>{user.login}</div>
            <div className="user-total-price">{user.totalPrice} ₽</div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Checkout;