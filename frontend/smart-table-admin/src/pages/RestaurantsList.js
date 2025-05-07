import React, { useState, useEffect } from "react";
import DefaultApi from "../api/restaurant_api/generated/src/api/DefaultApi";
import AdminV1RestaurantCreateRequest from "../api/restaurant_api/generated/src/model/AdminV1RestaurantCreateRequest";
import "../styles/RestaurantScreen.css";

export default function RestaurantsList() {
  const [restaurants, setRestaurants] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [newName, setNewName] = useState("");

  const userUuid = localStorage.getItem("user_uuid");
  const jwtToken = localStorage.getItem("jwt_token");
  const api = new DefaultApi();
  api.apiClient.basePath = "https://2663-2a01-4f9-c010-ecd2-00-1.ngrok-free.app";

  useEffect(() => {
    if (!userUuid) return;
    api.adminV1RestaurantListGet(userUuid, (err, data) => {
      if (err) console.error(err);
      else     setRestaurants(Array.isArray(data) ? data : []);
    });
  }, [userUuid]);

  const handleCreate = async () => {
    const name = newName.trim();
    if (!name) {
      alert("Введите название ресторана");
      return;
    }

    const req = AdminV1RestaurantCreateRequest.constructFromObject({
      restaurant_name: name,
    });

    try {
      const created = await new Promise((resolve, reject) => {
        api.apiClient.defaultHeaders["JWT-Token"] = jwtToken;
        api.adminV1RestaurantCreatePost(
          userUuid,
          req,
          (err, data) => err ? reject(err) : resolve(data)
        );
      });

      const newRest = {
        restaurant_uuid: created.restaurant_uuid,
        restaurant_name: name
      };

      setRestaurants((prev) => [...prev, newRest]);
      setNewName("");
      setShowModal(false);
    } catch (err) {
      console.error("Не удалось создать ресторан", err);
      alert(err.message || "Ошибка создания");
    }
  };

  return (
    <div className="rest-container">
      <div className="rest-header-bar">
        <button className="back-button">Выйти из аккаунта</button>
        <h1 className="header-title">Мои рестораны</h1>
        <button
          className="create-rest-button"
          onClick={() => setShowModal(true)}
        >
          Создать ресторан
        </button>
        <button className="profile-button">𓀡</button>
      </div>

      <div className="rest-list">
        {restaurants.length === 0 ? (
          <p className="no-rest">Нет ресторанов</p>
        ) : (
          restaurants.map((r) => (
            <button
              key={r.restaurant_uuid}
              className="rest-item"
              onClick={() => alert("Откроем плейсы")}
            >
              {r.restaurant_name}
            </button>
          ))
        )}
      </div>

      {showModal && (
        <div className="modal-backdrop">
          <div className="modal">
            <h3>Название ресторана</h3>
            <div class="input-container">
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
    </div>
  );
}