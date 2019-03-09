# gophgen
[![Go Report Card](https://goreportcard.com/badge/github.com/zelr0x/gophgen)](https://goreportcard.com/report/github.com/zelr0x/gophgen)
## A simple placeholder & identicon generator
Placeholders are a necessity when writing web apps. Gophgen gives you a simple API for fast and autonomous placeholder creation right in your html. 
### Usage
Simply run the server and type the `localhost:port/parameters` into `src` attribute of your `img` tags.
Default port is 8080 but you can set it to any unreserved port on startup: `gophgen 5555`.

Change `parameters` to what you want that image to be.
#### Identicon specifier
`http://localhost:8080` - random 64x64 image
`http://localhost:8080/i` - random 64x64 identicon
#### Size
Single number or two numbers separated by 'x' character:
`http://localhost:8080/250` - 250x250 image
`http://localhost:8080/240x100` - 240x100 image
`http://localhost:8080/i/85` - 85x85 identicon (use this format for identicons)
#### Color
One parameter for one color, two parameters for two colors (for identicons).
`http://localhost:8080/200/53f` - 200x200 image with #5533ffff as its color.
`http://localhost:8080/i/120/099b8c/bc3a1abb` - 120x120 identicon with background color #099b8cff and foreground color #bc3a1abb
`http://localhost:8080/i/64/53f/fuchsia` - 64x64 identicon with colors #5533ffff and "fuchsia" (from x11 color palette)
It's really flexible, just try other combinations.
#### Data
Want a unique identicon for a particular user? Just add some seed at the end:
`http://localhost:8080/i/128/099b8c/bc3a1a/d=UserName125` - 128x128 identicon with colors #099b8cff and #bc3a1aff for "UserName125"
