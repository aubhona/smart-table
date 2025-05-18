import React, { createContext, useContext, useState, useEffect } from "react";

const OrderContext = createContext();

export const useOrder = () => useContext(OrderContext);

export const OrderProvider = ({ children }) => {
  const [customer_uuid, setCustomerUuid] = useState(() => localStorage.getItem("customer_uuid"));
  const [order_uuid, setOrderUuid] = useState(() => localStorage.getItem("order_uuid"));
  const [room_code, setRoomCode] = useState(() => localStorage.getItem("room_code"));
  const [cart, setCart] = useState({});

  useEffect(() => {
    if (customer_uuid) localStorage.setItem("customer_uuid", customer_uuid);
  }, [customer_uuid]);

  useEffect(() => {
    if (order_uuid) localStorage.setItem("order_uuid", order_uuid);
  }, [order_uuid]);

  useEffect(() => {
    if (room_code) localStorage.setItem("room_code", room_code);
  }, [room_code]);

  return (
    <OrderContext.Provider value={{
      customer_uuid, setCustomerUuid,
      order_uuid, setOrderUuid,
      room_code, setRoomCode,
      cart, setCart,
    }}>
      {children}
    </OrderContext.Provider>
  );
};
