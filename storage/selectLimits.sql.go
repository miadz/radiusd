package storage

//generated by embd
const selectLimits = "SELECT 1\nFROM user\nJOIN product ON user.product_id = product.id\nWHERE user = ?"
