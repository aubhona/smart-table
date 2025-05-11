import React, { useState, useEffect } from "react";
import { handleMultipartResponse } from './multipartUtils';
import { useParams, useNavigate } from "react-router-dom";
import { QRCodeSVG } from "qrcode.react";
import { toPng } from "html-to-image";

import DefaultApi from "../api/place_api/generated/src/api/DefaultApi";
import AdminV1PlaceEmployeeAddRequest from "../api/place_api/generated/src/model/AdminV1PlaceEmployeeAddRequest";
import AdminV1PlaceMenuDishCreateRequest from "../api/place_api/generated/src/model/AdminV1PlaceMenuDishCreateRequest";
import AdminV1PlaceTableDeepLinksListRequest from "../api/place_api/generated/src/model/AdminV1PlaceTableDeepLinksListRequest";

import "../styles/PlaceScreen.css";

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

  const userUUID = localStorage.getItem("user_uuid");
  const jWTToken = localStorage.getItem("jwt_token");
  const api = new DefaultApi();

  api.apiClient.basePath = "https://87d6-2a01-4f9-c010-ecd2-00-1.ngrok-free.app";
  api.apiClient.defaultHeaders = {
    "User-UUID": userUUID,
    "JWT-Token": jWTToken,
    "ngrok-skip-browser-warning": "true",
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
    const resp = await fetch("https://87d6-2a01-4f9-c010-ecd2-00-1.ngrok-free.app/admin/v1/restaurant/dish/list", {
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
        "https://87d6-2a01-4f9-c010-ecd2-00-1.ngrok-free.app/admin/v1/place/employee/list",
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
    const resp = await fetch("https://87d6-2a01-4f9-c010-ecd2-00-1.ngrok-free.app/admin/v1/place/menu/dish/list", {
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

  return (
    <div className="ps-container">
      <div className="ps-header-bar">
        <button className="ps-back-button" onClick={() => navigate(-1)}>
          Назад
        </button>
        <h1 className="ps-title">Ресторан: {restaurantName},</h1>
        <h1 className="ps-title">Адрес: {placeName}</h1>

        {tab === "staff" && (
          <button
            className="ps-create-button"
            onClick={() => setShowAddModal(true)}
          >
            Добавить сотрудника
          </button>
        )}
        {tab === "menu" && (
          <button className="ps-create-button"
          onClick={() => setShowAddModal(true)}
          >
            Добавить блюдо
          </button>
        )}
        {tab === "orders" && (
          <button className="ps-create-button">
            Сгенерировать QR-код
          </button>
        )}

        <button className="ps-profile-button">𓀡</button>
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
                    className="ps-role-input"
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
                        <h3>{dish.name}</h3>
                        <p>{dish.description}</p>
                        <p>Категория: {dish.category}</p>
                        <p>{dish.calories} ккал, {dish.weight} г.</p>
                        <div className="price-tag">{dish.price} ₽</div>
                      </div>
                    </div>
                  ))
                )}
              </div>
            )}
          </>
        )}
      
        {tab === "orders" && (
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
                    className="ps-button"
                    onClick={() => downloadQR(index)}
                  >
                    Сохранить PNG
                  </button>
                </div>
              </div>
            ))}
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
    </div>
  );
}
