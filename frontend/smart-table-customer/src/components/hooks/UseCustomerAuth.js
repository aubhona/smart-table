import { useEffect, useState } from 'react';
import axios from 'axios';
import { useOrder } from '../OrderContext/OrderContext';
import { SERVER_URL } from '../../config';

const useCustomerAuth = () => {
  const { customer_uuid, setCustomerUuid, setJwtToken } = useOrder();
  const [loading, setLoading] = useState(true);
  const [showStartPrompt, setShowStartPrompt] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    const tgUser = window.Telegram?.WebApp?.initDataUnsafe?.user;
    if (!tgUser || !tgUser.id || !tgUser.username) {
      setShowStartPrompt(true);
      setLoading(false);
      return;
    }

    const payload = {
      tg_id: String(tgUser.id),
      tg_login: tgUser.username,
      init_data: window.Telegram?.WebApp?.initData || ''
    };

    axios.post(`${SERVER_URL}/customer/v1/sign-in`, payload)
      .then(res => {
        setCustomerUuid(res.data.customer_uuid);
        if (setJwtToken && res.data.jwt_token) {
          setJwtToken(res.data.jwt_token);
          localStorage.setItem('jwt_token', res.data.jwt_token);
        }
        setLoading(false);
      })
      .catch(err => {
        if (err.response?.status === 404) {
          setShowStartPrompt(true);
          setError('Пожалуйста, нажмите /start в боте, отсканируйте QR и введите Table ID заново.');
        } else {
          setError('Ошибка авторизации');
        }
        setLoading(false);
      });
  }, [setCustomerUuid, setJwtToken]);

  return { customer_uuid, loading, showStartPrompt, error };
};

export default useCustomerAuth;
