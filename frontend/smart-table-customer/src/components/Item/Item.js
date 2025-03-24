import React, { useState }from "react";
import { useParams, useNavigate } from "react-router-dom";
import "./Item.css";

const dishes = [
    { id: 1, category: "novinki", name: "Бурурброт", price: 10, calories: 250 },
    { id: 2, category: "novinki", name: "Плов", price: 12, calories: 300 },
    { id: 3, category: "pervoe", name: "Супчок", price: 15, calories: 400 },
    { id: 4, category: "vtoroe", name: "Котлета с пюре", price: 20, calories: 500 },
    { id: 5, category: "napitki", name: "Кофе", price: 17, calories: 200 },
    { id: 6, category: "napitki", name: "Чай", price: 10, calories: 200 },
    { id: 7, category: "deserty", name: "Блины", price: 50, calories: 300 },
    { id: 8, category: "pervoe", name: "Борщок", price: 100, calories: 200 },
    { id: 9, category: "novinki", name: "Яблоко", price: 6, calories: 100 },
    { id: 10, category: "holodnoe", name: "Сок", price: 6, calories: 100 },
  ];

function Item() {
    const { id } = useParams();
    const [cart, setCart] = useState({});
    const navigate = useNavigate();
    const dish = dishes.find(d => d.id === Number(id));
  
    if (!dish) return <div>Блюдо не найдено</div>;

    const updateQuantity = (id, change) => {
        setCart((prev) => {
          const newQuantity = (prev[id] || 1) + change;
          if (newQuantity <= 0) {
            const updatedCart = { ...prev };
            delete updatedCart[id]; 
            return updatedCart;
          }
          return { ...prev, [id]: newQuantity };
        });
      };
  
    return (
    <div className="item-container">
        <div className="top-bar">
          <button className="top-button" onClick={() => navigate(-1)}>назад</button>
          <button className="top-button">официант</button>
        </div>
  
        <div className="dish-image-item">Фотка {dish.name}</div>
  
        <div className="dish-info-item">
          <p className="description-item">{dish.description || "Описание не указано"}</p>
          <p className="composition-item">Состав: {dish.composition || "Не указан"}</p>
          
          <textarea placeholder="Комментарий к заказу" />
        </div>
  
        <div className="item-footer">
            <div className="item-summary">
                <div className="dish-name-item">{dish.name}</div>
                <div className="calories-item">{dish.calories} грамм</div>
                <div className="price-item">{dish.price} ₽</div>
            </div>
            <div className="item-actions">
                <div className="quantity-controls-item">
                    <button onClick={() => updateQuantity(dish.id, -1)}>-</button>
                    <span><strong>{cart[dish.id] || 1}</strong></span>
                    <button onClick={() => updateQuantity(dish.id, 1)}>+</button>
                </div>
                <button 
                    className="add-button" 
                    // onClick={() => addToCart(dish.id)} 
                    disabled={cart[dish.id] <= 0}
                >
                    Добавить
                </button>
                </div>
            </div>
        </div>
    );
}
  
  export default Item;