import React, { useState, useEffect } from "react";
import DefaultApi from "../api/restaurant_api/generated/src/api/DefaultApi";
import AdminV1RestaurantCreateRequest from "../api/restaurant_api/generated/src/model/AdminV1RestaurantCreateRequest";
import "../styles/RestaurantScreen.css";

export default function RestaurantsList() {
  const [restaurants, setRestaurants] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [newName, setNewName] = useState("");

  const userUuid = localStorage.getItem("user_uuid");
  const api = new DefaultApi();
  api.apiClient.basePath = "https://d193-2a12-5940-8a19-00-2.ngrok-free.app";

  useEffect(() => {
    if (userUuid) {
      api.adminV1RestaurantListGet(userUuid, (err, data) => {
        if (err) {
          console.error("Ошибка получения списка ресторанов:", err);
        } else {
          setRestaurants(Array.isArray(data) ? data : []);
        }
      });
    }
  }, [userUuid]);

  const handleCreate = async () => {
    if (!newName.trim()) return alert("Введите название ресторана");
    
    const req = AdminV1RestaurantCreateRequest.constructFromObject({
      name: newName.trim(),
    });

    const jwt = localStorage.getItem("jwt");

    console.log("jwt: ", jwt);

    if (jwt) {
      api.apiClient.defaultHeaders['Authorization'] = `Bearer ${jwt}`;
    }

    api.adminV1RestaurantCreatePost(userUuid, req, { withCredentials: true })
    .end((err, response) => {
      if (err) {
        alert("Не удалось создать ресторан");
        console.error(err);
      } else {
        setRestaurants((prev) => [...prev, response]);
        setNewName("");
        setShowModal(false);
      }
    });
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
              {r.name}
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