If you want to learn Go in the best possible way, this is the best place for you :D

To run this project, you can follow these steps:
1. git clone https://github.com/fajaaro/golearn
2. cd golearn
3. copy .env.example to .env
4. go mod tidy
5. go run main.go

With this base code, you can generate simple CRUD API (Get Many, Get One by ID, Store, Update, Delete & Restore) by follow these simple steps:
1. Create your model
2. Add your model to the auto migrate function on config/db.go
3. Create your controller by copy-paste from existing controller (for example product_controller). Then modify your controller as needed (you must modify function and model name)
4. Add routes on routes/routes.go
5. Done!

Key features to help your development:
1. Environment variable management (.env)
2. MVC architecture
3. Constant variable on config/constants.go
4. Middleware (JWT, Permission Management like Laravel Spatie, Checking Super Admin)
5. CRUD Generator (Create, Read, Update, Delete & Import Excel)
6. Function helpers (Upload file, send email, explode string to array, generate random string, etc)