import React, { useState, useEffect } from 'react';
import TableId from '../TableId/TableId';
import RoomCode from '../RoomCode/RoomCode';
import useCustomerAuth from './UseCustomerAuth';
import { useOrder } from '../OrderContext/OrderContext';
import { useNavigate } from 'react-router-dom';
import LoadingScreen from '../LoadingScreen/LoadingScreen';
import { SERVER_URL } from '../../config';

function OrderFlow() {
  const [error, setError] = useState('');
  const [step, setStep] = useState(1);
  const [creatingOrder, setCreatingOrder] = useState(false);
  const [pendingTableId, setPendingTableId] = useState('');

  const {
    setOrderUuid,
    setRoomCode,
    order_uuid,
    room_code,
    customer_uuid,
  } = useOrder();

  const { loading, showStartPrompt } = useCustomerAuth();
  const startParam = window.Telegram?.WebApp?.initDataUnsafe?.start_param;
  const navigate = useNavigate();

  useEffect(() => {
    if (!loading && customer_uuid && order_uuid) {
      navigate('/catalog', { replace: true });
    }
  }, [loading, customer_uuid, order_uuid, navigate]);

  useEffect(() => {
    if (loading || !customer_uuid || order_uuid) return;

    if (startParam) {
      setStep(2);
      setPendingTableId(startParam);
    } else {
      setStep(1);
    }
  }, [loading, customer_uuid, order_uuid, startParam]);

  if (creatingOrder) return <LoadingScreen message="Создаём заказ..." />;
  if (showStartPrompt)
    return (
      <div style={{ padding: '20%', textAlign: 'center', fontWeight: 'bold' }}>
        <h2>Пожалуйста, нажмите <b>Start</b> в Telegram-боте и перезапустите мини-приложение.</h2>
      </div>
    );
  if (loading || !customer_uuid)
    return <LoadingScreen message="Авторизация..." />;
  if (error)
    return (
      <div style={{ padding: '10%', textAlign: 'center', color: 'red', fontWeight: 'bold' }}>
        <h2>{error}</h2>
      </div>
    );

  const handleTableIdSubmit = async (tableId) => {
    setError('');
    setCreatingOrder(true);

    try {
      const res = await fetch(`${SERVER_URL}/customer/v1/order/create`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Customer-UUID': customer_uuid,
          'JWT-Token': 'bla-bla-bla',
        },
        body: JSON.stringify({ table_id: tableId }),
      });

      const data = await res.json();
      if (res.ok) {
        setOrderUuid(data.order_uuid);
        setRoomCode(data.room_code);
        navigate('/catalog');
      } else {
        if (data.need_room_code) {
          setStep(2);
          setPendingTableId(tableId);
        } else {
          setError('Ошибка: ' + (data.error || 'не удалось создать заказ'));
        }
      }
    } catch (e) {
      setError('Ошибка соединения с сервером');
    } finally {
      setCreatingOrder(false);
    }
  };

  const handleRoomCodeSubmit = async (inputRoomCode) => {
    setError('');
    setCreatingOrder(true);

    try {
      const res = await fetch(`${SERVER_URL}/customer/v1/order/create`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Customer-UUID': customer_uuid,
          'JWT-Token': 'bla-bla-bla',
        },
        body: JSON.stringify({
          table_id: pendingTableId,
          room_code: inputRoomCode,
        }),
      });

      const data = await res.json();
      if (res.ok) {
        setOrderUuid(data.order_uuid);
        setRoomCode(inputRoomCode); 
        navigate('/catalog');
      } else {
        setError('Неверный код комнаты');
      }
    } catch (e) {
      setError('Ошибка соединения с сервером');
    } finally {
      setCreatingOrder(false);
    }
  };

  return (
    <div>
      {step === 1 && (
        <TableId onSubmit={handleTableIdSubmit} error={error} />
      )}
      {step === 2 && (
        <RoomCode
          onSubmit={handleRoomCodeSubmit}
          error={error}
        />
      )}
    </div>
  );
}

export default OrderFlow;

