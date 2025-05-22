import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useOrder } from "../OrderContext/OrderContext";
import { SERVER_URL } from "../../config";
import { getAuthHeaders } from '../../utils/authHeaders';
import "./UsersList.css";

function UsersList() {
  const navigate = useNavigate();
  const { order_uuid, customer_uuid, room_code, jwt_token } = useOrder() || {};
  const [roomCode, setRoomCode] = useState(room_code || "");
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [selectedIdx, setSelectedIdx] = useState(null);

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
        ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
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
  }, [customer_uuid, order_uuid, jwt_token]);

  const goBack = () => {
    navigate("/catalog");
  };

  const handleUserClick = (user, idx) => {
    setSelectedIdx(idx);
    const userId = user.tg_login || user.username || user.login || user.customer_uuid || user.uuid;
    setTimeout(() => navigate(`/user-order/${userId}`), 120);
  };

  const sortedUsers = [...users].sort((a, b) => {
    const nameA = (a.username || a.login || a.tg_login || '').toLowerCase();
    const nameB = (b.username || b.login || b.tg_login || '').toLowerCase();
    return nameA.localeCompare(nameB);
  });

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
            <ul className="users-list">
              {sortedUsers.map((user, idx) => (
                <li
                  key={user.customer_uuid || user.uuid || idx}
                  className={`user-list-item${selectedIdx === idx ? ' selected' : ''}`}
                  onClick={() => handleUserClick(user, idx)}
                  style={{ cursor: 'pointer', background: 'transparent', boxShadow: 'none' }}
                >
                  <div className="user-list-bg">
                    <div className="user-list-inner">
                      {user.username || user.login || user.tg_login || `Пользователь #${idx + 1}`}
                    </div>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>

      <div className="ul-bottom-section">
        <button className="ul-go-back-button" onClick={goBack}>
          Назад
        </button>
      </div>
    </div>
  );
}

export default UsersList;
