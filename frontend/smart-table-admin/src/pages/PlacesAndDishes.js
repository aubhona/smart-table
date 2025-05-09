import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";

import PlaceApi from "../api/place_api/generated/src/api/DefaultApi";
import AdminV1PlaceListRequest from "../api/place_api/generated/src/model/AdminV1PlaceListRequest";
import AdminV1PlaceCreateRequest from "../api/place_api/generated/src/model/AdminV1PlaceCreateRequest";

import RestaurantApi from "../api/restaurant_api/generated/src/api/DefaultApi";

import "../styles/PlacesDishesScreen.css";  

function indexOf(buf, sub, from = 0) {
  for (let i = from; i <= buf.length - sub.length; i++) {
    let ok = true;
    for (let j = 0; j < sub.length; j++) {
      if (buf[i + j] !== sub[j]) { ok = false; break; }
    }
    if (ok) return i;
  }
  return -1;
}

function parseMixed(bodyBuf, boundary) {
  const enc = new TextEncoder();
  const bnd = enc.encode(`--${boundary}`);
  let pos = indexOf(bodyBuf, bnd);
  if (pos < 0) return [];
  pos += bnd.length;
  const parts = [];

  while (true) {
    if (bodyBuf[pos] === 45 && bodyBuf[pos+1] === 45) break;
    if (bodyBuf[pos] === 13 && bodyBuf[pos+1] === 10) pos += 2;

    const next = indexOf(bodyBuf, bnd, pos);
    if (next < 0) break;

    let chunk = bodyBuf.subarray(pos, next);
    if (chunk[chunk.length-2] === 13 && chunk[chunk.length-1] === 10) {
      chunk = chunk.subarray(0, chunk.length - 2);
    }

    const sep = indexOf(chunk, enc.encode('\r\n\r\n'));
    const headBuf = chunk.subarray(0, sep);
    const dataBuf = chunk.subarray(sep + 4);

    const headText = new TextDecoder().decode(headBuf);
    const headers = {};
    headText
      .split('\r\n')
      .filter(line => line.includes(':'))
      .forEach(line => {
        const [k, ...rest] = line.split(':');
        headers[k.trim().toLowerCase()] = rest.join(':').trim();
      });

    const cd = headers['content-disposition'] || '';
    const nameMatch = cd.match(/name="([^"]+)"/i);
    const fileMatch = cd.match(/filename="([^"]+)"/i);

    parts.push({
      name:     nameMatch?.[1]  || null,        
      filename: fileMatch?.[1]  || null,
      type:     headers['content-type'] || null,
      data:     dataBuf
    });

    pos = next + bnd.length;
  }

  return parts;
}

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

  const userUUID = localStorage.getItem("user_uuid");
  const jWTToken = localStorage.getItem("jwt_token");

  const placeApi = new PlaceApi();
  placeApi.apiClient.basePath = "https://b04d-2a01-4f9-c010-ecd2-00-1.ngrok-free.app";
  placeApi.apiClient.defaultHeaders = {
    "User-UUID": userUUID,
    "JWT-Token": jWTToken,
    "ngrok-skip-browser-warning": "true",
  };

  const restApi = new RestaurantApi();
  restApi.apiClient.basePath = "https://b04d-2a01-4f9-c010-ecd2-00-1.ngrok-free.app";
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
      const resp = await fetch("https://b04d-2a01-4f9-c010-ecd2-00-1.ngrok-free.app/admin/v1/restaurant/dish/list", {
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

      const ct = resp.headers.get("Content-Type") || "";
      const [, boundary] = ct.match(/boundary="?([^";]+)"?/) || [];
      if (!boundary) throw new Error("Не удалось вытащить boundary");

      const buf = new Uint8Array(await resp.arrayBuffer());
      const parts = parseMixed(buf, boundary);

      const jsonPart = parts.find(p => p.type === "application/json");
      if (!jsonPart) throw new Error("JSON часть не найдена");
      const json = JSON.parse(new TextDecoder().decode(jsonPart.data));
      const list = Array.isArray(json) ? json : json.dish_list || [];

      const imagesMap = {};
      parts.filter(p => p.filename).forEach(p => {
        const blob = new Blob([p.data], { type: p.type });
        const url = URL.createObjectURL(blob);
        const key = p.name.replace(/\.\w+$/, "");
        imagesMap[key] = url;
      });

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
        <button className="pd-profile-button">𓀡</button>
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
            <input
              value={category}
              onChange={(e) => setCategory(e.target.value)}
              placeholder="Введите категорию"
            />
  
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
