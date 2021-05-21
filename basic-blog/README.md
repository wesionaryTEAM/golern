# Basic Blog
A Basic blog app with Go `net/http` packages and no other dependencies
- The request and responses are written into a file and also read from the file.
- Each file has one blog post with `Title` as the filename and its content as body.

## Learning notes
- Fundamentals of web development with go
    - Using `net/http` package and routing with it
    - URL Path validation with `regexp`
    - Working with templates (template caching and rendering)
- Error handling and reducing code repetition with "function literals" and "closures"
- Working with files (Read and Write to file) with `io/ioutil`