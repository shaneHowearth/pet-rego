ALTER TABLE pet ADD COLUMN food TEXT;
UPDATE pet SET food='bones' WHERE LOWER(species)='dog';
UPDATE pet SET food='fish' WHERE LOWER(species)='cat';
UPDATE pet SET food='corn' WHERE LOWER(species)='chicken';
UPDATE pet SET food='mice' WHERE LOWER(species)='snake';
