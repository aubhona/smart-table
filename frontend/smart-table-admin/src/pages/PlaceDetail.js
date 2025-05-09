import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import DefaultApi from "../api/place_api/generated/src/api/DefaultApi";
import AdminV1PlaceEmployeeListRequest from "../api/place_api/generated/src/model/AdminV1PlaceEmployeeListRequest";
import AdminV1PlaceEmployeeAddRequest from "../api/place_api/generated/src/model/AdminV1PlaceEmployeeAddRequest";
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
  const [loadingStaff, setLoadingStaff] = useState(false);
  const [showAddModal, setShowAddModal] = useState(false);

  const [login, setLogin] = useState("");
  const [role, setRole] = useState("");
  const [error, setError] = useState("");

  const userUUID = localStorage.getItem("user_uuid");
  const jWTToken = localStorage.getItem("jwt_token");
  const api = new DefaultApi();

  api.apiClient.basePath = "https://b04d-2a01-4f9-c010-ecd2-00-1.ngrok-free.app";
  api.apiClient.defaultHeaders = {
    "User-UUID": userUUID,
    "JWT-Token": jWTToken,
    "ngrok-skip-browser-warning": "true",
  };

  async function loadStaff() {
    setLoadingStaff(true);
    try {
      const resp = await fetch(
        "https://b04d-2a01-4f9-c010-ecd2-00-1.ngrok-free.app/admin/v1/place/employee/list",
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
      setLoadingStaff(false);
    }
  }

  useEffect(() => {
    if (tab === "staff") {
      loadStaff();
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
          <button className="ps-create-button">Добавить блюдо</button>
        )}
        {tab === "orders" && (
          <button className="ps-create-button">Сгенерировать QR-код</button>
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
            {loadingStaff && <p className="ps-loading">Загрузка…</p>}
            {!loadingStaff && staff.length === 0 && (
              <p className="ps-empty">Нет сотрудников</p>
            )}
            {!loadingStaff &&
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
                  <input
                    className="ps-input"
                    placeholder="Роль"
                    value={role}
                    onChange={(e) => setRole(e.target.value)}
                  />
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

        {tab === "menu" && <p className="ps-placeholder">Menu content…</p>}
        {tab === "orders" && <p className="ps-placeholder">Orders content…</p>}
      </div>
    </div>
  );
}
