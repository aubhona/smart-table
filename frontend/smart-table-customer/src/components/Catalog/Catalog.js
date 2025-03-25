import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./Catalog.css";

const categories = [
  { id: "novinki", name: "Новинки" },
  { id: "pervoe", name: "Первое" },
  { id: "vtoroe", name: "Второе" },
  { id: "napitki", name: "Напитки" },
  { id: "deserty", name: "Десерты" },
  { id: "holodnoe", name: "холодное" },
];

const dishes = [
  { id: 1, category: "novinki", name: "Бурурброт", price: 10, calories: 250 },
  { id: 2, category: "novinki", name: "Плов", price: 1200, calories: 300 },
  { id: 3, category: "pervoe", name: "Супчок", price: 15, calories: 400 },
  { id: 4, category: "vtoroe", name: "Котлета с пюре", price: 20, calories: 500 },
  { id: 5, category: "napitki", name: "Кофе", price: 17, calories: 200 },
  { id: 6, category: "napitki", name: "Чай", price: 10, calories: 200 },
  { id: 7, category: "deserty", name: "Блины", price: 50, calories: 300 },
  { id: 8, category: "pervoe", name: "Борщок", price: 100, calories: 200 },
  { id: 9, category: "novinki", name: "Яблоко", price: 6, calories: 100 },
  { id: 10, category: "holodnoe", name: "Сок", price: 6, calories: 100 },
];
function Catalog() {
  const navigate = useNavigate();
  const [cart, setCart] = useState({});

  const updateQuantity = (id, change) => {
    setCart((prev) => {
      const newQuantity = (prev[id] || 0) + change;
      if (newQuantity <= 0) {
        const updatedCart = { ...prev };
        delete updatedCart[id]; 
        return updatedCart;
      }
      return { ...prev, [id]: newQuantity };
    });
  };

  const handleItemClick = (id) => {
    navigate(`/catalog/item/${id}`); 
  };

  const totalPrice = Object.keys(cart).reduce(
    (sum, id) => sum + (cart[id] || 0) * dishes.find((dish) => dish.id === Number(id)).price,
    0
  );

  const handleScroll = (id) => {
    const element = document.getElementById(id);
    if (element) {
      const offset = 150; 
      const elementPosition = element.getBoundingClientRect().top + window.scrollY;
      window.scrollTo({
        top: elementPosition - offset,
        behavior: "smooth", 
      });
    }
  };

  const handleGoToCart = () => {
    navigate('/cart', { state: { cart } }); 
  };

  const handleGoToUsers = () => {
    navigate('/catalog/users-list'); 
  };

  return (
    <div className="catalog-container">
      <div className="top-buttons">
        <button className="small-button">Удалить сессию</button>
        <button className="small-button" onClick={handleGoToUsers}>Код комнаты</button>
        <button className="small-button">Официант</button>
        <button className="small-button">Чекаут гоу</button>
      </div>

      <div className="category-tabs">
        <div className="category-scroll">
          {categories.map((cat) => (
            <button key={cat.id} onClick={() => handleScroll(cat.id)} className="category-link">
              {cat.name}
            </button>
          ))}
        </div>
      </div>

      {categories.map((cat) => (
          <div key={cat.id} id={cat.id} className="category-section">
            <h2 className="category-title">{cat.name}</h2>
                <div className="menu-grid">
                  {dishes
                    .filter((dish) => dish.category === cat.id)
                    .map((dish) => (
                      <div key={dish.id} className="menu-item" onClick={() => handleItemClick(dish.id)}>
                        <div className="dish-img">Фотка жоского блюда</div>
                        <div className="dish-info">
                          <p className="dish-price"><strong>{dish.price} ₽</strong></p>
                          <p className="dish-name">{dish.name}</p>
                          <p className="dish-calories">{dish.calories} грамм</p>
                      </div>
                    <div className="quantity-controls">
                      {cart[dish.id] ? (
                        <>
                          <button onClick={(e) => { e.stopPropagation(); updateQuantity(dish.id, -1); }}>-</button>
                          <span><strong>{cart[dish.id]}</strong></span>
                          <button onClick={(e) => { e.stopPropagation(); updateQuantity(dish.id, 1); }}>+</button>
                        </>
                      ) : (
                        <button className="add-button" onClick={(e) => { e.stopPropagation(); updateQuantity(dish.id, 1); }}>+</button>
                      )}
                    </div>
                  </div>
                ))}
            </div>
          </div>
        ))}
        <div className="scroll-padding"></div>
        
      <div className="total-price">
        <p>Итого: <strong>{totalPrice} ₽</strong></p>
        <button className="checkout-button" onClick={handleGoToCart}>Далее</button>
      </div>
    </div>
  );
}

export default Catalog;
