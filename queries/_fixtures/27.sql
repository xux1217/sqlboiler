SELECT "c".*, "d".* FROM cats as c, dogs as d FULL JOIN dogs d on d.cat_id = cats.id;