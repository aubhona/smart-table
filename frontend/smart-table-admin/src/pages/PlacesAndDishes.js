import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";

import PlaceApi from "../api/place_api/generated/src/api/DefaultApi";
import AdminV1PlaceListRequest from "../api/place_api/generated/src/model/AdminV1PlaceListRequest";
import AdminV1PlaceCreateRequest from "../api/place_api/generated/src/model/AdminV1PlaceCreateRequest";

import RestaurantApi from "../api/restaurant_api/generated/src/api/DefaultApi";
import AdminV1RestaurantDishListRequest from "../api/restaurant_api/generated/src/model/AdminV1RestaurantDishListRequest";

import "../styles/PlacesDishesScreen.css";  

export default function PlacesAndDishes() {
  const { restaurant_uuid } = useParams();

  const saved = JSON.parse(localStorage.getItem("current_restaurant") || "{}");
  const restaurantName = saved.restaurant_name || "Ресторан";

  const [tab, setTab] = useState("places");

  const [places, setPlaces] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);

  const [address, setAddress] = useState("");
  const [tableCount, setTableCount] = useState(1);
  const [openingTime, setOpeningTime] = useState("08:00");
  const [closingTime, setClosingTime] = useState("23:00");

  const [dishes, setDishes] = useState([]);

  const [dishName, setDishName] = useState("");
  const [description, setDescription] = useState("");
  const [category, setCategory] = useState("");
  const [calories, setCalories] = useState(0);
  const [weight, setWeight] = useState(0);
  const [pictureFile, setPictureFile] = useState(null);

  const userUUID = localStorage.getItem("user_uuid");
  const jWTToken = localStorage.getItem("jwt_token");

  const placeApi = new PlaceApi();
  placeApi.apiClient.basePath = "https://5506-135-181-37-249.ngrok-free.app";
  placeApi.apiClient.defaultHeaders = {
    "User-UUID": userUUID,
    "JWT-Token": jWTToken,
    "ngrok-skip-browser-warning": "true",
  };

  const restApi = new RestaurantApi();
  restApi.apiClient.basePath = "https://5506-135-181-37-249.ngrok-free.app";
  restApi.apiClient.defaultHeaders = {
    "User-UUID": userUUID,
    "JWT-Token": jWTToken,
    "ngrok-skip-browser-warning": "true",
  };

  useEffect(() => {
    (async () => {
      setLoading(true);
      try {
        const req = AdminV1PlaceListRequest.constructFromObject({
          restaurant_uuid: restaurant_uuid,
        });
        const list = await new Promise((res, rej) =>
            placeApi.adminV1PlaceListPost(userUUID, jWTToken, req, (err, data) =>
            err ? rej(err) : res(data.place_list)
          )
        );
        setPlaces(list);
      } catch (e) {
        console.error("Ошибка загрузки плейсов:", e);
      } finally {
        setLoading(false);
      }
    })();
  }, [restaurant_uuid]);

  useEffect(() => {
    if (tab !== "dishes") return;
  
    (async () => {
      setLoading(true);
      try {
        const resp = await fetch(
          "https://5506-135-181-37-249.ngrok-free.app/admin/v1/restaurant/dish/list",
          {
            method: "POST",
            headers: {
              "Accept": "multipart/mixed, application/json",
              "Content-Type": "application/json",
              "User-UUID": userUUID,
              "JWT-Token": jWTToken,
              "ngrok-skip-browser-warning": "true",
            },
            body: JSON.stringify({ restaurant_uuid: restaurant_uuid }),
          }
        );
        
        const ct = resp.headers.get("Content-Type") || "";
        const m = ct.match(/boundary="?([^";]+)"?/i);
        if (!m) throw new Error("Не найдена граница в Content-Type");

        const boundary = m[1];

        const buf = await resp.arrayBuffer();
        const text = new TextDecoder().decode(buf);

        const parts = text.split(`--${boundary}`);

        const jsonPart = parts.find((p) =>
            p.includes("Content-Type: application/json")
        );
        if (!jsonPart) throw new Error("Часть с JSON не найдена");

        const idx = jsonPart.indexOf("\r\n\r\n");
        const jsonText = jsonPart.slice(idx + 4).trim();
        const json = JSON.parse(jsonText);

        console.log(json);

        const list = Array.isArray(json)
        ? json
        : Array.isArray(json.dish_list)
        ? json.dish_list
        : [];

        setDishes(list);
      } catch (e) {
        console.error("Ошибка загрузки блюд:", e);
        setDishes([]);
      } finally {
        setLoading(false);
      }
    })();
  }, [restaurant_uuid, tab]);

  const handleCreatePlace = async () => {
    if (!address.trim()) return alert("Введите адрес плейса");
    try {
      const req = AdminV1PlaceCreateRequest.constructFromObject({
        restaurant_uuid: restaurant_uuid,
        address: address.trim(),
        table_count: tableCount,
        opening_time: openingTime,
        closing_time: closingTime,
      });
      const data = await new Promise((res, rej) =>
        placeApi.adminV1PlaceCreatePost(userUUID, jWTToken, req, (err, d) =>
          err ? rej(err) : res(d)
        )
      );
      setPlaces((prev) => [...prev, 
        { 
            uuid: data.place_uuid,
            address: req.address,
            table_count: req.table_count,
            opening_time: req.opening_time,
            closing_time: req.closing_time,
        }]);
      setAddress("");
      setTableCount(1);
      setOpeningTime("09:00");
      setClosingTime("21:00");
      setShowModal(false);
    } catch (e) {
      console.error("Ошибка создания плейса:", e);
      alert(e.body?.message || e.message);
    }
  };

  const handleCreateDish = async () => {
    if (
      !dishName.trim() ||
      !description.trim() ||
      !category.trim() ||
      calories <= 0 ||
      weight <= 0 ||
      !pictureFile
    ) {
      return alert("Заполните все поля и выберите фото");
    }
    try {
      const data = await new Promise((res, rej) =>
        restApi.adminV1RestaurantDishCreatePost(
          userUUID,
          jWTToken,
          restaurant_uuid,
          dishName.trim(),
          description.trim(),
          category.trim(),
          calories,
          weight,
          pictureFile,
          (err, d) => (err ? rej(err) : res(d))
        )
      );
      setDishes((prev) => [
        ...prev,
        {
          id: data.dish_uuid,
          name: dishName.trim(),
          description: description.trim(),
          calories,
          weight,
          category: category.trim(),
          picture_key: data.dish_uuid,
        },
      ]);
      setShowModal(false);
    } catch (e) {
      console.error("Ошибка создания блюда:", e);
      alert(e.body?.message || e.message);
    }
  };

  return (
    <div className="pd-container">
      <div className="pd-header-bar">
        <button
          className="pd-back-button"
          onClick={() => window.history.back()}
        >
          Назад
        </button>
        <h1 className="pd-title">Ресторан: {restaurantName}</h1>
        <button
          className="pd-create-button"
          onClick={() => setShowModal(true)}
        >
          {tab === "places" ? "Создать плейс" : "Создать блюдо"}
        </button>
        <button className="pd-profile-button">𓀡</button>
      </div>
  
      <div className="pd-tabs">
        <div
          className={`tab ${tab === "places" ? "active" : ""}`}
          onClick={() => setTab("places")}
        >
          places
        </div>
        <div
          className={`tab ${tab === "dishes" ? "active" : ""}`}
          onClick={() => setTab("dishes")}
        >
          dishes
        </div>
      </div>
  
      <div className={`pd-list ${tab}`}>
        {loading && <p className="pd-loading">Загрузка…</p>}
  
        {!loading && tab === "places" && places.length === 0 && (
          <p className="pd-empty">Нет плейсов</p>
        )}
        {!loading && tab === "places" && places.map((p) => (
            <div key={p.uuid} className="pd-item">
              <strong>{p.address}</strong>
              <br />
              столов: {p.table_count}, {p.opening_time}–{p.closing_time}
            </div>
          ))}
  
        {!loading && tab === "dishes" && dishes.length === 0 && (
          <p className="pd-empty">Нет блюд</p>
        )}
        {!loading && tab === "dishes" && dishes.map((d) => (
            <div key={d.id} className="pd-item pd-dish-card">
                 <div className="pd-dish-image">
                    <img
                        src={`https://5506-135-181-37-249.ngrok-free.app/files/${d.picture_key}`}
                    />
                </div>
                <div classname="pd-dish-info">
                    {d.name}
                </div>
            </div>
          ))}
      </div>
  
      {showModal && tab === "places" && (
        <div className="pd-modal-backdrop">
          <div className="pd-modal">
            <h3>Новый плейс</h3>
  
            <label>Адрес</label>
            <input
              value={address}
              onChange={(e) => setAddress(e.target.value)}
              placeholder="Введите адрес"
            />
  
            <label>Количество столов</label>
            <input
              type="number"
              min="1"
              value={tableCount}
              onChange={(e) => setTableCount(Number(e.target.value))}
            />
  
            <label>Время открытия</label>
            <input
              type="time"
              value={openingTime}
              onChange={(e) => setOpeningTime(e.target.value)}
            />
  
            <label>Время закрытия</label>
            <input
              type="time"
              value={closingTime}
              onChange={(e) => setClosingTime(e.target.value)}
            />
  
            <div className="pd-modal-buttons">
              <button onClick={handleCreatePlace}>Создать</button>
              <button onClick={() => setShowModal(false)}>Отмена</button>
            </div>
          </div>
        </div>
      )}
  
      {showModal && tab === "dishes" && (
        <div className="pd-modal-backdrop">
          <div className="pd-modal">
            <h3>Новое блюдо</h3>
  
            <label>Название</label>
            <input
              value={dishName}
              onChange={(e) => setDishName(e.target.value)}
              placeholder="Введите название"
            />
  
            <label>Описание</label>
            <input
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              placeholder="Введите описание"
            />
  
            <label>Категория</label>
            <input
              value={category}
              onChange={(e) => setCategory(e.target.value)}
              placeholder="Введите категорию"
            />
  
            <label>Калории</label>
            <input
              type="number"
              min="0"
              value={calories}
              onChange={(e) => setCalories(+e.target.value)}
            />
  
            <label>Вес</label>
            <input
              type="number"
              min="0"
              value={weight}
              onChange={(e) => setWeight(+e.target.value)}
            />
  
            <label>Фото</label>
            <input
              type="file"
              accept="image/*"
              onChange={(e) => setPictureFile(e.target.files[0])}
            />
  
            <div className="pd-modal-buttons">
              <button onClick={handleCreateDish}>Создать блюдо</button>
              <button onClick={() => setShowModal(false)}>Отмена</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
