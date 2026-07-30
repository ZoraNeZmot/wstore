DO $$
DECLARE
    v_user_id UUID;
    v_user2_id UUID;
    v_user_exists BOOLEAN;
    v_user2_exists BOOLEAN;
BEGIN
	v_user_exists := EXISTS (SELECT 1 FROM accounts.users WHERE email = 'demo@wstore.local');
	v_user2_exists := EXISTS (SELECT 1 FROM accounts.users WHERE email = 'demo2@wstore.local');

	IF NOT v_user_exists THEN
        	v_user_id := accounts.usf_create_user(
        		'demo@wstore.local',
        		'$argon2id$v=19$m=65536,t=4,p=4$cJ2EjXmmtbbeQmiP2TSS4A==$WJv8XkmOmN74NO4KxZDaRHdfLnYlElXhmbtJvCaE6hg=',
        		'First Seller'
        	);
    ELSE
        SELECT id INTO v_user_id FROM accounts.users WHERE email = 'demo@wstore.local';
	END IF;    

    IF NOT v_user2_exists THEN
        v_user2_id := accounts.usf_create_user(
        'demo2@wstore.local',
        '$argon2id$v=19$m=65536,t=4,p=4$PJv6XBU9giD6xmk4xois8A==$FjsytXt6pZhz3j0JprqQn00L1feOlZuW+Myb+1UDfWE=',
        'Second Seller'
        );
    ELSE
        SELECT id INTO v_user2_id FROM accounts.users WHERE email = 'demo2@wstore.local';
    END IF;

    -- First Seller: laptops, headphones, keyboards, mice, monitors
    PERFORM sales.usf_create_item(v_user_id, 'Laptop ASUS Vivobook 15', '15.6" FHD, Ryzen 5, 16GB RAM, 512GB SSD.', 649.00, 12, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Laptop ASUS Zenbook 14', '14" OLED, Intel Core Ultra 7, 16GB RAM, 1TB SSD.', 1099.00, 8, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Laptop Honor MagicBook 14', '14" IPS, Snapdragon X Elite, 16GB RAM, 512GB SSD.', 899.00, 10, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Laptop Lenovo IdeaPad Slim 5', '14" OLED, Ryzen 7, 16GB RAM, 512GB SSD.', 749.00, 14, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Laptop Dell XPS 13', '13.4" InfinityEdge, Intel Core Ultra 7, 32GB RAM.', 1299.00, 6, 'active');

    PERFORM sales.usf_create_item(v_user_id, 'Headphones Sony WH-1000XM5', 'Wireless noise-cancelling over-ear, 30h battery.', 349.99, 20, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Headphones Bose QuietComfort Ultra', 'Spatial audio ANC headphones with 24h playtime.', 379.00, 15, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Headphones Apple AirPods Max', 'Over-ear with Active Noise Cancellation and Spatial Audio.', 499.00, 8, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Headphones JBL Tune 760NC', 'Wireless over-ear with ANC and 35h battery.', 99.99, 35, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Headphones Samsung Galaxy Buds3 Pro', 'True wireless earbuds with ANC and IP57.', 199.00, 28, 'active');

    PERFORM sales.usf_create_item(v_user_id, 'Keyboard Keychron K2 Pro', 'Wireless mechanical, hot-swappable, RGB backlight.', 109.00, 22, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Keyboard Logitech MX Mechanical', 'Low-profile wireless mechanical for office use.', 159.00, 18, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Keyboard Razer BlackWidow V4', 'Full-size RGB mechanical gaming keyboard.', 189.00, 14, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Keyboard Corsair K70 RGB TKL', 'Tenkeyless mechanical with Cherry MX switches.', 149.00, 16, 'draft');
    PERFORM sales.usf_create_item(v_user_id, 'Keyboard Apple Magic Keyboard', 'Wireless scissor-switch keyboard for Mac.', 99.00, 20, 'active');

    PERFORM sales.usf_create_item(v_user_id, 'Mouse Logitech MX Master 3S', 'Ergonomic wireless mouse with MagSpeed scroll.', 99.99, 30, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Mouse Razer DeathAdder V3 Pro', 'Lightweight wireless gaming mouse, 30K DPI.', 149.00, 20, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Mouse Apple Magic Mouse', 'Multi-touch wireless mouse for Mac.', 79.00, 25, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Mouse Microsoft Surface Arc', 'Slim foldable Bluetooth mouse.', 69.99, 22, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Mouse Logitech G Pro X Superlight', 'Ultra-light wireless gaming mouse, 63g.', 129.00, 18, 'active');

    PERFORM sales.usf_create_item(v_user_id, 'Monitor LG UltraGear 27', '27" QHD 165Hz IPS gaming monitor.', 299.00, 10, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Monitor Dell UltraSharp 27', '27" 4K USB-C hub monitor for productivity.', 549.00, 7, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Monitor Samsung Odyssey G5', '32" QHD 165Hz curved gaming monitor.', 329.00, 9, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Monitor ASUS ProArt PA278QV', '27" WQHD color-accurate creative monitor.', 349.00, 8, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Monitor BenQ PD2705U', '27" 4K designer monitor with USB-C.', 449.00, 6, 'active');

    IF v_user_exists THEN
        RAISE NOTICE 'User % already exists', v_user_id;
    ELSE
        RAISE NOTICE 'Created user % with 25 items', v_user_id;
    END IF;

    -- Second Seller: same product types, overlapping + adjacent brands
    PERFORM sales.usf_create_item(v_user2_id, 'Laptop ASUS Vivobook 15', '15.6" FHD, Ryzen 5, 16GB RAM, 512GB SSD.', 639.00, 9, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Laptop ASUS TUF Gaming A15', '15.6" FHD 144Hz, Ryzen 7, RTX 4060, 16GB RAM. The best laptop for gaming.', 999.00, 7, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Laptop Honor MagicBook X16', '16" IPS, Intel Core i5, 16GB RAM, 512GB SSD.', 599.00, 11, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Laptop Lenovo ThinkPad E14', '14" FHD, Intel Core Ultra 5, 16GB RAM, 512GB SSD.', 849.00, 10, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Laptop HP Pavilion 16', '16" FHD, Intel Core i5, 16GB RAM, 512GB SSD.', 689.00, 13, 'active');

    PERFORM sales.usf_create_item(v_user2_id, 'Headphones Sony WH-1000XM5', 'Wireless noise-cancelling over-ear, 30h battery.', 339.99, 16, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Headphones Bose QuietComfort Ultra', 'Spatial audio ANC headphones with 24h playtime.', 369.00, 12, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Headphones Apple AirPods Pro 2', 'In-ear ANC earbuds with USB-C charging case.', 229.00, 22, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Headphones JBL Tune 760NC', 'Wireless over-ear with ANC and 35h battery.', 94.99, 40, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Headphones Sennheiser Momentum 4', 'Audiophile wireless ANC, up to 60h battery.', 289.00, 14, 'active');

    PERFORM sales.usf_create_item(v_user2_id, 'Keyboard Keychron K2 Pro', 'Wireless mechanical, hot-swappable, RGB backlight.', 104.00, 19, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Keyboard Logitech MX Mechanical', 'Low-profile wireless mechanical for office use.', 154.00, 15, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Keyboard Razer Huntsman V3 Pro', 'Analog optical gaming keyboard, adjustable actuation.', 249.00, 8, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Keyboard Corsair K70 RGB TKL', 'Tenkeyless mechanical with Cherry MX switches.', 144.00, 12, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Keyboard Microsoft Designer Compact', 'Slim Bluetooth keyboard for Windows.', 54.99, 30, 'draft');

    PERFORM sales.usf_create_item(v_user2_id, 'Mouse Logitech MX Master 3S', 'Ergonomic wireless mouse with MagSpeed scroll.', 94.99, 28, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Mouse Razer DeathAdder V3 Pro', 'Lightweight wireless gaming mouse, 30K DPI.', 139.00, 17, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Mouse Apple Magic Mouse', 'Multi-touch wireless mouse for Mac.', 74.00, 20, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Mouse Logitech G Pro X Superlight', 'Ultra-light wireless gaming mouse, 63g.', 119.00, 21, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Mouse Keychron M6', 'Wireless ergonomic mouse with 4K polling. Good for laptop users.', 79.00, 24, 'active');

    PERFORM sales.usf_create_item(v_user2_id, 'Monitor LG UltraGear 27', '27" QHD 165Hz IPS gaming monitor.', 289.00, 11, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Monitor Dell UltraSharp 27', '27" 4K USB-C hub monitor for productivity. Can be used as a laptop monitor.', 529.00, 5, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Monitor Samsung Odyssey G7', '32" QHD 240Hz curved gaming monitor.', 549.00, 6, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Monitor ASUS ProArt PA278QV', '27" WQHD color-accurate creative monitor.', 339.00, 9, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Monitor LG UltraFine 27', '27" 5K Thunderbolt monitor for Mac.', 1299.00, 4, 'active');

    IF v_user2_exists THEN
        RAISE NOTICE 'User % already exists', v_user2_id;
    ELSE
        RAISE NOTICE 'Created user % with 25 items', v_user2_id;
    END IF;
END;
$$;


/*
TMP:
SELECT title, description, price, quantity, status FROM sales.items where to_tsvector('english', title || ' ' || description) @@ to_tsquery('english', 'laptop');

SELECT title, description, price, quantity, status, ts_rank(to_tsvector('english', title || ' ' || description), to_tsquery('english', 'laptop')) AS rank FROM sales.items where to_tsvector('english', title || ' ' || description) @@ to_tsquery('english', 'laptop') ORDER BY rank DESC;

SELECT title, description, price, quantity, status, ts_rank(setweight(to_tsvector('english', title), 'A') || setweight(to_tsvector(description), 'b'), to_tsquery('english', 'laptop')) AS rank FROM sales.items where to_tsvector('english', title || ' ' || description) @@ to_tsquery('english', 'laptop') ORDER BY rank DESC;
*/