from http.server import BaseHTTPRequestHandler
from urllib.parse import urlparse, parse_qs
import logging


def fibonacci(n: int):
    """Return the first `n` Fibonacci numbers."""
    # TODO: Fill this function out
    n1, n2 = 0, 1
    count = 0
    result = []
    # check if the number of terms is valid
    if n <= 0:
      print("Please enter a positive integer")
    # if there is only one term, return n1
    elif n == 1:
      print("Fibonacci sequence upto",n,":")
      #print(n1)
      result.append(n1)
    # generate fibonacci sequence
    else:
      print("Fibonacci sequence:")
      while count < n:
        #print(n1)
        result.append(n1)
        nth = n1 + n2
        # update values
        n1 = n2
        n2 = nth
        count += 1
    return result
    pass


class GetFibs(BaseHTTPRequestHandler):
    def do_GET(self):
        query = urlparse(self.path).query
        params = parse_qs(query)
        print(query)
        print(params)

        if "n" not in params:
            self.send_response(422)
            return

        try:
            key = int(params["n"][0])
        except (IndexError, ValueError):
            self.send_response(422)

        nums = fibonacci(key)

        # convert nums from int to string list
        str_nums = [str(n) for n in nums]
        final_nums = ", ".join(str_nums)

        self.send_response(200)
        self.end_headers()
        self.wfile.write(bytes(final_nums, "UTF-8"))
        return


if __name__ == "__main__":
    from http.server import HTTPServer

    httpd = HTTPServer(("0.0.0.0", 80), GetFibs)
    httpd.serve_forever()
