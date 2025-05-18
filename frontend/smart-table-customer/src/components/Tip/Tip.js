import React from "react";
import { useNavigate } from "react-router-dom";
import "./Tip.css";

const Tip = () => {
  const navigate = useNavigate();

  // Заказы пользователей
  const users = [
    { name: "aubhona", totalPrice: 5000 },
    { name: "saundler", totalPrice: 3200 },
  ];

  // Моя общая сумма (например, из чекаута)
  const myTotalPrice = 5000; 

  // Общая сумма всех заказов
  const totalPrice = myTotalPrice + users.reduce((sum, user) => sum + user.totalPrice, 0);

  const handleSave = () => {
    // Логика для сохранения чека (например, скачивание, сохранение в файл)
    console.log("Сохранить чек!");
  };

  return (
    <div className="tip-container">
      <div className="tip-content">
        <h2>Чек</h2>
        <div className="order-list">
          {/* Мой заказ */}
          <div className="order-item">
            <div className="order-item-name">Мой заказ</div>
            <div className="order-item-total">{myTotalPrice} ₽</div>
          </div>

          {/* Заказы других пользователей */}
          {users.map((user, index) => (
            <div key={index} className="order-item">
              <div className="order-item-name">{user.name}</div>
              <div className="order-item-total">{user.totalPrice} ₽</div>
            </div>
          ))}
        </div>
      </div>

      <div className="tip-footer">
        <div className="tip-total-price">Итого: {totalPrice} ₽</div>
        <div className="buttons">
          <button className="save-button" onClick={handleSave}>
            Сохранить
          </button>
        </div>
      </div>
    </div>
  );
};

export default Tip;
