# 🚀 Project Zero

Welcome to **Project Zero** — cool name, right?

This project is a lightweight boilerplate that i created to kickstart new Go-Based web applications without having to rewrite the same code every time. <br>
Think of it as my __starter pack__ for spinning up HTTP servers with templating and styling baked in.

Whether you're prototyping fast or building something small and clean, this base setup will save you time and typing. ⌨️✨

---

## 🧱 Tech Stack Overview

### Backend
- Built with pure **[Go](https://go.dev/doc/effective_go) (Golang)** 🦫 - no frameworks, just standard libraries for full control and simplicity.
- Uses **[Templ](https://templ.guide/)** for HTML templating, enabling efficient and clean server-side rendering.

### Frontend
- Styled with **[SASS](https://sass-lang.com/documentation/)** for flexible and maintainable CSS.

## 🚀 Getting Started with `init.sh`
To customize Project Zero for your own project (e.g., updating module paths), run the included `init.sh` script. It automates replacing placeholders in your `go.mod` and source files.

### How to use
Open Git Bash or your WSL terminal and run:

```bash
./init.sh
```

> [!CAUTION]
> Windows CMD or PowerShell do not support this script natively — use Git Bash or WSL instead.

The script will then ask for your GitHub project link and replace it in all the necessary places.

## 🛠️ Build Instructions
Project Zero includes a build command to prepare your project for production by copying files, minifying JavaScript, and replacing development URLs.

To run the build:
```bash
./build.sh
```
> [!CAUTION]
> Like the [init](#-getting-started-with-initsh) command, it will run on `git bash` or `WSL`.

> [!WARNING] 
> The build process requires the `terser` JavaScript minifier to be installed and available in your system’s PATH.
> `terser` is a Node.js package, so to install it you need to have [Node.js and npm](https://nodejs.org/pt) installed first.
> Then install terser globally with:
> ```bash
> npm install -g terser
> ```
> Without terser installed, the minification step will raise an error, and will not copy the js folder to the build folder.

## ✨ Features
Project Zero comes with a small but powerful set of features to help you build web applications faster:

> [!NOTE]
> ```bash
> dev\
> ├── backend\  
> ├── features\ 
> └── frontend\  
> ```
> All features of **ProjectZero** are contained within this folder. You can easily update to the latest version by simply replacing this folder with the updated one.

- ⚙️ **[Router](#router)** <br>
    A clean and scalable way to group and register your routes using Go.
- 🧩 **[Middlewares](#middlewares)** <br>
    Easily plug in reusable logic before hitting your route handlers — fully composable using middleware chains.

## 📦 What's Inside?

This repo includes everything you need to get up and running:

- ✅ **Boilerplate HTTP server** — fully set up and ready to serve routes
- 📁 **Organized folder structure** — a clear project layout to keep things tidy
- 🔁 **Hot reload support** — auto-compile your Go and SASS files during development
- 🏗️ **Build command** — easily compile your app for production (see above for details)

# 📚 Documentation

This section provides more detailed information on how to use Project Zero.  

Whether you're extending the boilerplate or just trying to understand how things are wired together, this is your go-to reference. 🔧

## Router

In your `app`, the `router` is the component where you define and attach your routes so the application can listen and respond to incoming requests.

To keep things well-organized 🗂️, it's recommended to group related routes into logical folders. One way to do this is by creating a route group using the following structure:

```bash
routes\       # Main folder for all your routes.
│      
├── user\     # Custom folder for routes related to users.  
│   │   
│   │         # These are your defined routes.
│   │   
│   ├── registerRoute.go   
│   └── loginRoute.go
```

### Creating routes

To register routes, create an `init` function inside the file where your endpoint function is defined.

> In Go, `func init()` is a special function that runs automatically when the package is initialized.

Each endpoint should include its own `init` function for route registration.

```go
func init() {
    app.GetInstance().Router.RegisterRoutes("/test/route", "GET", TestPageRoute)
}
```

See more about [middlewares](#️-using-middlewares-in-project-zero).

For each route defined, you’ll need to call its init function in your `main.go` file by importing it like this:

```go
_ "placeholder/dev/backend/routes/your_route"
```

## Middlewares

> [!WARNING]
> Currently, middlewares can **only** be used in chains — even if you're applying just one.

### 🔍 What Are Middlewares?

Middlewares are small functions that run **before** your actual route logic. They're useful for handling common tasks like:

- Authentication ✅  
- CORS headers 🌍  
- Logging 📝  
- Input validation 📋

Let’s say you have several endpoints that require checking for an authentication cookie. Sure, you *could* call an auth function manually at the top of each handler — but that clutters your endpoint logic. Instead, you can use a middleware and attach it directly to the route. Clean and simple.

### 🛠️ Using Middlewares in Project Zero

Let’s revisit the standard route registration (as shown in [Router](#router)):

```go
func init(){
    app.GetInstance().Router.RegisterRoutes("/test/route", "GET", TestPageRoute)
}
```

Now let’s say we want to use a middleware, like CorsMiddleware.
> [!WARNING] 
> Currently, middlewares can **only** be used in chains — even if you're applying just one.

```go
func init() {
    app.GetInstance().Router.RegisterRoutes(
        "/test/route",
        "GET", // The allowed methods
        middlewares.Chain(
            http.HandlerFunc(TestPageRoute), // Your endpoint handler
            middlewares.Auth(), // Middleware(s)
        ).ServeHTTP
    )
}
```

### 🧱 Middleware breakdown

`middlewares.Chain()`

This function wraps your handler in one or more middleware layers. It accepts:

- http.Handler — your actual endpoint.
- http.Methods — the methods that your endpoint allow.
- One or more middleware functions that conform to the Middleware type.

Middleware Example: `CorsMiddleware`
```go
middlewares.CorsMiddleware(app.GetInstance().Router.Routes)
```
> [!NOTE]
> Some middlewares might require parameters (like allowed methods), so be sure to check their function signatures.

### 🧪 Creating Your Own Middleware

You can define custom middleware in the middlewares/ folder. Just follow this structure:

```go
func CustomMiddleware(allowedMethods ...string) Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Your custom logic here

            next.ServeHTTP(w, r.WithContext(r.Context()))
        })
    }
}
```
And that’s it — you've got a clean, reusable middleware 💪

# 👥 Who's Using Project Zero?
Project Zero is already helping developers kickstart their Go-based web applications with ease! Here are a few examples of how **Project Zero** is being used:

🧊 [Cubonauta](https://cubonauta.com): Cubonauta is a platform dedicated to teaching newcomers and helping veterans master the art of solving the Rubik's Cube.

✍️ [Xeubiart](https://xeubiart.com): Xeubiart is a talented tattoo artist based in Portugal, who uses **Project Zero** as the foundation for their personal website.

----

If you're using Project Zero in your project, feel free to share it with us! Drop a message in the issues or a pull request, and we'll be happy to showcase it here. Let’s grow the Project Zero community together! 🚀

<br>
<br>

![Works on My Machine](https://img.shields.io/badge/works-on%20my%20machine-green?style=flat&logo=probot)
