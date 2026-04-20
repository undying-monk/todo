import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, useNavigate, useLocation } from 'react-router-dom';
import SearchBar from './components/SearchBar';
import Feed from './components/Feed';
import HomeFeed from './components/HomeFeed';
import CategoryNavBar from './components/CategoryNavBar';

function AppContent() {
  const [searchQuery, setSearchQuery] = useState('');
  const navigate = useNavigate();
  const location = useLocation();

  const handleSearch = (query) => {
    setSearchQuery(query);
    if (query && location.pathname !== '/search') {
      navigate('/search');
    } else if (!query && location.pathname === '/search') {
      navigate('/');
    }
  };

  return (
    <div className="app-container">
      <header>
        <div className="header-top">
          <h1 style={{ cursor: 'pointer', margin: 0 }} onClick={() => { setSearchQuery(''); navigate('/'); }}>
            News Aggregator
          </h1>
          <SearchBar onSearch={handleSearch} initialQuery={searchQuery} />
        </div>
        <CategoryNavBar />
      </header>
      <main>
        <Routes>
          <Route path="/" element={<HomeFeed />} />
          <Route path="/search" element={<Feed searchQuery={searchQuery} />} />
          <Route path="/category/:categoryName" element={<Feed searchQuery="" />} />
        </Routes>
      </main>
    </div>
  );
}

function App() {
  return (
    <Router>
      <AppContent />
    </Router>
  );
}

export default App;
