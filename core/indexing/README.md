# Database Indexing

## Context
This is one of core principle to accelerate the performance of database queries, without scatter all rows in table. If you have 1 billion rows in a table, what strategy you will use to reduce the rows for query?


### B-TREE
- Column: each column value is a key in B-tree 
- Composite of multiple columns: column_A, column_B as a key in B-tree, this reduce the bitmap index execute AND bitwise between individual column indexing.

Good for range value (>= and <=) or exactly equal

### Hashing
- Find exactly equal what value is (==), good for random value such as uuid, hashing value

## Inverted index
- Use mapping structure for each word to row id, and do intersect between these mapping for a sentense.
- hello: [row1, row2]
- worlds: [row1]
=> intersection: [row1]

## Geospatial index