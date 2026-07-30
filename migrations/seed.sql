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


    PERFORM sales.usf_create_item(v_user_id, 'Wireless Headphones', 'Noise-cancelling over-ear headphones with 30h battery.', 79.99, 25, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Mechanical Keyboard', 'Hot-swappable RGB mechanical keyboard, tactile switches.', 129.50, 15, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'USB-C Hub', '7-in-1 hub with HDMI, SD card reader, and PD charging.', 49.00, 40, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Laptop Stand', 'Aluminum adjustable stand for 13–16 inch laptops.', 34.99, 30, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Webcam 1080p', 'Full HD webcam with built-in microphone and privacy shutter.', 59.90, 20, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Desk Lamp', 'LED desk lamp with adjustable brightness and color temperature.', 27.50, 50, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Mouse Pad XL', 'Extended cloth mouse pad with stitched edges.', 18.00, 100, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Portable SSD 1TB', 'USB 3.2 Gen 2 external SSD, up to 1050 MB/s.', 99.99, 12, 'active');
    PERFORM sales.usf_create_item(v_user_id, 'Bluetooth Speaker', 'Compact waterproof speaker with 12h playtime.', 44.95, 35, 'draft');
    PERFORM sales.usf_create_item(v_user_id, 'Phone Tripod', 'Flexible tripod with phone clamp and remote shutter.', 22.00, 45, 'active');

    IF v_user_exists THEN
        RAISE NOTICE 'User % already exists', v_user_id;
    ELSE
        RAISE NOTICE 'Created user % with 10 items', v_user_id;
    END IF;

    PERFORM sales.usf_create_item(v_user2_id, 'Ceramic Mug Set', 'Set of 4 matte ceramic mugs, dishwasher safe.', 32.00, 60, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Cotton Tote Bag', 'Reusable canvas tote with reinforced handles.', 14.50, 80, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Scented Candle', 'Soy wax candle with lavender and cedar notes, 40h burn.', 24.99, 45, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Notebook A5', 'Hardcover dotted notebook, 192 pages, lay-flat binding.', 16.00, 120, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Stainless Water Bottle', 'Insulated 750ml bottle, keeps drinks cold 24h.', 28.90, 55, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Wool Beanie', 'Soft merino wool beanie, one size fits most.', 19.99, 70, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Plant Pot Set', 'Trio of ceramic pots with drainage holes, 10/12/14 cm.', 39.00, 25, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Kitchen Towel Pack', 'Pack of 3 linen kitchen towels, machine washable.', 21.50, 90, 'active');
    PERFORM sales.usf_create_item(v_user2_id, 'Wall Calendar 2026', 'Illustrated wall calendar with monthly layouts.', 12.00, 40, 'draft');
    PERFORM sales.usf_create_item(v_user2_id, 'Bamboo Cutting Board', 'Large bamboo board with juice groove and handles.', 34.00, 30, 'active');

    IF v_user2_exists THEN
        RAISE NOTICE 'User % already exists', v_user2_id;
    ELSE
        RAISE NOTICE 'Created user % with 10 items', v_user2_id;
    END IF;
END;
$$;
