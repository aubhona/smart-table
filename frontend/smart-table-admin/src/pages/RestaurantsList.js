import React, { useState, useEffect } from "react";
import DefaultApi from "../restaurant_api/generated/src/api/DefaultApi";
import AdminV1RestaurantCreateRequest from "../restaurant_api/generated/src/model/AdminV1RestaurantCreateRequest";
import "../styles/RestaurantScreen.css";

export default function RestaurantsList() {
  const [restaurants, setRestaurants] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [newName, setNewName] = useState("");

  const userUuid = localStorage.getItem("userUuid");
  const api = new DefaultApi();
  api.apiClient.basePath = "https://8bb9-138-124-99-156.ngrok-free.app";

  // TODO: при наличии ручки «список ресторанов»:
  // useEffect(() => {
  //   api.adminV1RestaurantListGet(userUuid, { withCredentials: true })
  //     .then((data) => setRestaurants(data))
  //     .catch(console.error);
  // }, []);

  const handleCreate = async () => {
    if (!newName.trim()) {
      return alert("Введите название ресторана");
    }

    const req = AdminV1RestaurantCreateRequest.constructFromObject({
      name: newName.trim(),
    });

    try {
      const created = await new Promise((resolve, reject) => {
        api.adminV1RestaurantCreatePost(
          userUuid,
          req,
          (err, data, response) => {
            if (err) reject(err);
            else resolve(data);
          }
        );
      });
      setRestaurants((prev) => [...prev, created]);
      setNewName("");
      setShowModal(false);
    } catch (err) {
      console.error("Ошибка создания ресторана:", err);
      alert("Не удалось создать ресторан");
    }
  };

  return (
    <div className="rest-container">
      <h2 className="rest-header">Мои рестораны</h2>

      {restaurants.length === 0 ? (
        <p className="rest-header">Нет ресторанов</p>
      ) : (
        <ul className="rest-list">
          {restaurants.map((r) => (
            <li key={r.restaurant_uuid} className="rest-item">
              <button
                className="rest-button"
                onClick={() => alert("Откроем плейсы ресторана")}
              >
                {r.name}
              </button>
            </li>
          ))}
        </ul>
      )}

      <button
        className="create-button"
        onClick={() => setShowModal(true)}
      >
        Создать ресторан
      </button>

      {showModal && (
        <div className="modal-backdrop">
          <div className="modal">
            <h3>Название ресторана</h3>
            <input
              value={newName}
              onChange={(e) => setNewName(e.target.value)}
              placeholder="Введите название"
            />
            <div className="modal-buttons">
              <button className="rest-button" onClick={handleCreate}>
                Создать
              </button>
              <button
                className="rest-button"
                onClick={() => setShowModal(false)}
              >
                Отмена
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}