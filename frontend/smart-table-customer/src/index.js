import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import TwaSdk from "@twa-dev/sdk";
import { OrderProvider } from './components/OrderContext/OrderContext';

TwaSdk.ready();

if (!window.Telegram || !window.Telegram.WebApp) {
  alert("Ошибка: Telegram WebApp не загружен! Запустите приложение в Telegram.");
} else {
  window.Telegram.WebApp.ready();
  window.Telegram.WebApp.expand();
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <OrderProvider>
      <App />
    </OrderProvider>
  </React.StrictMode>
);
