import React, { useState, useEffect } from "react";
import { handleMultipartResponse } from '../components/multipartUtils';
import { useParams, useNavigate } from "react-router-dom";
import { QRCodeSVG } from "qrcode.react";
import { toPng } from "html-to-image";
import { v4 as uuidv4 } from 'uuid';

import DefaultApi from "../api/place_api/generated/src/api/DefaultApi";
import AdminV1PlaceEmployeeAddRequest from "../api/place_api/generated/src/model/AdminV1PlaceEmployeeAddRequest";
import AdminV1PlaceMenuDishCreateRequest from "../api/place_api/generated/src/model/AdminV1PlaceMenuDishCreateRequest";
import AdminV1PlaceTableDeepLinksListRequest from "../api/place_api/generated/src/model/AdminV1PlaceTableDeepLinksListRequest";
import AdminV1PlaceEmployeeDeleteRequest from "../api/place_api/generated/src/model/AdminV1PlaceEmployeeDeleteRequest";
import AdminV1PlaceEmployeeEditRequest from "../api/place_api/generated/src/model/AdminV1PlaceEmployeeEditRequest";
import AdminV1PlaceMenuDishDeleteRequest from "../api/place_api/generated/src/model/AdminV1PlaceMenuDishDeleteRequest";
import AdminV1PlaceMenuDishEditRequest from "../api/place_api/generated/src/model/AdminV1PlaceMenuDishEditRequest";

import "../styles/PlaceScreen.css";
import { SERVER_URL } from "../config";

export default function PlaceDetail() {
  const { restaurant_uuid, place_uuid } = useParams();
  const navigate = useNavigate();

  const savedRest = JSON.parse(localStorage.getItem("current_restaurant") || "{}");
  const restaurantName = savedRest.restaurant_name || restaurant_uuid;
  const savedPlace = JSON.parse(localStorage.getItem("current_place") || "{}");
  const placeName = savedPlace.place_name || place_uuid;

  const [tab, setTab] = useState("staff");

  const [staff, setStaff] = useState([]);
  const [loading, setLoading] = useState(false);
  const [showAddModal, setShowAddModal] = useState(false);

  const [login, setLogin] = useState("");
  const [role, setRole] = useState("");
  const [error, setError] = useState("");
  const [priceError, setPriceError] = useState("");

  const [menuDishes, setMenuDishes] = useState([]); 
  const [availableDishes, setAvailableDishes] = useState([]);
  const [showDishPicker, setShowDishPicker] = useState(false);
  const [selectedDish, setSelectedDish] = useState(null);
  const [price, setPrice] = useState("");

  const [deepLinks, setDeepLinks] = useState([]);
  const [loadingQR, setLoadingQR] = useState(false);
  const [qrError, setQrError] = useState("");

  const [showEditStaffModal, setShowEditStaffModal] = useState(false);
  const [editStaffData, setEditStaffData] = useState(null);

  const [showEditMenuDishModal, setShowEditMenuDishModal] = useState(false);
  const [editMenuDishData, setEditMenuDishData] = useState(null);
  const [editMenuDishPrice, setEditMenuDishPrice] = useState("");

  const STATUS_FLOW = ['В обработке', 'Принят', 'Готовится', 'Готов', 'Подан'];
  const ORDER_STATUSES = ['Открыт', 'Ожидает оплаты', 'Оплачен', 'Отменен'];

  const [orderSubTab, setOrderSubTab] = useState('open');
  const [orders, setOrders] = useState([
    {
      id: uuidv4(),
      status: 'Открыт',
      createdAt: new Date().toLocaleString(),
      tableId: `${place_uuid}_3`,
      guests: 4,
      totalPrice: 2150,
      customers: [
        {
          id: uuidv4(),
          name: 'Иван Петров',
          instagram: '@ivan_petrov',
          items: [
            {
              id: uuidv4(),
              name: 'Бургер',
              price: 500,
              status: 'В обработке',
              amount: 2
            },
            {
              id: uuidv4(),
              name: 'Кола',
              price: 150,
              status: 'Подан',
              amount: 1
            }
          ]
        },
        {
          id: uuidv4(),
          name: 'Мария Сидорова',
          instagram: '@maria_sid',
          items: [
            {
              id: uuidv4(),
              name: 'Салат Цезарь',
              price: 300,
              status: 'Готовится',
              amount: 1
            }
          ]
        }
      ]
    },
  {
    id: uuidv4(),
    status: 'Оплачен',
    createdAt: new Date().toLocaleString(),
    tableId: `${place_uuid}_2`,
    guests: 1,
    totalPrice: 1800,
    customers: [
      {
        id: uuidv4(),
        name: 'Александр Соловкин',
        instagram: '@l4sthope',
        items: [
          {
            id: uuidv4(),
            name: 'Суп',
            price: 400,
            status: 'Отменен',
            user: '@user3',
            amount: 2
          }
        ]
      }
    ],
  }
]);
const [selectedOrder, setSelectedOrder] = useState(null);
const [showCheckout, setShowCheckout] = useState(false);

  const userUUID = localStorage.getItem("user_uuid");
  const jWTToken = localStorage.getItem("jwt_token");
  const api = new DefaultApi();

  api.apiClient.basePath = SERVER_URL;
  api.apiClient.defaultHeaders = {
    "User-UUID": userUUID,
    "JWT-Token": jWTToken,
    "ngrok-skip-browser-warning": "true",
  };

  const updateOrderStatus = (orderId, newStatus) => {
  setOrders(prev => 
    prev.map(order => 
      order.id === orderId ? { ...order, status: newStatus } : order
    )
  );
};

const updateItemStatus = (orderId, customerId, itemId, newStatus) => {
    setOrders(prev => 
      prev.map(order => {
        if(order.id === orderId) {
          return {
            ...order,
            customers: order.customers.map(customer => {
              if(customer.id === customerId) {
                return {
                  ...customer,
                  items: customer.items.map(item => 
                    item.id === itemId ? { ...item, status: newStatus } : item
                  )
                };
              }
              return customer;
            })
          };
        }
        return order;
      })
    );
  };

  const handleGenerateQRCodes = async () => {
    try {
      await loadDeepLinks();
    } catch (e) {
      console.error("Ошибка генерации QR-кодов:", e);
    }
  };

  const downloadQR = async (id) => {
    const element = document.getElementById(`qrcode-${id}`);
    if (!element) return;

    try {
      const dataUrl = await toPng(element);
      const link = document.createElement("a");
      link.download = `table-${id + 1}.png`;
      link.href = dataUrl;
      link.click();
    } catch (error) {
      console.error("Ошибка сохранения QR-кода:", error);
    }
  };

  async function loadDeepLinks() {
    setLoadingQR(true);
    setQrError("");
    try {
      const req = AdminV1PlaceTableDeepLinksListRequest.constructFromObject({
        place_uuid: place_uuid,
      });
      const response = await new Promise((resolve, reject) => {
        api.adminV1PlaceTableDeeplinksListPost(
          userUUID,
          jWTToken,
          req,
          (error, data) => (error ? reject(error) : resolve(data))
        );
      });
      setDeepLinks(response.deeplinks || []);
    } catch (error) {
      console.error("Ошибка загрузки ссылок:", error);
      setQrError(error.body?.message || "Не удалось загрузить ссылки");
    } finally {
      setLoadingQR(false);
    }
  }

  async function loadAvailableDishes() {
    try {
      const resp = await fetch(`${SERVER_URL}/admin/v1/restaurant/dish/list`, {
        method: "POST",
        headers: {
          Accept: "multipart/mixed, application/json",
          "Content-Type": "application/json",
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true",
        },
        body: JSON.stringify({ restaurant_uuid: restaurant_uuid }),
      });
    
    const { list, imagesMap } = await handleMultipartResponse(resp, 'dish_list');
    
    setAvailableDishes(list.map(d => ({
      ...d,
      imageUrl: imagesMap[d.picture_key] || null
    })));
  } catch (e) {
    console.error("Ошибка загрузки блюд:", e);
    setAvailableDishes([]);
  }
}

  async function loadStaff() {
    setLoading(true);
    try {
      const resp = await fetch(
        `${SERVER_URL}/admin/v1/place/employee/list`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
            "User-UUID": userUUID,
            "JWT-Token": jWTToken,
            "ngrok-skip-browser-warning": "true",
          },
          body: JSON.stringify({ place_uuid: place_uuid }),
        }
      );
      if (!resp.ok) throw new Error(`HTTP ${resp.status}`);
      const data = await resp.json();
      setStaff(data.employee_list || []);
    } catch (e) {
      console.error("Ошибка загрузки сотрудников:", e);
      setStaff([]);
    } finally {
      setLoading(false);
    }
  }

  async function loadMenuDishes() {
  setLoading(true);
  try {
    const resp = await fetch(`${SERVER_URL}/admin/v1/place/menu/dish/list`, {
      method: "POST",
      headers: {
        "Accept": "multipart/mixed, application/json",
        "Content-Type": "application/json",
        "User-UUID": userUUID,
        "JWT-Token": jWTToken,
        "ngrok-skip-browser-warning": "true",
      },
      body: JSON.stringify({ place_uuid: place_uuid }),
    });

    const { list, imagesMap } = await handleMultipartResponse(resp);
    
    setMenuDishes(list.map(d => ({
      ...d,
      imageUrl: imagesMap[d.picture_key] || null,
      price: d.price
    })));
  } catch (e) {
    console.error("Ошибка загрузки блюд:", e);
    setMenuDishes([]);
  } finally {
    setLoading(false);
  }
}

  useEffect(() => {
    if (tab === "staff") {
      loadStaff();
    } else if (tab === "menu") {
      loadMenuDishes();
    } else if (tab === "orders") {
      loadDeepLinks();
    }
  }, [place_uuid, tab]);

  async function handleAddStaff() {
    if (!login.trim() || !role.trim()) {
      setError("Заполните все поля");
      return;
    }
    setError("");
    try {
      const req = AdminV1PlaceEmployeeAddRequest.constructFromObject({
        place_uuid: place_uuid,
        employee_login: login.trim(),
        employee_role: role.trim(),
      });
      console.log(req);
      await new Promise((res, rej) =>
        api.adminV1PlaceEmployeeAddPost(
          userUUID,
          jWTToken,
          req,
          (err) => (err ? rej(err) : res())
        )
      );

      await loadStaff();

      setShowAddModal(false);
      setLogin("");
      setRole("");
    } catch (e) {
      console.error("Ошибка добавления сотрудника:", e);
      setError(e.body?.message || e.message);
    }
  }

  const handleDeleteStaff = async (employee) => {
    console.log(employee);
    if (!window.confirm(`Удалить сотрудника ${employee.login}?`)) return;
    try {
      const req = AdminV1PlaceEmployeeDeleteRequest.constructFromObject({
        place_uuid: place_uuid,
        employee_uuid: employee.uuid,
      });
      await new Promise((res, rej) =>
        api.adminV1PlaceEmployeeDeletePost(userUUID, jWTToken, req, (err) =>
          err ? rej(err) : res()
        )
      );
      await loadStaff();
    } catch (e) {
      alert(e.body?.message || e.message || "Ошибка удаления сотрудника");
    }
  };

  const handleEditStaff = (employee) => {
  setEditStaffData(employee);
  setRole(employee.employee_role);
  setShowEditStaffModal(true);
};

  const handleUpdateStaff = async () => {
    if (!editStaffData || !role.trim()) return;
    try {
      const req = AdminV1PlaceEmployeeEditRequest.constructFromObject({
        place_uuid: place_uuid,
        employee_uuid: editStaffData.uuid,
        employee_role: role.trim(),
      });
      await new Promise((res, rej) =>
        api.adminV1PlaceEmployeeEditPost(userUUID, jWTToken, req, (err) =>
          err ? rej(err) : res()
        )
      );
      setShowEditStaffModal(false);
      setEditStaffData(null);
      await loadStaff();
    } catch (e) {
      alert(e.body?.message || e.message || "Ошибка редактирования");
    }
  };

  const handleAddMenuItem = async () => {
    if (!selectedDish) {
      setError("Выберите блюдо");
      return;
    }

    const pc = Number(price);
    if (!Number.isFinite(pc) || pc <= 0) {
      setPriceError("Укажите цену больше 0");
      return;
    }

    try {
      const req = AdminV1PlaceMenuDishCreateRequest.constructFromObject({
        place_uuid: place_uuid,
        dish_uuid: selectedDish.id,
        price: pc,
      });
      console.log(selectedDish);
      await new Promise((res, rej) =>
        api.adminV1PlaceMenuDishCreatePost(
          userUUID,
          jWTToken,
          req,
          (err) => (err ? rej(err) : res())
        )
      );

      await loadMenuDishes();
      setShowAddModal(false);
      setPrice("");
      setSelectedDish(null);
      setPriceError("");
    } catch (e) {
      console.error("Ошибка добавления блюда в меню:", e);
      setError(e.body?.message || e.message);
    }
  };

  const handleDeleteMenuDish = async (dish) => {
    if (!window.confirm(`Удалить блюдо "${dish.name}" из меню?`)) return;
    try {
      const req = AdminV1PlaceMenuDishDeleteRequest.constructFromObject({
        menu_dish_uuid: dish.uuid,
      });
      await new Promise((res, rej) =>
        api.adminV1PlaceMenuDishDeletePost(userUUID, jWTToken, req, (err) =>
          err ? rej(err) : res()
        )
      );
      await loadMenuDishes();
    } catch (e) {
      alert(e.body?.message || e.message || "Ошибка удаления блюда");
    }
  };

  const handleEditMenuDish = (dish) => {
  setEditMenuDishData(dish);
  setEditMenuDishPrice(dish.price);
  setShowEditMenuDishModal(true);
};

  const handleUpdateMenuDish = async () => {
    if (!editMenuDishData || !editMenuDishPrice) return;
    try {
      const req = AdminV1PlaceMenuDishEditRequest.constructFromObject({
        menu_dish_uuid: editMenuDishData.uuid,
        dish_uuid: editMenuDishData.dish_uuid,
        price: editMenuDishPrice,
      });
      await new Promise((res, rej) =>
        api.adminV1PlaceMenuDishEditPost(userUUID, jWTToken, req, (err) =>
          err ? rej(err) : res()
        )
      );
      setShowEditMenuDishModal(false);
      setEditMenuDishData(null);
      setEditMenuDishPrice("");
      await loadMenuDishes();
    } catch (e) {
      alert(e.body?.message || e.message || "Ошибка редактирования блюда");
    }
  };  

  return (
    <div className="ps-container">
      <div className="ps-header-bar">
        <button className="ps-back-button" onClick={() => navigate(-1)}>
          Назад
        </button>
        <h1 className="ps-title">Ресторан: {restaurantName},</h1>
        <h1 className="ps-title">Адрес: {placeName}</h1>

        {tab === "staff" && (
          <button className="ps-create-button" onClick={() => setShowAddModal(true)}>
            Добавить сотрудника
          </button>
        )}
        {tab === "menu" && (
          <button className="ps-create-button" onClick={() => setShowAddModal(true)}>
            Добавить блюдо
          </button>
        )}
        {tab === "tables" && (
          <button className="ps-create-button" onClick={handleGenerateQRCodes}>
            Сгенерировать QR-код
          </button>
        )}

        <button className="ps-profile-button">
          <span className="material-icons">person</span>
        </button>
      </div>

      <div className="ps-tabs">
        <div
          className={`tab ${tab === "staff" ? "active" : ""}`}
          onClick={() => setTab("staff")}
        >
          Сотрудники
        </div>
        <div
          className={`tab ${tab === "menu" ? "active" : ""}`}
          onClick={() => setTab("menu")}
        >
          Меню
        </div>
        <div
          className={`tab ${tab === "orders" ? "active" : ""}`}
          onClick={() => setTab("orders")}
        >
          Заказы
        </div>
        <div
          className={`tab ${tab === "tables" ? "active" : ""}`}
          onClick={() => setTab("tables")}
        >
          Столы
        </div>

      </div>

      <div className="ps-content">
        {tab === "staff" && (
          <>
            {loading && <p className="ps-loading">Загрузка…</p>}
            {!loading && staff.length === 0 && (
              <p className="ps-empty">Нет сотрудников</p>
            )}
            {!loading &&
              staff.map((u) => (
                <div key={u.uuid} className="ps-item">
                  <span>{u.first_name} {u.last_name}</span>
                  <span>{u.login}</span>
                  <span>{u.employee_role}</span>
                  <button
                    className="ps-button ps-edit-button"
                    style={{ marginRight: 8 }}
                    onClick={() => handleEditStaff(u)}
                  >
                    <span className="material-icons">edit</span>
                  </button>
                  <button
                    className="ps-button ps-button-cancel"
                    onClick={() => handleDeleteStaff(u)}
                  >
                    <span className="material-icons">delete</span>
                  </button>
                </div>
              ))}

            {showAddModal && (
              <div className="ps-backdrop">
                <div className="ps-modal">
                  <h3>Добавить сотрудника</h3>
                  <input
                    className="ps-input"
                    placeholder="Логин"
                    value={login}
                    onChange={(e) => setLogin(e.target.value)}
                  />
                  <select
                    className="ps-role-select"
                    value={role}
                    onChange={(e) => setRole(e.target.value)}
                  >
                    <option value="" disabled>
                      Выберите роль
                    </option>
                    <option value="Админ">Админ</option>
                    <option value="Официант">Официант</option>
                  </select>
                  
                  {error && <div className="ps-error-text">{error}</div>}
                  <div className="ps-modal-buttons">
                    <button
                      className="ps-button"
                      onClick={handleAddStaff}
                    >
                      Добавить
                    </button>
                    <button
                      className="ps-button ps-button-cancel"
                      onClick={() => {
                        setShowAddModal(false);
                        setError("");
                      }}
                    >
                      Отмена
                    </button>
                  </div>
                </div>
              </div>
            )}
          </>
        )}

        {tab === "menu" && (
          <>
            {loading ? (
              <p className="ps-loading">Загрузка меню...</p>
            ) : (
              <div className="menu-container">
                {menuDishes.length === 0 ? (
                  <p className="ps-empty">Меню пусто</p>
                ) : (
                  menuDishes.map(dish => (
                    <div key={dish.uuid} className="menu-item">
                      <div className="dish-image">
                        {dish.imageUrl ? (
                          <img src={dish.imageUrl} alt={dish.name} />
                        ) : (
                          <div className="no-image">Нет фото</div>
                        )}
                      </div>
                      <div className="dish-info">
                        <div className="dish-header">
                          <h3>{dish.name}</h3>
                        </div>
                        <p>{dish.description}</p>
                        <p>Категория: {dish.category}</p>
                        <p>{dish.calories} ккал, {dish.weight} г.</p>
                        <div className="price-tag">{dish.price} ₽</div>
                        <div className="dish-actions">
                          <button
                            className="ps-button ps-edit-button"
                            style={{ marginRight: 8 }}
                            onClick={() => handleEditMenuDish(dish)}
                          >
                            <span className="material-icons">edit</span>
                          </button>
                          <button
                            className="ps-button ps-button-cancel"
                            onClick={() => handleDeleteMenuDish(dish)}
                          >
                            <span className="material-icons">delete</span>
                          </button>    
                        </div>                               
                      </div>
                    </div>
                  ))
                )}
              </div>
            )}
          </>
        )}
      
        {tab === "tables" && (
          <div className="qr-container">
            {loadingQR && <p className="ps-loading">Загрузка QR-кодов...</p>}
            {qrError && <p className="ps-error-text">{qrError}</p>}
            
            {!loadingQR && deepLinks.length === 0 && (
              <p className="ps-empty">Нет доступных столов</p>
            )}

            {!loadingQR && deepLinks.map((link, index) => (
              <div key={index} className="qr-card">
                <div className="qr-code-wrapper" id={`qrcode-${index}`}>
                  <QRCodeSVG 
                    value={link} 
                    size={200}
                    fgColor="#2d2a2a"
                    bgColor="#ffffff"
                    level="H"
                  />
                </div>
                <div className="qr-meta">
                  <span>Стол {index + 1}</span>
                  <button 
                    onClick={() => downloadQR(index)}
                  >
                    Сохранить PNG
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}

        {tab === "orders" && (
    <div className="ps-orders-container">
      <div className="ps-order-subtabs">
        <button
          className={`subtab ${orderSubTab === 'open' ? 'active' : ''}`}
          onClick={() => setOrderSubTab('open')}
        >
          Открытые заказы
        </button>
        <button
          className={`subtab ${orderSubTab === 'closed' ? 'active' : ''}`}
          onClick={() => setOrderSubTab('closed')}
        >
          Закрытые заказы
        </button>
      </div>

      <div className="ps-orders-list">
        {orders
          .filter(order => {
            if (orderSubTab === 'open') {
              return ['Открыт', 'Ожидает оплаты'].includes(order.status);
            }
            return ['Оплачен', 'Отменен'].includes(order.status);
          })
          .map(order => (
            <div 
              key={order.id}
              className={`ps-order-card ${order.status.toLowerCase().replace(' ', '-')}`}
              onClick={() => setSelectedOrder(order)}ы
            >
              <div className="ps-order-header">
                <span>Заказ #{order.id.slice(0,6)}</span>
                <span>Стол: {order.tableId.split('_')[1]}</span>
              </div>
              <div className="ps-order-info">
                <span>{order.customers.length} посетителя</span>
                <span>{order.totalPrice}₽</span>
                <span className={`ps-status ps-status-${order.status.toLowerCase().replace(' ', '-')}`}>
                  {order.status}
                </span>
              </div>
            </div>
          ))}
      </div>

      {selectedOrder && (
        <div className="ps-order-modal">
          <div className="ps-modal-header">
            <div className="ps-modal-top-buttons">
              <button
                className="ps-action-btn"
                onClick={() => setShowCheckout(true)}
              >
                Состав
              </button>
              <button
                className="ps-action-btn"
                onClick={() => setShowCheckout(false)}
              >
                Параметры
              </button>
            </div>
            <h3>Заказ #{selectedOrder.id.slice(0,6)} (Стол {selectedOrder.tableId.split('_')[1]})</h3>
            <button onClick={() => {
              setSelectedOrder(null);
              setShowCheckout(false);
            }}>×</button>
          </div>

        {!showCheckout ? (
          <>
            <div className="ps-order-details">
              <div className="ps-detail-item">
                <span>Время создания:</span>
                <span>{selectedOrder.createdAt}</span>
              </div>
              <div className="ps-detail-item">
                <span>Количество гостей:</span>
                <span>{selectedOrder.guests}</span>
              </div>
              <div className="ps-detail-item">
                <span>Общая сумма:</span>
                <span>{selectedOrder.totalPrice}₽</span>
              </div>
              <span>Статус заказа:</span>
              <div className="ps-item-status">
                  <select
                    value={selectedOrder.status}
                    onChange={(e) => updateOrderStatus(selectedOrder.id, e.target.value)}
                  >
                    {ORDER_STATUSES.map(status => (
                      <option key={status} value={status}>{status}</option>
                    ))}
                  </select>
                  </div>
            </div>

            <div className="ps-modal-actions">
              <button
                className="ps-action-btn paid"
                onClick={() => {
                  updateOrderStatus(selectedOrder.id, 'Оплачен');
                  setSelectedOrder(null);
                }}
              >
                Оплачен
              </button>
              <button
                className="ps-action-btn cancel"
                onClick={() => {
                  updateOrderStatus(selectedOrder.id, 'Открыт');
                  setSelectedOrder(null);
                }}
              >
                Отмена
              </button>
            </div>
          </>
        ) : (
            <div className="ps-checkout-screen">
              {selectedOrder.customers.map(customer => (
                <div key={customer.id} className="ps-customer-section">
                  <div className="ps-customer-header">
                    <h4>{customer.name}</h4>
                    <span className="ps-customer-instagram">{customer.instagram}</span>
                  </div>
                  
                  {customer.items.map(item => (
                    <div key={item.id} className="ps-order-item">
                      <div className="ps-item-info">
                        <span className="ps-item-name">{item.name}</span>
                        <span className="ps-item-price">{item.amount}x {item.price}₽</span>
                      </div>
                      
                      <div className="ps-item-status">
                        <select
                          value={item.status}
                          onChange={(e) => 
                            updateItemStatus(
                              selectedOrder.id, 
                              customer.id,
                              item.id, 
                              e.target.value
                            )
                          }
                        >
                          {STATUS_FLOW.map(status => (
                            <option key={status} value={status}>{status}</option>
                          ))}
                        </select>
                        <button
                          className="ps-status-btn cancel"
                          onClick={() => 
                            updateItemStatus(
                              selectedOrder.id,
                              customer.id, 
                              item.id, 
                              'Отменен'
                            )
                          }
                        >
                          Отменить
                        </button>
                      </div>
                    </div>
                  ))}
                </div>
              ))}
            </div>
          )}
        </div>
      )}
    </div>
  )}
      </div>

      {showAddModal && tab === "menu" && (
        <div className="ps-backdrop">
          <div className="ps-modal">
            <h3>Добавить блюдо в меню</h3>
            
            {!selectedDish ? (
              <>
                <button 
                  className="ps-add-button"
                  onClick={() => {
                    loadAvailableDishes();
                    setShowDishPicker(true);
                  }}
                >
                  Выбрать блюдо
                </button>

                <div className="ps-modal-buttons">
                  <button 
                    className="ps-button"
                    onClick={handleAddMenuItem}
                  >
                    Добавить
                  </button>
                  <button
                    className="ps-button ps-button-cancel"
                    onClick={() => setShowAddModal(false)}
                  >
                    Отмена
                  </button>
                </div>
                
                {showDishPicker && (
                  <div className="dish-picker">
                    {availableDishes.map(d => (
                      <div 
                        key={d.uuid} 
                        className="dish-card"
                        onClick={() => {
                          setSelectedDish(d);
                          setShowDishPicker(false);
                        }}
                      >
                        <div className="preview-image">
                          {d.imageUrl && <img src={d.imageUrl} alt={d.name} />}
                        </div>
                        <div className="dish-details">
                          <h4>{d.name}</h4>
                          <p>{d.description}</p>
                          <p>{d.calories} ккал • {d.weight}г</p>
                        </div>
                      </div>
                    ))}
                  </div>
                )}
              </>
            ) : (
              <>
                <div className="selected-dish">
                  <span>{selectedDish.name}</span>
                  <button 
                    className="ps-clear-button"
                    onClick={() => setSelectedDish(null)}
                  >
                    ❌
                  </button>
                </div>
                
                <input
                  type="number"
                  className="ps-input"
                  placeholder="Цена"
                  value={price}
                  onChange={(e) => setPrice(e.target.value)}
                  min="1"
                />
                
                <div className="ps-modal-buttons">
                  <button 
                    className="ps-button"
                    onClick={handleAddMenuItem}
                  >
                    Добавить
                  </button>
                  <button
                    className="ps-button ps-button-cancel"
                    onClick={() => setShowAddModal(false)}
                  >
                    Отмена
                  </button>
                </div>
              </>
            )}
            {error && <div className="ps-error-text">{error}</div>}
          </div>
        </div>
      )}
      {showEditStaffModal && (
        <div className="ps-backdrop">
          <div className="ps-modal">
            <h3>Редактировать сотрудника</h3>
            <div>
              <span>{editStaffData?.first_name} {editStaffData?.last_name} ({editStaffData?.login})</span>
            </div>
            <select
              className="ps-role-select"
              value={role}
              onChange={e => setRole(e.target.value)}
            >
              <option value="" disabled>
                Выберите роль
              </option>
              <option value="Админ">Админ</option>
              <option value="Официант">Официант</option>
            </select>
            <div className="ps-modal-buttons">
              <button className="ps-button" onClick={handleUpdateStaff}>
                Сохранить
              </button>
              <button className="ps-button ps-button-cancel" onClick={() => setShowEditStaffModal(false)}>
                Отмена
              </button>
            </div>
          </div>
        </div>
      )}
      {showEditMenuDishModal && (
        <div className="ps-backdrop">
          <div className="ps-modal">
            <h3>Редактировать блюдо</h3>
            <div>
              <span>{editMenuDishData?.name}</span>
            </div>
            <input
              className="ps-input"
              type="number"
              placeholder="Цена"
              value={editMenuDishPrice}
              min="1"
              onChange={e => setEditMenuDishPrice(e.target.value)}
            />
            <div className="ps-modal-buttons">
              <button className="ps-button" onClick={handleUpdateMenuDish}>
                Сохранить
              </button>
              <button className="ps-button ps-button-cancel" onClick={() => setShowEditMenuDishModal(false)}>
                Отмена
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
