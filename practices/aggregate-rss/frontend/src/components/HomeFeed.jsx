import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import ArticleCard from './ArticleCard';

const HomeFeed = () => {
  const [categories, setCategories] = useState({});
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchArticles = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/articles?limit=200');
        const result = await response.json();
        
        if (result.data) {
          const grouped = {};
          result.data.forEach(article => {
            const cat = article.Category || 'Uncategorized';
            if (!grouped[cat]) {
              grouped[cat] = [];
            }
            if (grouped[cat].length < 10) {
              grouped[cat].push(article);
            }
          });
          setCategories(grouped);
        }
      } catch (error) {
        console.error("Failed to fetch articles:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchArticles();
  }, []);

  if (loading) {
    return <div className="loading">Loading feed...</div>;
  }

  return (
    <div className="home-feed">
      {Object.entries(categories).map(([category, articles]) => (
        <div key={category} className="category-section" style={{ marginBottom: '3rem' }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '1.5rem' }}>
            <h2>{category}</h2>
            <Link to={`/category/${encodeURIComponent(category)}`} style={{ color: 'var(--accent-color)', textDecoration: 'none', fontWeight: '600' }}>
              View All &rarr;
            </Link>
          </div>
          <div className="feed">
            {articles.map(article => (
              <ArticleCard key={article.ID} article={article} />
            ))}
          </div>
        </div>
      ))}
      {Object.keys(categories).length === 0 && (
        <div className="no-results">No articles found.</div>
      )}
    </div>
  );
};

export default HomeFeed;
