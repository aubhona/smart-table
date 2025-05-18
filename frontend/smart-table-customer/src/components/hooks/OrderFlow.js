import React, { useState, useEffect } from 'react';
import axios from 'axios';
import TableId from '../TableId/TableId';
import RoomCode from '../RoomCode/RoomCode';
import useCustomerAuth from './UseCustomerAuth';
import { useOrder } from '../OrderContext/OrderContext';
import { useNavigate } from 'react-router-dom';
import { SERVER_URL } from '../../config';
import LoadingScreen from '../LoadingScreen/LoadingScreen';

function OrderFlow() {
  const [error, setError] = useState('');
  const [step, setStep] = useState(1); 
  const [creatingOrder, setCreatingOrder] = useState(false);

  const { setOrderUuid, setRoomCode } = useOrder();
  const { customer_uuid, loading, showStartPrompt } = useCustomerAuth();

  const startParam = window.Telegram?.WebApp?.initDataUnsafe.start_param;

  const navigate = useNavigate();

  useEffect(() => {
    if (loading || !customer_uuid) return;

    if (startParam) {
      setCreatingOrder(true);
      axios.post(
          `${SERVER_URL}/customer/v1/order/create`,
        { 
          table_id: startParam
        },
        {
          headers: {
            'Customer-UUID': customer_uuid,
            'JWT-Token': 'bla-bla-bla'
          }
        }
      )
      .then(res => {
        setOrderUuid(res.data.order_uuid);
        setCreatingOrder(false);
        navigate('/catalog');
      })
      .catch(err => {
        setCreatingOrder(false);
        setError('Ошибка: ' + (err?.response?.data?.error || 'не удалось создать заказ'));
      });
    }
  }, [loading, customer_uuid]);

  if (creatingOrder) {
    return <LoadingScreen message="Создаём заказ..." />;
  }

  if (showStartPrompt) {
    return (
      <div style={{ padding: '20%', textAlign: 'center', fontWeight: 'bold' }}>
        <h2>Пожалуйста, нажмите <b>Start</b> в Telegram-боте и перезапустите мини-приложение.</h2>
      </div>
    );
  }

  if (loading || !customer_uuid) {
    return <LoadingScreen message="Авторизация..." />;
  }

  if (error) {
    return (
      <div style={{ padding: '10%', textAlign: 'center', color: 'red', fontWeight: 'bold' }}>
        <h2>{error}</h2>
      </div>
    );
  }

  const handleTableIdSubmit = async (tableId) => {
    try {
      const res = await axios.post(`${SERVER_URL}/customer/v1/order/create`, 
        {
          table_id: tableId
        },
        {
          headers: {
              'Customer-UUID': customer_uuid,
              'JWT-Token': 'bla-bla-bla'
          }
        }
      );
      setOrderUuid(res.data.order_uuid);
      navigate('/catalog');
    } catch (err) {
      if (err.response?.data?.need_room_code) {
        setStep(2); 
      } else {
        setError('Ошибка 2-ая: ' + (err.response?.data?.error || 'не удалось создать заказ'));
      }
    }
  };

  const handleRoomCodeSubmit = async (roomCode) => {
    try {
      const res = await axios.post(`${SERVER_URL}/customer/v1/order/create`, 
        {
          table_id: '',
          room_code: roomCode
        },
        {
          headers: {
            'Customer-UUID': customer_uuid,
            'JWT-Token': 'bla-bla-bla'
          }
        }
      );
      setOrderUuid(res.data.order_uuid);
      setRoomCode(roomCode);
      navigate('/catalog');
    } catch (err) {
      setError('Неверный код комнаты');
    }
  };

  return (
    <div>
      {step === 1 && <TableId onSubmit={handleTableIdSubmit} error={error} />}
      {step === 2 && <RoomCode onSubmit={handleRoomCodeSubmit} error={error} />}
    </div>
  );
}

export default OrderFlow;
