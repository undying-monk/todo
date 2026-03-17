# Technical Requirements: Real-Time News Aggregation Platform

## Overview
A real-time news aggregation platform designed to collect, store, and display articles from multiple RSS feeds (e.g., newspaper sites). The platform provides a seamless reading experience with chronological sorting, infinite scrolling, and cross-site text search.

---

## Backend (BE) Agent Requirements

### 1. RSS Feed Collector & Scheduler
- **Job Scheduler**: Implement a scheduling system (e.g., Cron job, Celery, or background worker) that runs at a configurable interval (e.g., every 1 hour).
- **RSS Parsing**: Capable of reading and parsing standard RSS/Atom feeds from multiple configured sources.
- **Data Deduplication**: Ensure that the same article from a feed is not saved multiple times (use GUID or article URL as a unique identifier).

### 2. Data Storage
- **Database**: Use a suitable database (PostgreSQL, MongoDB, etc.) to store article metadata.
- **Schema/Model**:
  - `content_snippet`: Brief description or snippet of the article
  - `category`: Category of the article (e.g., News, Sports, Tech, etc.)
  - `favicon`: URL to the source's favicon
- **Indices**: Ensure indices on `published_at` (for sorting) and text-search indices on `title`/`content_snippet` (for fast searching).

### 3. RESTful API Endpoints
- **List Articles API**: 
  - Endpoint: `GET /api/articles`
  - Features: Returns articles sorted by `published_at` in descending order. Must support pagination (e.g., `limit` and `offset` or cursor-based pagination).
- **Search Articles API**: 
  - Endpoint: `GET /api/articles/search`
  - Query Params: `q` (search term), pagination parameters.
  - Features: Performs text search across `title` and `content_snippet` for the given keyword across all sources.

---e

## Frontend (FE) Agent Requirements

### 1. Feed / Home Page
- **Article List**: Display the fetched articles in a clean, readable feed layout, sorted by their published timestamp (newest first).
- **Article Card Component**: Each article should display:
  - Title
  - Source site (e.g., "The New York Times") and its favicon
  - Publishing timestamp (e.g., "2 hours ago")
  - Snippet/Description summary
- **Redirection**: Clicking on an article card MUST open the original article link in a new tab (`target="_blank"`).

### 2. Scrolling & Pagination
- **Infinite Scrolling**: Automatically fetch and append the next page of articles from the BE as the user scrolls to the bottom of the current list.

### 3. Search Functionality
- **Global Search Bar**: A prominently placed search bar at the top of the interface.
- **Real-Time or Submit Search**: Allow the user to type keywords to search for articles. 
- **Search Results**: Replace the main feed with the search results corresponding to the keyword, while maintaining the same article card layout and click-through functionality.

---

## Task Lists

### Phase 1: Planning & Setup
- [x] Requirements Gathering & Setup (Go/go-gin)
- [x] Initialize BE project
- [x] Initialize FE project
- [x] Setup Database (PostgreSQL/MongoDB)
- [x] Configure environment variables for API keys and DB connection
- [x] Create Dockerfile for Backend
- [x] Create Dockerfile for Frontend
- [x] Create docker-compose.yml for orchestration (FE, BE, Database)

### Phase 2: Backend (BE) Implementation
- [x] Setup Database Schema/Models
  - `id`, `title`, `link`, `source`, `favicon`, `published_at`, `content_snippet`, `category`
  - Add indexing for `published_at` and text search
- [x] Implement RSS Feed Collector Service
  - [x] Add parsing logic for `https://vnexpress.net/rss`
  - [x] Add parsing logic for `https://tuoitre.vn/rss.htm`
  - [x] Add parsing logic for `https://thanhnien.vn/rss.html`
- [x] Implement Job Scheduler to run collector every 1 hour
- [x] Implement Data Deduplication logic based on article link/GUID
- [x] Create RESTful API Endpoints
  - [x] `GET /api/articles` (List with pagination & sorting)
  - [x] `GET /api/articles/search` (Text search with pagination)
- [x] Implement Article Categorization logic
  - [x] Categorize common topics (e.g., World, Politics, Sports, Technology, Entertainment)
  - [x] Map RSS feed categories to the internal categorization system

### Phase 3: Frontend (FE) Implementation
- [x] Setup core layout, styling (CSS/Tailwind) and routing
- [x] Build Article Card Component
- [x] Implement Feed / Home Page to display articles
- [x] Integrate API `GET /api/articles`
- [x] Implement Infinite Scrolling / Pagination
- [x] Build Global Search Bar component
- [x] Integrate API `GET /api/articles/search`
- [x] Ensure click-through functionality (`target="_blank"`)

### Phase 4: Verification & Polish
- [ ] Test BE RSS parsing and deduplication
- [ ] Test FE infinite scrolling and search
- [ ] End-to-end integration testing
- [ ] Test Docker orchestration
- [ ] UI/UX polishing
