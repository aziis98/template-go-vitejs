# Go & ViteJS Boilerplate

This will repo is a template for a larger application based on ideas from <https://slides.com/aziis98/go-and-vite-js>, this uses the following technologies

- The **frontend** is made using **ViteJS** and for now is just a SPA that uses <https://github.com/preactjs/preact-router> (I have a working prototype of a multi-page version with full integration between the dev-client and the server but its still a bit rough)

- The **backend** is made using **Go** and <https://github.com/gofiber/fiber>. Fiber is very similar to ExpressJS and uses the fasthttp Go stack that should be a bit better that the Go default HTTP library.

## Usage

### Development

The following command will start the backend (by default on `:4000`) and the ViteJS development server (by default on `:3000`).

```bash shell
$ go run -v ./cmd/develop
```

### Production

```bash shell
$ npm run build
$ go run -v ./cmd/backend

# or just run...
$ make
```

## Future

- As sad before, polish the ViteJS plugin for the dev-server for server side rendering.

    - Define routes once in Go by just referencing files from `frontend/pages` folder (internally this works by serving a special route in the backend while developing that lists all routes to mount for the ViteJS server)
    
    - Javascript-less pages and view engine support (as before when ViteJS is asked to render a server-dynamic page it will first do its processing to the HTML file and then ask the backend to render just that page on demand)

    - Remove Preact by-default

- Add SQLite by default and more boilerplate for the DB

- Go now supports embedded files in the binary so this could actually be shipped to a single binary with all frontend assets bundled in



