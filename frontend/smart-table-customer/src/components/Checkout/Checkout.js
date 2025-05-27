import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useOrder } from "../OrderContext/OrderContext";
import LoadingScreen from "../LoadingScreen/LoadingScreen";
import { SERVER_URL } from "../../config";
import { getAuthHeaders } from '../hooks/authHeaders';
import "./Checkout.css";

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

const Checkout = () => {
  const navigate = useNavigate();
  const { customer_uuid, order_uuid, jwt_token } = useOrder();
  const [users, setUsers] = useState([]);
  const [orderDetails, setOrderDetails] = useState([]);
  const [isHost, setIsHost] = useState(false);
  const [loading, setLoading] = useState(true);
  const [finishing, setFinishing] = useState(false);
  const [error, setError] = useState("");
  const [selectedFriendIdx, setSelectedFriendIdx] = useState(null);

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

  useEffect(() => {
    if (!customer_uuid || !order_uuid) return;
    setLoading(true);
  
    let isMounted = true;
    let pollingInterval;
  
    const fetchUsersAndOrders = async () => {
      try {
        const res = await fetch(`${SERVER_URL}/customer/v1/order/customer/list`, {
          method: "GET",
          headers: {
            Accept: "multipart/mixed, application/json",
            "ngrok-skip-browser-warning": "true",
            "Content-Type": "application/json",
            ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
          },
        });
        if (!res.ok) throw new Error("Не удалось загрузить данные пользователей");
        const data = await res.json();
  
        if (!isMounted) return;
  
        const customerList = Array.isArray(data.customer_list) ? data.customer_list : [];
        setUsers(customerList);
        const myOrder = customerList.find(
          (u) => u.customer_uuid === customer_uuid || u.uuid === customer_uuid
        );
        if (myOrder?.is_active) {
          setOrderDetails(
            Array.isArray(myOrder.item_list) ? myOrder.item_list.filter((i) => i.count > 0) : []
          );
        } else {
          setOrderDetails([]);
        }
        setIsHost(myOrder?.is_host === true);
        setLoading(false);
      } catch (e) {
        if (!isMounted) return;
        setError(e.message || "Ошибка при загрузке пользователей");
        setUsers([]);
        setOrderDetails([]);
        setLoading(false);
      }
    };
  
    fetchUsersAndOrders();
  
    pollingInterval = setInterval(fetchUsersAndOrders, 4000); 
  
    return () => {
      isMounted = false;
      clearInterval(pollingInterval);
    };
  }, [customer_uuid, order_uuid, jwt_token]);

  const handleGoCatalog = () => navigate("/catalog");
  const handleUserClick = (user, idx) => {
    setSelectedFriendIdx(idx);
    const userId = user.tg_login || user.username || user.login || user.customer_uuid || user.uuid;
    setTimeout(() => navigate(`/user-order/${userId}`), 120);
  };

  const handleFinishOrder = async () => {
    setFinishing(true);
    setError("");
    try {
      const res = await fetch(`${SERVER_URL}/customer/v1/order/finish`, {
        method: "POST",
        headers: {
          ...getAuthHeaders({ customer_uuid, jwt_token, order_uuid }),
        },
      });
  
      if (res.status === 204) {
        navigate("/tip");
      } else {
        let err = "Ошибка при завершении заказа";
        try {
          const data = await res.json();
          if (data?.error) err = data.error;
        } catch {}
        setError(err);
      }
    } catch (e) {
      setError(e.message || "Ошибка завершения заказа");
    }
    setFinishing(false);
  };

  if (loading) return <LoadingScreen message="Загрузка заказа..." />;
  if (error)
    return (
      <div className="checkout-container">
        <div className="checkout-content">
          <div className="checkout-error">{error}</div>
        </div>
      </div>
    );

  const friends = users.filter(
    (u) => (u.customer_uuid || u.uuid) !== customer_uuid && u.is_active && (u.customer_uuid || u.uuid)
  );

  return (
    <div className="checkout-container">
      <div className="checkout-content wide-block">
        <h2 style={{display:'flex',justifyContent:'space-between',alignItems:'center'}}>
          Ваш заказ
          {(() => {
            const me = users.find(u => u.customer_uuid === customer_uuid || u.uuid === customer_uuid);
            if (me && (me.result_price || me.total_price)) {
              return (
                <span style={{fontSize:'1.0em',fontWeight:'bold',color:'#5e7643',marginLeft:'20px'}}>
                  {(me.result_price || me.total_price)}<RubleIcon />
                </span>
              );
            }
            return null;
          })()}
        </h2>
        <div className="order-list">
          {orderDetails.length === 0 && <div className="empty-order">Нет блюд в заказе</div>}
          {orderDetails.map((item, index) => (
            <div key={index} className="order-item">
              <div className="order-item-name">{item.name}</div>
              <div className="order-item-qty">x{item.count}</div>
              <div className="order-item-total">{(item.result_price || (item.price * item.count))}<RubleIcon /></div>
              <div className="order-item-status">{renderStatus(item.status)}</div>
            </div>
          ))}
        </div>
      </div>

      <div className="friends-orders-block wide-block">
        <h3>Заказы друзей</h3>
        <div className="user-orders-list">
          {friends.length === 0 && (
            <div className="empty-friends">Нет заказов друзей</div>
          )}
          {friends.map((user, index) => (
            <div
              key={user.customer_uuid || user.uuid || index}
              className={`user-order user-order-summary${selectedFriendIdx === index ? ' selected' : ''}`}
              onClick={() => handleUserClick(user, index)}
              style={{cursor: 'pointer', width: '100%'}}
            >
              <div className="user-order-row" style={{display:'flex',justifyContent:'space-between',alignItems:'center',fontSize:'20px',padding:'10px 0'}}>
                <span className="user-login" style={{fontWeight:'bold'}}>{user.username || user.login || user.tg_login}</span>
                <span className="user-total-price">{user.total_price || 0}<RubleIcon /></span>
              </div>
            </div>
          ))}
        </div>
      </div>

      <div className="checkout-footer">
        <div className="ch-total-price">
          Итого: {users.reduce((sum, user) => sum + (Number(user.total_price)|| 0), 0)}<RubleIcon /></div>
        <div className="buttons">
          <button className="ch-go-back-button" onClick={handleGoCatalog}>
            Вернуться в каталог
          </button>
          {isHost && (
            <button className="button" onClick={handleFinishOrder} disabled={finishing}>
              {finishing ? "Завершаем заказ..." : "Завершить заказ"}
            </button>
          )}
        </div>
      </div>
    </div>
  );
};

export default Checkout;
