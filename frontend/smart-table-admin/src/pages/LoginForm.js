import React, { useState } from "react";
import { Navigate } from "react-router-dom";
import DefaultApi from "../api/user_api/generated/src/api/DefaultApi";
import AdminV1UserSignInRequest from "../api/user_api/generated/src/model/AdminV1UserSignInRequest";
import "../styles/AuthScreens.css";
import { SERVER_URL } from "../config";
import { ToastContainer } from "../components/Toast/Toast";
import useToast from "../components/hooks/useToast";

export default function LoginForm() {
  const { toasts, addToast, removeToast } = useToast();
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [redirect, setRedirect] = useState(false);

  const handleLogin = async (e) => {
    e.preventDefault();

    const payload = AdminV1UserSignInRequest.constructFromObject({ login, password });
    const api = new DefaultApi();
    api.apiClient.basePath = SERVER_URL;

    try {
      const response = await new Promise((resolve, reject) => {
        api.adminV1UserSignInPost(payload, (err, data, response) => {
          if (err) {
            reject(err);
          } else {
            resolve({
              user_uuid: data.user_uuid,
              jwt_token: data.jwt_token
            });
          }
        });
      });

      const { user_uuid, jwt_token } = response;

      localStorage.setItem("user_uuid", user_uuid);
      localStorage.setItem("jwt_token", jwt_token);

      console.log(user_uuid, jwt_token);

      addToast("Успешная авторизация!", "success");
      setTimeout(() => setRedirect(true), 1000); 
    } catch (err) {
      const code = err.response?.body?.code;
      if (code === "not_found") {
        addToast("Пользователь не найден", "error");
      } else if (code === "incorrect_password") {
        addToast("Неверный пароль", "error");
      } else {
        addToast("Ошибка авторизации", "error");
      }
      console.error("Ошибка авторизации:", err);
    }
  };

  if (redirect) {
    return <Navigate to="/restaurants" replace />;
  }

  return (
    <div className="auth-container">
      <ToastContainer toasts={toasts} removeToast={removeToast} />
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
        <button type="submit" className="auth-button">
          Войти
        </button>
      </form>
    </div>
  );
}