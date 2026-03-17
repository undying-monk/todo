import React, { useState, useEffect } from 'react';
import { NavLink } from 'react-router-dom';

const CategoryNavBar = () => {
  const categories = [
    { id: '1', name: 'Thế giới', label: 'World' },
    { id: '2', name: 'Thời sự', label: 'Politics' },
    { id: '3', name: 'Kinh doanh', label: 'Business' },
    { id: '4', name: 'Giải trí', label: 'Entertainment' },
    { id: '5', name: 'Thể thao', label: 'Sports' },
    { id: '6', name: 'Pháp luật', label: 'Law' },
    { id: '7', name: 'Giáo dục', label: 'Education' },
    { id: '8', name: 'Sức khỏe', label: 'Health' },
    { id: '9', name: 'Đời sống', label: 'Life' },
    { id: '10', name: 'Du lịch', label: 'Travel' },
    { id: '11', name: 'Khoa học', label: 'Science' },
    { id: '12', name: 'Công nghệ', label: 'Technology' },
    { id: '13', name: 'Gia đình', label: 'Family' },
    { id: '14', name: 'Khác', label: 'Others' }
  ];

  return (
    <nav className="category-nav">
      <ul className="category-list">
        <li>
          <NavLink 
            to="/" 
            className={({ isActive }) => isActive ? 'category-link active' : 'category-link'}
            end
          >
            Home
          </NavLink>
        </li>
        {categories.map((cat) => (
          <li key={cat.id}>
            <NavLink 
              to={`/category/${encodeURIComponent(cat.label)}`}
              className={({ isActive }) => isActive ? 'category-link active' : 'category-link'}
            >
              {cat.label}
            </NavLink>
          </li>
        ))}
      </ul>
    </nav>
  );
};

export default CategoryNavBar;
