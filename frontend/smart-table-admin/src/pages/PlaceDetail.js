import React, { useState, useEffect, useCallback, useMemo } from "react";
import { handleMultipartResponse } from '../components/multipartUtils';
import { useParams, useNavigate } from "react-router-dom";
import { QRCodeSVG } from "qrcode.react";
import { toPng } from "html-to-image";

import DefaultApi from "../api/place_api/generated/src/api/DefaultApi";
import AdminV1PlaceEmployeeAddRequest from "../api/place_api/generated/src/model/AdminV1PlaceEmployeeAddRequest";
import AdminV1PlaceMenuDishCreateRequest from "../api/place_api/generated/src/model/AdminV1PlaceMenuDishCreateRequest";
import AdminV1PlaceTableDeepLinksListRequest from "../api/place_api/generated/src/model/AdminV1PlaceTableDeepLinksListRequest";

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
  const [bulkStatus, setBulkStatus] = useState("");
  const isLocked = LOCKED_ORDER_STATUSES.includes(orderStatus) || LOCKED_DISH_STATUSES.includes(item.status) || isClosedOrder;

  const renderComment = () =>
    item.comment ? (
      <div className="ps-item-comment">
        <span>Комментарий: {item.comment}</span>
      </div>
    ) : null;

  const uuids = item.item_uuid_list;
  const handleBulkStatusChange = async (status) => {
    setBulkStatus(status);
    for (let uuid of uuids) {
      await editOrderItemStatus(uuid, status);
    }
  };

  if (item.count === 1) {
    return (
      <div className="ps-order-item">
        <div className="ps-item-info">
          <span className="ps-item-name">{item.name}</span>
          <span className="ps-item-price">{item.count}x {item.item_price}&nbsp;&#8381;</span>
          <span className="ps-item-total-price">Итого: {item.result_price}&nbsp;&#8381;</span>
          {renderComment()}
        </div>
        <div className="ps-item-status">
          <select
            value={item.status}
            className="ps-status-select-wide"
            disabled={isLocked}
            onChange={e =>
              !isLocked && editOrderItemStatus(item.item_uuid_list[0], e.target.value)
            }
          >
            {DISH_STATUS_FLOW.map(status => (
              <option key={status} value={status}>
                {DISH_STATUS_MAP[status]}
              </option>
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
        <span className="ps-item-price ps-order-group-price">
          {item.count}x {item.item_price}&nbsp;&#8381;
        </span>
        <div className="ps-order-group-right">
          <span className="ps-item-total-price ps-order-group-total">
            Итого: {item.result_price}&nbsp;&#8381;
          </span>
          <div className="ps-order-group-actions">
            <div className="ps-bulk-status-selector">
              <select
                value={bulkStatus}
                className="ps-status-select-wide"
                disabled={isLocked}
                onChange={e => handleBulkStatusChange(e.target.value)}
              >
                <option value="">Изменить статус всем...</option>
                {DISH_STATUS_FLOW.map(status => (
                  <option key={status} value={status}>
                    {DISH_STATUS_MAP[status]}
                  </option>
                ))}
              </select>
            </div>
            <button
              className="ps-expand-btn ps-order-group-expand"
              onClick={() => setOpen((o) => !o)}
            >
              {open ? "Скрыть блюда" : "Показать все"}
            </button>
          </div>
        </div>
      </div>
      {renderComment()}

      {open && (
        <div className="ps-order-item-list">
          {item.item_uuid_list.map((uuid, idx) => (
            <div
              key={`${item.menu_dish_uuid}-${uuid}-${idx}`}
              className="ps-order-item-single ps-order-group-single"
            >
              <span className="ps-item-name">{item.name} #{idx + 1}</span>
              <span>{item.item_price}&nbsp;&#8381;</span>
              <div>
                <select
                  value={item.status}
                  className="ps-status-select-wide"
                  disabled={isLocked}
                  onChange={e =>
                    !isLocked && editOrderItemStatus(uuid, e.target.value)
                  }
                >
                  {DISH_STATUS_FLOW.map(status => (
                    <option key={status} value={status}>
                      {DISH_STATUS_MAP[status]}
                    </option>
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

  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [dishToDelete, setDishToDelete] = useState(null);

  const [orderSubTab, setOrderSubTab] = useState('open');
  const [orders, setOrders] = useState([]);
  const [selectedOrder, setSelectedOrder] = useState(null);
  const [showCheckout, setShowCheckout] = useState(false);

  const [selectedOrderUuid, setSelectedOrderUuid] = useState(null);

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
    let fastList = [];
    let imagesMap = {};
    try {
      const fastResp = await fetch(`${SERVER_URL}/admin/v1/restaurant/dish/info/list`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true",
        },
        body: JSON.stringify({ restaurant_uuid: restaurant_uuid }),
      });
      if (!fastResp.ok) throw new Error("Ошибка получения блюд (fast)");
      const fastData = await fastResp.json();
      fastList = (fastData.dish_list || []).map(d => ({ ...d, imageUrl: null }));
      setAvailableDishes(fastList);

      const slowResp = await fetch(`${SERVER_URL}/admin/v1/restaurant/dish/list`, {
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
      const { list, imagesMap: slowImagesMap } = await handleMultipartResponse(slowResp, 'dish_list');
      imagesMap = slowImagesMap;
      setAvailableDishes(list.map(d => ({ ...d, imageUrl: imagesMap[d.picture_key] || null })));
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
    let fastList = [];
    let imagesMap = {};
    try {
      const fastResp = await fetch(`${SERVER_URL}/admin/v1/place/menu/dish/info/list`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true",
        },
        body: JSON.stringify({ place_uuid: place_uuid }),
      });
      if (!fastResp.ok) throw new Error("Ошибка получения меню (fast)");
      const fastData = await fastResp.json();
      fastList = (fastData.menu_dish_list || []).map(d => ({ ...d, imageUrl: null }));
      setMenuDishes(fastList);
      setLoading(false);

      const slowResp = await fetch(`${SERVER_URL}/admin/v1/place/menu/dish/list`, {
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
      const { list, imagesMap: slowImagesMap } = await handleMultipartResponse(slowResp);
      imagesMap = slowImagesMap;
      setMenuDishes(list.map(d => ({ ...d, imageUrl: imagesMap[d.picture_key] || null, price: d.price })));
    } catch (e) {
      console.error("Ошибка загрузки блюд:", e);
      setMenuDishes([]);
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

  const loadOrderDetails = useCallback(async (order_uuid, place_uuid, cancelToken = { canceled: false }) => {
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
      if (!cancelToken.canceled) {
        setSelectedOrder(prev =>
          !prev || !prev.order_main_info || prev.order_main_info.uuid === (data.order_info.order_main_info?.uuid)
            ? { ...data.order_info, order_main_info: { ...data.order_info.order_main_info, status: ORDER_STATUS_MAP[data.order_info.order_main_info.status] || data.order_info.order_main_info.status } }
            : prev
        );
      }
    } catch (e) {
      let msg = e.body?.message || e.message || "Ошибка получения деталей заказа";
      setError(msg);
    } finally {
      setLoading(false);
    }
  }, [userUUID, jWTToken, setSelectedOrder, setLoading, setError]);

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
      await loadOrderDetails(order_uuid, place_uuid);
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

  useEffect(() => {
    if (tab !== "orders") return;
    loadOrders();
    const interval = setInterval(() => {
      loadOrders();
    }, 5000);
    return () => clearInterval(interval);
  }, [tab, place_uuid, orderSubTab, loadOrders]);

  useEffect(() => {
    if (tab !== "orders" || !selectedOrderUuid) return;
    const cancelToken = { canceled: false };
    loadOrderDetails(selectedOrderUuid, place_uuid, cancelToken);
    const interval = setInterval(() => {
      loadOrderDetails(selectedOrderUuid, place_uuid, cancelToken);
    }, 5000);
    return () => {
      cancelToken.canceled = true;
      clearInterval(interval);
    };
  }, [tab, selectedOrderUuid, loadOrderDetails, place_uuid]);

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
    setDishToDelete(dish);
    setShowDeleteConfirm(true);
  };

  const confirmDelete = async () => {
    if (!dishToDelete) return;
  
    try {
      const resp = await fetch(`${SERVER_URL}/admin/v1/place/menu/dish/delete`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Accept": "application/json",
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true",
        },
        body: JSON.stringify({
          place_uuid: place_uuid,           
          menu_dish_uuid: dishToDelete.id,        
        }),
      });
  
      if (resp.status === 204) {
        await loadMenuDishes();
        setShowDeleteConfirm(false);
        setDishToDelete(null);
        return;
      }
  
      let errMsg = "Ошибка удаления блюда";
      try {
        const errData = await resp.json();
        errMsg = errData?.message || errMsg;
      } catch (_) {}
      alert(errMsg);
    } catch (e) {
      alert(e.message || "Ошибка удаления блюда");
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
            setShowDeleteConfirm(false);
            setSelectedDish(null);
            setPrice("");
            setError("");
          }}>
            Сгенерировать QR-код
          </button>
        )}
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
                    <div key={dish.uuid || dish.id} className="menu-item">
                      <div className="dish-image">
                        {dish.imageUrl ? (
                          <img src={dish.imageUrl} alt={dish.name} />
                        ) : (
                          <div className="shimmer shimmer-rect" />
                        )}
                      </div>
                      <div className="dish-info">
                        <div className="dish-header">
                          <h3>{dish.name}</h3>
                        </div>
                        <p>{dish.description}</p>
                        <p>Категория: {dish.category}</p>
                        <p>{dish.calories} ккал, {dish.weight} г.</p>
                        <span className="price-tag">{dish.price}&nbsp;&#8381;</span>
                        <div className="dish-actions">
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
                        onClick={() => {
                          setSelectedOrderUuid(order.uuid);
                          setShowCheckout(false);
                        }}
                      >
                        <div className="ps-order-header">
                          <span>Заказ #{order.uuid.slice(0,6)}</span>
                          <span>Стол: {order.table_number}</span>
                        </div>
                        <div className="ps-order-info">
                          <span>{order.guests_count} посетителя</span>
                          <span>{order.total_price}&nbsp;&#8381;</span>
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
                    setSelectedOrderUuid(null);
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
                    <span>{selectedOrder.order_main_info.total_price}&nbsp;&#8381;</span>
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
                              <span className="ps-item-total-price">Итоговая цена: {customer.total_price}&nbsp;&#8381;</span>
                              <span className="ps-customer-instagram" style={{ marginTop: 4 }}>{customer.tg_id}</span>
                            </div>
                          </div>
                          {customer.item_group_list.map((item, idx) => (
                            <OrderItemGroup
                            key={`${selectedOrder.order_main_info.uuid}-${customer.uuid}-${item.menu_dish_uuid}-${idx}`}
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
                    {availableDishes.length === 0 && (
                      <div style={{display:'flex',gap:'1rem'}}>
                        {[...Array(4)].map((_,i) => (
                          <div key={i} className="dish-card">
                            <div className="preview-image shimmer shimmer-rect" style={{height:80}} />
                            <div className="dish-details" />
                          </div>
                        ))}
                      </div>
                    )}
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
                          {d.imageUrl
                            ? <img src={d.imageUrl} alt={d.name} />
                            : <div className="shimmer shimmer-rect" style={{height:80}} />}
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

      {showDeleteConfirm && dishToDelete && (
        <div className="ps-backdrop">
          <div className="ps-modal delete-confirm-modal">
            <h3>Подтверждение удаления</h3>
            <p>Вы уверены, что хотите удалить блюдо <strong>"{dishToDelete.name}"</strong> из меню?</p>
            <div className="ps-modal-buttons">
              <button
                className="ps-button ps-button-delete-confirm"
                onClick={confirmDelete}
              >
                Удалить
              </button>
              <button
                className="ps-button ps-button-cancel"
                onClick={() => {
                  setShowDeleteConfirm(false);
                  setDishToDelete(null);
                }}
              >
                Отмена
              </button>
            </div>
          </div>
        </div>
      )}

      {selectedOrderUuid && !selectedOrder && (
        <div className="ps-order-modal"><div className="ps-loading">Загрузка заказа...</div></div>
      )}
    </div>
  );
}
