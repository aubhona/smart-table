import { useEffect, useState } from 'react';
import axios from 'axios';

const useCustomerAuth = () => {
  const [customerUuid, setCustomerUuid] = useState(null);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);
  const [showStartPrompt, setShowStartPrompt] = useState(false);

  useEffect(() => {
    const tgUser = window.Telegram?.WebApp?.initDataUnsafe?.user;

    // const mockUser = tgUser || {
    //   id: '123456',
    //   username: 'testuser',
    //   photo_url: '',
    // };
  
    // const payload = {
    //   tg_id: String(mockUser.id),
    //   tg_login: mockUser.username,
    //   chat_id: String(mockUser.id),
    // };

    if (!tgUser || !tgUser.id || !tgUser.username) {
      setShowStartPrompt(true);
      setLoading(false);
      return;
    }

    const payload = {
      tg_id: String(tgUser.id),
      tg_login: tgUser.username,
      chat_id: String(tgUser.id),
      avatar: tgUser.photo_url || '',
    };

    axios.post('https://d53f-138-124-112-105.ngrok-free.app/customer/v1/sign-in', payload)
      .then((res) => {
        setCustomerUuid(res.data.customer_uuid);
        setLoading(false);
      })
      .catch((err) => {
        if (err.response && err.response.status === 404) {
          axios.post('https://d53f-138-124-112-105.ngrok-free.app/customer/v1/sign-up', payload)
            .then((res) => {
              setCustomerUuid(res.data.customer_uuid);
            })
            .catch(() => {
              setError('Ошибка регистрации');
            })
            .finally(() => setLoading(false));
        } else {
          setError('Ошибка авторизации');
          setLoading(false);
        }
      });
  }, []);


  return { customerUuid, loading, error, showStartPrompt };
};

export default useCustomerAuth;
