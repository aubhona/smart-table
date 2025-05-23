export function getAuthHeaders({ customer_uuid, jwt_token, order_uuid }, withOrder = true) {
  const headers = {
    'Customer-UUID': customer_uuid || localStorage.getItem('customer_uuid') || '',
    'JWT-Token': jwt_token || localStorage.getItem('jwt_token') || '',
  };
  if (withOrder && (order_uuid || localStorage.getItem('order_uuid'))) {
    headers['Order-UUID'] = order_uuid || localStorage.getItem('order_uuid');
  }
  return headers;
} 