# README #

Basic screen savers. Bootstrapped by https://github.com/sdbondi/go-glfw3-play/blob/master/main.go

### TODO ###

* Figure out 

* Suggestions from Gerald
    btw, for super fast speeds, look at making a fragment shader.. you only need to render one box, and have that shader draw on that box as a texture
    there's also shadertoy which is an online shader experimentation thing that runs in WebGL

### Docs ###

* http://godoc.org/github.com/go-gl/glfw3
* http://godoc.org/github.com/go-gl/gl
* http://en.wikipedia.org/wiki/Cartesian_coordinate_system
* https://www.youtube.com/watch?v=NSI4KKjYS_w

### Algorithm ###

* 1 - Pixel colour depends on distance from a point (0.0) 
* 2 - Pixel colour is the average of the 3 pixels on the row below + the previous colour of the pixel. See rule 2 at https://www.youtube.com/watch?v=NSI4KKjYS_w

### How do I get set up? ###

* brew install --build-bottle --static glfw3
* brew install glew
* go get

### Descriptions ###

* hello - a hello world
* square - like hello but try a square instead of a triangle

* 8bit-fire - an 8 bit version of a fire.
* fire - The ultimate goal. A realistic looking fire