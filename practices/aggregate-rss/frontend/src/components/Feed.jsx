import React, { useState, useEffect, useCallback, useRef } from 'react';
import ArticleCard from './ArticleCard';

const Feed = ({ searchQuery }) => {
  const [articles, setArticles] = useState([]);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const observer = useRef();

  const lastArticleRef = useCallback(node => {
    if (loading) return;
    if (observer.current) observer.current.disconnect();
    observer.current = new IntersectionObserver(entries => {
      if (entries[0].isIntersecting && hasMore) {
        setPage(prevPage => prevPage + 1);
      }
    });
    if (node) observer.current.observe(node);
  }, [loading, hasMore]);

  const fetchArticles = useCallback(async () => {
    setLoading(true);
    try {
      const endpoint = searchQuery 
        ? `http://localhost:8080/api/articles/search?q=${encodeURIComponent(searchQuery)}&page=${page}&limit=10`
        : `http://localhost:8080/api/articles?page=${page}&limit=10`;
      
      const response = await fetch(endpoint);
      const result = await response.json();
      
      if (result.data) {
        setArticles(prev => page === 1 ? result.data : [...prev, ...result.data]);
        setHasMore(result.data.length === 10);
      } else {
        setHasMore(false);
      }
    } catch (error) {
      console.error("Failed to fetch articles:", error);
    } finally {
      setLoading(false);
    }
  }, [searchQuery, page]);

  useEffect(() => {
    setArticles([]);
    setPage(1);
    setHasMore(true);
  }, [searchQuery]);

  useEffect(() => {
    fetchArticles();
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
