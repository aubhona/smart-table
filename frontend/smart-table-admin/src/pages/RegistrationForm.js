import React, { useState } from "react";
import DefaultApi from "../api/generated/src/api/DefaultApi";
import AdminV1UserSignUpRequest from "../api/generated/src/model/AdminV1UserSignUpRequest";
import "../styles/AuthScreens.css";

const api = new DefaultApi();

function RegistrationForm() {
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
      alert("Пароли не совпадают!");
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
      await api.adminV1UserSignUpPost(payload, { withCredentials: true });
      alert("Регистрация успешна!");
    } catch (err) {
      if (err.response?.body?.code === "already_exist") {
        alert("Такой пользователь уже существует!");
      } else {
        alert("Ошибка при регистрации");
      }
      console.error("Ошибка регистрации:", err);
    }
  };

  return (
    <div className="auth-container">
      <h2>Регистрация</h2>
      <form className="auth-form" onSubmit={handleSubmit}>
        <input name="login" placeholder="Логин" onChange={handleChange} required />
        <input name="tg_login" placeholder="Telegram логин" onChange={handleChange} />
        <input name="first_name" placeholder="Имя (латиницей)" onChange={handleChange} required />
        <input name="last_name" placeholder="Фамилия (латиницей)" onChange={handleChange} required />
        <input name="password" type="password" placeholder="Пароль" onChange={handleChange} required />
        <input name="password_confirm" type="password" placeholder="Повторите пароль" onChange={handleChange} required />
        <button type="submit" className="auth-button">Зарегистрироваться</button>
      </form>
    </div>
  );
}

export default RegistrationForm;
