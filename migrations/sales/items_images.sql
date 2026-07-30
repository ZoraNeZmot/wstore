CREATE TABLE IF NOT EXISTS sales.items_images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_id UUID NOT NULL,
    image_path TEXT NOT NULL,
    -- TODO: sort_order INT NOT NULL DEFAULT 0,
    is_primary BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_item_id FOREIGN KEY (item_id) REFERENCES sales.items(id) ON DELETE CASCADE
);

-- FUNCTIONS

CREATE OR REPLACE FUNCTION sales.usf_create_item_image(
    p_id UUID,
    p_item_id UUID,
    p_image_path TEXT,
    p_is_primary BOOLEAN DEFAULT FALSE
) RETURNS UUID
AS $$
DECLARE v_item_image_id UUID;
BEGIN
    INSERT INTO sales.items_images (id, item_id, image_path, is_primary)
    VALUES (p_id, p_item_id, p_image_path, p_is_primary)
    RETURNING id INTO v_item_image_id;
    RETURN v_item_image_id;
END;
$$ LANGUAGE plpgsql;

-- GET FUNCTIONS

CREATE OR REPLACE FUNCTION sales.usf_get_item_images(
    p_item_id UUID
)
RETURNS SETOF text
AS $$
DECLARE v_image_path TEXT;
BEGIN
    RETURN QUERY SELECT image_path FROM sales.items_images WHERE item_id = p_item_id;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION sales.usf_get_item_image_primary(
    p_item_id UUID
)
RETURNS TEXT
AS $$
DECLARE v_image_path TEXT;
BEGIN
    SELECT image_path INTO v_image_path FROM sales.items_images WHERE item_id = p_item_id AND is_primary = TRUE;
    RETURN v_image_path;
END;
$$ LANGUAGE plpgsql;