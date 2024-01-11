# List Ops

A small utility web app to make certain list operations useful. 

## Description

This utility currently supports Chunking and set operations between two lists. I wanted to create something simple 
to experiment with HTMX and Templ.

## Getting Started

### Dependencies

* Compiling
  * Go 1.21
  * Make
  * templ
    * Templ can be installed with:
     ```
      go install github.com/a-h/templ/cmd/templ@latest
      ```

### Installing

* You can download the binary for your platform or build with ```make all```

### Executing program

* ``` ./listops```
* Or if you run it with a port 
  * ```
    ADDRESS="localhost:8080" ./listops
    ```

## Authors

Contributors names and contact info

Dex Wood

## Version History

* 2023-01-11
    * Initial release

## License

This project is licensed under the MIT License - see the LICENSE.md file for details

## Acknowledgments

* [Templ](https://github.com/a-h/templ)
* [Gin](https://github.com/gin-gonic/gin)
* [htmx](https://htmx.org/)