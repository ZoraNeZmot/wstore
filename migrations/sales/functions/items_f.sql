CREATE OR REPLACE FUNCTION sales.usf_create_item(
    p_seller_id UUID, 
    p_title VARCHAR(255),
    p_description TEXT,
    p_price DECIMAL(10, 2),
    p_quantity INT,
    p_status sales.item_status DEFAULT 'draft'
)
RETURNS UUID
AS $$
DECLARE v_item_id UUID;
BEGIN
    INSERT INTO sales.items (seller_id, title, description, price, quantity, status) VALUES (p_seller_id, p_title, p_description, p_price, p_quantity, p_status) RETURNING id INTO v_item_id;
    RETURN v_item_id;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE sales.usp_update_item(
    p_item_id UUID,
    p_title VARCHAR(255),
    p_description TEXT,
    p_price DECIMAL(10, 2),
    p_quantity INT,
    p_status sales.item_status
)
AS $$
BEGIN
    UPDATE sales.items SET title = p_title, description = p_description, price = p_price, quantity = p_quantity, status = p_status WHERE id = p_item_id;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE sales.usp_delete_item(
    p_item_id UUID
)
AS $$
BEGIN
    UPDATE sales.items SET status = 'deleted' WHERE id = p_item_id;
END;
$$ LANGUAGE plpgsql;

-- Get item full (seller_id, title, description, price, quantity, status)
CREATE OR REPLACE FUNCTION sales.usf_get_item(
    p_item_id UUID
)
RETURNS TABLE(
    seller_id UUID,
    title VARCHAR(255),
    description TEXT,
    price DECIMAL(10, 2),
    quantity INT,
    status sales.item_status
)
AS $$
BEGIN
    RETURN QUERY SELECT i.seller_id, i.title, i.description, i.price, i.quantity, i.status FROM sales.items AS i WHERE i.id = p_item_id;
END;
$$ LANGUAGE plpgsql;


-- Get short items (id, title, price, status)
CREATE OR REPLACE FUNCTION sales.usf_get_items_by_seller(
    p_seller_id UUID,
    p_amount INT DEFAULT 10
)
RETURNS TABLE(
    id UUID,
    title VARCHAR(255),
    price DECIMAL(10, 2),
    status sales.item_status
)
AS $$
BEGIN
    RETURN QUERY 
        SELECT i.id, i.title, i.price, i.status FROM sales.items AS i
        WHERE i.seller_id = p_seller_id 
        LIMIT p_amount;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION sales.usf_get_items(
    p_amount INT DEFAULT 10
)
RETURNS TABLE(
    id UUID,
    seller_id UUID,
    title VARCHAR(255),
    price DECIMAL(10, 2),
    status sales.item_status
)
AS $$
BEGIN
    RETURN QUERY SELECT i.id, i.seller_id, i.title, i.price, i.status FROM sales.items AS i LIMIT p_amount;
END;
$$ LANGUAGE plpgsql;