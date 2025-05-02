import React, { useState } from "react";
import { Navigate } from "react-router-dom";
import DefaultApi from "../api/generated/src/api/DefaultApi";
import AdminV1UserSignInRequest from "../api/generated/src/model/AdminV1UserSignInRequest";
import "../styles/AuthScreens.css";

export default function LoginForm() {
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [redirect, setRedirect] = useState(false);

  const handleLogin = async (e) => {
    e.preventDefault();
    setError("");  // сброс предыдущей ошибки

    const payload = AdminV1UserSignInRequest.constructFromObject({ login, password });
    const api = new DefaultApi();
    api.apiClient.basePath = "https://8bb9-138-124-99-156.ngrok-free.app";

    try {
      const { user_uuid } = await api.adminV1UserSignInPost(payload, { withCredentials: true });
      localStorage.setItem("userUuid", user_uuid);
      setRedirect(true);  // переключимся на экран ресторанов
    } catch (err) {
      const code = err.response?.body?.code;
      if (code === "not_found") {
        setError("Пользователь не найден");
      } else if (code === "incorrect_password") {
        setError("Неверный пароль");
      } else {
        setError("Ошибка авторизации");
      }
      console.error("Ошибка авторизации:", err);
    }
  };

  if (redirect) {
    return <Navigate to="/restaurants" replace />;
  }

  return (
    <div className="auth-container">
      <h2>Вход в систему</h2>
      <form className="auth-form" onSubmit={handleLogin}>
        <input
          type="text"
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
        {error && <p className="auth-error">{error}</p>}
        <button type="submit" className="auth-button">
          Войти
        </button>
      </form>
    </div>
  );
}