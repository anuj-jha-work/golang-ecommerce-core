# Go Commerce API 🛒

A RESTful E-Commerce Backend built with Go (Golang). This project covers the core backend building blocks including authentication, product discovery, address management, and shopping cart operations.

## 🚀 Core Features

* **User Authentication:** Secure User Signup and Login using JWT & password hashing.
* **Product Catalogue:** Public APIs to view all products and search products via query parameters, plus protected routes to add new products.
* **Cart Management:** Add items to cart, view/list cart items with total calculation, and remove items from the cart.
* **Address Management:** Full CRUD capability to add, edit, and delete user delivery addresses.
* **Checkout Flow:** Process cart items for final order checkout.

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Router/Framework:** Gin / Chi *(jo aap use kar rahe ho wo likho)*
* **Database:** MongoDB / PostgreSQL *(apna DB likho)*
* **Authentication:** JWT (JSON Web Tokens) & Bcrypt

## 📌 API Endpoints Overview

### Auth & User
* `POST /signup` - Register a new user
* `POST /login` - Authenticate user & return token

### Products
* `GET /products` - View all products (Public)
* `GET /products/search?q=name` - Search products (Public)
* `POST /products` - Add a new product (Protected)

### Cart & Checkout (Protected)
* `POST /cart` - Add item to cart
* `GET /cart` - List all items in user's cart
* `DELETE /cart/:id` - Remove item from cart
* `POST /checkout` - Checkout current cart

### Address (Protected)
* `POST /address` - Add a new delivery address
* `PUT /address/:id` - Edit an existing address
* `DELETE /address/:id` - Delete an address

## ⚙️ How to Run Locally

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/your-username/go-commerce](https://github.com/your-username/go-commerce)