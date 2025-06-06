import React, { useState, useEffect, useCallback, useMemo } from "react";
import { useParams } from "react-router-dom";
import { handleMultipartResponse } from '../components/multipartUtils';
import { SERVER_URL } from "../config";

import PlaceApi from "../api/place_api/generated/src/api/DefaultApi";
import AdminV1PlaceListRequest from "../api/place_api/generated/src/model/AdminV1PlaceListRequest";
import AdminV1PlaceCreateRequest from "../api/place_api/generated/src/model/AdminV1PlaceCreateRequest";

import RestaurantApi from "../api/restaurant_api/generated/src/api/DefaultApi";

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
  const [tableCountError, setTableCountError] = useState("");
  const [openingTime, setOpeningTime] = useState("08:00");
  const [closingTime, setClosingTime] = useState("23:00");

  const [dishes, setDishes] = useState([]);
  const [dishName, setDishName] = useState("");
  const [description, setDescription] = useState("");
  const [category, setCategory] = useState("");
  const [calories, setCalories] = useState(1);
  const [caloriesError, setCaloriesError] = useState("");
  const [weight, setWeight] = useState(1);
  const [weightError, setWeightError] = useState("");
  const [pictureFile, setPictureFile] = useState(null);

  const userUUID = localStorage.getItem("user_uuid");
  const jWTToken = localStorage.getItem("jwt_token");

   const categories = [
    "Завтраки",
    "Супы",
    "Второе",
    "Салаты",
    "Закуски",
    "Десерты",
    "Напитки"
  ];

  const placeApi = useMemo(() => {
    const api = new PlaceApi();
    api.apiClient.basePath = SERVER_URL;
    api.apiClient.defaultHeaders = {
      "User-UUID": userUUID,
      "JWT-Token": jWTToken,
      "ngrok-skip-browser-warning": "true",
    };
    return api;
  }, [userUUID, jWTToken]);

  const restApi = new RestaurantApi();
  restApi.apiClient.basePath = SERVER_URL;
  restApi.apiClient.defaultHeaders = {
    "User-UUID": userUUID,
    "JWT-Token": jWTToken,
    "ngrok-skip-browser-warning": "true",
  };

  const loadPlaces = useCallback(async () => {
    setLoading(true);
    try {
      const req = AdminV1PlaceListRequest.constructFromObject({ restaurant_uuid });
      const list = await new Promise((res, rej) =>
        placeApi.adminV1PlaceListPost(userUUID, jWTToken, req, (e,d) => e ? rej(e) : res(d.place_list))
      );
      setPlaces(list);
    } catch (e) {
      console.error("Ошибка загрузки плейсов:", e);
    } finally {
      setLoading(false);
    }
  }, [restaurant_uuid, placeApi, userUUID, jWTToken]);

  const loadDishes = useCallback(async () => {
    setLoading(true);
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
        body: JSON.stringify({ restaurant_uuid }),
      });
      if (!fastResp.ok) throw new Error("Ошибка получения блюд (fast)");
      const fastData = await fastResp.json();
      fastList = (fastData.dish_list || []).map(d => ({ ...d, imageUrl: null }));
      setDishes(fastList);
      setLoading(false); 

      const slowResp = await fetch(`${SERVER_URL}/admin/v1/restaurant/dish/list`, {
        method: "POST",
        headers: {
          Accept: "multipart/mixed, application/json",
          "Content-Type": "application/json",
          "User-UUID": userUUID,
          "JWT-Token": jWTToken,
          "ngrok-skip-browser-warning": "true",
        },
        body: JSON.stringify({ restaurant_uuid }),
      });
      const { list, imagesMap: slowImagesMap } = await handleMultipartResponse(slowResp, 'dish_list');
      imagesMap = slowImagesMap;
      setDishes(list.map(d => ({ ...d, imageUrl: imagesMap[d.picture_key] || null })));
    } catch (e) {
      console.error("Ошибка загрузки блюд:", e);
      setDishes([]);
      setLoading(false);
    }
  }, [restaurant_uuid, userUUID, jWTToken]);

   useEffect(() => {
    loadPlaces();
  }, [restaurant_uuid, loadPlaces]);

   useEffect(() => {
    if (tab === "dishes") {
      loadDishes();
    }
  }, [restaurant_uuid, tab, loadDishes]);

  const handleCreatePlace = async () => {
    const tc = Number(tableCount);
    if (!Number.isFinite(tc) || tc <= 0) {
      setTableCountError("Введите целое число большее или равное 1");
      return;
    }

    if (!address.trim()) return alert("Введите адрес плейса");

    try {
      const req = AdminV1PlaceCreateRequest.constructFromObject({
        restaurant_uuid: restaurant_uuid,
        address: address.trim(),
        table_count: tc,
        opening_time: openingTime,
        closing_time: closingTime,
      });
      await new Promise((res, rej) =>
        placeApi.adminV1PlaceCreatePost(userUUID, jWTToken, req, (err, d) =>
          err ? rej(err) : res(d)
        )
      );
      await loadPlaces();

      setAddress("");
      setTableCount(1);
      setOpeningTime("09:00");
      setClosingTime("21:00");
      setTableCountError("");
      setShowModal(false);
    } catch (e) {
      console.error("Ошибка создания плейса:", e);
      alert(e.body?.message || e.message);
    }
  };

  const handleCreateDish = async () => {
    let ok = true;

    const cal = Number(calories);
    if (!Number.isFinite(cal) || cal <= 0) {
      setCaloriesError("Введите число большее или равное 1");
      ok = false;
    }

    const wt = Number(weight);
    if (!Number.isFinite(wt) || wt <= 0) {
      setWeightError("Введите число большее или равное 1");
      ok = false;
    }

    if (
      !dishName.trim() ||
      !description.trim() ||
      !category.trim() ||
      calories <= 0 ||
      weight <= 0 ||
      !pictureFile
    ) {
      alert("Заполните все поля и выберите фото");
      ok = true;
    }

    if(!ok) {
      return;
    }
    
    try {
      await new Promise((res, rej) =>
        restApi.adminV1RestaurantDishCreatePost(
          userUUID,
          jWTToken,
          restaurant_uuid,
          dishName.trim(),
          description.trim(),
          category.trim(),
          cal,
          wt,
          pictureFile,
          (err, d) => {
            (err ? rej(err) : res(d))
          }
        )
      );

      await loadDishes();

      setDishName("");
      setDescription("");
      setCategory("");
      setCalories(1);
      setWeight(1);
      setPictureFile(null);
      setCaloriesError("");
      setWeightError("");
      setShowModal(false);
    } catch (e) {
      console.error("Ошибка создания блюда:", e);
      const msg = e.body?.message || e.message;

      setCaloriesError(msg);
      setWeightError(msg);
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
      </div>
  
      <div className="pd-tabs">
        <div
          className={`tab ${tab === "places" ? "active" : ""}`}
          onClick={() => setTab("places")}
        >
          Плейсы
        </div>
        <div
          className={`tab ${tab === "dishes" ? "active" : ""}`}
          onClick={() => setTab("dishes")}
        >
          Блюда
        </div>
      </div>
  
      <div className={`pd-list ${tab}`}>
        {loading && <p className="pd-loading">Загрузка…</p>}
  
        {!loading && tab === "places" && places.length === 0 && (
          <p className="pd-empty">Нет плейсов</p>
        )}
        {!loading && tab === "places" && places.map((p) => (
          <div key={p.uuid} className="pd-item" onClick={() => {
            localStorage.setItem("current_place", JSON.stringify({
              place_uuid: p.uuid,
              place_name: p.address
            }));
            window.location.href =`/restaurants/${restaurant_uuid}/places-dishes/${p.uuid}`
          }}>
            <div className="pd-place-info">
              <div className="pd-place-address">{p.address}</div>
              <div className="pd-place-details">
                столов: {p.table_count}, {p.opening_time}–{p.closing_time}
              </div>
            </div>
          </div>
        ))}
  
        {!loading && tab === "dishes" && dishes.length === 0 && (
          <p className="pd-empty">Нет блюд</p>
        )}
        {loading && tab === "dishes" && (
          <div style={{display:'flex',flexWrap:'wrap',gap:'1rem'}}>
            {[...Array(6)].map((_,i) => (
              <div key={i} className="pd-dish-card">
                <div className="pd-dish-image shimmer shimmer-rect" style={{height:120}} />
                <div className="pd-dish-info pd-dish-info-left" />
              </div>
            ))}
          </div>
        )}
        {!loading && tab === "dishes" && dishes.map((d) => (
            <div key={d.id} className="pd-dish-card">
                 <div className="pd-dish-image">
                  {d.imageUrl
                    ? <img src={d.imageUrl} alt={d.name} />
                    : <div className="shimmer shimmer-rect" style={{width:370, height:260}} />}
                </div>
                <div className="pd-dish-info pd-dish-info-left">
                  <div className="pd-dish-title"><b>{d.name}</b></div>
                  <div className="pd-dish-desc"><span>Описание:</span> {d.description}</div>
                  <div className="pd-dish-category"><span>Категория:</span> {d.category}</div>
                  <div className="pd-dish-weight"><span>Вес:</span> {d.weight} г</div>
                  <div className="pd-dish-calories"><span>Калории:</span> {d.calories} ккал</div>
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
              type="text"
              inputMode="numeric"
              value={tableCount}
              onChange={(e) => {
                  setTableCount(e.target.value);
                  setTableCountError("");
                }}
                placeholder="Введите число"
            />
            {tableCountError && <div className="pd-error-text">{tableCountError}</div>}
  
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
            <select
              value={category}
              onChange={(e) => setCategory(e.target.value)}
              className="pd-category-select"
            >
              <option value="">Выберите категорию</option> {
                categories.map((category) => (
                  <option key={category} value={category}> {
                    category
                  }
                </option>
              ))}
            </select>
  
            <label>Калории</label>
            <input
              type="text"
              inputMode="numeric"
              value={calories}
              onChange={(e) => {
                setCalories(e.target.value);
                setCaloriesError("");
              }}
              placeholder="Введите калории"
            />
            {caloriesError && <div className="pd-error-text">{caloriesError}</div>}
            
            <label>Вес</label>
            <input
              type="text"
              inputMode="numeric"
              value={weight}
              onChange={(e) => {
                setWeight(e.target.value);
                setWeightError("");
              }}
              placeholder="Введите вес (граммы)"
            />
            {weightError && <div className="pd-error-text">{weightError}</div>}

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
