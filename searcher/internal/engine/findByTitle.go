package engine

// SQL req

// for single row (url, title, html-code)
// SELECT FROM sites * WHERE lower(<row>) LIKE "%<substring>%";

// for row with array (keywords)
// SELECT * FROM sites WHERE EXISTS (SELECT 1 FROM unnest(keywords) AS element WHERE lower(element::text) LIKE lower('%<substring>%'));
