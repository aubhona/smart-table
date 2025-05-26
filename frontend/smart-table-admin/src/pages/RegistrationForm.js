import React, { useState } from "react";
import { Navigate } from "react-router-dom";
import DefaultApi from "../api/user_api/generated/src/api/DefaultApi";
import AdminV1UserSignUpRequest from "../api/user_api/generated/src/model/AdminV1UserSignUpRequest";
import "../styles/AuthScreens.css";
import { SERVER_URL } from "../config";
import { ToastContainer } from "../components/Toast/Toast";
import useToast from "../components/hooks/useToast";

const api = new DefaultApi();
api.apiClient.basePath = SERVER_URL;

export default function RegistrationForm() {
  const { toasts, addToast, removeToast } = useToast();
  const [redirect, setRedirect] = useState(false);
  const [form, setForm] = useState({
    login: "",
    tg_login: "",
    first_name: "",
    last_name: "",
    password: "",
    password_confirm: "",
  });
  
  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (form.password !== form.password_confirm) {
      addToast("Пароли не совпадают!", "error");
      return;
    }

    const payload = AdminV1UserSignUpRequest.constructFromObject({
      login: form.login,
      tg_login: form.tg_login,
      first_name: form.first_name,
      last_name: form.last_name,
      password: form.password,
    });

    try {
      const response = await new Promise((resolve, reject) => {
        api.adminV1UserSignUpPost(payload, (err, data, response) => {
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

      addToast("Регистрация успешна!", "success");
      
      setForm({
        login: "",
        tg_login: "",
        first_name: "",
        last_name: "",
        password: "",
        password_confirm: "",
      });

      setTimeout(() => setRedirect(true), 1000);
    } catch (err) {
      const code = err.response?.body?.code;
      if (code === "already_exist") {
        addToast("Пользователь с таким логином уже существует", "error");
      } else {
        addToast("Ошибка при регистрации", "error");
      }
      console.error("Ошибка регистрации:", err);
    }
  };

  if (redirect) {
    return <Navigate to="/restaurants" replace />;
  }

  return (
    <div className="auth-container">
      <ToastContainer toasts={toasts} removeToast={removeToast} />
      <h2>Регистрация</h2>
      <form className="auth-form" onSubmit={handleSubmit}>
        <input 
          name="login" 
          placeholder="Логин" 
          value={form.login}
          onChange={handleChange} 
          required 
        />
        <input 
          name="tg_login" 
          placeholder="Telegram логин" 
          value={form.tg_login}
          onChange={handleChange} 
        />
        <input
          name="first_name"
          placeholder="Имя (латиницей)"
          value={form.first_name}
          onChange={handleChange}
          required
        />
        <input
          name="last_name"
          placeholder="Фамилия (латиницей)"
          value={form.last_name}
          onChange={handleChange}
          required
        />
        <input
          name="password"
          type="password"
          placeholder="Пароль"
          value={form.password}
          onChange={handleChange}
          required
        />
        <input
          name="password_confirm"
          type="password"
          placeholder="Повторите пароль"
          value={form.password_confirm}
          onChange={handleChange}
          required
        />
        <button type="submit" className="auth-button">
          Зарегистрироваться
        </button>
      </form>
    </div>
  );
}
