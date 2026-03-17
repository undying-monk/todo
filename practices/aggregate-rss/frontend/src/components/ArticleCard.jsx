import React from 'react';

const ArticleCard = ({ article }) => {
  const formatDate = (dateString) => {
    const date = new Date(dateString);
    const now = new Date();
    const diffInSeconds = Math.floor((now - date) / 1000);

    if (diffInSeconds < 60) return 'just now';
    if (diffInSeconds < 3600) return `${Math.floor(diffInSeconds / 60)}m ago`;
    if (diffInSeconds < 86400) return `${Math.floor(diffInSeconds / 3600)}h ago`;
    if (diffInSeconds < 604800) return `${Math.floor(diffInSeconds / 86400)}d ago`;

    return date.toLocaleDateString();
  };

  return (
    <a 
      href={article.Link} 
      target="_blank" 
      rel="noopener noreferrer" 
      className="article-card"
    >
      {article.ThumbnailURL && (
        <div className="article-thumbnail-container">
          <img 
            src={article.ThumbnailURL} 
            alt={article.Title} 
            className="article-thumbnail"
            onError={(e) => e.target.style.display = 'none'}
          />
        </div>
      )}
      <div className="article-content">
        <div className="article-header">
          <span className="article-source">{article.Source}</span>
          <span className="article-date">{formatDate(article.PublishedAt)}</span>
        </div>
        <h2 className="article-title">{article.Title}</h2>
      </div>
    </a>
  );
};

export default ArticleCard;
