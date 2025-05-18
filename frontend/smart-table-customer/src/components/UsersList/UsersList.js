import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useOrder } from "../OrderContext/OrderContext";
import { SERVER_URL } from "../../config";
import "./UsersList.css";

function UsersList() {
  const navigate = useNavigate();
  const { order_uuid, customer_uuid, room_code } = useOrder() || {};
  const [roomCode, setRoomCode] = useState(room_code || "");
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!customer_uuid || !order_uuid) {
      setError("Не хватает данных для запроса. Проверьте контекст заказа.");
      setLoading(false);
      return;
    }

    setLoading(true);

    fetch(`${SERVER_URL}/customer/v1/order/customer/list`, {
      method: "GET",
      headers: {
        Accept: "multipart/mixed, application/json",
        "Content-Type": "application/json",
        "ngrok-skip-browser-warning": "69420",
        "Customer-UUID": customer_uuid,
        "Order-UUID": order_uuid,
        "JWT-Token": "bla-bla-bla"
      }
    })
      .then(async (res) => {
        setLoading(false);
        if (!res.ok) {
          throw new Error("Ошибка при получении пользователей заказа");
        }
        const data = await res.json();

        if (data.room_code) setRoomCode(data.room_code);
        if (Array.isArray(data.users)) setUsers(data.users);
        else if (Array.isArray(data.customer_list)) setUsers(data.customer_list);
        else setUsers([]);
      })
      .catch((err) => {
        setError(err.message || "Ошибка загрузки списка пользователей");
        setLoading(false);
      });
  }, [customer_uuid, order_uuid]);

  const goBack = () => {
    navigate("/catalog");
  };

  return (
    <div className="users-container">
      <div className="users-header">
        <div className="room-code">
          {roomCode ? `Код комнаты: ${roomCode}` : "Код не получен"}
        </div>
      </div>
      
      <div className="users-content">
        <h2 className="users-title">Список пользователей</h2>
        <div className="users-list">
          {loading ? (
            <div className="users-message loading">Загрузка...</div>
          ) : error ? (
            <div className="users-message error">{error}</div>
          ) : users.length === 0 ? (
            <div className="users-message empty">Нет друзей!</div>
          ) : (
            users.map((user) => (
              <div key={user.id || user.customer_uuid} className="user-item">
                <div className="user-avatar">
                  {(user.name || user.username || user.login || "?").charAt(0).toUpperCase()}
                </div>
                <div className="user-name">{user.name || user.username || user.login}</div>
              </div>
            ))
          )}
        </div>
      </div>

      <div className="bottom-section">
        <button className="go-back-button" onClick={goBack}>
          Назад
        </button>
      </div>
    </div>
  );
}

export default UsersList;
