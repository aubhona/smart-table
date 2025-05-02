import React, { useState } from "react";
import DefaultApi from "../api/generated/src/api/DefaultApi";
import AdminV1UserSignInRequest from "../api/generated/src/model/AdminV1UserSignInRequest";
import "../styles/AuthScreens.css";

export default function LoginForm() {
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const payload = AdminV1UserSignInRequest.constructFromObject({ login, password });

    const api = new DefaultApi();
    api.apiClient.basePath = "https://8bb9-138-124-99-156.ngrok-free.app";

    try {
      const response = await api.adminV1UserSignInPost(payload, { withCredentials: true });

      const { user_uuid } = response;
      localStorage.setItem("userUuid", user_uuid);

      alert(`Авторизация успешна. Ваш UUID: ${user_uuid}`);
    } catch (err) {
      const code = err.response?.body?.code;
      if (code === "not_found") {
        alert("Пользователь не найден");
      } else if (code === "incorrect_password") {
        alert("Неверный пароль");
      } else {
        alert("Ошибка авторизации");
      }
      console.error("Ошибка авторизации:", err);
    }
  };

  return (
    <div className="auth-container">
      <h2>Вход в систему</h2>
      <form className="auth-form" onSubmit={handleSubmit}>
        <input
          placeholder="Логин"
          value={login}
          onChange={(e) => setLogin(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Пароль"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button type="submit" className="auth-button">
          Войти
        </button>
      </form>
    </div>
  );
}