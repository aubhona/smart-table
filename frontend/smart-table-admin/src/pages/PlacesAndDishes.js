import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { handleMultipartResponse } from '../components/multipartUtils';
import { SERVER_URL } from "../config";

import PlaceApi from "../api/place_api/generated/src/api/DefaultApi";
import AdminV1PlaceListRequest from "../api/place_api/generated/src/model/AdminV1PlaceListRequest";
import AdminV1PlaceCreateRequest from "../api/place_api/generated/src/model/AdminV1PlaceCreateRequest";
import AdminV1PlaceDeleteRequest from "../api/place_api/generated/src/model/AdminV1PlaceDeleteRequest";
import AdminV1PlaceEditRequest from "../api/place_api/generated/src/model/AdminV1PlaceEditRequest";

import RestaurantApi from "../api/restaurant_api/generated/src/api/DefaultApi";

import "../styles/PlacesDishesScreen.css";  

export default function PlacesAndDishes() {
  const { restaurant_uuid } = useParams();
  const navigate = useNavigate();

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

  const [showEditPlaceModal, setShowEditPlaceModal] = useState(false);
  const [editPlace, setEditPlace] = useState(null);
  const [editAddress, setEditAddress] = useState("");
  const [editTableCount, setEditTableCount] = useState(1);

  const userUUID = localStorage.getItem("user_uuid");
  const jWTToken = localStorage.getItem("jwt_token");

   const categories = [
    "Завтрак",
    "Супы",
    "Второе",
    "Салаты",
    "Десерты",
    "Напитки"
  ];

  const placeApi = new PlaceApi();
  placeApi.apiClient.basePath = SERVER_URL;
  placeApi.apiClient.defaultHeaders = {
    "User-UUID": userUUID,
    "JWT-Token": jWTToken,
    "ngrok-skip-browser-warning": "true",
  };

  const restApi = new RestaurantApi();
  restApi.apiClient.basePath = SERVER_URL;
  restApi.apiClient.defaultHeaders = {
    "User-UUID": userUUID,
    "JWT-Token": jWTToken,
    "ngrok-skip-browser-warning": "true",
  };

  async function loadPlaces() {
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
  }

  async function loadDishes() {
    setLoading(true);
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
        body: JSON.stringify({ restaurant_uuid }),
      });

    const { list, imagesMap } = await handleMultipartResponse(resp, 'dish_list');
    
    setDishes(list.map(d => ({
      ...d,
      imageUrl: imagesMap[d.picture_key] || null
    })));
    } catch (e) {
      console.error("Ошибка загрузки блюд:", e);
      setDishes([]);
    } finally {
      setLoading(false);
    }
  }

   useEffect(() => {
    loadPlaces();
  }, [restaurant_uuid]);

   useEffect(() => {
    if (tab === "dishes") {
      loadDishes();
    }
  }, [restaurant_uuid, tab]);

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
      const data = await new Promise((res, rej) =>
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

  const handleDeletePlace = async (place_uuid) => {
    if (!window.confirm("Удалить этот плейс?")) return;
    try {
      const req = AdminV1PlaceDeleteRequest.constructFromObject({ place_uuid });
      await new Promise((res, rej) =>
        placeApi.adminV1PlaceDeletePost(userUUID, jWTToken, req, (err) =>
          err ? rej(err) : res()
        )
      );
      await loadPlaces();
    } catch (e) {
      alert(e.body?.message || e.message || "Ошибка удаления плейса");
    }
  };

  const openEditPlaceModal = (place) => {
  setEditPlace(place);
  setEditAddress(place.address);
  setEditTableCount(place.table_count);
  setShowEditPlaceModal(true);
};

  const handleSaveEditPlace = async () => {
    if (!editPlace) return;
    try {
      const req = AdminV1PlaceEditRequest.constructFromObject({
        place_uuid: editPlace.uuid,
        address: editAddress,
        table_count: editTableCount,
        opening_time: openingTime,
        closingTime: closingTime
      });
      await new Promise((res, rej) =>
        placeApi.adminV1PlaceEditPost(userUUID, jWTToken, req, (err) =>
          err ? rej(err) : res()
        )
      );
      setShowEditPlaceModal(false);
      await loadPlaces();
    } catch (e) {
      alert(e.body?.message || e.message || "Ошибка редактирования плейса");
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
      const data = await new Promise((res, rej) =>
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
        <button className="pd-profile-button">
          <span className="material-icons">person</span>
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
              <span className="pd-place-address">{p.address}</span>
              столов: {p.table_count}, {p.opening_time}–{p.closing_time}            
              <div className="pd-place-actions">
                <button
                  className="pd-button pd-edit-button"
                  onClick={e => {
                    e.stopPropagation();
                    openEditPlaceModal(p)
                  }}
                  title="Редактировать"
                >
                  <span className="material-icons">edit</span>
                </button>
                <button
                  className="pd-button pd-button-cancel"
                  onClick={e => {
                    e.stopPropagation();
                    handleDeletePlace(p)
                  }}
                  title="Удалить"
                >
                  <span className="material-icons">delete</span>
                </button>
              </div>
            </div>
          ))}
  
        {!loading && tab === "dishes" && dishes.length === 0 && (
          <p className="pd-empty">Нет блюд</p>
        )}
        {!loading && tab === "dishes" && dishes.map((d) => (
            <div key={d.id} className="pd-dish-card">
                 <div className="pd-dish-image">
                  {d.imageUrl
                    ? <img src={d.imageUrl} alt={d.name} />
                    : <div className="pd-no-image">нет фото</div>}
                </div>
                <div className="pd-dish-info">
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
      {showEditPlaceModal && (
        <div className="pd-modal-backdrop">
          <div className="pd-modal">
            <h3>Редактировать плейс</h3>
            <input
              value={editAddress}
              onChange={e => setEditAddress(e.target.value)}
              placeholder="Адрес"
            />
            <input
              value={editTableCount}
              onChange={e => setEditTableCount(e.target.value)}
              placeholder="Количество столов"
              type="number"
              min={1}
            />
            <div className="pd-modal-buttons">
              <button onClick={handleSaveEditPlace}>Сохранить</button>
              <button onClick={() => setShowEditPlaceModal(false)}>Отмена</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
