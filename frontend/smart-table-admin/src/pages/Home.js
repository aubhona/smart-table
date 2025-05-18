import React from "react";
import { Link } from "react-router-dom";
import "../styles/AuthScreens.css";

function Home() {
  return (
    <div className="auth-container">
      <h1>Добро пожаловать!</h1>
      <div className="auth-button-group">
        <Link to="/register">
          <button className="auth-button">Зарегистрироваться</button>
        </Link>
        <Link to="/login">
          <button className="auth-button">Войти в систему</button>
        </Link>
      </div>
    </div>
  );
}
export default Home;

