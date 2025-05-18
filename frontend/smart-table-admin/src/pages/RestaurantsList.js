import React, { useState, useEffect } from "react";
import DefaultApi from "../api/restaurant_api/generated/src/api/DefaultApi";
import AdminV1RestaurantCreateRequest from "../api/restaurant_api/generated/src/model/AdminV1RestaurantCreateRequest";
import "../styles/RestaurantScreen.css";
import AdminV1RestaurantDeleteRequest from "../api/restaurant_api/generated/src/model/AdminV1RestaurantDeleteRequest";
import AdminV1RestaurantEditRequest from "../api/restaurant_api/generated/src/model/AdminV1RestaurantEditRequest";
import { SERVER_URL } from "../config";

export default function RestaurantsList() {
  const [restaurants, setRestaurants] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [newName, setNewName] = useState("");
  const [, setError] = useState("");
  const [, setLoading] = useState(true);

  const [showEditModal, setShowEditModal] = useState(false);
  const [editRestaurantUuid, setEditRestaurantUuid] = useState("");
  const [editRestaurantName, setEditRestaurantName] = useState("");

  const userUUID = localStorage.getItem("user_uuid");
  const jWTToken = localStorage.getItem("jwt_token");
  const api = new DefaultApi();
  api.apiClient.basePath = SERVER_URL;

  async function fetchRestaurants() {
    const resp = await fetch(
      `${api.apiClient.basePath}/admin/v1/restaurant/list`,
      {
        method: "GET",
        headers: {
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true"
        },
      }
    );
    if (!resp.ok) throw new Error(`HTTP ${resp.status}`);
    const data = await resp.json();
    return data.restaurant_list;
  }
  
  useEffect(() => {
    (async () => {
      try {
        const list = await fetchRestaurants();
        setRestaurants(list.map(r => ({
          restaurant_uuid: r.uuid,
          restaurant_name: r.name,
        })));
      } catch (e) {
        console.error(e);
        setError("Ошибка загрузки списка ресторанов");
      } finally {
        setLoading(false);
      }
    })();
  }, []);

  const handleCreate = async () => {
    const name = newName.trim();
    if (!name) {
      setError("Введите название ресторана");
      return;
    }
  
    try {
      const request = AdminV1RestaurantCreateRequest.constructFromObject({
        restaurant_name: name
      });
  
      api.apiClient.defaultHeaders = {
        ...api.apiClient.defaultHeaders,
        "User-UUID": userUUID,
        "JWT-Token": jWTToken,
        "Content-Type": "application/json",
        "ngrok-skip-browser-warning": "true"
      };
  
      const data = await new Promise((resolve, reject) => {
        api.adminV1RestaurantCreatePost(
          userUUID,
          jWTToken,
          request,
          (err, data, res) => {
            if (err) return reject(err);
            resolve(data);   
          }
        );
      });
  
      if (data.restaurant_uuid) {
        setRestaurants(prev => [
          ...prev,
          {
            restaurant_uuid: data.restaurant_uuid,
            restaurant_name: name, 
          },
        ]);
        setNewName("");
        setShowModal(false);
        setError("");
      }
  
    } catch (err) {
      console.error("Ошибка создания:", err);
      const errorMsg = err.body?.message || err.message || "Ошибка при создании ресторана";
      setError(errorMsg);
  
      if (err.body?.code === "already_exist") {
        setNewName("");
      }
    }
  };

  const handleDeleteRestaurant = async (restaurant_uuid) => {
    if (!window.confirm("Удалить ресторан? Это действие необратимо!")) return;
    try {
      const req = AdminV1RestaurantDeleteRequest.constructFromObject({
        restaurant_uuid
      });
      await new Promise((res, rej) =>
        api.adminV1RestaurantDeletePost(userUUID, jWTToken, req, (err) =>
          err ? rej(err) : res()
        )
      );
      setRestaurants(prev => prev.filter(r => r.restaurant_uuid !== restaurant_uuid));
    } catch (e) {
      alert(e.body?.message || e.message || "Ошибка удаления ресторана");
    }
  };

  const handleEditRestaurant = async () => {
    if (!editRestaurantName.trim()) {
      setError("Введите название");
      return;
    }
    try {
      const req = AdminV1RestaurantEditRequest.constructFromObject({
        restaurant_uuid: editRestaurantUuid,
        restaurant_name: editRestaurantName
      });
      await new Promise((res, rej) =>
        api.adminV1RestaurantEditPost(userUUID, jWTToken, req, (err) =>
          err ? rej(err) : res()
        )
      );
      setRestaurants(prev =>
        prev.map(r =>
          r.restaurant_uuid === editRestaurantUuid
            ? { ...r, restaurant_name: editRestaurantName }
            : r
        )
      );
      setShowEditModal(false);
    } catch (e) {
      alert(e.body?.message || e.message || "Ошибка редактирования ресторана");
    }
  };

  const openEditModal = (r) => {
    setEditRestaurantUuid(r.restaurant_uuid);
    setEditRestaurantName(r.restaurant_name);
    setShowEditModal(true);
  };

  const handleLogout = () => {
    localStorage.removeItem("user_uuid");
    localStorage.removeItem("jwt_token");
    window.location.href = "/";
  };
  
  return (
    <div className="rest-container">
      <div className="rest-header-bar">
        <button className="back-button" onClick={handleLogout}>Выйти из аккаунта</button>
        <h1 className="header-title">Мои рестораны</h1>
        <button
          className="create-rest-button"
          onClick={() => setShowModal(true)}
        >
          Создать ресторан
        </button>
        <button className="icon-button profile">
          <span className="material-icons">account_circle</span>
        </button>
      </div>

      <div className="rest-list">
        {restaurants.length === 0 ? (
          <p className="no-rest">Нет ресторанов</p>) : (
            restaurants.map((r) => (
              <button
                key={r.restaurant_uuid}
                className="rest-item"
                onClick={() => {
                  localStorage.setItem("current_restaurant", JSON.stringify({
                    restaurant_uuid: r.restaurant_uuid,
                    restaurant_name: r.restaurant_name
                  }));
                  window.location.href = `/restaurants/${r.restaurant_uuid}/places-dishes`
                }}
              >
                <span className="rest-name">{r.restaurant_name}</span>
                <div className="rest-actions">
                  <button
                    className="icon-button edit"
                    onClick={(e) => {
                      e.stopPropagation();
                      openEditModal(r);
                    }}
                  >
                    <span className="material-icons">edit</span>
                  </button>
                  <button
                    className="icon-button delete"
                    onClick={(e) => {
                      e.stopPropagation();
                      handleDeleteRestaurant(r.restaurant_uuid);
                    }}
                  >
                    <span className="material-icons">delete</span>
                  </button>
                </div>
              </button>
            ))
          )}
      </div>

      {showModal && (
        <div className="modal-backdrop">
          <div className="modal">
            <h3>Название ресторана</h3>
            <div className="input-container">
              <input
                value={newName}
                onChange={(e) => setNewName(e.target.value)}
                placeholder="Введите название"
              />
              </div>
            <div className="modal-buttons">
              <button className="pill-button" onClick={handleCreate}>
                Создать
              </button>
              <button className="pill-button" onClick={() => setShowModal(false)}>
                Отмена
              </button>
            </div>
          </div>
        </div>
      )}
      {showEditModal && (
        <div className="modal-backdrop">
          <div className="modal">
            <h3>Редактировать ресторан</h3>
            <div className="input-container">
              <input
                value={editRestaurantName}
                onChange={e => setEditRestaurantName(e.target.value)}
                placeholder="Введите новое название"
              />
            </div>
            <div className="modal-buttons">
              <button className="pill-button" onClick={handleEditRestaurant}>Сохранить</button>
              <button className="pill-button" onClick={() => setShowEditModal(false)}>Отмена</button>
            </div>
            </div>
      </div>
      )}
    </div>
  );
}