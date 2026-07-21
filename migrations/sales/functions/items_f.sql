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
    p_status sales.item_status DEFAULT 'draft'
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

-- Get item
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
    RETURN QUERY SELECT seller_id, title, description, price, quantity, status FROM sales.items WHERE id = p_item_id;
END;
$$ LANGUAGE plpgsql;

-- Get all items
CREATE OR REPLACE FUNCTION sales.usf_get_items_by_seller(
    p_seller_id UUID
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
    RETURN QUERY SELECT seller_id, title, description, price, quantity, status FROM sales.items WHERE seller_id = p_seller_id;
END;
$$ LANGUAGE plpgsql;
