import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import TwaSdk from "@twa-dev/sdk";

console.log("TwaSdk exports:", TwaSdk);
TwaSdk.ready();

console.log("Telegram WebApp:", window.Telegram.WebApp);

if (!window.Telegram || !window.Telegram.WebApp) {
  alert("Ошибка: Telegram WebApp не загружен! Запустите приложение в Telegram.");
} else {
  console.log("Telegram WebApp загружен:", window.Telegram.WebApp);
  window.Telegram.WebApp.ready();
  window.Telegram.WebApp.expand();
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
