import React from "react";
import { useNavigate } from "react-router-dom";
import "./UserOrder.css";

const UserOrder = () => {
  const navigate = useNavigate();

  const userOrders = [
    { name: "Бургер", price: 1000, quantity: 2, total: 2000, status: 1 },
    { name: "Плов", price: 1200, quantity: 1, total: 1200, status: 2 },
    { name: "Яблоко", price: 600, quantity: 3, total: 1800, status: 3 },
  ];

  const totalPrice = userOrders.reduce((sum, item) => sum + item.total, 0);

  const renderStatus = (status) => {
    const statusColors = ["white", "blue", "yellow", "green", "cyan"];
    return <span className={`status-circle ${statusColors[status - 1]}`}></span>;
  };

  const handleGoBack = () => navigate(-1); 

  return (
    <div className="user-order-container">
      <div className="user-order-header">
        <button className="back-button" onClick={handleGoBack}>
          Назад
        </button>
      </div>

      <div className="user-order-content">
        <h2>Заказ пользователя aubhona</h2>
        <div className="user-order-list">
          {userOrders.map((item, index) => (
            <div key={index} className="user-order-item">
              <div className="user-order-item-name">{item.name}</div>
              <div className="user-order-item-total">{item.total} ₽</div>
              <div className="user-order-item-status">
                {renderStatus(item.status)}
              </div>
            </div>
          ))}
        </div>
      </div>

      <div className="user-order-footer">
        <div className="user-order-total-price">Итого: {totalPrice} ₽</div>
        <div className="user-order-actions">
        </div>
      </div>
    </div>
  );
};

export default UserOrder;
