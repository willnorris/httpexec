httpexec executes a specified command in response to HTTP requests and return
the combined stdout and stderr streams.  Requests can optionally require a
password using HTTP BasicAuth.

httpexec is only designed to run a single, specific command and does not take
any input from request.  For example, I use this to [trigger rebuilding][] my
personal hugo-powered website from GitHub Actions.

[trigger rebuilding]: https://github.com/willnorris/willnorris.com/blob/main/.github/workflows/deploy.yml

For example, run httpexec:

    $ httpexec -command "echo hello world" -password secret

A request with the required password receives a 403:

    $ curl -i http://localhost:8080
    
    HTTP/1.1 403 Forbidden
    Date: Thu, 20 May 2021 16:21:34 GMT
    Content-Length: 0

Including the password returns the response of the command:

    $ curl -u :secret -i http://localhost:8080
    
    HTTP/1.1 200 OK
    Date: Thu, 20 May 2021 16:20:00 GMT
    Content-Length: 12
    Content-Type: text/plain; charset=utf-8
    
    hello world

Commands that return a non-zero status result in a 500 response, along with
the command output.

    $ httpexec -command "ls /does/not/exist"
    $ curl -i http://localhost:8080

    HTTP/1.1 500 Internal Server Error
    Date: Thu, 20 May 2021 16:46:32 GMT
    Content-Length: 63
    Content-Type: text/plain; charset=utf-8

    ls: cannot access '/does/not/exist': No such file or directory
