import React, { useEffect, useState } from "react";
import { useOrder } from "../OrderContext/OrderContext";
import { useNavigate } from "react-router-dom";
import LoadingScreen from "../LoadingScreen/LoadingScreen";
import { SERVER_URL } from "../../config";
import { getAuthHeaders } from '../hooks/authHeaders';
import "./Tip.css";

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

const Tip = () => {
  const { customer_uuid, order_uuid, jwt_token } = useOrder();
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [saving, setSaving] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const [selectedIdx, setSelectedIdx] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    if (!customer_uuid || !order_uuid) return;
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
        setLoading(false);
        if (!res.ok) throw new Error("Не удалось загрузить чек");
        const data = await res.json();
        setUsers(Array.isArray(data.customer_list) ? data.customer_list : []);
      })
      .catch((e) => {
        setError(e.message || "Ошибка при загрузке чека");
        setUsers([]);
        setLoading(false);
      });
  }, [customer_uuid, order_uuid, jwt_token]);

  const totalPrice = users.reduce((sum, user) => sum + (Number(user.total_price )|| 0), 0);

  const sortedUsers = [...users].sort((a, b) => {
    const nameA = (a.username || a.login || a.tg_login || '').toLowerCase();
    const nameB = (b.username || b.login || b.tg_login || '').toLowerCase();
    return nameA.localeCompare(nameB);
  });

  const handleSave = () => {
    if (!customer_uuid || !order_uuid) return;
    setSaving(true);
    fetch(`${SERVER_URL}/customer/v1/order/tip/save`, {
      method: "POST",
      headers: {
        "Accept": "application/json",
        "ngrok-skip-browser-warning": "true",
        "Content-Type": "application/json",
        ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
      },
      body: JSON.stringify({})
    })
      .then(async (res) => {
        setSaving(false);
        if (!res.ok) throw new Error("Ошибка при сохранении чека");
        setShowModal(true);
      })
      .catch((e) => {
        setSaving(false);
        alert(e.message || "Ошибка при сохранении чека");
      });
  };

  const handleUserClick = (user, idx) => {
    setSelectedIdx(idx);
    const userId = user.tg_login || user.username || user.login || user.customer_uuid || user.uuid;
    setTimeout(() => navigate(`/user-order/${userId}`), 120);
  };

  if (loading) return <LoadingScreen message="Загрузка чека..." />;
  if (error)
    return (
      <div className="tip-container">
        <div className="tip-error">{error}</div>
      </div>
    );

  return (
    <div className="tip-container">
      <div className="tip-content">
        <h2>Чек заказа</h2>
        <div className="order-list">
          {sortedUsers.map((user, idx) => (
            <div
              key={user.customer_uuid || idx}
              className={`order-item${selectedIdx === idx ? ' selected' : ''}`}
              onClick={() => handleUserClick(user, idx)}
              style={{ cursor: 'pointer' }}
            >
              <div className="order-item-name">
                {user.username || user.login || user.tg_login || `Пользователь #${idx + 1}`}
              </div>
              <div className="order-item-total">{user.total_price || 0}<RubleIcon /></div>
            </div>
          ))}
        </div>
      </div>
      <div className="tip-footer">
        <div className="tip-total-price">Итого: {totalPrice}<RubleIcon /></div>
        <div className="buttons">
          <button className="save-button" onClick={handleSave} disabled={saving}>
            {saving ? "Сохраняем..." : "Сохранить"}
          </button>
        </div>
      </div>
      {showModal && (
        <div className="tip-modal-overlay">
          <div className="tip-modal">
            <div className="tip-modal-message">Ваш чек в чате!</div>
            <button className="tip-modal-close" onClick={() => {
              setShowModal(false);
              if (window.Telegram?.WebApp?.close) {
                window.Telegram.WebApp.close();
              }
            }}>Закрыть</button>
          </div>
        </div>
      )}
    </div>
  );
};

export default Tip;
