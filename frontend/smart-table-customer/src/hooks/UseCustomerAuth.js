import { useEffect, useState } from 'react';
import axios from 'axios';

const useCustomerAuth = () => {
  const [customerUuid, setCustomerUuid] = useState(null);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  const [showStartPrompt, setShowStartPrompt] = useState(false);

  useEffect(() => {
    const tgUser = window.Telegram?.WebApp?.initDataUnsafe?.user;

    const mockUser = tgUser || {
      id: '123456',
      username: 'testuser',
      photo_url: '',
    };
  
    const payload = {
      tg_id: String(mockUser.id),
      tg_login: mockUser.username,
      chat_id: String(mockUser.id),
    };

    // if (!tgUser || !tgUser.id || !tgUser.username) {
    //   setShowStartPrompt(true);
    //   setLoading(false);
    //   return;
    // }

    // const payload = {
    //   tg_id: String(tgUser.id),
    //   tg_login: tgUser.username,
    //   chat_id: String(tgUser.id),
    // };

  //   axios.post('https://d53f-138-124-112-105.ngrok-free.app/customer/v1/order/customer/sign-in', payload)
  //     .then((res) => {
  //       console.log("Sign-in success", res.data);
  //       setCustomerUuid(res.data.customer_uuid);
  //       setLoading(false);
  //     })
  //     .catch((err) => {
  //       if (err.response && err.response.status === 404) {
  //         const formData = new FormData();
  //         formData.append('tg_id', payload.tg_id);
  //         formData.append('tg_login', payload.tg_login);
  //         formData.append('chat_id', payload.chat_id);
  //         formData.append('avatar', mockUser.photo_url || '');

  //         axios.post('https://d53f-138-124-112-105.ngrok-free.app/customer/v1/order/customer/sign-up', formData)
  //           .then((res) => {
  //             console.log("Sign-up success", res.data);
  //             setCustomerUuid(res.data.customer_uuid)
  //           })
  //           .catch((err) => {
  //             console.error("Sign-up error:", err);
  //             setError('Ошибка регистрации');
  //           })
  //           .finally(() => setLoading(false));
  //       } else {
  //         setError('Ошибка авторизации');
  //         setLoading(false);
  //       }
  //     });
  }, []);

  return { customerUuid, loading, error, showStartPrompt };
};

export default useCustomerAuth;
