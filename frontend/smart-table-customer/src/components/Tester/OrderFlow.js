// import React, { useState } from 'react';
// import axios from 'axios';
// import TableId from '../TableId/TableId';
// import RoomCode from '../RoomCode/RoomCode';
// import useCustomerAuth from '../../hooks/UseCustomerAuth';

// function OrderFlow() {
//   const [tableId, setTableId] = useState('');
//   const [roomCode, setRoomCode] = useState('');
//   const [isTableValid, setIsTableValid] = useState(false);
//   const [error, setError] = useState('');

//   const { customerUuid, loading, showStartPrompt } = useCustomerAuth();

//   const handleTableIdSubmit = async (id) => {
//     try {
//       await axios.post('https://d53f-138-124-112-105.ngrok-free.app/customer/v1/order/create', {
//         table_id: id,
//         room_code: '',
//         customer_uuid: customerUuid
//       });
//       setTableId(id);
//       setIsTableValid(true);
//       setError('');
//     } catch {
//       setError('Неверный номер стола');
//     }
//   };

//   const handleRoomCodeSubmit = async (code) => {
//     try {
//       await axios.post('https://d53f-138-124-112-105.ngrok-free.app/customer/v1/order/create', {
//         table_id: tableId,
//         room_code: code,
//         customer_uuid: customerUuid
//       });
//       setRoomCode(code);
//       window.location.href = '/catalog';
//     } catch {
//       setError('Неверный код комнаты');
//     }
//   };

//   // if (showStartPrompt) {
//   //   return <div style={{ padding: '20%', textAlign: 'center' }}><h2>Нажмите "Start" в боте и перезапустите мини-приложение</h2></div>;
//   // }

//   if (loading || !customerUuid) {
//     return <div style={{ padding: '10%', textAlign: 'center' }}><h2>Загрузка...</h2></div>;
//   }

//   return (
//     <div>
//       {!isTableValid
//         ? <TableId onSubmit={handleTableIdSubmit} error={error} />
//         : <RoomCode onSubmit={handleRoomCodeSubmit} error={error} />}
//     </div>
//   );
// }

// export default OrderFlow;

import React, { useState } from 'react';
import axios from 'axios';
import TableId from '../TableId/TableId';
import RoomCode from '../RoomCode/RoomCode';

// Заглушка вместо авторизации
const customerUuid = 'mock-customer-uuid-123';

function OrderFlow() {
  const [tableId, setTableId] = useState('');
  const [roomCode, setRoomCode] = useState('');
  const [isTableValid, setIsTableValid] = useState(false);
  const [error, setError] = useState('');

  const handleTableIdSubmit = async (id) => {
    try {
      await axios.post('https://d53f-138-124-112-105.ngrok-free.app/customer/v1/order/create', {
        table_id: id,
        room_code: '', // room_code пока пустой
        customer_uuid: customerUuid
      });
      setTableId(id);
      setIsTableValid(true);
      setError('');
    } catch {
      setError('Неверный номер стола');
    }
  };

  const handleRoomCodeSubmit = async (code) => {
    try {
      await axios.post('https://d53f-138-124-112-105.ngrok-free.app/customer/v1/order/create', {
        table_id: tableId,
        room_code: code,
        customer_uuid: customerUuid
      });
      setRoomCode(code);
      window.location.href = '/catalog';
    } catch {
      setError('Неверный код комнаты');
    }
  };

  return (
    <div>
      {!isTableValid
        ? <TableId onSubmit={handleTableIdSubmit} error={error} />
        : <RoomCode onSubmit={handleRoomCodeSubmit} error={error} />}
    </div>
  );
}

export default OrderFlow;
