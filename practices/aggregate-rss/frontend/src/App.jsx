import React, { useState } from 'react';
import SearchBar from './components/SearchBar';
import Feed from './components/Feed';

function App() {
  const [searchQuery, setSearchQuery] = useState('');

  return (
    <div className="app-container">
      <header>
        <h1>News Aggregator</h1>
        <SearchBar onSearch={setSearchQuery} />
      </header>
      <main>
        <Feed searchQuery={searchQuery} />
      </main>
    </div>
  );
}

export default App;
