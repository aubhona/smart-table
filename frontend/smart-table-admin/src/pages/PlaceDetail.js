import React, { useState, useEffect, useCallback, useMemo } from "react";
import { handleMultipartResponse } from '../components/multipartUtils';
import { useParams, useNavigate } from "react-router-dom";
import { QRCodeSVG } from "qrcode.react";
import { toPng } from "html-to-image";

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

const ORDER_STATUS_MAP = {
  new: "Открыт",
  payment_waiting: "Ожидает оплаты",
  paid: "Оплачен",
  canceled_by_service: "Отменен"
};

const ORDER_REVERSE_STATUS_MAP = Object.fromEntries(
  Object.entries(ORDER_STATUS_MAP).map(([eng, rus]) => [rus, eng])
);

const DISH_STATUS_FLOW = ['accepted', 'cooking', 'cooked', 'served', 'canceled_by_service'];
const DISH_STATUS_MAP = {
  accepted: 'Принят',
  cooking: 'Готовится',
  cooked: 'Готов',
  served: 'Подан',
  canceled_by_service: 'Отменен'
};

const LOCKED_ORDER_STATUSES = ['paid', 'canceled_by_service'];
const LOCKED_DISH_STATUSES = ['served', 'canceled_by_service'];

const OrderItemGroup = ({ item, orderStatus, editOrderItemStatus, isClosedOrder }) => {
  const [open, setOpen] = useState(false);
  const isLocked = LOCKED_ORDER_STATUSES.includes(orderStatus) || LOCKED_DISH_STATUSES.includes(item.status) || isClosedOrder;

  if (item.count === 1) {
    return (
      <div className="ps-order-item">
        <div className="ps-item-info">
          <span className="ps-item-name">{item.name}</span>
          <span className="ps-item-price">{item.count}x {item.item_price}₽</span>
          <span className="ps-item-total-price">Итого: {item.result_price}₽</span>
        </div>
        <div className="ps-item-status">
          <select
            value={item.status}
            className="ps-status-select-wide"
            disabled={isLocked}
            onChange={e => !isLocked && editOrderItemStatus(item.item_uuid_list[0], e.target.value)}
          >
            {DISH_STATUS_FLOW.map(status => (
              <option key={status} value={status}>{DISH_STATUS_MAP[status]}</option>
            ))}
          </select>
        </div>
      </div>
    );
  }

  return (
    <div className="ps-order-item-group">
      <div className="ps-item-info ps-order-group-info">
        <span className="ps-item-name ps-order-group-name">{item.name}</span>
        <span className="ps-item-price ps-order-group-price">{item.count}x {item.item_price}₽</span>
        <div className="ps-order-group-right">
          <span className="ps-item-total-price ps-order-group-total">Итого: {item.result_price}₽</span>
          <button
            className="ps-expand-btn ps-order-group-expand"
            onClick={() => setOpen(o => !o)}
          >
            {open ? "Скрыть блюда" : "Показать все"}
          </button>
        </div>
      </div>
      {open && (
        <div className="ps-order-item-list">
          {item.item_uuid_list.map((uuid, idx) => (
            <div 
              key={uuid} 
              className="ps-order-item-single ps-order-group-single"
            >
              <span className="ps-item-name">{item.name} #{idx + 1}</span>
              <span>{item.item_price}₽</span>
              <div>
                <select
                  value={item.status}
                  className="ps-status-select-wide"
                  disabled={isLocked}
                  onChange={e => !isLocked && editOrderItemStatus(uuid, e.target.value)}
                >
                  {DISH_STATUS_FLOW.map(status => (
                    <option key={status} value={status}>{DISH_STATUS_MAP[status]}</option>
                  ))}
                </select>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

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

  const [orderSubTab, setOrderSubTab] = useState('open');
  const [orders, setOrders] = useState([]);
  const [selectedOrder, setSelectedOrder] = useState(null);
  const [showCheckout, setShowCheckout] = useState(false);

  const userUUID = localStorage.getItem("user_uuid");
  const jWTToken = localStorage.getItem("jwt_token");
  
  const api = useMemo(() => {
    const apiInstance = new DefaultApi();
    apiInstance.apiClient.basePath = SERVER_URL;
    apiInstance.apiClient.defaultHeaders = {
      "User-UUID": userUUID,
      "JWT-Token": jWTToken,
      "ngrok-skip-browser-warning": "true",
    };
    return apiInstance;
  }, [userUUID, jWTToken]);

  const loadDeepLinks = useCallback(async () => {
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
  }, [place_uuid, userUUID, jWTToken, api]);

  const handleGenerateQRCodes = useCallback(async () => {
    try {
      await loadDeepLinks();
    } catch (e) {
      console.error("Ошибка генерации QR-кодов:", e);
    }
  }, [loadDeepLinks]);

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

  const loadOrders = useCallback(async () => {
    setLoading(true);
    setError("");
    try {
      const resp = await fetch(`${SERVER_URL}/admin/v1/place/order/list`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Accept": "application/json",
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true"
        },
        body: JSON.stringify({ 
          is_active: orderSubTab === 'open',
          place_uuid 
        })
      });
      if (!resp.ok) throw resp;

      const data = await resp.json();
      setOrders(
        (data.order_list || []).map(order => ({
          ...order,
          status: ORDER_STATUS_MAP[order.status] || order.status
        }))
      );
    } catch (e) {
      let msg = e.body?.message || e.message || "Ошибка получения заказов";
      setError(msg);
      setOrders([]);
    } finally {
      setLoading(false);
    }
  }, [place_uuid, userUUID, jWTToken, orderSubTab]);

  const loadMenuDishes = useCallback(async () => {
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
  }, [place_uuid, userUUID, jWTToken]);

  const loadStaff = useCallback(async () => {
    setLoading(true);
    try {
      const resp = await fetch(`${SERVER_URL}/admin/v1/place/employee/list`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Accept": "application/json",
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true",
        },
        body: JSON.stringify({ place_uuid: place_uuid }),
      });
      if (!resp.ok) throw new Error(`HTTP ${resp.status}`);
      const data = await resp.json();
      setStaff(data.employee_list || []);
    } catch (e) {
      console.error("Ошибка загрузки сотрудников:", e);
      setStaff([]);
    } finally {
      setLoading(false);
    }
  }, [place_uuid, userUUID, jWTToken]);

  async function loadOrderDetails(order_uuid, place_uuid) {
    setLoading(true);
    setError("");
    try {
      const resp = await fetch(`${SERVER_URL}/admin/v1/place/order/info`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Accept": "application/json",
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true"
        },
        body: JSON.stringify({ order_uuid, place_uuid })
      });
      if (!resp.ok) throw resp;

      const data = await resp.json();
      setSelectedOrder({
        ...data.order_info,
        order_main_info: {
          ...data.order_info.order_main_info,
          status: ORDER_STATUS_MAP[data.order_info.order_main_info.status] || data.order_info.order_main_info.status
        }
      });
    } catch (e) {
      let msg = e.body?.message || e.message || "Ошибка получения деталей заказа";
      setError(msg);
    } finally {
      setLoading(false);
    }
  }

  async function editOrder(order_uuid, order_status, place_uuid, extraParams = {}) {
    setLoading(true);
    setError("");

    if (LOCKED_ORDER_STATUSES.includes(selectedOrder?.order_main_info.status)) {
      setLoading(false);
      return;
    }

    try {
      const payload = {
        place_uuid: place_uuid,
        order_uuid: order_uuid,
        table_number: selectedOrder?.order_main_info?.table_number,
        order_status,
        ...extraParams, 
      };
      const resp = await fetch(`${SERVER_URL}/admin/v1/place/order/edit`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Accept": "application/json",
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true"
        },
        body: JSON.stringify(payload)
      });
      if (!resp.ok) {
        let errText = await resp.text();
        throw new Error(errText || "Ошибка изменения заказа");
      }

      await loadOrders();
      await loadOrderDetails(order_uuid);
    } catch (e) {
      setError(e.body?.message || e.message || "Ошибка редактирования заказа");
    } finally {
      setLoading(false);
    }
  }

  async function editOrderItemStatus(order_uuid, item_uuid, place_uuid, newStatus) {
    if (!selectedOrder) return;

    if (LOCKED_ORDER_STATUSES.includes(selectedOrder.order_main_info.status)) return;

    setLoading(true);
    setError("");
    try {
      const payload = {
        place_uuid: place_uuid,
        order_uuid: order_uuid,
        table_number: selectedOrder?.order_main_info?.table_number,
        item_group: {
          item_uuid_list: [item_uuid],
          item_status: newStatus,
        },
      };

      console.log(payload);

      const resp = await fetch(`${SERVER_URL}/admin/v1/place/order/edit`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Accept": "application/json",
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true",
        },
        body: JSON.stringify(payload),
      });

      if (!resp.ok) throw new Error(`Ошибка изменения статуса позиции: ${resp.status}`);

      await loadOrderDetails(selectedOrder.order_main_info.uuid, place_uuid);
      await loadOrders();
    } catch (e) {
      setError(e.message || "Ошибка при изменении статуса позиции");
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
      loadOrders();
    } else if (tab === "tables") {
      handleGenerateQRCodes();
    }
  }, [place_uuid, tab, handleGenerateQRCodes, loadMenuDishes, loadOrders, loadStaff]);

  // POLLING: для заказов
  useEffect(() => {
    if (tab !== "orders") return;
    // Первый вызов сразу
    loadOrders();
    // Интервал
    const interval = setInterval(() => {
      loadOrders();
    }, 5000); // 5 секунд
    return () => clearInterval(interval);
  }, [tab, place_uuid, orderSubTab, loadOrders]);

  // POLLING: для деталей заказа
  useEffect(() => {
    if (tab !== "orders" || !selectedOrder) return;
    const orderUuid = selectedOrder.order_main_info.uuid;
    loadOrderDetails(orderUuid, place_uuid);
    const interval = setInterval(() => {
      loadOrderDetails(orderUuid, place_uuid);
    }, 5000); // 5 секунд
    return () => clearInterval(interval);
  }, [tab, selectedOrder?.order_main_info?.uuid, place_uuid]);

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
      setError("Укажите цену больше 0");
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
      setError("");
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
          <button className="ps-create-button" onClick={() => {
            setShowAddModal(true);
            setLogin("");
            setRole("");
            setError("");
          }}>
            Добавить сотрудника
          </button>
        )}
        {tab === "menu" && (
          <button className="ps-create-button" onClick={() => {
            setShowAddModal(true);
            setSelectedDish(null);
            setPrice("");
            setError("");
          }}>
            Добавить блюдо
          </button>
        )}
        {tab === "tables" && (
          <button className="ps-create-button" onClick={() => {
            setShowAddModal(false);
            setShowEditStaffModal(false);
            setShowEditMenuDishModal(false);
            setSelectedDish(null);
            setPrice("");
            setError("");
          }}>
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
                        <span className="price-tag">{dish.price}₽</span>
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
                .length === 0 ? (
                  <div className="ps-no-orders">
                    Нет заказов
                  </div>
                ) : (
                  orders
                    .filter(order => order && order.uuid)
                    .filter(order => {
                      if (orderSubTab === 'open') {
                        return ['Открыт', 'Ожидает оплаты'].includes(order.status);
                      }
                      return ['Оплачен', 'Отменен'].includes(order.status);
                    })
                    .map(order => (
                      <div 
                        key={order.uuid}
                        className={`ps-order-card ${order.status.toLowerCase().replace(' ', '-')}`}
                        onClick={async () => {
                          await loadOrderDetails(order.uuid, place_uuid);
                          setShowCheckout(false);
                        }}
                      >
                        <div className="ps-order-header">
                          <span>Заказ #{order.uuid.slice(0,6)}</span>
                          <span>Стол: {order.table_number}</span>
                        </div>
                        <div className="ps-order-info">
                          <span>{order.guests_count} посетителя</span>
                          <span>{order.total_price}₽</span>
                          <span className={`ps-status ps-status-${order.status.toLowerCase().replace(' ', '-')}`}>
                            {order.status}
                          </span>
                        </div>
                      </div>
                    ))
                )}
            </div>

            {selectedOrder && (
              <div className="ps-order-modal">
                <div className="ps-modal-header">
                  <div className="ps-modal-top-buttons">
                    <button
                      className={`ps-action-btn${showCheckout ? ' active' : ''}`}
                      onClick={() => setShowCheckout(true)}
                    >
                      <span className="ps-modal-tab-label">Состав заказа</span>
                    </button>
                    <button
                      className={`ps-action-btn${!showCheckout ? ' active' : ''}`}
                      onClick={() => setShowCheckout(false)}
                    >
                    <span className="ps-modal-tab-label">Параметры заказа</span>
                    </button>
                  </div>
                    <h3>Заказ #{selectedOrder.order_main_info.uuid.slice(0,6)} (Стол {selectedOrder.order_main_info.table_number})</h3>
                  <button onClick={() => {
                    setSelectedOrder(null);
                    setShowCheckout(false);
                  }}>×</button>
                </div>

              {!showCheckout && (
                <>
                  <div className="ps-order-details">
                    <div className="ps-detail-item">
                      <span>Время создания:</span>
                      <span>{selectedOrder.order_main_info.created_at}</span>
                    </div>
                    <div className="ps-detail-item">
                      <span>Количество гостей:</span>
                      <span>{selectedOrder.order_main_info.guests_count}</span>
                    </div>
                    <div className="ps-detail-item">
                      <span>Общая сумма:</span>
                    <span>{selectedOrder.order_main_info.total_price}₽</span>
                    </div>
                    <span>Статус заказа:</span>
                    <div className="ps-item-status">
                      <select
                        value={ORDER_REVERSE_STATUS_MAP[selectedOrder.order_main_info.status] || selectedOrder.order_main_info.status}
                        onChange={async (e) => {
                          if (LOCKED_ORDER_STATUSES.includes(selectedOrder.order_main_info.status) || orderSubTab === 'closed') return;
                          const newRusStatus = e.target.value; 
                          const newEngStatus = ORDER_REVERSE_STATUS_MAP[newRusStatus] || newRusStatus;
                          await editOrder(selectedOrder.order_main_info.uuid, newEngStatus, place_uuid);
                          setSelectedOrder(null);
                          setShowCheckout(false);
                        }}
                        disabled={LOCKED_ORDER_STATUSES.includes(selectedOrder.order_main_info.status) || orderSubTab === 'closed'}
                      >
                        {Object.entries(ORDER_STATUS_MAP).map(([key, label]) => (
                          <option key={key} value={key}>
                            {label}
                          </option>
                        ))}
                      </select>
                    </div>
                  </div>
                </>
              )} 

                {showCheckout && selectedOrder.customer_list && (
                  <div className="ps-checkout-screen">
                    {selectedOrder.customer_list.every(customer => !customer.item_group_list || customer.item_group_list.length === 0) ? (
                      <div className="ps-no-orders">Еще ничего не заказано!</div>
                    ) : (
                      selectedOrder.customer_list.map(customer => (
                        <div key={customer.uuid} className="ps-customer-section">
                          <div className="ps-customer-header">
                          <h4>{customer.tg_login}</h4>
                            <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-end', minWidth: 100 }}>
                              <span className="ps-item-total-price">Итоговая цена: {customer.total_price}₽</span>
                              <span className="ps-customer-instagram" style={{ marginTop: 4 }}>{customer.tg_id}</span>
                            </div>
                          </div>
                          {customer.item_group_list.map(item => (
                            <OrderItemGroup
                            key={item.menu_dish_uuid}
                            item={item}
                            orderStatus={selectedOrder.order_main_info.status}
                            editOrderItemStatus={(item_uuid, newStatus) => 
                              editOrderItemStatus(
                                selectedOrder.order_main_info.uuid,
                                item_uuid,
                                place_uuid,
                                newStatus
                              )
                            }
                            isClosedOrder={orderSubTab === 'closed'}
                          />
                          ))}
                        </div>
                      ))
                    )}
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
                    setSelectedDish(null);
                    setPrice("");
                    setError("");
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
                    onClick={() => {
                      setShowDishPicker(false);
                      setShowAddModal(false);
                      setSelectedDish(null);
                      setPrice("");
                      setError("");
                    }}
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
                    onClick={() => {
                      setShowAddModal(false);
                      setSelectedDish(null);
                      setPrice("");
                      setError("");
                    }}
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
