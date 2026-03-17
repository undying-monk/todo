import React, { useState, useEffect, useCallback, useRef } from 'react';
import ArticleCard from './ArticleCard';

const Feed = ({ searchQuery }) => {
  const [articles, setArticles] = useState([]);
  const [cursor, setCursor] = useState(null);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const observer = useRef();

  const fetchArticles = useCallback(async (currentCursor) => {
    setLoading(true);
    try {
      const endpoint = searchQuery 
        ? `http://localhost:8080/api/articles/search?q=${encodeURIComponent(searchQuery)}&limit=10${currentCursor ? `&cursor=${currentCursor}` : ''}`
        : `http://localhost:8080/api/articles?limit=10${currentCursor ? `&cursor=${currentCursor}` : ''}`;
      
      const response = await fetch(endpoint);
      const result = await response.json();
      
      if (result.data && result.data.length > 0) {
        setArticles(prev => currentCursor ? [...prev, ...result.data] : result.data);
        setCursor(result.cursor || null);
        setHasMore(!!result.cursor);
      } else {
        if (!currentCursor) {
          setArticles([]);
        }
        setHasMore(false);
      }
    } catch (error) {
      console.error("Failed to fetch articles:", error);
    } finally {
      setLoading(false);
    }
  }, [searchQuery]);

  const lastArticleRef = useCallback(node => {
    if (loading) return;
    if (observer.current) observer.current.disconnect();
    observer.current = new IntersectionObserver(entries => {
      if (entries[0].isIntersecting && hasMore && cursor) {
        fetchArticles(cursor);
      }
    });
    if (node) observer.current.observe(node);
  }, [loading, hasMore, cursor, fetchArticles]);

  useEffect(() => {
    setArticles([]);
    setCursor(null);
    setHasMore(true);
    fetchArticles(null);
  }, [fetchArticles]);

  return (
    <div className="feed">
      {articles.map((article, index) => {
        if (articles.length === index + 1) {
          return (
            <div ref={lastArticleRef} key={article.ID}>
              <ArticleCard article={article} />
            </div>
          );
        } else {
          return <ArticleCard key={article.ID} article={article} />;
        }
      })}
      
      {loading && <div className="loading">Fetching more news...</div>}
      {!loading && articles.length === 0 && (
        <div className="no-results">No articles found matching your search.</div>
      )}
      {!hasMore && articles.length > 0 && (
        <div className="no-results">You've reached the end of the feed.</div>
      )}
    </div>
  );
};

export default Feed;
