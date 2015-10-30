# Go-Albums

## Description
A hypothetical web app written in [Go](https://golang.org/doc/install) on top of [Martini](http://martini.codegangsta.io/).

How does this differ from most other `go-martini` repos? It integrates [Grunt](http://gruntjs.com/getting-started), [Bower](http://bower.io/), and [AngularJS](https://angularjs.org/).

## System Requirements
1. [Git](https://git-scm.com/downloads): Duh
1. [Nodejs](https://nodejs.org/en/): Make sure you use the following version (4.2.1):
```
C:\Users\Ivan>npm --version
2.14.7
C:\Users\Ivan>node --version
v4.2.1
```
1. [Grunt](http://gruntjs.com/getting-started): `npm install -g grunt-cli`
1. [Bower](http://bower.io/): `npm install -g bower`
1. [Go](https://golang.org/doc/install)
	- We're using [Martini](http://martini.codegangsta.io/), [Docs](https://github.com/go-martini/martini): `go get github.com/go-martini/martini`
	- Auth for Martini: `go get github.com/martini-contrib/auth`

### Dude Why So many freakin' tools???
1. Nodejs: Even though we're using `Go` for our app, several of the client-side toolset uses command line apps that are "extensions" of node. `npm` is Node's 'Node Package Manager' and is the accessor to Node's enormous repository of ready-to-use libraries. You need this to install `Grunt` and `Bower`. Competitors: None (not for what we're using it for)? Maybe Ruby has something similar with gem? Maybe there's a Go task runner?
1. Grunt: Awesome task runner for anything. See the `Gruntfile.js` for an example config. It is used to define everything on our project - How to build our Go code, compile SASS into minified CSS, minify HTML, run Jasmine/Karma tests (unit testing framework for javascript), etc... Competitors: Gulp, etc...
1. Bower: Similar to `npm`, but only for client-side dependencies. This is where we define the version of AngularJS we're using, Bourbon Neat, etc...
1. Go: Uh... we need `Go` to actually go write this `Go` app.
	- Martini: I tried doing HTTP routing in pure Go but there's alot of annoying syntax to deal with. For now, I think Martini is the way to go. See `main.go` for examples on HTTP routes. We can easily apply auth to specific routes which is really what we want.
1. [Bourbon Neat](http://neat.bourbon.io/): This is the `Node` port for the Ruby library. It's a complete framework for building anything the `SASS` way... highly recommended by the best Ruby on Rails guy I know. Fortunately there are `Node` ports for these: `node-sass`, `node-bourbon`, `node-neat`. It provides a grid and several other things required from a UI standpoint to get us building an awesome front-end. This replacement will significantly reduce our SCSS bloat and remove our need for Ruby/Compass.

## Instructions
Pull the project down from git:
```
git clone https://github.com/goalbums.git
```

Navigate to the folder and **Get Everything**
```
npm install
bower update
```

## Other System Changes
Verify that the following System Variables exist:
1. `GOPATH = (path\to\directory\with\all\go\projects)`
1. `GOROOT = C:\Go` (should already exist from Go installation)

Verify that you have the appropriate executables on your PATH (I think this can be the user PATH. Shouldn't have to be a System PATH - and you *don't* need `Python` for some of the `npm` packages - you can ignore those errors. Remember, we're not actually *running* `Node` on production, it's just helping us develop and package our app via `Grunt` and `Bower`):

```
C:\Users\ivan_portugal\AppData\Roaming\npm;C:\Python34;C:\Program Files\nodejs\;C:\Go\bin;C:\Program Files\Git\cmd
```

## Run
```
grunt server
```

## Kudos
Random stuff that has helped me build this project with ease:

- [Yeoman](http://yeoman.io/)
- [Yeoman Angular-Go-Martini](https://github.com/rayokota/generator-angular-go-martini) - This is a big one, duh.
- [martini-api-example](https://github.com/PuerkitoBio/martini-api-example)