INSERT INTO restaurants (id, name, address, phone_number, description)
VALUES
('550e8400-e29b-41d4-a716-446655440000', 'Bella Italia', '123 Main St, Cityville, CA', '+1-555-123-4567', 'Authentic Italian cuisine'),
('710b962e-041c-11e1-9234-0123456789ab', 'Sushi Sora', '456 Oak Ave, Townsville, NY', '+1-555-987-6543', 'Fresh sushi and Japanese dishes'),
('d3d29d70-1c75-40a7-a2e6-d1c2f3e3f3b3', 'El Camino Cantina', '789 Elm Rd, Villagetown, TX', '+1-555-222-3333', 'Tex-Mex favorites and margaritas'),
('810b962e-041c-11e1-9234-0123456789ab', 'Thai Orchid', '567 Pine Blvd, Hamlet City, FL', '+1-555-888-9999', 'Delicious Thai food and curries'),
('910b962e-041c-11e1-9234-0123456789ab', 'The Steakhouse', '890 Cedar Ln, Mountainview, AZ', '+1-555-444-5555', 'Premium steaks and fine dining'),
('a0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', 'Café de Paris', '234 Birch St, Lakeside, OR', '+1-555-777-8888', 'French bistro and pastries'),
('b0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', 'Mama Mia Pizzeria', '789 Maple Ave, Riverdale, NJ', '+1-555-666-7777', 'Authentic Italian pizzas and pasta'),
('c0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', 'The Burger Joint', '345 Oakwood Dr, Brookside, MA', '+1-555-333-4444', 'Gourmet burgers and fries'),
('f0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', 'Vege Delight', '456 Spruce Ave, Green Valley, NV', '+1-555-777-2222', 'Vegetarian and vegan specialties'),
('c34cd720-4721-4a37-a2f5-a605021528e2', 'Seafood Harbor', '789 Coral Rd, Beachside, CA', '+1-555-444-3333', 'Fresh seafood and sushi bar');

INSERT INTO reservations (id, user_id, restaurant_id, reservation_time, status)
VALUES
('550e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440000', '2024-07-10 18:00:00', 'confirmed'),
('550e8400-e29b-41d4-a716-446655440002', '123e4567-e89b-12d3-a456-426614174000', '710b962e-041c-11e1-9234-0123456789ab', '2024-07-11 19:30:00', 'confirmed'),
('550e8400-e29b-41d4-a716-446655440003', '789abcde-f012-3456-789a-bcdef0123456', 'd3d29d70-1c75-40a7-a2e6-d1c2f3e3f3b3', '2024-07-12 20:00:00', 'cancelled'),
('550e8400-e29b-41d4-a716-446655440004', 'abcdef12-3456-789a-bcde-f01234567890', '810b962e-041c-11e1-9234-0123456789ab', '2024-07-13 12:30:00', 'confirmed'),
('550e8400-e29b-41d4-a716-446655440005', '6dfa25a3-3221-46b1-a163-2e398640fd56', '910b962e-041c-11e1-9234-0123456789ab', '2024-07-14 14:00:00', 'pending'),
('550e8400-e29b-41d4-a716-446655440006', 'a577b9de-0b98-43ce-9133-4f08a0f966c3', 'a0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', '2024-07-15 17:00:00', 'pending'),
('550e8400-e29b-41d4-a716-446655440007', '13579acf-2468-ace0-9753-02468ace9753', 'b0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', '2024-07-16 18:30:00', 'confirmed'),
('550e8400-e29b-41d4-a716-446655440008', 'b39a5c7a-2a2e-4b1b-bc65-444cf80e0590', 'c0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', '2024-07-17 19:00:00', 'confirmed'),
('550e8400-e29b-41d4-a716-446655440009', '3f2b9c86-9f25-4224-93c3-41c1a4b6f0e9', 'f0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', '2024-07-18 13:00:00', 'pending'),
('550e8400-e29b-41d4-a716-446655440010', 'a1b2c3d4-e5f6-7a8b-9c0d-1e2f3a4b5c6d', 'c34cd720-4721-4a37-a2f5-a605021528e2', '2024-07-19 15:30:00', 'confirmed');

INSERT INTO menu (id, restaurant_id, name, description, price)
VALUES
('550e8400-e29b-41d4-a716-446655441001', '550e8400-e29b-41d4-a716-446655440000', 'Margherita Pizza', 'Classic Italian pizza with tomato sauce, mozzarella, and basil', 12.99),
('550e8400-e29b-41d4-a716-446655441002', '810b962e-041c-11e1-9234-0123456789ab', 'Sashimi Platter', 'Assorted slices of fresh raw fish served with soy sauce and wasabi', 24.99),
('550e8400-e29b-41d4-a716-446655441003', 'a0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', 'Pad Thai Noodles', 'Stir-fried rice noodles with shrimp, tofu, bean sprouts, and peanuts', 16.50),
('550e8400-e29b-41d4-a716-446655441004', 'b0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', 'BBQ Bacon Burger', 'Beef patty topped with crispy bacon, cheddar cheese, BBQ sauce, lettuce, and tomato', 14.75),
('550e8400-e29b-41d4-a716-446655441005', 'c0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', 'General Tso Chicken', 'Deep-fried chicken pieces coated in a sweet and spicy sauce', 18.25),
('550e8400-e29b-41d4-a716-446655441006', 'd3d29d70-1c75-40a7-a2e6-d1c2f3e3f3b3', 'Fish Tacos', 'Soft corn tortillas filled with grilled fish, cabbage slaw, avocado, and salsa', 13.50),
('550e8400-e29b-41d4-a716-446655441007', '910b962e-041c-11e1-9234-0123456789ab', 'Vegetarian Curry', 'Mixed vegetables cooked in a creamy coconut curry sauce', 15.99),
('550e8400-e29b-41d4-a716-446655441008', 'f0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44', 'Seafood Paella', 'Traditional Spanish dish with saffron rice, mixed seafood, and chorizo', 28.50),
('550e8400-e29b-41d4-a716-446655441009', 'c34cd720-4721-4a37-a2f5-a605021528e2', 'Caprese Salad', 'Fresh mozzarella, tomatoes, basil, olive oil, and balsamic vinegar', 10.50),
('550e8400-e29b-41d4-a716-446655441010', '550e8400-e29b-41d4-a716-446655440000', 'Tiramisu', 'Classic Italian dessert made with layers of coffee-soaked ladyfingers and mascarpone cheese', 8.99);
